package adams

import (
	"crypto/ecdsa"
	"math/big"

	btcRPC "github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/ethclient"
	solRPC "github.com/gagliardetto/solana-go/rpc"
	kasrpc "github.com/kaspanet/kaspad/infrastructure/network/rpcclient"
	"google.golang.org/grpc"

	"go.uber.org/zap"

	partymodulekeeper "github.com/TeaPartyCrypto/partychain/x/party/keeper"
	partyTypes "github.com/TeaPartyCrypto/partychain/x/party/types"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

// ErrorEvent represents the expected information in an emitted error event
// "tea.party.error"| ERROREVENT
type ErrorEvent struct {
	Err     string
	Context string
	Data    interface{}
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

type envAccessor struct {
	pkgadapter.EnvConfig
	Development bool `envconfig:"DEV" default:"false"`
	Watch       bool `envconfig:"WATCH" default:"true"`

	CELORPC1 string `envconfig:"CELO_RPC_1" default:"" required:"true"`
	CELORPC2 string `envconfig:"CELO_RPC_2" default:"" required:"true"`

	MineOnliumRPC1 string `envconfig:"MO_RPC_1" required:"true"`
	MineOnliumRPC2 string `envconfig:"MO_RPC_2" required:"true"`

	ETHRPC1 string `envconfig:"ETH_RPC_1" default:"" required:"true"`
	ETHRPC2 string `envconfig:"ETH_RPC_2" default:"" required:"true"`

	POLYRPC1 string `envconfig:"POLY_RPC_1" default:"" required:"true"`
	POLYRPC2 string `envconfig:"POLY_RPC_2" default:"" required:"true"`

	SOLRPC1 string `envconfig:"SOL_RPC_1" default:"" required:"true"`
	SOLRPC2 string `envconfig:"SOL_RPC_2" default:"" required:"true"`

	PARTYCHAIN1 string `envconfig:"PARTY_CHAIN_1" default:"" required:"true"`
}

// // SellerNotification represents the information that is to be sent to the seller
// // once a buyer appears for a posted order
// type SellerNotification struct {
// 	Address string `json:"address"` // the address of the seller
// 	Amount  string `json:"amount"`  // the amount of the order
// 	Network string `json:"network"` // the network the order is on
// }

// ExchangeServer holds the state of the exchange server.
type ExchangeServer struct {
	celoNode    EthereumNode
	mineOnlium  EthereumNode
	ethNode     EthereumNode
	polygonNode PolygonNode
	solNode     SOLNode
	partyNode   *grpc.ClientConn

	// temporary state for the database
	orders        []SellOrder
	completOrders []CompletedOrder `json:"completed_orders"`

	partyChainOrders         *partyTypes.QueryAllTradeOrdersResponse
	partyChainCompleteOrders *partyTypes.QueryAllPendingOrdersResponse

	// ordersCollection contains the MongoDB collection for the sell orders.
	// ordersCollection *mongo.Collection
	ordersInProgress []partyTypes.PendingOrders

	// nknClient is the client used to interact with the NKN network.
	// nknClient *nkn.MultiClient

	ceClient cloudevents.Client
	logger   *zap.SugaredLogger
	dev      bool
	watch    bool

	PartyKeeper *partymodulekeeper.Keeper
	sdkContext  sdk.Context
}

// BTCNode hold all the information and interfaces we need to interact
// with the a bitcoin node.
type BTCNode struct {
	rpcClient    *btcRPC.Client
	rpcClientTwo *btcRPC.Client
	rpcConfig    *btcRPC.ConnConfig
	rpcConfigTwo *btcRPC.ConnConfig
}

type SOLNode struct {
	rpcClient    *solRPC.Client
	rpcClientTwo *solRPC.Client
}

type EthereumNode struct {
	rpcClient    *ethclient.Client
	rpcClientTwo *ethclient.Client
}

type PolygonNode struct {
	rpcClient    *ethclient.Client
	rpcClientTwo *ethclient.Client
}

type KaspaNode struct {
	rpcClient    *kasrpc.RPCClient
	rpcClientTwo *kasrpc.RPCClient
}

type AccountGenResponse struct {
	PrivateKey string `json:"privateKey"`
	PubKey     string `json:"publicKey"`
	Address    string `json:"address"`
}

// type AccountDelivery struct {
// 	PrivateKey string `json:"privateKey"`
// 	Chain      string `json:"chain"`
// }

// BuyOrder is a struct that contains the information expected in a buy order
type BuyOrder struct {
	TXID string `json:"txid"`
	// BuyerShippingAddress represents the public key of the account the buyer wants to receive on
	BuyerShippingAddress string `json:"buyerShippingAddress"`
	// BuyerNKNAddress reflects the  publicly address of the buyer.
	BuyerNKNAddress string `json:"buyerNKNAddress"`
	// PaymentTransactionID reflects the transaction ID of the payment made in MO.
	PaymentTransactionID string `json:"paymentTransactionID"`
	// RefundAddress reflects the address of which the funds will be refunded in case of a failure.
	RefundAddress string `json:"refundAddress"`
	// TradeAsset reflects the asset the buyer elected to trade for (mineonlium, bitcoin, USDT, etc).
	// this is an optional field. only avalible when the seller lists "ANY" as the trade asset.
	TradeAsset string `json:"tradeAsset"`
}

// SellOrder contains the information expected in a sell order.
type SellOrder struct {
	// TradeAsset reflects the asset that the SELLER wishes to obtain. (bitcoin, mineonlium, USDT, etc).
	TradeAsset string `json:"tradeAsset"`
	// Price reflects the ammount of TradeAsset the SELLER requires.
	Price *big.Int `json:"price"`
	// Currency reflects the currency that the SELLER wishes to trade. (bitcoin, mineonlium, USDT, etc).
	Currency string `json:"currency"`
	// Amount reflects the ammount of Currency the SELLER wishes to trade.
	Amount *big.Int `json:"amount"`
	// TXID reflects the Transaction ID of the SELL order to be created.
	TXID string `json:"txid"`
	// Locked tells us if this transaction is pending/proccessing another payment.
	Locked bool `json:"locked" default:false`
	// SellerShippingAddress reflects the public key of the account the seller wants to receive on
	SellerShippingAddress string `json:"sellerShippingAddress"`
	// SellerNKNAddress reflects the  public NKN address of the seller.
	SellerNKNAddress string `json:"sellerNKNAddress"`
	// RefundAddress reflects the address of which the funds will be refunded in case of a failure.
	RefundAddress string `json:"refundAddress"`
}

// CompletedTransactionInformation represents data expected when
// describing a transaction that has been completed on-chain.
type CompletedTransactionInformation struct {
	// the transaction id of the completed transaction
	TXID string `json:"txid"`
	// the amount of the transaction
	Amount *big.Int `json:"amount"`
	// the blockchain the transaction was completed on
	Blockchain string `json:"blockchain"`
}

type EscrowWallet struct {
	PublicAddress string            `json:"publicAddress"`
	PrivateKey    string            `json:"privateKey"`
	Chain         string            `json:"chain"`
	ECDSA         *ecdsa.PrivateKey `json:"ecdsa"`
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

// Query contains the information expected in a transaction query
type Query struct {
	TXID string `json:"txid"`
}

type BlockScoutTxQueryResponse struct {
	Message string `json:"message"`
	Result  struct {
		BlockNumber    string        `json:"blockNumber"`
		Confirmations  string        `json:"confirmations"`
		From           string        `json:"from"`
		GasLimit       string        `json:"gasLimit"`
		GasPrice       string        `json:"gasPrice"`
		GasUsed        string        `json:"gasUsed"`
		Hash           string        `json:"hash"`
		Input          string        `json:"input"`
		Logs           []interface{} `json:"logs"`
		NextPageParams interface{}   `json:"next_page_params"`
		RevertReason   string        `json:"revertReason"`
		Success        bool          `json:"success"`
		TimeStamp      string        `json:"timeStamp"`
		To             string        `json:"to"`
		Value          string        `json:"value"`
	} `json:"result"`
	Status string `json:"status"`
}
