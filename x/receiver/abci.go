package receiver

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mconcat/hackatom-ibc-app/x/receiver/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	return k.GetValidatorUpdates(ctx)
}
