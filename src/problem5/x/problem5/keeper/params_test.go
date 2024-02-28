package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "problem5/testutil/keeper"
	"problem5/x/problem5/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.Problem5Keeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
