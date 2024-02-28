package keeper

import (
	"github.com/alvinjiang1/blockchain/x/blockchain/types"
)

var _ types.QueryServer = Keeper{}
