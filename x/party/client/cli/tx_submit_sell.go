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

func CmdSubmitSell() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-sell [trade-asset] [price] [currency] [amount] [seller-shipping-addr] [seller-nkn-addr] [refund-addr]",
		Short: "Broadcast message submit-sell",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTradeAsset := args[0]
			argPrice := args[1]
			argCurrency := args[2]
			argAmount := args[3]
			argSellerShippingAddr := args[4]
			argSellerNknAddr := args[5]
			argRefundAddr := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitSell(
				clientCtx.GetFromAddress().String(),
				argTradeAsset,
				argPrice,
				argCurrency,
				argAmount,
				argSellerShippingAddr,
				argSellerNknAddr,
				argRefundAddr,
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
