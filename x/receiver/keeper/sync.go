package keeper

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	cryptoencoding "github.com/tendermint/tendermint/crypto/encoding"

	"github.com/mconcat/hackatom-ibc-app/x/receiver/types"
)

func (k Keeper) GetValidatorUpdate(ctx sdk.Context, address []byte) (res abci.ValidatorUpdate, ok bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NewValidatorUpdateKey(address))
	if bz == nil {
		return
	}
	return k.MustUnmarshalValidatorUpdate(bz), true
}

func (k Keeper) HasValidatorUpdate(ctx sdk.Context, address []byte) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NewValidatorUpdateKey(address))
}

func (k Keeper) SetValidatorUpdate(ctx sdk.Context, address []byte, val abci.ValidatorUpdate) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NewValidatorUpdateKey(address), k.MustMarshalValidatorUpdate(val))
}

func (k Keeper) DeleteValidatorUpdate(ctx sdk.Context, address []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NewValidatorUpdateKey(address))
}

func (k Keeper) GetValidatorUpdates(ctx sdk.Context) (res []abci.ValidatorUpdate) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ValidatorUpdatesKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		res = append(res, k.MustUnmarshalValidatorUpdate(iter.Value()))
	}
	return
}

func (k Keeper) UpdateValidators(ctx sdk.Context, updates types.ValidatorUpdates) {
	sort.Sort(updates)
	for _, update := range updates.ValidatorUpdates {
		pubkey, err := cryptoencoding.PubKeyFromProto(update.PubKey)
		if err != nil {
			panic(err) // TODO move to ValidateBasic
		}
		address := pubkey.Address()
		if update.Power == 0 {
			k.DeleteValidatorUpdate(ctx, address)
		} else {
			k.SetValidatorUpdate(ctx, address, *update)
		}
	}
}

func (k Keeper) CheckValidatorUpdateCommitment(ctx sdk.Context, commit []byte) error {
	// TODO:
	return nil
}
