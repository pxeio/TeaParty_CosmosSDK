package keeper

import (
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPendingOrders set a specific pendingOrders in the store from its index
func (k Keeper) SetPendingOrders(ctx sdk.Context, pendingOrders types.PendingOrders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingOrdersKeyPrefix))
	b := k.cdc.MustMarshal(&pendingOrders)
	store.Set(types.PendingOrdersKey(
		pendingOrders.Index,
	), b)
}

// GetPendingOrders returns a pendingOrders from its index
func (k Keeper) GetPendingOrders(
	ctx sdk.Context,
	index string,

) (val types.PendingOrders, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingOrdersKeyPrefix))

	b := store.Get(types.PendingOrdersKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePendingOrders removes a pendingOrders from the store
func (k Keeper) RemovePendingOrders(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingOrdersKeyPrefix))
	store.Delete(types.PendingOrdersKey(
		index,
	))
}

// GetAllPendingOrders returns all pendingOrders
func (k Keeper) GetAllPendingOrders(ctx sdk.Context) (list []types.PendingOrders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PendingOrdersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PendingOrders
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
