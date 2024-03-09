package keeper_test

import (
	"testing"

	testkeeper "github.com/akiroinu/akiro/testutil/keeper"
	"github.com/akiroinu/akiro/x/akiro/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AkiroKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
