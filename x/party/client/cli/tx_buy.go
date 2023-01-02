package cli

import (
	"strconv"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBuy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy [tx-id] [buyer-shipping-address] [buyer-nkn-address] [refund-address]",
		Short: "Broadcast message buy",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTxID := args[0]
			argBuyerShippingAddress := args[1]
			argBuyerNKNAddress := args[2]
			argRefundAddress := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuy(
				clientCtx.GetFromAddress().String(),
				argTxID,
				argBuyerShippingAddress,
				argBuyerNKNAddress,
				argRefundAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
