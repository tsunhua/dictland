package keeper

import (
	"dictland/x/dictland/types"
)

var _ types.QueryServer = Keeper{}
