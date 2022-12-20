package keeper

import (
	"github.com/yoshidan/cosmos-trustless-swap/x/swap/types"
)

var _ types.QueryServer = Keeper{}
