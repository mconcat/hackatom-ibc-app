package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"
)

func (k Keeper) MustMarshalValidatorUpdate(val abci.ValidatorUpdate) []byte {
	return k.cdc.MustMarshalBinaryBare(&val)
}

func (k Keeper) MustUnmarshalValidatorUpdate(bz []byte) abci.ValidatorUpdate {
	var val abci.ValidatorUpdate
	k.cdc.MustUnmarshalBinaryBare(bz, &val)
	return val
}
