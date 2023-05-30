package app

import (
	snapshot "github.com/cosmos/cosmos-sdk/snapshots/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	protoio "github.com/gogo/protobuf/io"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/CosmWasm/wasmd/x/wasm/keeper"
)

var _ snapshot.ExtensionSnapshotter = &GridironSnapshotter{}

// GridironSnapshotter custom cosmos-sdk snapshotter that extends
// the wasmd snapshotter to restore poe contract cache
type GridironSnapshotter struct {
	*keeper.WasmSnapshotter
	app *GridironApp
}

// NewGridironSnapshotter constructor
func NewGridironSnapshotter(app *GridironApp, cms sdk.MultiStore, wasm *keeper.Keeper) *GridironSnapshotter {
	return &GridironSnapshotter{
		WasmSnapshotter: keeper.NewWasmSnapshotter(cms, wasm),
		app:             app,
	}
}

// Restore restores wasm files and caches
func (t *GridironSnapshotter) Restore(
	height uint64, format uint32, protoReader protoio.Reader,
) (snapshot.SnapshotItem, error) {
	item, err := t.WasmSnapshotter.Restore(height, format, protoReader)
	if err == nil {
		// pinned contract are initialized in WasmSnapshotter already
		ctx := t.app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		t.app.poeKeeper.InitContractAddressCache(ctx)
	}
	return item, err
}
