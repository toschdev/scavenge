package keeper

import (
	"context"

	"crypto/sha256"
	"encoding/hex"
	"scavenge/x/scavenge/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CommitAnswer(goCtx context.Context, msg *types.MsgCommitAnswer) (*types.MsgCommitAnswerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the question
	question, exists := k.GetScavengeQuestion(ctx, msg.QuestionId)
	if !exists {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "question not found")
	}

	solutionScavengerBytes := []byte(question.EncryptedAnswer + msg.Creator)
	solutionScavengerHash := sha256.Sum256(solutionScavengerBytes)
	commitHash := hex.EncodeToString(solutionScavengerHash[:])

	// Hash of solution and scavenger combined - this prevents front-running
	commit := types.CommittedAnswer{
		QuestionId: msg.QuestionId,
		HashAnswer: commitHash, // This should be hash(solution + creator)
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
