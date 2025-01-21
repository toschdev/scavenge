package keeper

import (
	"context"

	"scavenge/x/scavenge/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RevealAnswer(goCtx context.Context, msg *types.MsgRevealAnswer) (*types.MsgRevealAnswerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRevealAnswerResponse{}, nil
}
