package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCommitAnswer{}

func NewMsgCommitAnswer(creator string, questionId uint64, hashAnswer string) *MsgCommitAnswer {
	return &MsgCommitAnswer{
		Creator:    creator,
		QuestionId: questionId,
		HashAnswer: hashAnswer,
	}
}

func (msg *MsgCommitAnswer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
