package cli

import (
	"strconv"

	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdShowNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-nft [seller] [id]",
		Short: "Query ShowNFT",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqSeller := args[0]
			reqId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowNFTRequest{
				Seller: reqSeller,
				Id:     reqId,
			}

			res, err := queryClient.ShowNFT(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
