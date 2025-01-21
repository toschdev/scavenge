package keeper

import (
	"encoding/binary"
	"scavenge/x/scavenge/types"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetScavengeQuestion store a specific scavengeQuestion in the store
func (k Keeper) SetScavengeQuestion(ctx sdk.Context, scavengeQuestion types.ScavengeQuestion) uint64 {
	store := prefix.NewStore(k.getStore(ctx), types.KeyPrefix(types.ScavengeKey))
	appendedValue := k.cdc.MustMarshal(&scavengeQuestion)

	// Get the current count
	count := k.GetScavengeQuestionCount(ctx)

	// Store using the count as the ID
	store.Set(GetScavengeQuestionIDBytes(count), appendedValue)

	// Update the count
	k.SetScavengeQuestionCount(ctx, count+1)

	return count
}

// GetScavengeQuestion returns a scavengeQuestion from its id
func (k Keeper) GetScavengeQuestion(ctx sdk.Context, id uint64) (val types.ScavengeQuestion, found bool) {
	store := prefix.NewStore(k.getStore(ctx), types.KeyPrefix(types.ScavengeKey))
	b := store.Get(GetScavengeQuestionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetScavengeQuestionCount get the total number of scavengeQuestion
func (k Keeper) GetScavengeQuestionCount(ctx sdk.Context) uint64 {
	store := k.getStore(ctx)
	byteKey := types.KeyPrefix(types.ScavengeCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// SetScavengeQuestionCount set the total number of scavengeQuestion
func (k Keeper) SetScavengeQuestionCount(ctx sdk.Context, count uint64) {
	store := k.getStore(ctx)
	byteKey := types.KeyPrefix(types.ScavengeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// Helper function to convert ID to bytes
func GetScavengeQuestionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAllScavengeQuestion returns all scavenge questions
func (k Keeper) GetAllScavengeQuestion(ctx sdk.Context) (list []types.ScavengeQuestion) {
	store := prefix.NewStore(k.getStore(ctx), types.KeyPrefix(types.ScavengeKey))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ScavengeQuestion
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}

func (k Keeper) GetAllCommittedAnswer(ctx sdk.Context) (list []types.CommittedAnswer) {
	store := prefix.NewStore(k.getStore(ctx), types.KeyPrefix(types.CommitKeyPrefix))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CommittedAnswer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
