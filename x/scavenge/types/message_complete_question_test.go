package types

import (
	"testing"

	"scavenge/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCompleteQuestion_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCompleteQuestion
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCompleteQuestion{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCompleteQuestion{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
