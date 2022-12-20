package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/yoshidan/cosmos-trustless-swap/testutil/keeper"
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SaleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
