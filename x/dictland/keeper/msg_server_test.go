package keeper_test

import (
	"context"
	"testing"

	keepertest "dictland/testutil/keeper"
	"dictland/x/dictland/keeper"
	"dictland/x/dictland/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DictlandKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
