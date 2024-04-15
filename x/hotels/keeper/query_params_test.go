package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/r3al4f/nestchain/testutil/keeper"
	"github.com/r3al4f/nestchain/x/hotels/keeper"
	"github.com/r3al4f/nestchain/x/hotels/types"
)

func TestParamsQuery(t *testing.T) {
	k, ctx, _ := keepertest.HotelsKeeper(t)

	qs := keeper.NewQueryServerImpl(k)
	params := types.DefaultParams()
	require.NoError(t, k.Params.Set(ctx, params))

	response, err := qs.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
