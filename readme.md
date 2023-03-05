... good job leaking other peoples code


GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build . 

ignite scaffold chain github.com/TeaPartyCrypto/partychainalpha --no-module

cd partychain

ignite scaffold module party -y

ignite scaffold message submit-sell tradeAsset price currency amount sellerShippingAddr sellerNknAddr refundAddr  --module party -y

ignite scaffold message buy txID buyerShippingAddress buyerNKNAddress refundAddress  --module party -y

ignite scaffold map trade-orders tradeAsset price currency amount sellerShippingAddr sellerNknAddr refundAddr --no-message  --module party -y

ignite scaffold map pending-orders buyerEscrowWalletPublicKey buyerEscrowWalletPrivateKey sellerEscrowWalletPublicKey sellerEscrowWalletPrivateKey sellerPaymentComplete:bool sellerPaymentCompleteBlockHeight:int buyerPaymentComplete:bool buyerPaymentCompleteBlockHeight:int amount buyerShippingAddress buyerRefundAddress buyerNKNAddress sellerRefundAddress sellerShippingAddress sellerNKNAddress tradeAsset currency price blockHeight:int --no-message  --module party -y

ignite scaffold map finalizing-orders buyerEscrowWalletPublicKey buyerEscrowWalletPrivateKey sellerEscrowWalletPublicKey sellerEscrowWalletPrivateKey sellerPaymentComplete:bool sellerPaymentCompleteBlockHeight:int buyerPaymentComplete:bool buyerPaymentCompleteBlockHeight:int amount buyerShippingAddress buyerRefundAddress buyerNKNAddress sellerRefundAddress sellerShippingAddress sellerNKNAddress tradeAsset currency price blockHeight:int --no-message  --module party -y


ignite scaffold map orders-awaiting-finalizer nknAddress walletPrivateKey walletPublicKey shippingAddress  refundAddress amount chain --module party  --no-message -y  

ignite scaffold map orders-under-watch nknAddress walletPrivateKey walletPublicKey shippingAddress  refundAddress amount chain startBlockHeight:int paymentComplete:bool  --module party  --no-message -y  

ignite scaffold message account-watch-outcome txID buyer:bool paymentOutcome  --module party -y

ignite scaffold message transaction-result txID outcome --module party -y 

ignite scaffold map transcation-results txID outcome blockHeight:int --module party -y 


ignite scaffold map paid-transactions orderID seller:bool --module party -y 

## Building from cli
ignite scaffold chain github.com/TeaPartyCrypto/partychain --no-module

cd partychain

ignite scaffold module party -y

// submit a sell order
ignite scaffold message submit-sell tradeAsset price currency amount sellerShippingAddr sellerNknAddr refundAddr  --module party -y

// buy a submited order
ignite scaffold message buy txID buyerShippingAddress buyerNKNAddress refundAddress  --module party -y

// update the outcome of an account watch request. This is called from warren/bifrost
ignite scaffold message account-watch-outcome txID buyer:bool paymentOutcome  --module party -y

// account-watch-failure is called when a warren fails to finish watching his account. 
ignite scaffold message account-watch-failure txID  --module party -y


// a map to store open trade orders
ignite scaffold map trade-orders tradeAsset price currency amount sellerShippingAddr sellerNknAddr refundAddr --no-message  --module party -y

// a map to store the pending (or complete) orders awaiting processing
ignite scaffold map pending-orders buyerEscrowWalletPublicKey buyerEscrowWalletPrivateKey sellerEscrowWalletPublicKey sellerEscrowWalletPrivateKey sellerPaymentComplete:bool sellerPaymentCompleteBlockHeight:int buyerPaymentComplete:bool buyerPaymentCompleteBlockHeight:int amount buyerShippingAddress buyerRefundAddress buyerNKNAddress sellerRefundAddress sellerShippingAddress sellerNKNAddress tradeAsset currency price blockHeight:int --no-message  --module party -y

// a map to store the orders that have been approved for payment processing 
ignite scaffold map orders-awaiting-finalizer nknAddress walletPrivateKey walletPublicKey shippingAddress  refundAddress amount chain --module party  --no-message -y  

// ?
ignite scaffold message transaction-result txID outcome --module party -y 

// DO I NEED THESE?
ignite scaffold map payment-updates txID paymentTransactionID buyer:bool paymentOutcome  --module party
// claim-account-watch is called from Warren. this allows for us to know if a order is under watch or not and at what time the claim started.
ignite scaffold message claim-account-watch txID   --module party
// a map to store both "who" is watching an account pair and "when" they started watching it. 
ignite scaffold map account-watch txID timeStamp  --module party



// edit x/party/keeper/msg_server_submit_sell.go
```
package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitSell(goCtx context.Context, msg *types.MsgSubmitSell) (*types.MsgSubmitSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// atempt to find a open sell order from the same seller
	// if found, deny the sell order
	// if not found, create a new sell order
	order, found := k.GetTradeOrders(ctx, msg.Creator)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Account already has an existing open sell order.")
	}

	// create a new sell order
	order = types.TradeOrders{
		Index:              msg.Creator,
		TradeAsset:         msg.TradeAsset,
		Price:              msg.Price,
		Currency:           msg.Currency,
		Amount:             msg.Amount,
		SellerShippingAddr: msg.SellerShippingAddr,
		SellerNknAddr:      msg.SellerNknAddr,
		RefundAddr:         msg.RefundAddr,
	}

	// store the sell order
	k.SetTradeOrders(ctx, order)
	return &types.MsgSubmitSellResponse{}, nil
}

```


// edit x/party/keeper/msg_server_buy.go
```
package keeper

import (
	"context"
	"crypto/ecdsa"

	"github.com/TeaPartyCrypto/partychain/x/party/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (k msgServer) Buy(goCtx context.Context, msg *types.MsgBuy) (*types.MsgBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	// atempt to find a matching sell order
	// if found, start the trade
	// if not found, deny the buy order
	tradeOrder, found := k.GetTradeOrders(ctx, msg.TxID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No matching sell order found.")
	}

	// create a new escrow wallet for the buyer
	// buyerEscrowWallet := "0x0000000000000000000000000000000000000001"
	err, buyerPrivateKey, buyerPublicKey := generateEVMAccount()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Failed to generate escrow wallet for buyer.")
	}

	// create a new escrow wallet for the seller
	// sellerEscrowWallet := "0x0000000000000000000000000000000000000002"
	err, sellerPrivateKey, sellerPublicKey := generateEVMAccount()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Failed to generate escrow wallet for seller.")
	}

	// create a pending-order object
	po := types.PendingOrders{
		Index:                        tradeOrder.Index,
		BuyerEscrowWalletPublicKey:   buyerPublicKey,
		BuyerEscrowWalletPrivateKey:  buyerPrivateKey.D.String(),
		SellerEscrowWalletPublicKey:  sellerPublicKey,
		SellerEscrowWalletPrivateKey: sellerPrivateKey.D.String(),
		SellerPaymentComplete:        false,
		BuyerPaymentComplete:         false,
		Amount:                       tradeOrder.Amount,
		BuyerShippingAddress:         msg.BuyerShippingAddress,
		BuyerRefundAddress:           msg.RefundAddress,
		BuyerNKNAddress:              msg.BuyerNKNAddress,
		SellerRefundAddress:          tradeOrder.RefundAddr,
		SellerShippingAddress:        tradeOrder.SellerShippingAddr,
		SellerNKNAddress:             tradeOrder.SellerNknAddr,
		TradeAsset:                   tradeOrder.TradeAsset,
		Currency:                     tradeOrder.Currency,
		Price:                        tradeOrder.Price,
	}

	// store the pending order
	k.SetPendingOrders(ctx, po)

	// remove the trade order from the store
	k.RemoveTradeOrders(ctx, po.Index)

	return &types.MsgBuyResponse{}, nil
}

// generateEVMAccount generates a new Ethereum account
// returning error, private key, and public address
func generateEVMAccount() (error, *ecdsa.PrivateKey, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return err, nil, ""
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := hexutil.Encode(privateKeyBytes)[2:]
	return nil, privateKey, publicKey
}


```


msg_server_account_watch_outcome
```
package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	OUTCOME_SUCCESS = "success"
	OUTCOME_FAILURE = "failure"
	OUTCOME_TIMEOUT = "timeout"
)

func (k msgServer) AccountWatchOutcome(goCtx context.Context, msg *types.MsgAccountWatchOutcome) (*types.MsgAccountWatchOutcomeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	po := k.GetAllPendingOrders(ctx)
	for _, p := range po {
		if p.Index == msg.TxID {
			switch msg.PaymentOutcome {
			case OUTCOME_SUCCESS:
				if msg.Buyer {
					p.BuyerPaymentComplete = true
					p.BuyerPaymentCompleteBlockHeight = int32(ctx.BlockHeight())
				} else {
					p.SellerPaymentComplete = true
					p.SellerPaymentCompleteBlockHeight = int32(ctx.BlockHeight())
				}
			case OUTCOME_FAILURE:
				if msg.Buyer {
					p.BuyerPaymentComplete = false
				} else {
					p.SellerPaymentComplete = false
				}
			case OUTCOME_TIMEOUT:
				if msg.Buyer {
					p.BuyerPaymentComplete = false
				} else {
					p.SellerPaymentComplete = false
				}
			}

			k.RemovePendingOrders(ctx, p.Index)
			k.SetPendingOrders(ctx, p)
		}
	}
	return &types.MsgAccountWatchOutcomeResponse{}, nil
}
```


//msg_server_transaction_result.go
```
package keeper

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) TransactionResult(goCtx context.Context, msg *types.MsgTransactionResult) (*types.MsgTransactionResultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// remove the order from the list of accounts awaiting finalization
	for _, order := range k.GetAllOrdersAwaitingFinalizer(ctx) {
		if order.Index == msg.TxID {
			k.RemoveOrdersAwaitingFinalizer(ctx, order.Index)
			break
		}
	}
	return &types.MsgTransactionResultResponse{}, nil
}

```

# Test from cli

ignite chain serve


partychaind tx party submit-sell  "ethereum" "1000000000000000000" "polygon" "40000000000000000000" "0x5bbfa5724260Cb175cB39b24802A04c3bfe72eb3" "196845ec8da9904068db299ffba2e037773e319bd648d002695207a64ba5e159" "0x5bbfa5724260Cb175cB39b24802A04c3bfe72eb3" --from alice   -y

partychaind q party list-trade-orders --output json

./partychaind tx party buy "a34eba303ef9734b95710e2b073f57bd0b045e5f7c6f844baf7d91d796318fbe" "0x5bbfa5724260Cb175cB39b24802A04c3bfe72eb3" "1c765347ee40622f90646c1619af9f41377a98a361c917bacfd35d37bbef6538"  "0x5bbfa5724260Cb175cB39b24802A04c3bfe72eb3" --from alice --keyring-backend test -y

partychaind  tx party account-watch-outcome party1w062pa09dqvcrlk6tyl2yreaac7ttgq8avsmxa false success --from alice -y
partychaind  tx party account-watch-outcome party1w062pa09dqvcrlk6tyl2yreaac7ttgq8avsmxa true success --from alice -y

partychaind q party list-pending-orders

partychaind q party list-orders-awaiting-finalizer




## programtaically interact 

```
package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"

	types "github.com/teapartycrypto/testcliimport/types"
)

func main() {
	queryState()
}

func queryState() error {

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",    // your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		return err
	}
	defer grpcConn.Close()

	queryClient := types.NewQueryClient(grpcConn)

	queryRes, err := queryClient.OrdersAll(context.Background(), &types.QueryAllOrdersRequest{})
	if err != nil {
		return err
	}
	fmt.Println(queryRes)

	return nil
}
```

env GOOS=linux GOARCH=amd64 go build -o ./build/partychaind-linux-amd64 ./cmd/partychaind/main.go



## Starting a Boot Node and configuring 

./partychaind init party 

./partychaind keys --keyring-backend file add mac  
K!1poiv^F6fs8XtOdxwqRpp6
- address: party1hedrjppvwuzvpxj2w5rp6xg3gqc0mx7klk8dwl
  name: mac
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A/R1DH0GrxEM+LF5t0mZA6Ol4TaLkZaZNPZClo/PTyO4"}'
  type: local


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

pretty cricket rifle adapt number home chief kick ready tilt silly end promote thing true north fresh autumn try word vacant canoe place unable



./partychaind keys add alice  --keyring-backend test
- address: party1tvddzkq3yvq6q0kmg0hlaszlwwtn74hr5ehu36
  name: alice
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Ah51X9Zxduik5agilBFOcgs6h7M66jIfTUukZkyU7MZU"}'
  type: local


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

note rescue often vocal wedding quality velvet rude trouble trophy humble salute display help tuition scorpion bid level swap upper art lift option sort


./partychaind add-genesis-account party1tvddzkq3yvq6q0kmg0hlaszlwwtn74hr5ehu36 5000000000000000000000stake 
./partychaind add-genesis-account party190z2fe0g96cx63ewqfashuszdcpupaleg7l5lm 5000000000000000000000stake



// Move the gen file to the other Validators

./partychaind gentx mac  100000000stake --chain-id party-001  --keyring-backend file
K!1poiv^F6fs8XtOdxwqRpp6
./partychaind gentx alice  100000000stake   --keyring-backend test



./partychaind collect-gentxs

// Move the gen file to the other Validators

// Edit the config.toml file and add the external_address of the machines ie `IP:PORT`

./partychaind start

// Node the Node ID
// Edit the config.toml file of the peer nodes and add the seed ie <id>@IP:PORT
// also enable the API for the block explorer



## Running in docker

docker run --rm -it \
    -v $(pwd)/docker/bob/:/root/.partychaind/ \
    tmjeff/partychaind  \
    init party-1


mkdir -p docker/bob/keys
echo -n password > docker/bob/keys/passphrase.txt
docker run --rm -it \
    -v $(pwd)/docker/bob:/root/.partychaind \
    tmjeff/partychaind  \
    keys \
    --keyring-backend file --keyring-dir /root/.partychaind/keys \
    add bob --home /root/.partychaind/

- address: cosmos18n4nyucwu436nx4czhg43kdzyh64lfv4wj6cqa
  name: bob
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A7JbUfXrrO1f+F8LokSv4i4nK4f1Uv+jvOy/bSNq15q7"}'
  type: local


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

estate mobile afford toilet love useless few clock exhaust apple taste organ shrimp rally develop cover hunt lift slogan gate hair wonder move rabbit

docker run --rm -it \f
    -v $(pwd)/docker/bob/:/root/.partychaind/ \
    tmjeff/partychaind  \
    gentx cosmos18n4nyucwu436nx4czhg43kdzyh64lfv4wj6cqa 500000000000000tea  --account-number 0 --sequence 0  --chain-id party --pubkey '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A7JbUfXrrO1f+F8LokSv4i4nK4f1Uv+jvOy/bSNq15q7"}'   --gas 1000000 --gas-prices 0.1ngram  --keyring-backend file --home /root/.partychaind/




## Import/Collect Genesis Transactions

partychaind collect-gentxs 


Save / Publish the Genesis file



### Basic client impementation 


```
package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/TeaPartyCrypto/adamsclient/types"
	"github.com/TeaPartyCrypto/partychain/x/party/types"

	"github.com/ignite/cli/ignite/pkg/cosmosclient"
)

func main() {
	addressPrefix := "party"

	// Create a Cosmos client instance
	cosmos, err := cosmosclient.New(
		context.Background(),
		cosmosclient.WithAddressPrefix(addressPrefix),
		// cosmosclient.WithKeyringBackend("file"),
		// cosmosclient.WithKeyringDir("/home/jeff/.partychain"),
		// cosmosclient.WithNodeAddress("http://127.0.0.1:1317"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(cosmos)

	// // Account `alice` was initialized during `ignite chain serve`
	accountName := "alice"

	// Get account from the keyring
	account, err := cosmos.Account(accountName)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := account.Address(addressPrefix)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(addr)

	msg := &types.MsgAccountWatchOutcome{
		Creator:        addr,
		TxID:           "0x1",
		Buyer:          true,
		PaymentOutcome: "success",
	}

	txResp, err := cosmos.BroadcastTx(context.Background(), account, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("MsgAccountWatchOutcome:\n\n")
	fmt.Println(txResp)

}

```

replace github.com/TeaPartyCrypto/partychain => ../partychain

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1


rm -rf ~/.partychain
./partychaind init party 
./partychaind keys add alice  --keyring-backend test 
./partychaind add-genesis-account alice 5000000000000000000000stake 
./partychaind gentx alice  100000000stake   --keyring-backend test
./partychaind collect-gentxs
go run .  start --log_level error
