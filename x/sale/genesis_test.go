package sale_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "swap/testutil/keeper"
	"swap/testutil/nullify"
	"swap/x/sale"
	"swap/x/sale/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SaleKeeper(t)
	sale.InitGenesis(ctx, *k, genesisState)
	got := sale.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
