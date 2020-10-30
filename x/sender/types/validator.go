package types

import (
  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/mconcat/hackatom-ibc-app/x/sender/types"
)

func (k Keeper) IterateSyncEntry(
  ctx sdk.Context, fn func(types.SyncEntry) error 
) error {
  store := ctx.KVStore(k.storeKey)
  iter := sdk.KVStorePrefixIterator(store, types.SyncsKey)
  defer iter.Close()
  for ; iter.Valid(); iter.Next() {
    value := k.MustUnmarshalSync(

    )
  }
}
