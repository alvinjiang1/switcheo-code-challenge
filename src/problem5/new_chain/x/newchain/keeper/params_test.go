package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

    keepertest "new_chain/testutil/keeper"
    "new_chain/x/newchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.NewchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
