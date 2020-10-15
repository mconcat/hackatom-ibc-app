package keeper

import (
	"github.com/mconcat/hackatom-ibc-app/x/receiver/types"
)

var _ types.QueryServer = Keeper{}
