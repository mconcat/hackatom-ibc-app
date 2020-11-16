package sender

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mconcat/hackatom-ibc-app/x/sender/keeper"
	"github.com/mconcat/hackatom-ibc-app/x/sender/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgRegisterEntry:
			return handleMsgRegisterEntry(ctx, k, msg)
    case *types.MsgSyncEntry:
      return handleMsgSyncEntry(ctx, k, msg)
    default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgRegisterEntry(ctx sdk.Context, k keeper.Keeper, msg *types.MsgRegisterEntry) (*sdk.Result, error) {
	if err := k.RegisterEntry(ctx, types.Entry{
    msg.EntryId,
    msg.ChannelId,
    msg.PortId,
    msg.ValidatorSetProvider,
  }); err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}

func handleMsgSyncEntry(ctx sdk.Context, k keeper.Keeper, msg *types.MsgSyncEntry) (*sdk.Result, error) {
  if err := k.SyncEntry(ctx, msg.EntryId); err != nil {
    return nil, err
  }
  return &sdk.Result{}, nil
}
