package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"

	"github.com/mconcat/hackatom-ibc-app/x/receiver/types"
	sendertypes "github.com/mconcat/hackatom-ibc-app/x/sender/types"
)

func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet, data sendertypes.ValidatorUpdatesPacketData) error {
	if err := data.ValidateBasic(); err != nil {
		return err
	}

	k.UpdateValidators(ctx, types.NewValidatorUpdates(data.ValidatorUpdates))

	if err := k.CheckValidatorUpdateCommitment(ctx, data.ValidatorsCommit); err != nil {
		return err
	}

	return nil
}
