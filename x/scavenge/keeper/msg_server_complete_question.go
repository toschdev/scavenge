package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CompleteQuestion(goCtx context.Context, msg *types.MsgCompleteQuestion) (*types.MsgCompleteQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCompleteQuestionResponse{}, nil
}
