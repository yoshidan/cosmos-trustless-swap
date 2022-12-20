package swap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/yoshidan/cosmos-trustless-swap/testutil/keeper"
	"github.com/yoshidan/cosmos-trustless-swap/testutil/nullify"
	"github.com/yoshidan/cosmos-trustless-swap/x/swap"
	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SwapKeeper(t)
	swap.InitGenesis(ctx, *k, genesisState)
	got := swap.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
