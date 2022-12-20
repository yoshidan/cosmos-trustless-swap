package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"swap/x/sale/types"
)

var _ = strconv.Itoa(0)

func CmdSellNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sell-nft [class-id] [nft-id] [price]",
		Short: "Broadcast message SellNFT",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argClassId := args[0]
			argNftId := args[1]
			argPrice := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSellNFT(
				clientCtx.GetFromAddress().String(),
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
