package adams

import (
	"context"
	"os"
	"strconv"

	zap "go.uber.org/zap"

	"github.com/ethereum/go-ethereum/ethclient"
	solRPC "github.com/gagliardetto/solana-go/rpc"
)

func NewAdams() (*ExchangeServer, error) {
	e := &ExchangeServer{}
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	e.logger = logger.Sugar()
	// initialize the ethereum nodes.
	ethClient1, err := ethclient.Dial(os.Getenv("ETH_RPC_1"))
	if err != nil {
		return nil, err
	}
	ethClient2, err := ethclient.Dial(os.Getenv("ETH_RPC_1"))
	if err != nil {
		return nil, err
	}

	// initialize the polygon nodes.
	polyclient, err := ethclient.Dial(os.Getenv("POLY_RPC_1"))
	if err != nil {
		return nil, err
	}
	polyclientTwo, err := ethclient.Dial(os.Getenv("POLY_RPC_2"))
	if err != nil {
		return nil, err
	}

	celo, err := ethclient.Dial(os.Getenv("CELO_RPC_1"))
	if err != nil {
		return nil, err
	}
	celo2, err := ethclient.Dial(os.Getenv("CELO_RPC_1"))
	if err != nil {
		return nil, err
	}

	// initialize the solana nodes.
	solClient := solRPC.New(os.Getenv("SOL_RPC_1"))

	// test connection
	_, err = solClient.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
	if err != nil {
		return nil, err
	}

	solClient2 := solRPC.New(os.Getenv("SOL_RPC_2"))

	// test connection
	_, err = solClient2.GetRecentBlockhash(context.Background(), solRPC.CommitmentRecent)
	if err != nil {
		return nil, err
	}

	// ptyChain := os.Getenv("PARTY_CHAIN_1")
	// // Create a connection to the gRPC server.
	// grpcConn, err := grpc.Dial(
	// 	ptyChain,            // your gRPC server address.
	// 	grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	// 	// This instantiates a general gRPC codec which handles proto bytes. We pass in a nil interface registry
	// 	// if the request/response types contain interface instead of 'nil' you should pass the application specific codec.
	// 	grpc.WithDefaultCallOptions(grpc.ForceCodec(codec.NewProtoCodec(nil).GRPCCodec())),
	// )
	// if err != nil {
	// 	return nil, err
	// }

	e.celoNode.rpcClient = celo
	e.celoNode.rpcClientTwo = celo2
	e.ethNode.rpcClient = ethClient1
	e.ethNode.rpcClientTwo = ethClient2
	e.polygonNode.rpcClient = polyclient
	e.polygonNode.rpcClientTwo = polyclientTwo
	e.solNode.rpcClient = solClient
	e.solNode.rpcClientTwo = solClient2
	e.watch, _ = strconv.ParseBool(os.Getenv("WATCH"))
	e.dev, _ = strconv.ParseBool(os.Getenv("DEV"))

	return e, nil
}

func (e *ExchangeServer) Dispatch(awrr *AccountWatchRequestResult) error {
	// remove the order from ordersInProgress list
	for i, order := range e.ordersInProgress {
		if order.Index == awrr.AccountWatchRequest.TransactionID {
			e.ordersInProgress = append(e.ordersInProgress[:i], e.ordersInProgress[i+1:]...)
			break
		}
	}

	// notify the party chain of the transaction outcome
	if err := e.notifyPartyChainOfWatchResult(awrr); err != nil {
		e.logger.Errorw("failed to notify the party chain of the watch result", err)
		return err
	}

	return nil
}
