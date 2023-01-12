package cli

import (
    "context"
	
    "github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/TeaPartyCrypto/partychain/x/party/types"
)

func CmdListOrdersUnderWatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-orders-under-watch",
		Short: "list all orders-under-watch",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllOrdersUnderWatchRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.OrdersUnderWatchAll(context.Background(), params)
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

func CmdShowOrdersUnderWatch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-orders-under-watch [index]",
		Short: "shows a orders-under-watch",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argIndex := args[0]
            
            params := &types.QueryGetOrdersUnderWatchRequest{
                Index: argIndex,
                
            }

            res, err := queryClient.OrdersUnderWatch(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
