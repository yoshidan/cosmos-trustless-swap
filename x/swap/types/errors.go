package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

// x/swap module sentinel errors
var (
	ErrSwapNotFound           = sdkerrors.RegisterWithGRPCCode(ModuleName, 1100, codes.NotFound, "swap not found")
	ErrInsufficientPermission = sdkerrors.RegisterWithGRPCCode(ModuleName, 1101, codes.PermissionDenied, "insufficient permission")
	ErrInvalidSwapData        = sdkerrors.RegisterWithGRPCCode(ModuleName, 1102, codes.DataLoss, "invalid swap data. discard this swap")
	ErrSwapExists             = sdkerrors.RegisterWithGRPCCode(ModuleName, 1103, codes.AlreadyExists, "swap already exists")
)
