package keeper_test

import (
	"context"
	"testing"

	keepertest "swap/testutil/keeper"
	"swap/x/sale/keeper"
	"swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SaleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
