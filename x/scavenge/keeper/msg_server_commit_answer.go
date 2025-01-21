package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CommitAnswer(goCtx context.Context, msg *types.MsgCommitAnswer) (*types.MsgCommitAnswerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Hash of solution and scavenger combined - this prevents front-running
	commit := types.CommittedAnswer{
		QuestionId: msg.QuestionId,
		HashAnswer: msg.HashAnswer, // This should be hash(solution + creator)
		Creator:    msg.Creator,
	}

	// Check if a commit already exists
	_, found := k.GetCommittedAnswer(ctx, msg.QuestionId, msg.Creator)
	if found {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "commit already exists for this question and creator")
	}

	k.SetCommittedAnswer(ctx, commit)

	return &types.MsgCommitAnswerResponse{}, nil
}
