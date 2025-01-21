package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "scavenge/testutil/keeper"
	"scavenge/x/scavenge/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ScavengeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
