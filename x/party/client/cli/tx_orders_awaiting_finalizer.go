package cli

import (
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdCreateOrdersAwaitingFinalizer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-orders-awaiting-finalizer [index] [nkn-address] [wallet-private-key] [wallet-public-key] [shipping-address] [refund-address] [amount]",
		Short: "Create a new orders-awaiting-finalizer",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argNknAddress := args[1]
			argWalletPrivateKey := args[2]
			argWalletPublicKey := args[3]
			argShippingAddress := args[4]
			argRefundAddress := args[5]
			argAmount := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOrdersAwaitingFinalizer(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argNknAddress,
				argWalletPrivateKey,
				argWalletPublicKey,
				argShippingAddress,
				argRefundAddress,
				argAmount,
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

func CmdUpdateOrdersAwaitingFinalizer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-orders-awaiting-finalizer [index] [nkn-address] [wallet-private-key] [wallet-public-key] [shipping-address] [refund-address] [amount]",
		Short: "Update a orders-awaiting-finalizer",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexIndex := args[0]

			// Get value arguments
			argNknAddress := args[1]
			argWalletPrivateKey := args[2]
			argWalletPublicKey := args[3]
			argShippingAddress := args[4]
			argRefundAddress := args[5]
			argAmount := args[6]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateOrdersAwaitingFinalizer(
				clientCtx.GetFromAddress().String(),
				indexIndex,
				argNknAddress,
				argWalletPrivateKey,
				argWalletPublicKey,
				argShippingAddress,
				argRefundAddress,
				argAmount,
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

func CmdDeleteOrdersAwaitingFinalizer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-orders-awaiting-finalizer [index]",
		Short: "Delete a orders-awaiting-finalizer",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexIndex := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteOrdersAwaitingFinalizer(
				clientCtx.GetFromAddress().String(),
				indexIndex,
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
