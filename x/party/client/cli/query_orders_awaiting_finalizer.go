package cli

import (
	"context"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListOrdersAwaitingFinalizer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-orders-awaiting-finalizer",
		Short: "list all orders-awaiting-finalizer",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrdersAwaitingFinalizerRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OrdersAwaitingFinalizerAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowOrdersAwaitingFinalizer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-orders-awaiting-finalizer [index]",
		Short: "shows a orders-awaiting-finalizer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argIndex := args[0]

			params := &types.QueryGetOrdersAwaitingFinalizerRequest{
				Index: argIndex,
			}

			res, err := queryClient.OrdersAwaitingFinalizer(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
