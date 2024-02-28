package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/alvinjiang1/blockchain/testutil/keeper"
	"github.com/alvinjiang1/blockchain/x/blockchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.BlockchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
