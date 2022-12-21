package cli

import (
	"github.com/spf13/cast"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"
)

var _ = strconv.Itoa(0)

func CmdSell() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell [id] [amount] [price]",
		Short: "Broadcast message sell",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argAmount := args[1]
			argPrice := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSell(
				clientCtx.GetFromAddress().String(),
				argId,
				argAmount,
				argPrice,
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
