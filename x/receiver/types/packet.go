package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type ValidatorUpdatesPacketData struct {
	ValidatorsHash   []byte
	ValidatorUpdates []abci.ValidatorUpdate
}

// NewValidatorUpdatesPacketData constructs a new ValsetUpdatePacketData instance
func NewValidatorUpdatesPacketData(validatorsHash []byte, validatorUpdates []abci.ValidatorUpdate) ValidatorUpdatesPacketData {
	return ValidatorUpdatesPacketData{
		ValidatorsHash:   validatorsHash,
		ValidatorUpdates: validatorUpdates,
	}
}

func (vpd ValidatorUpdatesPacketData) ValidateBasic() error {
	if len(vpd.ValidatorsHash) == 0 && len(vpd.ValidatorUpdates) == 0 {
		return sdkerrors.Wrap(ErrInvalidValsetPacket, "both commitment hash or validator updates cannot be blank")
	}

	return nil
}

func (vpd ValidatorUpdatesPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&vpd))
}
