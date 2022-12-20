package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/yoshidan/cosmos-trustless-swap/testutil/keeper"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/keeper"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SaleKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
