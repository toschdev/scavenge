package keeper

import (
	"scavenge/x/scavenge/types"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCommittedAnswer store a specific committedAnswer
func (k Keeper) SetCommittedAnswer(ctx sdk.Context, committedAnswer types.CommittedAnswer) {
	store := prefix.NewStore(k.getStore(ctx), types.KeyPrefix(types.CommitKeyPrefix))
	key := types.GetCommittedAnswerKey(
		committedAnswer.QuestionId,
		committedAnswer.Creator,
	)
	b := k.cdc.MustMarshal(&committedAnswer)
	store.Set(key, b)
}

// GetCommittedAnswer returns a specific committedAnswer
func (k Keeper) GetCommittedAnswer(ctx sdk.Context, questionId uint64, creator string) (val types.CommittedAnswer, found bool) {
	store := prefix.NewStore(k.getStore(ctx), types.KeyPrefix(types.CommitKeyPrefix))
	key := types.GetCommittedAnswerKey(questionId, creator)
	b := store.Get(key)
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
