package sender

import (
  abci "github.com/tendermint/tendermint/abci/types"

  sdk "github.com/cosmos/cosmos-sdk/types"

  "github.com/mconcat/hackatom-ibc-app/x/sender/keeper"
  "github.com/mconcat/hackatom-ibc-app/x/sender/types"
)

// TODO: check how endblocker handles errors, unify to panic or return
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.Validator {
  k.IterateSyncEntry(func(ctx sdk.Context, sync types.SyncEntry) error {
    provider, err := k.GetProviderFromSyncEntry(ctx, entry)
   
		if err != nil {
			panic(err) // entries should be valid if they were successfully stored, so this should not return an error
		}

		// Checking whether to send a packet or not inside EndBlocker is fragile against sybil attack
		// Try to make it invoked by transaction
		if k.HeartbeatHeightPassed(ctx, entry) || k.PowerDiffThresholdPassed(ctx, entry, provider) {
			if err = k.UpdateHeartbeat(ctx, entry); err != nil {
				return err
			}
			if err = k.UpdateValidatorset(ctx, entry, provider); err != nil {
				return err
			}
		

			updates := k.GetValidatorSetUpdates(ctx, sync, provider)
			err = k.SendValidatorUpdates(
				ctx,
				sync.PortID,
				sync.ChannelID,
				sync.TimeoutHeight,
				sync.TimeoutTimestamp,
				sync.GetCommitmentBytes(),
				updates,
			)
			if err != nil {
				panic(err) // should we just return an error?
			}
		}
  })
}
