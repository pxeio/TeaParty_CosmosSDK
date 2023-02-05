package party

import (
	"crypto/ecdsa"
	"math/big"
)

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

type AccountResponse struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

type EscrowWallet struct {
	PublicAddress string            `json:"publicAddress"`
	PrivateKey    string            `json:"privateKey"`
	Chain         string            `json:"chain"`
	ECDSA         *ecdsa.PrivateKey `json:"ecdsa"`
}

const (
	OUTCOME_SUCCESS = "success"
	OUTCOME_FAILURE = "failure"
	OUTCOME_TIMEOUT = "timeout"
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

type NKNNotification struct {
	Address    string `json:"address"`
	Amount     string `json:"amount"`
	Network    string `json:"network"`
	PrivateKey string `json:"privateKey"`
	Chain      string `json:"chain"`
	Error      string `json:"error"`
}
