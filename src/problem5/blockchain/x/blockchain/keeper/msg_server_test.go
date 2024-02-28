package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/alvinjiang1/blockchain/testutil/keeper"
	"github.com/alvinjiang1/blockchain/x/blockchain/keeper"
	"github.com/alvinjiang1/blockchain/x/blockchain/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.BlockchainKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
