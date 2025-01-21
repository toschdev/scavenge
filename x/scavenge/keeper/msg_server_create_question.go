package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateQuestion(goCtx context.Context, msg *types.MsgCreateQuestion) (*types.MsgCreateQuestionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create a new scavenge question
	question := types.ScavengeQuestion{
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

	return &types.MsgCreateQuestionResponse{}, nil
}
