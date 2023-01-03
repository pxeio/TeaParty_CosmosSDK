package keeper

import (
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetOrdersAwaitingFinalizer set a specific ordersAwaitingFinalizer in the store from its index
func (k Keeper) SetOrdersAwaitingFinalizer(ctx sdk.Context, ordersAwaitingFinalizer types.OrdersAwaitingFinalizer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersAwaitingFinalizerKeyPrefix))
	b := k.cdc.MustMarshal(&ordersAwaitingFinalizer)
	store.Set(types.OrdersAwaitingFinalizerKey(
		ordersAwaitingFinalizer.Index,
	), b)
}

// GetOrdersAwaitingFinalizer returns a ordersAwaitingFinalizer from its index
func (k Keeper) GetOrdersAwaitingFinalizer(
	ctx sdk.Context,
	index string,

) (val types.OrdersAwaitingFinalizer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersAwaitingFinalizerKeyPrefix))

	b := store.Get(types.OrdersAwaitingFinalizerKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrdersAwaitingFinalizer removes a ordersAwaitingFinalizer from the store
func (k Keeper) RemoveOrdersAwaitingFinalizer(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersAwaitingFinalizerKeyPrefix))
	store.Delete(types.OrdersAwaitingFinalizerKey(
		index,
	))
}

// GetAllOrdersAwaitingFinalizer returns all ordersAwaitingFinalizer
func (k Keeper) GetAllOrdersAwaitingFinalizer(ctx sdk.Context) (list []types.OrdersAwaitingFinalizer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersAwaitingFinalizerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OrdersAwaitingFinalizer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
