package adams

import (
	"context"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	solRPC "github.com/gagliardetto/solana-go/rpc"
	"google.golang.org/grpc"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	pkgadapter "knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	"github.com/cosmos/cosmos-sdk/codec"
)

// EnvAccessorCtor for configuration parameters
func EnvAccessorCtor() pkgadapter.EnvConfigAccessor {
	return &envAccessor{}
}

var _ pkgadapter.Adapter = (*ExchangeServer)(nil)

// NewAdapter adapter implementation
func NewAdapter(ctx context.Context, envAcc pkgadapter.EnvConfigAccessor, ceClient cloudevents.Client) pkgadapter.Adapter {
	env := envAcc.(*envAccessor)
	e := &ExchangeServer{}
	e.logger = logging.FromContext(ctx)

	// initialize the MO nodes.
	Moclient, err := ethclient.Dial(env.MineOnliumRPC1)
	if err != nil {
		e.logger.Errorw("holy fucknuts batman!! The RPC for Mineonlium is down!!")
		if !env.Development {
			panic(err)
			return nil
		}
	}
	MoclientTwo, err := ethclient.Dial(env.MineOnliumRPC2)
	if err != nil {
		e.logger.Errorw("holy fucknuts batman!! The Secondary RPC for Mineonlium is down!!")
		if !env.Development {
			panic(err)
			return nil
		}
	}

	// initialize the ethereum nodes.
	ethClient1, err := ethclient.Dial(env.ETHRPC1)
	if err != nil {
		e.logger.Errorw("damm son no connection to eth rpc 1 ")
		if !env.Development {
			panic(err)
			return nil
		}
	}
	ethClient2, err := ethclient.Dial(env.ETHRPC2)
	if err != nil {
		e.logger.Errorw("damm son no connection to eth rpc 2 ")
		if !env.Development {
			panic(err)
			return nil
		}
	}

	// initialize the polygon nodes.
	polyclient, err := ethclient.Dial(env.POLYRPC1)
	if err != nil {
		e.logger.Errorw("damm son no connection to polygon rpc 1 ")
		if !env.Development {
			panic(err)
			return nil
		}
	}
	polyclientTwo, err := ethclient.Dial(env.POLYRPC2)
	if err != nil {
		e.logger.Errorw("damm son no connection to polygon rpc 2 ")
		if !env.Development {
			panic(err)
			return nil
		}
	}

	celo, err := ethclient.Dial(env.CELORPC1)
	if err != nil {
		e.logger.Errorw("damm son no connection to celo rpc 1 ")
		if !env.Development {
			panic(err)
			return nil
		}
	}
	celo2, err := ethclient.Dial(env.CELORPC2)
	if err != nil {
		e.logger.Errorw("damm son no connection to celo rpc 2 ")
		if !env.Development {
			panic(err)
			return nil
		}
	}

	// initialize the solana nodes.
	solClient := solRPC.New(env.SOLRPC1)

	// test connection
	_, err = solClient.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
	if err != nil {
		e.logger.Errorw("damm son no connection to solana rpc 1 ")
		if !env.Development {
			panic(err)
			return nil
		}
	}

	solClient2 := solRPC.New(env.SOLRPC2)

	// test connection
	_, err = solClient2.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
	if err != nil {
		e.logger.Errorw("damm son no connection to solana rpc 1 ")
		if !env.Development {
			panic(err)
			return nil
		}
	}

	// Create a connection to the gRPC server.
	grpcConn, err := grpc.Dial(
		env.PARTYCHAIN1,     // your gRPC server address.
		grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
		// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
		// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
		grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	)
	if err != nil {
		e.logger.Errorw("damm son no connection to party chain 1 ")
		if !env.Development {
			panic(err)
			return nil
		}
	}

	e.celoNode.rpcClient = celo
	e.celoNode.rpcClientTwo = celo2
	e.mineOnlium.rpcClient = Moclient
	e.mineOnlium.rpcClientTwo = MoclientTwo
	e.ethNode.rpcClient = ethClient1
	e.ethNode.rpcClientTwo = ethClient2
	e.polygonNode.rpcClient = polyclient
	e.polygonNode.rpcClientTwo = polyclientTwo
	e.solNode.rpcClient = solClient
	e.solNode.rpcClientTwo = solClient2
	e.partyNode = grpcConn

	e.watch = env.Watch
	e.dev = env.Development
	e.ceClient = ceClient

	orders, err := fetchTradeOrdersFromPartyChain(grpcConn)
	if err != nil {
		e.logger.Errorw("failed to fetch orders from party chain")
		if !env.Development {
			panic(err)
			return nil
		}
	}

	e.logger.Info("Fetched Current Orders from the Party Chain:")
	e.logger.Infof("Orders: %+v", orders.TradeOrders)
	e.partyChainOrders = orders

	completeOrders, err := fetchPendingOrdersFromPartyChain(grpcConn)
	if err != nil {
		e.logger.Errorw("failed to fetch complete orders from party chain")
		if !env.Development {
			panic(err)
			return nil
		}
	}

	e.logger.Info("Fetched Current Pending Orders from the Party Chain:")
	e.logger.Infof("Pending Orders: %+v", completeOrders.PendingOrders)
	e.partyChainCompleteOrders = completeOrders

	return e
}

func (e *ExchangeServer) Start(ctx context.Context) error {
	e.logger.Info("starting warren..")

	e.logger.Info("starting the blockscanners")
	e.Watch(ctx)

	http.HandleFunc("/listorders", e.FetchTradeOrders)
	e.logger.Info("starting sentry service on port :8080")
	return http.ListenAndServe(":8080", nil)
}
