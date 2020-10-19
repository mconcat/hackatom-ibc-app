package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

const (
	TypeMsgRegisterSync = "register_sync"
)

func NewMsgRegisterSync(
	channelID string,
	sender sdk.AccAddress,
) *MsgRegisterSync {
	return &MsgRegisterSync{
		ChannelId: channelID,
		Sender:    sender.String(),
	}
}

func (MsgRegisterSync) Route() string {
	return RouterKey
}

func (MsgRegisterSync) Type() string {
	return TypeMsgRegisterSync
}

func (msg MsgRegisterSync) ValidateBasic() error {
	if err := host.ChannelIdentifierValidator(msg.ChannelId); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}
	return nil
}

func (msg MsgRegisterSync) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRegisterSync) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddr}
}
