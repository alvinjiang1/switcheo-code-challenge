package newchain_test

import (
	"testing"

	keepertest "new_chain/testutil/keeper"
	"new_chain/testutil/nullify"
	newchain "new_chain/x/newchain/module"
	"new_chain/x/newchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:	types.DefaultParams(),
		
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NewchainKeeper(t)
	newchain.InitGenesis(ctx, k, genesisState)
	got := newchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	

	// this line is used by starport scaffolding # genesis/test/assert
}
