package keeper

import (
	"github.com/decspeed/hello/x/hello/types"
)

var _ types.QueryServer = Keeper{}
