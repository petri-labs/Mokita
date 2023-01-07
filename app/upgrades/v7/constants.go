package v7

import (
	"github.com/CosmWasm/wasmd/x/wasm"

	"github.com/petri-labs/mokita/app/upgrades"
	superfluidtypes "github.com/petri-labs/mokita/x/superfluid/types"

	store "github.com/cosmos/cosmos-sdk/store/types"
)

// UpgradeName defines the on-chain upgrade name for the Mokisis v7 upgrade.
const UpgradeName = "v7"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{wasm.ModuleName, superfluidtypes.ModuleName},
	},
}
