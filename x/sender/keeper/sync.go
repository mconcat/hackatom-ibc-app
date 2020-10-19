package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mconcat/hackatom-ibc-app/x/sender/types"
)

func (k Keeper) GetSyncs(ctx sdk.Context) (syncs []types.Sync) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.SyncsKey)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		sync := k.MustUnmarshalSync(iter.Value())
		syncs = append(syncs, sync)
	}

	return
}

func (k Keeper) RegisterSync(ctx sdk.Context, sync types.Sync) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has(types.NewSyncKey(sync.ChannelId)) {
		return types.ErrSyncAlreadyExists
	}
	store.Set(types.NewSyncKey(sync.ChannelId), k.MustMarshalSync(sync))
	return nil
}
