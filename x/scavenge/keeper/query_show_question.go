package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowQuestion(goCtx context.Context, req *types.QueryShowQuestionRequest) (*types.QueryShowQuestionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	question, found := k.GetScavengeQuestion(ctx, req.QuestionId)
	if !found {
		return nil, status.Error(codes.NotFound, "question not found")
	}

	return &types.QueryShowQuestionResponse{
		ScavengeQuestion: question,
	}, nil
}
