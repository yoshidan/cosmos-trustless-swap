package cli

import (
	"strconv"

	"github.com/spf13/cast"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdReceiveNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "receive-nft [sender] [id]",
		Short: "Broadcast message ReceiveNFT",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argSender := args[0]
			argId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgReceiveNFT(
				clientCtx.GetFromAddress().String(),
				argSender,
				argId,
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
