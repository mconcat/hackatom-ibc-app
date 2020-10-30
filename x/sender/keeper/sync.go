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

func (k Keeper) IterateSyncEntry(
  ctx sdk.Context,
  fn func(types.SyncEntry) error,
) error {
  store := ctx.KVStore(k.storeKey)
  iter := sdk.KVStorePrefixIterator(store, types.SyncsKey)
  defer iter.Close()
  for ; iter.Valid(); iter.Next() {
    entry := k.MustUnmarshalSyncEntry(iter.Value())
    err := fn(entry)
    if err != nil {
      return err
    }
  }
  return nil
}

func (k Keeper) GetProviderFromSyncEntry(ctx sdk.Context, sync types.SyncEntry) (types.ValidatorSetProvider, error) {
  // TODO: retrieve provider from syncentry using dynamic capability

  return k.defaultValidatorSetProvider
}

func (k Keeper) CheckUpdateNeeded(ctx sdk.Context, sync types.SyncEntry) bool {
  // TODO: check voting power difference threshold
  // TODO: check block height difference threshold

  return true
}
