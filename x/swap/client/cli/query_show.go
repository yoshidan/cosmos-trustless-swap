package cli

import (
	"strconv"

	"github.com/spf13/cast"

	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [sender] [id]",
		Short: "Query show",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqSender := args[0]
			reqId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryShowRequest{
				Sender: reqSender,
				Id:     reqId,
			}

			res, err := queryClient.Show(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
