package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	"crypto/sha256"
	"encoding/hex"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RevealAnswer(goCtx context.Context, msg *types.MsgRevealAnswer) (*types.MsgRevealAnswerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the question exists and is not completed
	question, found := k.GetScavengeQuestion(ctx, msg.QuestionId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "question not found")
	}
	if question.Completed {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "question already completed")
	}

	// Check if a commit exists
	commit, found := k.GetCommittedAnswer(ctx, msg.QuestionId, msg.Creator)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "commit not found: must commit before reveal")
	}

	// Verify the hash of the answer matches the commit
	answer := msg.PlainText + msg.Creator
	answerHash := sha256.Sum256([]byte(answer))
	hashedAnswer := hex.EncodeToString(answerHash[:])
	if hashedAnswer != commit.HashAnswer {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "submitted answer hash does not match commit")
	}

	// Verify the answer is correct
	solutionHash := sha256.Sum256([]byte(msg.PlainText))
	solutionHashString := hex.EncodeToString(solutionHash[:])
	if solutionHashString != question.EncryptedAnswer {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "incorrect answer")
	}

	// Send the bounty to the winner
	winner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid winner address")
	}

	coins := sdk.NewCoins(sdk.NewCoin("stake", sdkmath.NewInt(int64(question.Bounty))))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winner, coins); err != nil {
		return nil, errorsmod.Wrap(err, "failed to send bounty to winner")
	}

	// Update the question as completed
	question.Completed = true
	question.Winner = msg.Creator
	k.SetScavengeQuestion(ctx, question)

	return &types.MsgRevealAnswerResponse{}, nil
}
