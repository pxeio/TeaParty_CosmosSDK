package cli

import (
	"strconv"

	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdAccountWatchOutcome() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "account-watch-outcome [tx-id] [buyer] [payment-outcome]",
		Short: "Broadcast message account-watch-outcome",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTxID := args[0]
			argBuyer, err := cast.ToBoolE(args[1])
			if err != nil {
				return err
			}
			argPaymentOutcome := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgAccountWatchOutcome(
				clientCtx.GetFromAddress().String(),
				argTxID,
				argBuyer,
				argPaymentOutcome,
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
