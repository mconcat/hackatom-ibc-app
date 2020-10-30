package sender

import (
  abci "github.com/tendermint/tendermint/abci/types"

  sdk "github.com/cosmos/cosmos-sdk/types"

  "github.com/mconcat/hackatom-ibc-app/x/sender/keeper"
  "github.com/mconcat/hackatom-ibc-app/x/sender/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.Validator {
  k.IterateSyncEntry(func(ctx sdk.Context, sync types.SyncEntry) error {
    provider, err := k.GetProviderFromSyncEntry(ctx, entry)
    
    if !k.CheckUpdateNeeded(ctx, sync, provider) {
      return nil
    }
    updates := k.GetValidatorSetUpdates(ctx, sync, provider)
    return k.SendValidatorUpdates(
      ctx,
      sync.PortID,
      sync.ChannelID,
      sync.TimeoutHeight,
      sync.TimeoutTimestamp,
      sync.GetCommitmentBytes(),
      updates,
    )
  })
}
