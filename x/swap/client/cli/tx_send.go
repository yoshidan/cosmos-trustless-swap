package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"swap/x/swap/types"
)

var _ = strconv.Itoa(0)

func CmdSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [receiver] [amount] [amount-to-receive]",
		Short: "Broadcast message send",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argAmount := args[1]
			argAmountToReceive := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSend(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argAmount,
				argAmountToReceive,
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
