package keeper

import (
	"swap/x/swap/types"
)

var _ types.QueryServer = Keeper{}
