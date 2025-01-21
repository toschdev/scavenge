package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListQuestions(goCtx context.Context, req *types.QueryListQuestionsRequest) (*types.QueryListQuestionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sdkCtx := sdk.UnwrapSDKContext(goCtx)
	questions := k.GetAllScavengeQuestion(sdkCtx)

	return &types.QueryListQuestionsResponse{
		ScavengeQuestion: questions,
	}, nil
}
