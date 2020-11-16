package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewPacketDataValidatorSet(commit []byte, updates []*abci.ValidatorUpdate) PacketDataValidatorSet {
	return PacketDataValidatorSet{
		ValidatorsCommit: commit,
		ValidatorUpdates: updates,
	}
}

func (vupd PacketDataValidatorSet) ValidateBasic() error {
	if len(vupd.ValidatorsCommit) == 0 && len(vupd.ValidatorUpdates) == 0 {
		return sdkerrors.Wrap(ErrInvalidUpdates, "both commitment or updates cannot be blank")
	}

	return nil
}

func (vupd PacketDataValidatorSet) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&vupd))
}
