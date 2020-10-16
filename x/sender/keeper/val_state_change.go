package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/exported"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/encoding"
)

// Implemented as sending all last validator information
// It works as intended, but for sake of efficiency, make it send only the updated validators
func (k Keeper) ReturnValidatorSetUpdates(ctx sdk.Context) (updates []abci.ValidatorUpdate) {
	k.valset.IterateLastValidators(ctx, func(_ int64, validator exported.ValidatorI) (stop bool) {
		pk, err := encoding.PubKeyToProto(validator.GetConsPubKey())
		if err != nil {
			panic(err)
		}

		updates = append(updates, abci.ValidatorUpdate{
			PubKey: pk,
			Power:  validator.GetConsensusPower(),
		})

		return
	})

	return
}
