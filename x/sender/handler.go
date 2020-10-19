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
		case *types.MsgRegisterSync:
			return handleMsgRegisterSync(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgRegisterSync(ctx sdk.Context, k keeper.Keeper, msg *types.MsgRegisterSync) (*sdk.Result, error) {
	if err := k.RegisterSync(ctx, types.Sync{msg.ChannelId}); err != nil {
		return nil, err
	}
	return &sdk.Result{}, nil
}
