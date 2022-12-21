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

func CmdSellNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-nft [id] [class-id] [nft-id] [price]",
		Short: "Broadcast message SellNFT",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argClassId := args[1]
			argNftId := args[2]
			argPrice := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSellNFT(
				clientCtx.GetFromAddress().String(),
				argId,
				argClassId,
				argNftId,
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
