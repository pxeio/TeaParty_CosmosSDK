package party

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	// this line is used by starport scaffolding # 1

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/TeaPartyCrypto/partychain/x/party/client/cli"
	"github.com/TeaPartyCrypto/partychain/x/party/keeper"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	partyTypes "github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/shopspring/decimal"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	solRPC "github.com/gagliardetto/solana-go/rpc"

	nkn "github.com/nknorg/nkn-sdk-go"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

// CompletedOrder contains all the required elements to complete an order
type CompletedOrder struct {
	// BuyerEscrowWallet the escrow wallet that the buyer will be inserting the
	// TradeAsset into.
	BuyerEscrowWallet EscrowWallet `json:"buyerEscrowWallet"`
	// SellerEscrowWallet the escrow wallet that the seller will be inserting the
	// Currency into.
	SellerEscrowWallet EscrowWallet `json:"sellerEscrowWallet"`
	// SellerPaymentComplete is a boolean that tells us if the seller has completed
	// the payment.
	SellerPaymentComplete bool `json:"sellerPaymentComplete"`
	// BuyerPaymentComplete is a boolean that tells us if the buyer has completed
	// the payment.
	BuyerPaymentComplete bool `json:"buyerPaymentComplete"`
	// Amount the amount of funds that we are sending to the buyer.
	Amount *big.Int `json:"amount"`
	// OrderID the orderID that we are completing.
	OrderID string `json:"orderID"`
	// BuyerShippingAddress the public key of the account the buyer wants to receive on
	BuyerShippingAddress string `json:"buyerShippingAddress"`
	// BuyerRefundAddress
	BuyerRefundAddress string `json:"buyerRefundAddress"`
	// SellerRefundAddress
	SellerRefundAddress string `json:"sellerRefundAddress"`
	// SellerShippingAddress the public key of the account the seller wants to receive on
	SellerShippingAddress string `json:"sellerShippingAddress"`
	// BuyerNKNAddress the public NKN address of the buyer.
	BuyerNKNAddress string `json:"buyerNKNAddress"`
	// SellerNKNAddress the public NKN address of the seller.
	SellerNKNAddress string `json:"sellerNKNAddress"`
	// TradeAsset is the asset that we are sending to the buyer.
	TradeAsset string `json:"tradeAsset"`
	// Currency the currency that we are sending to the seller.
	Currency string `json:"currency"`
	// Price the price of the trade. (how much of the TradeAsset we are asking
	// from the seller for the Currency)
	Price *big.Int `json:"price"`
	// Timeout the amount of time that we are willing to wait for the transaction to be mined.
	Timeout int64 `json:"timeout"`
	// Stage reflects the stage of the order.
	Stage int `json:"stage"`
}

// AccountWatchRequest is the information we need to watch a new account
// this type is associated with the "tea.party.watch.account" | IOWATCHACCOUNTREQUEST event type
type AccountWatchRequest struct {
	Seller        bool     `json:"seller"`
	Account       string   `json:"account"`
	Chain         string   `json:"chain"`
	Amount        *big.Int `json:"amount"`
	TransactionID string   `json:"transaction_id"`
	TimeOut       int64    `json:"timeout"`
}

// AccountWatchRequestResult is the result of the watch request
// this type is associated with the "tea.party.watch.result" | IOWATCHRESULT event type
type AccountWatchRequestResult struct {
	AccountWatchRequest AccountWatchRequest `json:"account_watch_request"`
	Result              string              `json:"result"`
}

type EscrowWallet struct {
	PublicAddress string            `json:"publicAddress"`
	PrivateKey    string            `json:"privateKey"`
	Chain         string            `json:"chain"`
	ECDSA         *ecdsa.PrivateKey `json:"ecdsa"`
}

const (
	ETH = "ethereum"
	MO  = "mineonlium"
	POL = "polygon"
	KAS = "kaspa"
	RXD = "radiant"
	CEL = "celo"
	SOL = "solana"

	// unsupported
	BTC  = "bitcoin"
	ALP  = "alephium"
	LTC  = "litecoin"
	NEAR = "near"
)

type NKNNotification struct {
	Address    string `json:"address"`
	Amount     string `json:"amount"`
	Network    string `json:"network"`
	PrivateKey string `json:"privateKey"`
	Chain      string `json:"chain"`
	Error      string `json:"error"`
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage. The default GenesisState need to be defined by the module developer and is primarily used for testing
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root Tx command for the module. The subcommands of this root command are used by end-users to generate new transactions containing messages defined in the module
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the root query command for the module. The subcommands of this root command are used by end-users to generate new queries to the subset of the state defined by the module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	ordersInWatch []string
	mx            sync.Mutex
	wg            sync.WaitGroup
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// Deprecated: use RegisterServices
func (am AppModule) Route() sdk.Route { return sdk.Route{} }

// Deprecated: use RegisterServices
func (AppModule) QuerierRoute() string { return types.RouterKey }

// Deprecated: use RegisterServices
func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module. It should be incremented on each consensus-breaking change introduced by the module. To avoid wrong/empty versions, the initial version should be set to 1
func (AppModule) ConsensusVersion() uint64 { return 1 }

func (am AppModule) initMonitor(ctx sdk.Context, order partyTypes.PendingOrders) error {
	ta := order.TradeAsset
	const productionTimeLimit = 7200 // 2 hours
	const devTimelimit = 300         // 300 second
	var timeLimit int64
	timeLimit = devTimelimit
	// if e.dev {
	// 	timeLimit = devTimelimit
	// } else {
	// 	timeLimit = productionTimeLimit
	// }

	biPrice := new(big.Int)
	d, _ := decimal.NewFromString(order.Price)
	// if err != nil {
	// 	e.logger.Error("Error converting price to decimal")
	// 	return
	// }
	biPrice = d.BigInt()

	biAmount := new(big.Int)
	dA, _ := decimal.NewFromString(order.Amount)
	// if err != nil {
	// 	e.logger.Error("Error converting amount to decimal")
	// 	return
	// }

	biAmount = dA.BigInt()
	taAmount := biAmount

	// if order.TradeAsset == "ANY" {
	// 	// fetch the market price of the trade asset
	// 	marketPrice, err := FetchMarketPriceInUSD(ta)
	// 	if err != nil {
	// 		e.logger.Error("error fetching market price: " + err.Error())
	// 		return
	// 	}

	// 	e.logger.Infof("market price of %s is: %s", ta, marketPrice)

	// 	// calcuate the amount to send to the buyer
	// 	pgto := big.NewFloat(0).SetInt(biPrice)
	// 	bito := big.NewFloat(0).SetInt(marketPrice)

	// 	// convert to big.int

	// 	fl, _ := pgto.Quo(pgto, bito).Float64()

	// 	// taAmount = FloatToBigInt(fl)
	// 	// TODO: TEST THIS!!! this is a big change from how it was before
	// 	taAmount = big.NewInt(int64(fl * 100000000))
	// 	// taAmount = big.NewInt(int64(fl * 100000000))
	// 	e.logger.Infof("calculated amount to send to buyer: %s", fl)
	// }

	co := &CompletedOrder{
		OrderID:               order.Index,
		BuyerShippingAddress:  order.BuyerShippingAddress,
		SellerShippingAddress: order.SellerShippingAddress,
		TradeAsset:            ta,
		Price:                 biPrice,
		Currency:              order.Currency,
		Amount:                taAmount,
		Timeout:               timeLimit,
		SellerNKNAddress:      order.SellerNKNAddress,
		BuyerNKNAddress:       order.BuyerNKNAddress,
	}

	buyersAccountWatchRequest := &AccountWatchRequest{}
	sellersAccountWatchRequest := &AccountWatchRequest{}

	switch order.Currency {
	case SOL:
		// generate a new solana account for the buyer
		acc := createSolanaAccount()

		co.SellerEscrowWallet = EscrowWallet{
			PublicAddress: acc.PublicKey,
			PrivateKey:    acc.PrivateKey,
			Chain:         SOL,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         SOL,
			Amount:        co.Amount,
			TransactionID: co.OrderID,
			Seller:        true,
		}

	case CEL:
		acc := generateEVMAccount()
		co.SellerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         CEL,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         CEL,
			Amount:        co.Amount,
			TransactionID: co.OrderID,
			Seller:        true,
		}

	case ETH:
		acc := generateEVMAccount()
		co.SellerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         ETH,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         ETH,
			Amount:        co.Amount,
			TransactionID: co.OrderID,
			Seller:        true,
		}

	case POL:
		acc := generateEVMAccount()
		co.SellerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         POL,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         POL,
			Amount:        co.Amount,
			TransactionID: co.OrderID,
			Seller:        true,
		}

	case MO:
		acc := generateEVMAccount()
		co.SellerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         MO,
		}

		if err := notifySellerOfBuyer(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		sellersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.SellerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         MO,
			Amount:        co.Amount,
			TransactionID: co.OrderID,
			Seller:        true,
		}
	default:
		return errors.New("invalid currency")
	}

	switch ta {
	case SOL:
		acc := createSolanaAccount()
		co.BuyerEscrowWallet = EscrowWallet{
			PublicAddress: acc.PublicKey,
			PrivateKey:    acc.PrivateKey,
			Chain:         SOL,
		}

		if err := sendBuyerPayInfo(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         SOL,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.OrderID,
		}

	case MO:
		acc := generateEVMAccount()
		co.BuyerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         MO,
		}

		if err := sendBuyerPayInfo(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}
		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         MO,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.OrderID,
		}

	case ETH:
		acc := generateEVMAccount()
		co.BuyerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         ETH,
		}

		if err := sendBuyerPayInfo(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         ETH,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.OrderID,
		}
	case CEL:
		acc := generateEVMAccount()
		co.BuyerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         CEL,
		}

		if err := sendBuyerPayInfo(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         CEL,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.OrderID,
		}

	case POL:
		acc := generateEVMAccount()
		co.BuyerEscrowWallet = EscrowWallet{
			ECDSA:         acc,
			PublicAddress: crypto.PubkeyToAddress(acc.PublicKey).String(),
			PrivateKey:    hex.EncodeToString(acc.D.Bytes()),
			Chain:         POL,
		}

		if err := sendBuyerPayInfo(*co); err != nil {
			// TODO:: Cancle the order
			return err
		}

		// emit a new event to let Warren know that we need to start watching a new account
		buyersAccountWatchRequest = &AccountWatchRequest{
			Account:       co.BuyerEscrowWallet.PublicAddress,
			TimeOut:       co.Timeout,
			Chain:         POL,
			Amount:        co.Price,
			Seller:        false,
			TransactionID: co.OrderID,
		}
	default:
		return errors.New("invalid currency")
	}

	bouw := partyTypes.OrdersUnderWatch{
		Index:            co.BuyerEscrowWallet.PublicAddress,
		NknAddress:       co.BuyerNKNAddress,
		WalletPrivateKey: co.BuyerEscrowWallet.PrivateKey,
		WalletPublicKey:  co.BuyerEscrowWallet.PublicAddress,
		ShippingAddress:  co.BuyerShippingAddress,
		Amount:           order.Price,
		Chain:            buyersAccountWatchRequest.Chain,
		PaymentComplete:  false,
	}
	am.keeper.SetOrdersUnderWatch(ctx, bouw)

	souw := partyTypes.OrdersUnderWatch{
		Index:            co.SellerEscrowWallet.PublicAddress,
		NknAddress:       co.SellerNKNAddress,
		WalletPrivateKey: co.SellerEscrowWallet.PrivateKey,
		WalletPublicKey:  co.SellerEscrowWallet.PublicAddress,
		ShippingAddress:  co.SellerShippingAddress,
		Amount:           order.Amount,
		Chain:            sellersAccountWatchRequest.Chain,
		PaymentComplete:  false,
	}

	am.keeper.SetOrdersUnderWatch(ctx, souw)

	fmt.Println("checking that the order was set in the store")

	o, ok := am.keeper.GetOrdersUnderWatch(ctx, co.BuyerEscrowWallet.PublicAddress)
	if !ok {
		fmt.Println("order not found in store")
	}

	fmt.Println("order from store: ", o)
	fmt.Println("order from store: ", o)
	fmt.Println("order from store: ", o)
	fmt.Println("order from store: ", o)
	fmt.Println("order from store: ", o)
	fmt.Println("order from store: ", o)
	fmt.Println("order from store: ", o)
	fmt.Println("order from store: ", o)
	fmt.Println("order from store: ", o)
	fmt.Println("order from store: ", o)

	go am.watchAccount(ctx, buyersAccountWatchRequest)
	go am.watchAccount(ctx, sellersAccountWatchRequest)
	return nil
}

// BeginBlock contains the logic that is automatically triggered at the beginning of each block
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	fmt.Println("CURRENT STATE:")
	fmt.Println("trade orders in the store: ", am.keeper.GetAllTradeOrders(ctx))
	fmt.Println("Pending Orders: ", am.keeper.GetAllPendingOrders(ctx))
	fmt.Println("Orders Awaiting Finalizer: ", am.keeper.GetAllOrdersAwaitingFinalizer(ctx))
	fmt.Println("Orders Under Watch: ", am.keeper.GetAllOrdersUnderWatch(ctx))

	po := am.keeper.GetAllPendingOrders(ctx)
	for _, order := range po {
		fmt.Println("order: ", order)
		// TODO:: check if order has expired
		// if it has, then refund the escrowed funds
		// and remove the order from the list of pending orders

		// check the status of every active order
		// TODO:: This does not scale. Come up with a better solution
		// add this to a queue and process it in a separate go routine

		am.keeper.RemovePendingOrders(ctx, order.Index)
		if err := am.initMonitor(ctx, order); err != nil {
			fmt.Println("error: ", err)
		}
	}

	// oaf := am.keeper.GetAllOrdersAwaitingFinalizer(ctx)
	// for _, order := range oaf {
	// 	go am.finalizeOrder(ctx, order)
	// }

}

// sendPrivateKey is called to send the private key of an escrow wallet.
func (am AppModule) sendPrivateKey(order partyTypes.OrdersAwaitingFinalizer) error {
	account, err := nkn.NewAccount(nil)
	if err != nil {
		return err
	}

	// initialize the nkn client.
	nknClient, err := nkn.NewMultiClient(account, "", 4, false, nil)
	if err != nil {
		return err
	}

	ac := &NKNNotification{
		PrivateKey: order.WalletPrivateKey,
		Address:    order.WalletPublicKey,
		Chain:      order.Chain,
	}

	bytes, err := json.Marshal(ac)
	if err != nil {
		return err
	}

	<-nknClient.OnConnect.C
	onReply, err := nknClient.Send(nkn.NewStringArray(order.NknAddress), bytes, nil)
	if err != nil {
		return err
	}
	// create a timeout of 2 minutes
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	select {
	case reply := <-onReply.C:
		switch string(reply.Data) {
		case "ok":
			return nil
		default:
			return fmt.Errorf("buyer has encountered an error receiving sellers escrow wallet private key for order: " + order.Index)
		}
	case <-ctx.Done():
		return fmt.Errorf("buyer has not responded to notification of new PK for order: " + order.Index)
	}
	return nil
}

func (am AppModule) finalizeOrder(ctx sdk.Context, order partyTypes.OrdersAwaitingFinalizer) {
	am.mx.Lock()
	defer am.mx.Unlock()
	if err := am.sendPrivateKey(order); err != nil {
		if err := am.sendFunds(order); err != nil {
			// am.keeper.SetFailedOrders()
			// TODO: we need to notify the party chain that this has happend &|
			// we need to build a reconciler to adjust the parameters in the order
			// and try to force it through again.
			return
		}
	}

	// TODO:: notify the party chain that the order has been finished and the funds have been sent
	// and confirmed

	// if err := notifyPartyChainOfTransactionResult(order.Index, "success"); err != nil {
	// 	e.logger.Error("error notifying the party chain that the transaction was successful: " + err.Error())
	// }

	am.keeper.RemoveOrdersAwaitingFinalizer(ctx, order.Index)

}

func (am AppModule) sendFunds(order partyTypes.OrdersAwaitingFinalizer) error {
	// convert the order.Amount to a big int
	dA, err := decimal.NewFromString(order.Amount)
	if err != nil {
		return err
	}
	biAmount := dA.BigInt()

	curencyFee := new(big.Int).Div(biAmount, big.NewInt(100))
	biAmount.Sub(biAmount, curencyFee)
	assetFee := new(big.Int).Div(biAmount, big.NewInt(100))
	biAmount.Sub(biAmount, assetFee)

	// send the funds
	switch order.Chain {
	case SOL:
		solClient := solRPC.New("https://api.testnet.solana.com")
		_, err := solClient.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
		if err != nil {
			return err
		}

		if err := am.sendCoreSOLAsset(order.WalletPrivateKey, order.ShippingAddress, order.Index, biAmount, solClient); err != nil {
			return err
		}
	case CEL:
		// create a new client for the second node
		// initialize the ethereum nodes.
		celClient1, err := ethclient.Dial("https://celo-alfajores.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		err = am.sendCoreEVMAsset(order.WalletPrivateKey, order.WalletPublicKey, order.ShippingAddress, biAmount, order.Index, celClient1)
		if err != nil {
			return err
		}
	case POL:
		polyClient1, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		err = am.sendCoreEVMAsset(order.WalletPrivateKey, order.WalletPublicKey, order.ShippingAddress, biAmount, order.Index, polyClient1)
		if err != nil {
			return err
		}
	case ETH:
		ethClient1, err := ethclient.Dial("https://goerli.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}

		err = am.sendCoreEVMAsset(order.WalletPrivateKey, order.WalletPublicKey, order.ShippingAddress, biAmount, order.Index, ethClient1)
		if err != nil {
			return err
		}
	}
	return nil
}

// EndBlock contains the logic that is automatically triggered at the end of each block
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func generateEVMAccount() *ecdsa.PrivateKey {
	privateKey, _ := crypto.GenerateKey()
	return privateKey
}

type AccountResponse struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

func createSolanaAccount() AccountResponse {
	account := solana.NewWallet()
	ar := &AccountResponse{
		PrivateKey: account.PrivateKey.String(),
		PublicKey:  account.PublicKey().String(),
	}
	return *ar
}

func notifySellerOfBuyer(co CompletedOrder) error {
	account, err := nkn.NewAccount(nil)
	if err != nil {
		return err
	}

	nknClient, err := nkn.NewMultiClient(account, "", 4, false, nil)
	if err != nil {
		return err
	}
	defer nknClient.Close()

	sn := &NKNNotification{
		Address: co.SellerEscrowWallet.PublicAddress,
		Amount:  co.Amount.String(),
		Network: co.Currency,
	}
	bytes, err := json.Marshal(sn)
	if err != nil {
		return err
	}

	<-nknClient.OnConnect.C
	onReply, err := nknClient.Send(nkn.NewStringArray(co.SellerNKNAddress), bytes, nil)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	select {
	case reply := <-onReply.C:
		switch string(reply.Data) {
		case "ok":
			return nil
		default:
			return fmt.Errorf("seller has encountered an error for order: " + co.OrderID)
		}
	case <-ctx.Done():
		return fmt.Errorf("seller has not responded to notification of new buyer for order: " + co.OrderID)
	}
}

func sendBuyerPayInfo(co CompletedOrder) error {
	account, err := nkn.NewAccount(nil)
	if err != nil {
		return err
	}
	nknClient, err := nkn.NewMultiClient(account, "", 4, false, nil)
	if err != nil {
		return err
	}
	defer nknClient.Close()

	sn := &NKNNotification{
		Address: co.BuyerEscrowWallet.PublicAddress,
		Amount:  co.Price.String(),
		Network: co.TradeAsset,
	}
	bytes, err := json.Marshal(sn)
	if err != nil {
		return err
	}

	<-nknClient.OnConnect.C
	onReply, err := nknClient.Send(nkn.NewStringArray(co.BuyerNKNAddress), bytes, nil)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	select {
	case reply := <-onReply.C:
		switch string(reply.Data) {
		case "ok":
			return nil
		default:
			return fmt.Errorf("buyer has encountered an error for order: " + co.OrderID)
		}
	case <-ctx.Done():
		return fmt.Errorf("buyer has not responded to notification of escrow information for order: " + co.OrderID)
	}
}

func (am AppModule) watchAccount(ctx sdk.Context, awr *AccountWatchRequest) error {
	switch awr.Chain {
	case SOL:
		solClient := solRPC.New("https://api.testnet.solana.com")
		_, err := solClient.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
		if err != nil {
			return err
		}

		solClient2 := solRPC.New("https://api.testnet.solana.com")
		_, err = solClient2.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
		if err != nil {
			return err
		}
		go am.waitAndVerifySOLChain(ctx, *awr, solClient, solClient2)
	case CEL:
		// create a new client for the second node
		// initialize the ethereum nodes.
		celClient1, err := ethclient.Dial("https://celo-alfajores.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		celClient2, err := ethclient.Dial("https://celo-alfajores.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		go am.waitAndVerifyEVMChain(ctx, celClient1, celClient2, *awr)
	case ETH:
		// initialize the ethereum nodes.
		ethClient1, err := ethclient.Dial("https://goerli.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		ethClient2, err := ethclient.Dial("https://goerli.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		go am.waitAndVerifyEVMChain(ctx, ethClient1, ethClient2, *awr)
	case POL:
		// initialize the ethereum nodes.
		polyClient1, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		polyClient2, err := ethclient.Dial("https://polygon-mumbai.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e")
		if err != nil {
			return err
		}
		go am.waitAndVerifyEVMChain(ctx, polyClient1, polyClient2, *awr)
	}
	return nil
}

func (am AppModule) waitAndVerifySOLChain(ctx sdk.Context, request AccountWatchRequest, rpcClient, rpcClientTwo *solRPC.Client) error {
	awrr := &AccountWatchRequestResult{
		AccountWatchRequest: request,
		Result:              OUTCOME_SUCCESS,
	}

	am.dispatch(ctx, awrr)
	return nil

	// the request.Amount is currently in ETH big.Int format convert to uint64
	amount, err := strconv.ParseUint(request.Amount.String(), 10, 64)
	if err != nil {
		return err
	}

	// convert from wei to lamports
	amount = amount / 1000000000

	// create a ticker that ticks every 30 seconds
	// ticker := time.NewTicker(time.Second * 30)
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	// create a timer that times out after the specified timeout
	timer := time.NewTimer(time.Second * time.Duration(request.TimeOut))
	defer timer.Stop()
	// start a for loop that checks the balance of the address
	canILive := true
	for canILive {
		select {
		case <-ticker.C:
			// create new solana public key from string
			pk, err := solana.PublicKeyFromBase58(request.Account)
			if err != nil {
				break
			}

			balance, err := rpcClient.GetBalance(context.Background(), pk, solRPC.CommitmentFinalized)
			if err != nil {
				break
			}

			// if the balance is equal to the amount, verify with the
			// second RPC server.
			if balance.Value >= amount {
				verifiedBalance, err := rpcClientTwo.GetBalance(context.Background(), pk, solRPC.CommitmentFinalized)
				if err != nil {
					break
				}

				if verifiedBalance.Value >= amount {
					// send a complete order event
					awrr := &AccountWatchRequestResult{
						AccountWatchRequest: request,
						Result:              OUTCOME_SUCCESS,
					}

					am.dispatch(ctx, awrr)
					canILive = false
					return nil
				} else {
					break
				}
			}
		case <-timer.C:
			// if the timer times out, return an error
			awrr := &AccountWatchRequestResult{
				AccountWatchRequest: request,
				Result:              OUTCOME_TIMEOUT,
			}

			am.dispatch(ctx, awrr)
			canILive = false
			return nil
		}
	}
	return nil
}

const (
	OUTCOME_SUCCESS = "success"
	OUTCOME_FAILURE = "failure"
	OUTCOME_TIMEOUT = "timeout"
)

func (am AppModule) dispatch(ctx sdk.Context, awrr *AccountWatchRequestResult) {
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	fmt.Println("dispatched : " + awrr.AccountWatchRequest.TransactionID)
	ouw, ok := am.keeper.GetOrdersUnderWatch(ctx, awrr.AccountWatchRequest.Account)
	if !ok {
		return
	}

	fmt.Println("ORDRES UNDER WATCH")
	fmt.Println(ouw)

	switch awrr.Result {
	case OUTCOME_SUCCESS:
		fmt.Println("success")
		ouw.PaymentComplete = true
		// p.PaymentCompleteBlockHeigh = int32(ctx.BlockHeight())
	case OUTCOME_FAILURE:
		ouw.PaymentComplete = false
	case OUTCOME_TIMEOUT:
		ouw.PaymentComplete = false
	}

	oaf := partyTypes.OrdersAwaitingFinalizer{
		Index:            ouw.Index,
		NknAddress:       ouw.NknAddress,
		WalletPrivateKey: ouw.WalletPrivateKey,
		WalletPublicKey:  ouw.WalletPublicKey,
		ShippingAddress:  ouw.ShippingAddress,
		RefundAddress:    ouw.RefundAddress,
		Amount:           ouw.Amount,
		Chain:            ouw.Chain,
	}

	fmt.Println(oaf)

	// goctx := sdk.UnwrapSDKContext(ctx)
	am.keeper.SetOrdersAwaitingFinalizer(ctx, oaf)
	am.keeper.RemoveOrdersUnderWatch(ctx, ouw.Index)

	// po := am.keeper.GetAllOrdersUnderWatch(ctx)
	// fmt.Printf("orders under watch: %+v", po)
	// for _, p := range po {
	// 	if p.Index == awrr.AccountWatchRequest.Account {
	// 		switch awrr.Result {
	// 		case OUTCOME_SUCCESS:
	// 			p.PaymentComplete = true
	// 			// p.PaymentCompleteBlockHeigh = int32(ctx.BlockHeight())
	// 		case OUTCOME_FAILURE:
	// 			p.PaymentComplete = false
	// 		case OUTCOME_TIMEOUT:
	// 			p.PaymentComplete = false
	// 		}

	// 		oaf := partyTypes.OrdersAwaitingFinalizer{
	// 			Index:            p.WalletPublicKey,
	// 			NknAddress:       p.NknAddress,
	// 			WalletPrivateKey: p.WalletPrivateKey,
	// 			WalletPublicKey:  p.WalletPublicKey,
	// 			ShippingAddress:  p.ShippingAddress,
	// 			RefundAddress:    p.RefundAddress,
	// 			Amount:           p.Amount,
	// 			Chain:            p.Chain,
	// 		}
	// 		am.keeper.SetOrdersAwaitingFinalizer(ctx, oaf)
	// 		am.keeper.RemoveOrdersUnderWatch(ctx, p.Index)
	// 	}
	// }
}

func (am AppModule) waitAndVerifyEVMChain(ctx sdk.Context, client, client2 *ethclient.Client, request AccountWatchRequest) {
	awrr := &AccountWatchRequestResult{
		AccountWatchRequest: request,
		Result:              OUTCOME_SUCCESS,
	}

	am.dispatch(ctx, awrr)

	return

	// create a ticker that ticks every 30 seconds
	// ticker := time.NewTicker(time.Second * 30)

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	// create a timer that times out after the specified timeout
	timer := time.NewTimer(time.Second * time.Duration(request.TimeOut))
	defer timer.Stop()

	account := common.HexToAddress(request.Account)

	// start a for loop that checks the balance of the address
	canILive := true
	for canILive {
		select {
		case <-ticker.C:
			balance, err := client.BalanceAt(context.Background(), account, nil)
			if err != nil {
				continue
			}
			// if the balance is equal to the amount, verify with the
			// second RPC server.
			if balance.Cmp(request.Amount) == 0 || balance.Cmp(request.Amount) == 1 {
				verifiedBalance, err := client2.BalanceAt(context.Background(), account, nil)
				if err != nil {
					continue
				}

				if verifiedBalance.Cmp(request.Amount) == 0 || verifiedBalance.Cmp(request.Amount) == 1 {
					// send a complete order event
					awrr := &AccountWatchRequestResult{
						AccountWatchRequest: request,
						Result:              OUTCOME_SUCCESS,
					}

					am.dispatch(ctx, awrr)
					canILive = false
					return
				} else {
					return
				}
			}
		case <-timer.C:
			// if the timer times out, return an error
			// if the timer times out, return an error
			awrr := &AccountWatchRequestResult{
				AccountWatchRequest: request,
				Result:              OUTCOME_TIMEOUT,
			}

			am.dispatch(ctx, awrr)
			canILive = false
			return
		}
	}
	return
}

func (am AppModule) sendCoreSOLAsset(fromWalletPrivateKey, toAddress, txid string, amount *big.Int, rpcClient *solRPC.Client) error {
	privateKey, err := solana.PrivateKeyFromBase58(fromWalletPrivateKey)
	if err != nil {
		return err
	}

	toAddressPublicKey, err := solana.PublicKeyFromBase58(toAddress)
	if err != nil {
		return err
	}

	recent, err := rpcClient.GetRecentBlockhash(context.Background(), solRPC.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	// big.int to lamports
	amountLamparts := amount.Mul(amount, big.NewInt(1000000000))

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amountLamparts.Uint64(),
				privateKey.PublicKey(),
				toAddressPublicKey,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(privateKey.PublicKey()),
	)
	if err != nil {
		return err
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if privateKey.PublicKey().Equals(key) {
				return &privateKey
			}
			return nil
		},
	)
	if err != nil {
		return err
	}

	// TODO: Migrate to ws client so we can use the sendandconfirmtransaction method
	// Send transaction, and wait for confirmation:
	opts := solRPC.TransactionOpts{}
	_, err = rpcClient.SendTransactionWithOpts(
		context.Background(),
		tx,
		opts,
	)
	if err != nil {
		return err
	}
	return nil
}

func (am AppModule) sendCoreEVMAsset(walletPrivK, walletPubK, toAddress string, amount *big.Int, txid string, rpcClient *ethclient.Client) error {
	// view the current balance of the paying wallet
	ecpk := ecdsa.PublicKey{}
	ecpk.X, ecpk.Y = elliptic.Unmarshal(crypto.S256(), common.FromHex(walletPubK))
	account := crypto.PubkeyToAddress(ecpk)
	// send the currency to the buyer
	// read nonce
	nonce, err := rpcClient.PendingNonceAt(context.Background(), account)
	if err != nil {
		return err
	}

	// create gas params
	gasLimit := uint64(31000) // in units
	gasPrice, err := rpcClient.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	// convert the string address to an address
	qualifiedAddress := common.HexToAddress(toAddress)

	// create a transaction
	tx := ethTypes.NewTransaction(nonce, qualifiedAddress, amount, gasLimit, gasPrice, nil)

	// fetch chain id
	chainID, err := rpcClient.NetworkID(context.Background())
	if err != nil {
		return err
	}

	//convert from a private key in a string to a *ecdsa.PrivateKey
	fromWallet, err := crypto.HexToECDSA(walletPrivK)
	if err != nil {
		return err
	}

	// sign the transaction
	signedTx, err := ethTypes.SignTx(tx, ethTypes.NewEIP155Signer(chainID), fromWallet)
	if err != nil {
		return err
	}

	// send the transaction
	err = rpcClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	return nil
}
