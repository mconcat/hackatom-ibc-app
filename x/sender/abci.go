package sender

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/mconcat/hackatom-ibc-app/x/sender/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) (_ []abci.ValidatorUpdate) {
	vals := k.ReturnValidatorSetUpdates(ctx)

}
