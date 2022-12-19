package swap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "swap/testutil/keeper"
	"swap/testutil/nullify"
	"swap/x/swap"
	"swap/x/swap/types"
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
