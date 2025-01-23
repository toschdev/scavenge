package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "scavenge"

	ChainDenom = "stake"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_scavenge"

	// Key prefixes for different store types
	ScavengeKey      = "Scavenge/value/"
	ScavengeCountKey = "Scavenge/count/"
	CommitKeyPrefix  = "Commit/value/"

	AttributeKeyQuestionId = "question_id"
	AttributeKeyCreator    = "creator"
	AttributeKeyBounty     = "bounty"
	AttributeKeyCommitter  = "committer"
	AttributeKeyCommitHash = "commit_hash"
	AttributeKeyWinner     = "winner"
)

var (
	ParamsKey = []byte("p_scavenge")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// CommitKey returns the store key to retrieve a Commit
func CommitKey(questionId uint64, creator string) []byte {
	questionIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(questionIdBytes, questionId)
	return append(append(KeyPrefix(CommitKeyPrefix), questionIdBytes...), []byte(creator)...)
}

func GetCommittedAnswerKey(questionId uint64, creator string) []byte {
	questionIdBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(questionIdBytes, questionId)

	var key []byte
	key = append(key, KeyPrefix(CommitKeyPrefix)...)
	key = append(key, questionIdBytes...)
	key = append(key, []byte(creator)...)

	return key
}
