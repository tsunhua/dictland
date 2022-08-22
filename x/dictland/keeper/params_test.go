package keeper_test

import (
	"testing"

	testkeeper "dictland/testutil/keeper"
	"dictland/x/dictland/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DictlandKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
