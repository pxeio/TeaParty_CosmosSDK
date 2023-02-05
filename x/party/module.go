package party

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"sync"

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

	"github.com/ethereum/go-ethereum/crypto"
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
	// timeLimit = devTimelimit
	// if e.dev {
	// 	timeLimit = devTimelimit
	// } else {
	timeLimit = productionTimeLimit
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
			Account:       co.SellerNKNAddress,
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
			Account:       co.SellerNKNAddress,
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
			Account:       co.SellerNKNAddress,
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
			Account:       co.SellerNKNAddress,
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
			Account:       co.SellerNKNAddress,
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
			Account:       co.BuyerNKNAddress,
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
			Account:       co.BuyerNKNAddress,
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
			Account:       co.BuyerNKNAddress,
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
			Account:       co.BuyerNKNAddress,
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
			Account:       co.BuyerNKNAddress,
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
		Index:            co.BuyerNKNAddress,
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
		Index:            co.SellerNKNAddress,
		NknAddress:       co.SellerNKNAddress,
		WalletPrivateKey: co.SellerEscrowWallet.PrivateKey,
		WalletPublicKey:  co.SellerEscrowWallet.PublicAddress,
		ShippingAddress:  co.SellerShippingAddress,
		Amount:           order.Amount,
		Chain:            sellersAccountWatchRequest.Chain,
		PaymentComplete:  false,
	}

	am.keeper.SetOrdersUnderWatch(ctx, souw)
	go am.watchAccount(ctx, buyersAccountWatchRequest)
	go am.watchAccount(ctx, sellersAccountWatchRequest)
	return nil
}

// BeginBlock contains the logic that is automatically triggered at the beginning of each block
func (am AppModule) BeginBlock(ctx sdk.Context, _ abci.RequestBeginBlock) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("CURRENT STATE:")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("trade orders in the store: ")
	// fmt.Println("trade orders in the store: ", am.keeper.GetAllTradeOrders(ctx))
	for _, order := range am.keeper.GetAllTradeOrders(ctx) {
		b, err := json.MarshalIndent(order, "", "  ")
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println(string(b))
	}
	fmt.Println("")
	fmt.Println("")
	// fmt.Println("Pending Orders: ", am.keeper.GetAllPendingOrders(ctx))
	fmt.Println("Pending Orders: ")
	for _, order := range am.keeper.GetAllPendingOrders(ctx) {
		b, err := json.MarshalIndent(order, "", "  ")
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println(string(b))
	}
	fmt.Println("")
	fmt.Println("")
	// fmt.Println("Orders Awaiting Finalizer: ", am.keeper.GetAllOrdersAwaitingFinalizer(ctx))
	fmt.Println("Orders Awaiting Finalizer: ")
	for _, order := range am.keeper.GetAllOrdersAwaitingFinalizer(ctx) {
		b, err := json.MarshalIndent(order, "", "  ")
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println(string(b))
	}
	fmt.Println("")
	fmt.Println("")
	// fmt.Println("Orders Under Watch: ", am.keeper.GetAllOrdersUnderWatch(ctx))
	fmt.Println("Orders Under Watch: ")
	// pretty print the orders under watch
	for _, order := range am.keeper.GetAllOrdersUnderWatch(ctx) {
		b, err := json.MarshalIndent(order, "", "  ")
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println(string(b))
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Complete orders in Finalizing Orders: ")
	for _, order := range am.keeper.GetAllFinalizingOrders(ctx) {
		b, err := json.MarshalIndent(order, "", "  ")
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println(string(b))
	}
	fmt.Println("")
	fmt.Println("")

	po := am.keeper.GetAllPendingOrders(ctx)
	for _, order := range po {
		// TODO:: check if order has expired
		// if it has, then refund the escrowed funds
		// and remove the order from the list of pending orders

		// check the status of every active order
		// TODO:: This does not scale. Come up with a better solution
		// add this to a queue and process it in a separate go routine
		// convert the order to an order awaiting finalizer
		// and add it to the list of orders awaiting finalizer

		am.keeper.SetFinalizingOrders(ctx, partyTypes.FinalizingOrders{
			Index:                            order.Index,
			BuyerEscrowWalletPublicKey:       order.BuyerEscrowWalletPublicKey,
			BuyerEscrowWalletPrivateKey:      order.BuyerEscrowWalletPrivateKey,
			SellerEscrowWalletPublicKey:      order.SellerEscrowWalletPublicKey,
			SellerEscrowWalletPrivateKey:     order.SellerEscrowWalletPrivateKey,
			SellerPaymentComplete:            false,
			BuyerPaymentComplete:             false,
			Amount:                           order.Amount,
			Price:                            order.Price,
			Currency:                         order.Currency,
			TradeAsset:                       order.TradeAsset,
			BlockHeight:                      order.BlockHeight,
			SellerPaymentCompleteBlockHeight: order.SellerPaymentCompleteBlockHeight,
			BuyerPaymentCompleteBlockHeight:  order.BuyerPaymentCompleteBlockHeight,
			BuyerRefundAddress:               order.BuyerRefundAddress,
			SellerRefundAddress:              order.SellerRefundAddress,
			BuyerNKNAddress:                  order.BuyerNKNAddress,
			SellerNKNAddress:                 order.SellerNKNAddress,
			BuyerShippingAddress:             order.BuyerShippingAddress,
			SellerShippingAddress:            order.SellerShippingAddress,
		})
		am.keeper.RemovePendingOrders(ctx, order.Index)
		if err := am.initMonitor(ctx, order); err != nil {
			fmt.Println("error: ", err)
		}
	}

	oaf := am.keeper.GetAllOrdersAwaitingFinalizer(ctx)
	for _, order := range oaf {
		fmt.Printf("Order %s is awaiting finalization", order.Index)
		fmt.Println("sending order to finalizer")
		if err := am.finalizeOrder(ctx, order); err != nil {
			// TODO: handle error
			fmt.Println("error: ", err)
		}
	}

}

// EndBlock contains the logic that is automatically triggered at the end of each block
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
