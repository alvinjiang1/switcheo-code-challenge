package problem5_test

import (
	"testing"

	keepertest "problem5/testutil/keeper"
	"problem5/testutil/nullify"
	problem5 "problem5/x/problem5/module"
	"problem5/x/problem5/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Problem5Keeper(t)
	problem5.InitGenesis(ctx, k, genesisState)
	got := problem5.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
