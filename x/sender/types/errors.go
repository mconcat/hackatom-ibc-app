package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/sender module sentinel errors
var (
	ErrInvalidUpdates    = sdkerrors.Register(ModuleName, 1100, "invalid validator updates")
	ErrSyncAlreadyExists = sdkerrors.Register(ModuleName, 1101, "sync already exists for the channel id")
)
