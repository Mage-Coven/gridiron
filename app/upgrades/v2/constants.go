package v2

import (
	"github.com/mage-coven/gridiron/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Gridiron v2 upgrade.
const UpgradeName = "v2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
}
