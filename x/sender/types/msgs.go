package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

const (
	TypeMsgRegisterEntry = "register_sync"
)

func NewMsgRegisterEntry(
	channelID string,
	sender sdk.AccAddress,
) *MsgRegisterEntry {
	return &MsgRegisterEntry{
		ChannelId: channelID,
		Sender:    sender.String(),
	}
}

func (MsgRegisterEntry) Route() string {
	return RouterKey
}

func (MsgRegisterEntry) Type() string {
	return TypeMsgRegisterEntry
}

func (msg MsgRegisterEntry) ValidateBasic() error {
	if err := host.ChannelIdentifierValidator(msg.ChannelId); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}
	return nil
}

func (msg MsgRegisterEntry) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgRegisterEntry) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddr}
}
