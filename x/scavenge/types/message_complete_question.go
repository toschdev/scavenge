package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCompleteQuestion{}

func NewMsgCompleteQuestion(creator string, questionId uint64, winner string) *MsgCompleteQuestion {
	return &MsgCompleteQuestion{
		Creator:    creator,
		QuestionId: questionId,
		Winner:     winner,
	}
}

func (msg *MsgCompleteQuestion) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
