package dictland_test

import (
	"testing"

	keepertest "dictland/testutil/keeper"
	"dictland/testutil/nullify"
	"dictland/x/dictland"
	"dictland/x/dictland/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DictlandKeeper(t)
	dictland.InitGenesis(ctx, *k, genesisState)
	got := dictland.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
