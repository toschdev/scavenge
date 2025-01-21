package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowCommit(goCtx context.Context, req *types.QueryShowCommitRequest) (*types.QueryShowCommitResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	commit, found := k.GetCommittedAnswer(ctx, req.QuestionId, req.Creator)
	if !found {
		return nil, status.Error(codes.NotFound, "commit not found")
	}

	return &types.QueryShowCommitResponse{
		CommittedAnswer: commit,
	}, nil
}
