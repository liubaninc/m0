package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/liubaninc/m0/x/mibc/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
)

// ChanCloseInit defines a wrapper function for the channel Keeper's function
func (k Keeper) ChanCloseInit(ctx sdk.Context, portID, channelID string) error {
	capName := host.ChannelCapabilityPath(portID, channelID)
	chanCap, ok := k.scopedKeeper.GetCapability(ctx, capName)
	if !ok {
		return sdkerrors.Wrapf(channeltypes.ErrChannelCapabilityNotFound, "could not retrieve channel capability at: %s", capName)
	}
	return k.channelKeeper.ChanCloseInit(ctx, portID, channelID, chanCap)
}

// IsBound checks if the module is already bound to the desired port
func (k Keeper) IsBound(ctx sdk.Context, portID string) bool {
	_, ok := k.scopedKeeper.GetCapability(ctx, host.PortPath(portID))
	return ok
}

// BindPort defines a wrapper function for the ort Keeper's function in
// order to expose it to module's InitGenesis function
func (k Keeper) BindPort(ctx sdk.Context, portID string) error {
	cap := k.portKeeper.BindPort(ctx, portID)
	return k.ClaimCapability(ctx, cap, host.PortPath(portID))
}

// GetPort returns the portID for the module. Used in ExportGenesis
func (k Keeper) GetPort(ctx sdk.Context) string {
	store := ctx.KVStore(k.storeKey)
	return string(store.Get(types.PortKey))
}

// SetPort sets the portID for the module. Used in InitGenesis
func (k Keeper) SetPort(ctx sdk.Context, portID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.PortKey, []byte(portID))
}

// AuthenticateCapability wraps the scopedKeeper's AuthenticateCapability function
func (k Keeper) AuthenticateCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) bool {
	return k.scopedKeeper.AuthenticateCapability(ctx, cap, name)
}

// ClaimCapability allows the module that can claim a capability that IBC module passes to it
func (k Keeper) ClaimCapability(ctx sdk.Context, cap *capabilitytypes.Capability, name string) error {
	return k.scopedKeeper.ClaimCapability(ctx, cap, name)
}

// GetDenomTrace retreives the full identifiers trace and base denomination from the store.
func (k Keeper) GetDenomTrace(ctx sdk.Context, denomTraceHash tmbytes.HexBytes) (types.DenomTrace, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DenomTraceKey)
	bz := store.Get(denomTraceHash)
	if bz == nil {
		return types.DenomTrace{}, false
	}

	denomTrace := k.MustUnmarshalDenomTrace(bz)
	return denomTrace, true
}

// HasDenomTrace checks if a the key with the given denomination trace hash exists on the store.
func (k Keeper) HasDenomTrace(ctx sdk.Context, denomTraceHash tmbytes.HexBytes) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DenomTraceKey)
	return store.Has(denomTraceHash)
}

// SetDenomTrace sets a new {trace hash -> denom trace} pair to the store.
func (k Keeper) SetDenomTrace(ctx sdk.Context, denomTrace types.DenomTrace) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.DenomTraceKey)
	bz := k.MustMarshalDenomTrace(denomTrace)
	store.Set(denomTrace.Hash(), bz)
}

// GetAllDenomTraces returns the trace information for all the denominations.
func (k Keeper) GetAllDenomTraces(ctx sdk.Context) types.Traces {
	traces := types.Traces{}
	k.IterateDenomTraces(ctx, func(denomTrace types.DenomTrace) bool {
		traces = append(traces, denomTrace)
		return false
	})

	return traces.Sort()
}

// IterateDenomTraces iterates over the denomination traces in the store
// and performs a callback function.
func (k Keeper) IterateDenomTraces(ctx sdk.Context, cb func(denomTrace types.DenomTrace) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DenomTraceKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {

		denomTrace := k.MustUnmarshalDenomTrace(iterator.Value())
		if cb(denomTrace) {
			break
		}
	}
}

// UnmarshalDenomTrace attempts to decode and return an DenomTrace object from
// raw encoded bytes.
func (k Keeper) UnmarshalDenomTrace(bz []byte) (types.DenomTrace, error) {
	var denomTrace types.DenomTrace
	if err := k.cdc.UnmarshalBinaryBare(bz, &denomTrace); err != nil {
		return types.DenomTrace{}, err
	}
	return denomTrace, nil
}

// MustUnmarshalDenomTrace attempts to decode and return an DenomTrace object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalDenomTrace(bz []byte) types.DenomTrace {
	var denomTrace types.DenomTrace
	k.cdc.MustUnmarshalBinaryBare(bz, &denomTrace)
	return denomTrace
}

// MarshalDenomTrace attempts to encode an DenomTrace object and returns the
// raw encoded bytes.
func (k Keeper) MarshalDenomTrace(denomTrace types.DenomTrace) ([]byte, error) {
	return k.cdc.MarshalBinaryBare(&denomTrace)
}

// MustMarshalDenomTrace attempts to encode an DenomTrace object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalDenomTrace(denomTrace types.DenomTrace) []byte {
	return k.cdc.MustMarshalBinaryBare(&denomTrace)
}

// DenomPathFromHash returns the full denomination path prefix from an ibc denom with a hash
// component.
func (k Keeper) DenomPathFromHash(ctx sdk.Context, denom string) (string, error) {
	// trim the denomination prefix, by default "ibc/"
	hexHash := denom[len(types.DenomPrefix+"/"):]

	hash, err := types.ParseHexHash(hexHash)
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrInvalidDenomForTransfer, err.Error())
	}

	denomTrace, found := k.GetDenomTrace(ctx, hash)
	if !found {
		return "", sdkerrors.Wrap(types.ErrTraceNotFound, hexHash)
	}

	fullDenomPath := denomTrace.GetFullDenomPath()
	return fullDenomPath, nil
}