package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
)

// SetOrdersUnderWatch set a specific ordersUnderWatch in the store from its index
func (k Keeper) SetOrdersUnderWatch(ctx sdk.Context, ordersUnderWatch types.OrdersUnderWatch) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersUnderWatchKeyPrefix))
	b := k.cdc.MustMarshal(&ordersUnderWatch)
	store.Set(types.OrdersUnderWatchKey(
        ordersUnderWatch.Index,
    ), b)
}

// GetOrdersUnderWatch returns a ordersUnderWatch from its index
func (k Keeper) GetOrdersUnderWatch(
    ctx sdk.Context,
    index string,
    
) (val types.OrdersUnderWatch, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersUnderWatchKeyPrefix))

	b := store.Get(types.OrdersUnderWatchKey(
        index,
    ))
    if b == nil {
        return val, false
    }

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOrdersUnderWatch removes a ordersUnderWatch from the store
func (k Keeper) RemoveOrdersUnderWatch(
    ctx sdk.Context,
    index string,
    
) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersUnderWatchKeyPrefix))
	store.Delete(types.OrdersUnderWatchKey(
	    index,
    ))
}

// GetAllOrdersUnderWatch returns all ordersUnderWatch
func (k Keeper) GetAllOrdersUnderWatch(ctx sdk.Context) (list []types.OrdersUnderWatch) {
    store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OrdersUnderWatchKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.OrdersUnderWatch
		k.cdc.MustUnmarshal(iterator.Value(), &val)
        list = append(list, val)
	}

    return
}
