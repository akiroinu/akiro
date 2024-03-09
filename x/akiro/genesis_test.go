package akiro_test

import (
	"testing"

	keepertest "github.com/akiroinu/akiro/testutil/keeper"
	"github.com/akiroinu/akiro/testutil/nullify"
	"github.com/akiroinu/akiro/x/akiro"
	"github.com/akiroinu/akiro/x/akiro/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AkiroKeeper(t)
	akiro.InitGenesis(ctx, *k, genesisState)
	got := akiro.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
