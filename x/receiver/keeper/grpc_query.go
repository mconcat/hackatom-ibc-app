package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mconcat/hackatom-ibc-app/x/receiver/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (q Keeper) Validator(c context.Context, req *types.QueryValidatorRequest) (*types.QueryValidatorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	address, err := sdk.AccAddressFromBech32(req.ValidatorAddr)
	if err != nil {
		return nil, err
	}

	val, found := q.GetValidatorUpdate(ctx, address)
	if !found {
		return nil, status.Error(
			codes.NotFound,
			sdkerrors.Wrap(types.ErrValidatorNotFound, req.ValidatorAddr).Error(),
		)
	}

	return &types.QueryValidatorResponse{
		Validator: types.ToABCIValidator(val),
	}, nil
}

func (q Keeper) Validators(c context.Context, req *types.QueryValidatorsRequest) (*types.QueryValidatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	updates := q.GetValidatorUpdates(ctx)

	var vals []abci.Validator
	for _, update := range updates {
		vals = append(vals, types.ToABCIValidator(update))
	}

	return &types.QueryValidatorsResponse{
		Validators: vals,
	}, nil
}
