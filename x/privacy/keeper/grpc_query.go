package keeper

import (
	"privacy/x/privacy/types"
)

var _ types.QueryServer = Keeper{}
