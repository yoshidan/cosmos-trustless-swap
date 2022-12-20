package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/swap module sentinel errors
var (
	ErrSwapNotFound           = sdkerrors.Register(ModuleName, 1100, "swap not found")
	ErrInsufficientPermission = sdkerrors.Register(ModuleName, 1101, "insufficient permission")
)
