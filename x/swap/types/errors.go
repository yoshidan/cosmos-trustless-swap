package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/swap module sentinel errors
var (
	ErrSwapNotFound           = sdkerrors.Register(ModuleName, 1100, "swap not found")
	ErrInvalidSwapStatus      = sdkerrors.Register(ModuleName, 1101, "swap status not active")
	ErrInsufficientPermission = sdkerrors.Register(ModuleName, 1102, "insufficient permission")
)
