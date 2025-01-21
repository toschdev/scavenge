package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateQuestion(goCtx context.Context, msg *types.MsgCreateQuestion) (*types.MsgCreateQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateQuestionResponse{}, nil
}
