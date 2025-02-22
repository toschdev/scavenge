package keeper

import (
	"context"
	"strconv"

	"scavenge/x/scavenge/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateQuestion(goCtx context.Context, msg *types.MsgCreateQuestion) (*types.MsgCreateQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the next question ID
	count := k.GetScavengeQuestionCount(ctx)

	// Create a new scavenge question
	question := types.ScavengeQuestion{
		Id:              count,
		Creator:         msg.Creator,
		Question:        msg.Question,
		EncryptedAnswer: msg.EncryptedAnswer,
		Bounty:          msg.Bounty,
		Completed:       false,
		Winner:          "",
	}

	// Lock the bounty in the module account
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator address")
	}

	coins := sdk.NewCoins(sdk.NewCoin(types.ChainDenom, sdkmath.NewInt(int64(msg.Bounty))))
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, coins); err != nil {
		return nil, errorsmod.Wrap(err, "failed to lock bounty")
	}

	k.SetScavengeQuestion(ctx, question)
	k.SetScavengeQuestionCount(ctx, count+1)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, "create_question"),
			sdk.NewAttribute(types.AttributeKeyQuestionId, strconv.FormatUint(count, 10)),
			sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyBounty, strconv.FormatUint(msg.Bounty, 10)),
		),
	)

	return &types.MsgCreateQuestionResponse{
		Id: count,
	}, nil
}
