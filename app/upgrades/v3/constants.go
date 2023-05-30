package v3

import (
	"github.com/mage-coven/gridiron/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Gridiron v3 upgrade.
const UpgradeName = "v3"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
}
