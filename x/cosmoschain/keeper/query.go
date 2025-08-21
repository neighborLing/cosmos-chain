package keeper

import (
	"cosmos-chain/x/cosmoschain/types"
)

var _ types.QueryServer = Keeper{}
