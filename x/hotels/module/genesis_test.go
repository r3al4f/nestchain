package hotels_test

import (
	"testing"

	keepertest "github.com/r3al4f/nestchain/testutil/keeper"
	"github.com/r3al4f/nestchain/testutil/nullify"
	hotels "github.com/r3al4f/nestchain/x/hotels/module"
	"github.com/r3al4f/nestchain/x/hotels/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.HotelsKeeper(t)
	hotels.InitGenesis(ctx, k, genesisState)
	got := hotels.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
