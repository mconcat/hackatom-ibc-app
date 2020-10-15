package keeper

import (
	"github.com/mconcat/hackatom-ibc-app/x/sender/types"
)

var _ types.QueryServer = Keeper{}
