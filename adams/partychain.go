package adams

import (
	"context"

	"google.golang.org/grpc"

	partyTypes "github.com/TeaPartyCrypto/partychain/x/party/types"
)

// fetchTradeOrdersFromPartyChain fetches the current list of orders from the TeaParty chain.
func fetchTradeOrdersFromPartyChain(grpc *grpc.ClientConn) (*partyTypes.QueryAllTradeOrdersResponse, error) {
	queryClient := partyTypes.NewQueryClient(grpc)
	queryRes, err := queryClient.TradeOrdersAll(context.Background(), &partyTypes.QueryAllTradeOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return queryRes, nil
}

// fetchPendingOrdersFromPartyChain fetches the current list of pending orders from the TeaParty chain.
func fetchPendingOrdersFromPartyChain(grpc *grpc.ClientConn) (*partyTypes.QueryAllPendingOrdersResponse, error) {
	queryClient := partyTypes.NewQueryClient(grpc)
	queryRes, err := queryClient.PendingOrdersAll(context.Background(), &partyTypes.QueryAllPendingOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return queryRes, nil
}

// notifyPartyChainOfWatchResult calls the update-payment method on the TeaParty chain.
func notifyPartyChainOfWatchResult(grpc *grpc.ClientConn, result *AccountWatchRequestResult) error {
	// import an account

	// create a cosmos client
	// cosmosClient := cosmos.NewClient("http://localhost:1317")

	mup := partyTypes.NewMsgAccountWatchOutcome("party14pjyxktr7tfpwhe4zygwl7lq8znyslacrvqplx", result.AccountWatchRequest.TransactionID, "paymenttransactionID", !result.AccountWatchRequest.Seller, result.Result)
	_, err := partyTypes.NewMsgClient(grpc).AccountWatchOutcome(context.Background(), mup)
	if err != nil {
		return err
	}

	return nil
}
