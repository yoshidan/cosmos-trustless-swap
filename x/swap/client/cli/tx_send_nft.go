package cli

import (
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
		Use:   "send-nft [receiver] [class-id] [nft-id] [amount-to-receive]",
		Short: "Broadcast message SendNFT",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argClassId := args[1]
			argNftId := args[2]
			argAmountToReceive := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendNFT(
				clientCtx.GetFromAddress().String(),
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
