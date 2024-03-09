package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/akiroinu/akiro/testutil/keeper"
	"github.com/akiroinu/akiro/x/akiro/keeper"
	"github.com/akiroinu/akiro/x/akiro/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AkiroKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
