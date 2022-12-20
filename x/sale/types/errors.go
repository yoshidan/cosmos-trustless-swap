package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/sale module sentinel errors
var (
	ErrSaleNotFound           = sdkerrors.Register(ModuleName, 1200, "sale not found")
	ErrInsufficientPermission = sdkerrors.Register(ModuleName, 1201, "insufficient permission")
)
