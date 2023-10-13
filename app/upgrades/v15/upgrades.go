// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)
package v15

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	strideoutpost "github.com/evmos/evmos/v14/precompiles/outposts/stride"
	evmkeeper "github.com/evmos/evmos/v14/x/evm/keeper"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v15.0.0
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		// we are depecrating crisis module since it is not being used
		logger.Debug("deleting crisis module from version map...")
		delete(vm, "crisis")

		// Leave modules are as-is to avoid running InitGenesis.
		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, vm)
	}
}

// EnableOutposts appends the address of the Stride and Osmosis Outposts
// to the list of active precompiles.
// NOTE: the osmosis outpost address need to be added
func EnableOutposts(ctx sdk.Context, evmKeeper *evmkeeper.Keeper) error {
	// Get the list of active precompiles from the genesis state
	params := evmKeeper.GetParams(ctx)
	activePrecompiles := params.ActivePrecompiles
	activePrecompiles = append(activePrecompiles, strideoutpost.Precompile{}.Address().String())
	params.ActivePrecompiles = activePrecompiles

	return evmKeeper.SetParams(ctx, params)
}