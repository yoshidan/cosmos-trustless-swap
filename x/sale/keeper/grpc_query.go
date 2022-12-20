package keeper

import (
	"github.com/yoshidan/cosmos-trustless-swap/x/sale/types"
)

var _ types.QueryServer = Keeper{}
