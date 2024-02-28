package keeper

import (
	"new_chain/x/newchain/types"
)

var _ types.QueryServer = Keeper{}
