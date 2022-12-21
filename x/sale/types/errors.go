package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
	"google.golang.org/grpc/codes"
)

// x/sale module sentinel errors
var (
	ErrSaleNotFound           = sdkerrors.RegisterWithGRPCCode(ModuleName, 1200, codes.NotFound, "sale not found")
	ErrInsufficientPermission = sdkerrors.RegisterWithGRPCCode(ModuleName, 1201, codes.PermissionDenied, "insufficient permission")
	ErrInvalidSaleData        = sdkerrors.RegisterWithGRPCCode(ModuleName, 1202, codes.DataLoss, "invalid sale data.")
	ErrSaleExists             = sdkerrors.RegisterWithGRPCCode(ModuleName, 1203, codes.AlreadyExists, "sale already exists")
)
