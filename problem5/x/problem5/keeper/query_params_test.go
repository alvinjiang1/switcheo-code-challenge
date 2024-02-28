package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "problem5/testutil/keeper"
	"problem5/x/problem5/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.Problem5Keeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
