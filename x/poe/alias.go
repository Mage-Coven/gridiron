// nolint

package poe

import (
	"github.com/mage-coven/gridiron/x/poe/keeper"
	"github.com/mage-coven/gridiron/x/poe/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey
)

type DeliverTxfn = keeper.DeliverTxFn
