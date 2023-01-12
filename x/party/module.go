package party

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gagliardetto/solana-go"
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

func (am AppModule) initMonitor(ctx sdk.Context, order partyTypes.PendingOrders) {
	am.mx.Lock()
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
		return
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
		return
	}

	go am.watchAccount(ctx, buyersAccountWatchRequest)
	go am.watchAccount(ctx, sellersAccountWatchRequest)
	am.mx.Unlock()
}

// BeginBlock contains the logic that is automatically triggered at the beginning of each block
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	po := am.keeper.GetAllPendingOrders(ctx)
	for _, order := range po {
		fmt.Println("order: ", order)
		// TODO:: check if order has expired
		// if it has, then refund the escrowed funds
		// and remove the order from the list of pending orders

		// check the status of every active order
		// TODO:: This does not scale. Come up with a better solution
		// add this to a queue and process it in a separate go routine

		if len(am.ordersInWatch) == 0 {
			// am.keeper.RemovePendingOrders(ctx, order.Index)
			am.wg.Add(1)
			go am.initMonitor(ctx, order)
			am.wg.Wait()
			am.ordersInWatch = append(am.ordersInWatch, order.Index)
		} else {
			found := false
			for _, oiw := range am.ordersInWatch {
				if oiw == order.Index {
					found = true
				}
			}
			if !found {
				// am.keeper.RemovePendingOrders(ctx, order.Index)
				go am.initMonitor(ctx, order)
				am.ordersInWatch = append(am.ordersInWatch, order.Index)
			}
		}

		// look at the list of pending-orders and see if any of them
		// have both buyer and seller payments made
		// if they have, then we need to transfer the funds from the
		// escrow account to the seller and buyer || send the Private keys via NKN
		// and remove the order from the list of pending orders
		if order.BuyerPaymentComplete && order.SellerPaymentComplete {
			// check that the number of blocks since the payments were completed
			// if it has been more then 24 blocks, then create the finalizers
			// if less then 24 blocks, then return
			// currentBlockHeight := int32(ctx.BlockHeight())
			// sellerPaymentCompleteBlockHeight := order.SellerPaymentCompleteBlockHeight
			// buyerPaymentCompleteBlockHeight := order.BuyerPaymentCompleteBlockHeight
			// blocksSinceSellerPaymentComplete := currentBlockHeight - sellerPaymentCompleteBlockHeight
			// blocksSinceBuyerPaymentComplete := currentBlockHeight - buyerPaymentCompleteBlockHeight
			// if blocksSinceSellerPaymentComplete < 24 && blocksSinceBuyerPaymentComplete < 24 {
			// 	// return because the order needs more confirmations
			// 	return
			// }

			// TODO::  Double check that the buyer and seller have the correct amount of funds in their escrow accounts

			// create a new order awaiting finalizer for the buyer
			buyeroaf := partyTypes.OrdersAwaitingFinalizer{
				Index:            order.SellerEscrowWalletPublicKey,
				NknAddress:       order.BuyerNKNAddress,
				WalletPrivateKey: order.SellerEscrowWalletPrivateKey,
				WalletPublicKey:  order.SellerEscrowWalletPublicKey,
				Amount:           order.Amount,
				RefundAddress:    order.BuyerRefundAddress,
				ShippingAddress:  order.BuyerShippingAddress,
				Chain:            order.Currency,
			}

			// create a new order awaiting finalizer for the seller
			selleroaf := partyTypes.OrdersAwaitingFinalizer{
				Index:            order.BuyerEscrowWalletPublicKey,
				NknAddress:       order.SellerNKNAddress,
				WalletPrivateKey: order.BuyerEscrowWalletPrivateKey,
				WalletPublicKey:  order.BuyerEscrowWalletPublicKey,
				Amount:           order.Price,
				RefundAddress:    order.SellerRefundAddress,
				ShippingAddress:  order.SellerShippingAddress,
				Chain:            order.TradeAsset,
			}

			am.keeper.SetOrdersAwaitingFinalizer(ctx, buyeroaf)
			am.keeper.SetOrdersAwaitingFinalizer(ctx, selleroaf)
			am.keeper.RemovePendingOrders(ctx, order.Index)
		}

	}

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
			e := fmt.Sprintf("timeout occured waiting for " + request.Account + " to have a payment of " + request.Amount.String())
			awrr := &AccountWatchRequestResult{
				AccountWatchRequest: request,
				Result:              e,
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
	// update the order in the database
	// po := am.keeper.GetAllPendingOrders(ctx)
	// for _, p := range po {
	// 	if p.Index == awrr.AccountWatchRequest.TransactionID {
	// 		switch awrr.Result {
	// 		case OUTCOME_SUCCESS:
	// 			if !awrr.AccountWatchRequest.Seller {
	// 				p.BuyerPaymentComplete = true
	// 				p.BuyerPaymentCompleteBlockHeight = int32(ctx.BlockHeight())
	// 			} else {
	// 				p.SellerPaymentComplete = true
	// 				p.SellerPaymentCompleteBlockHeight = int32(ctx.BlockHeight())
	// 			}
	// 		case OUTCOME_FAILURE:
	// 			if !awrr.AccountWatchRequest.Seller {
	// 				p.BuyerPaymentComplete = false
	// 			} else {
	// 				p.SellerPaymentComplete = false
	// 			}
	// 		case OUTCOME_TIMEOUT:
	// 			if !awrr.AccountWatchRequest.Seller {
	// 				p.BuyerPaymentComplete = false
	// 			} else {
	// 				p.SellerPaymentComplete = false
	// 			}
	// 		}

	// 		am.keeper.RemovePendingOrders(ctx, p.Index)
	// 		am.keeper.SetPendingOrders(ctx, p)
	// 	}
	// }

	return
}

func (am AppModule) waitAndVerifyEVMChain(ctx sdk.Context, client, client2 *ethclient.Client, request AccountWatchRequest) {

	awrr := &AccountWatchRequestResult{
		AccountWatchRequest: request,
		Result:              "suceess",
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
