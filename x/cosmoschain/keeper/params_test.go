package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "cosmos-chain/testutil/keeper"
	"cosmos-chain/x/cosmoschain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CosmoschainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
