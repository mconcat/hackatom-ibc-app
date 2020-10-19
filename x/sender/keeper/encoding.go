package keeper

import "github.com/mconcat/hackatom-ibc-app/x/sender/types"

func (k Keeper) MustMarshalSync(sync types.Sync) []byte {
	return k.cdc.MustMarshalBinaryBare(&sync)
}

func (k Keeper) MustUnmarshalSync(bz []byte) types.Sync {
	var sync types.Sync
	k.cdc.MustUnmarshalBinaryBare(bz, &sync)
	return sync
}
