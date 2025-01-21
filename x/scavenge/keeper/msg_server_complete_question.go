package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CompleteQuestion(goCtx context.Context, msg *types.MsgCompleteQuestion) (*types.MsgCompleteQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the question
	question, found := k.GetScavengeQuestion(ctx, msg.QuestionId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "question not found")
	}

	// Check if the creator is completing their own question
	if question.Creator != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "only the creator can complete the question")
	}

	// Check if already completed
	if question.Completed {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "question already completed")
	}

	// Update the question
	question.Completed = true
	question.Winner = msg.Winner
	k.SetScavengeQuestion(ctx, question)

	return &types.MsgCompleteQuestionResponse{}, nil
}
