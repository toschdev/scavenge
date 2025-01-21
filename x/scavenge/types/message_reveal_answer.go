package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRevealAnswer{}

func NewMsgRevealAnswer(creator string, questionId uint64, answer string, plainText string) *MsgRevealAnswer {
	return &MsgRevealAnswer{
		Creator:    creator,
		QuestionId: questionId,
		Answer:     answer,
		PlainText:  plainText,
	}
}

func (msg *MsgRevealAnswer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
