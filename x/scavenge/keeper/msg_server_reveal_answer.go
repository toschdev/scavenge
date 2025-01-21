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

	// First verify the commit exists by recreating the hash of solution + creator
	solutionScavengerBytes := []byte(msg.PlainText + msg.Creator)
	solutionScavengerHash := sha256.Sum256(solutionScavengerBytes)
	commitHash := hex.EncodeToString(solutionScavengerHash[:])

	// Get the commit
	commit, found := k.GetCommittedAnswer(ctx, msg.QuestionId, msg.Creator)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "no commit found: must commit before reveal")
	}

	// Verify the hash matches their commit
	if commitHash != commit.HashAnswer {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "committed hash does not match")
	}

	// Now hash just the solution to find the matching question
	solutionHash := sha256.Sum256([]byte(msg.PlainText))
	solutionHashString := hex.EncodeToString(solutionHash[:])

	// Get the question
	question, found := k.GetScavengeQuestion(ctx, msg.QuestionId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "question not found")
	}

	// Verify the solution hash matches the question's encrypted answer
	if solutionHashString != question.EncryptedAnswer {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "incorrect answer")
	}

	if question.Completed {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "question already completed")
	}

	// Award the bounty
	winner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid winner address")
	}

	coins := sdk.NewCoins(sdk.NewCoin("token", sdkmath.NewInt(int64(question.Bounty))))
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winner, coins); err != nil {
		return nil, errorsmod.Wrap(err, "failed to send bounty")
	}

	// Update the question
	question.Completed = true
	question.Winner = msg.Creator
	k.SetScavengeQuestion(ctx, question)

	return &types.MsgRevealAnswerResponse{}, nil
}
