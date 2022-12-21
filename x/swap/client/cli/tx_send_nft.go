package cli

import (
	"github.com/spf13/cast"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"
)

var _ = strconv.Itoa(0)

func CmdSendNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send-nft [id] [receiver] [class-id] [nft-id] [amount-to-receive]",
		Short: "Broadcast message SendNFT",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argReceiver := args[1]
			argClassId := args[2]
			argNftId := args[3]
			argAmountToReceive := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendNFT(
				clientCtx.GetFromAddress().String(),
				argId,
				argReceiver,
				argClassId,
				argNftId,
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
