package keeper

import (
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetFinalizingOrders set a specific finalizingOrders in the store from its index
func (k Keeper) SetFinalizingOrders(ctx sdk.Context, finalizingOrders types.FinalizingOrders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FinalizingOrdersKeyPrefix))
	b := k.cdc.MustMarshal(&finalizingOrders)
	store.Set(types.FinalizingOrdersKey(
		finalizingOrders.Index,
	), b)
}

// GetFinalizingOrders returns a finalizingOrders from its index
func (k Keeper) GetFinalizingOrders(
	ctx sdk.Context,
	index string,

) (val types.FinalizingOrders, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FinalizingOrdersKeyPrefix))

	b := store.Get(types.FinalizingOrdersKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFinalizingOrders removes a finalizingOrders from the store
func (k Keeper) RemoveFinalizingOrders(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FinalizingOrdersKeyPrefix))
	store.Delete(types.FinalizingOrdersKey(
		index,
	))
}

// GetAllFinalizingOrders returns all finalizingOrders
func (k Keeper) GetAllFinalizingOrders(ctx sdk.Context) (list []types.FinalizingOrders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FinalizingOrdersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FinalizingOrders
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
