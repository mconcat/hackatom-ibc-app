package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mconcat/hackatom-ibc-app/x/sender/types"
)

func (k Keeper) GetEntrys(ctx sdk.Context) (syncs []types.Entry) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.EntrysKey)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		sync := k.MustUnmarshalEntry(iter.Value())
		syncs = append(syncs, sync)
	}

	return
}

func (k Keeper) SetEntry(ctx sdk.Context, sync types.Entry) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has(types.NewEntryKey(sync.ChannelId)) {
		return types.ErrEntryAlreadyExists
	}
	store.Set(types.NewEntryKey(sync.ChannelId), k.MustMarshalEntry(sync))
	return nil
}

func (k Keeper) IterateEntryEntry(
  ctx sdk.Context,
  fn func(types.EntryEntry) error,
) error {
  store := ctx.KVStore(k.storeKey)
  iter := sdk.KVStorePrefixIterator(store, types.EntrysKey)
  defer iter.Close()
  for ; iter.Valid(); iter.Next() {
    entry := k.MustUnmarshalEntryEntry(iter.Value())
    err := fn(entry)
    if err != nil {
      return err
    }
  }
  return nil
}

func (k Keeper) UpdateHeartbeat(ctx sdk.Context, entry types.EntryEntry) error {

}
