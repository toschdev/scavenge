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

	// Check if the question exists and is not completed
	question, found := k.GetScavengeQuestion(ctx, msg.QuestionId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "question not found")
	}
	if question.Completed {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "question already completed")
	}

	// Create and store the committed answer
	commit := types.CommittedAnswer{
		QuestionId: msg.QuestionId,
		HashAnswer: msg.HashAnswer,
		Creator:    msg.Creator,
	}

	// Check if a commit already exists for this creator and question
	_, found = k.GetCommittedAnswer(ctx, msg.QuestionId, msg.Creator)
	if found {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "commit already exists for this question")
	}

	k.SetCommittedAnswer(ctx, commit)

	return &types.MsgCommitAnswerResponse{}, nil
}
