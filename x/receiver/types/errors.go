package types

// DONTCOVER
import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/receiver module sentinel errors
var (
	ErrValidatorNotFound = sdkerrors.Register(ModuleName, 1100, "validator not found")
)
