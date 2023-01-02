package keeper

import (
	"github.com/TeaPartyCrypto/partychain/x/party/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTradeOrders set a specific tradeOrders in the store from its index
func (k Keeper) SetTradeOrders(ctx sdk.Context, tradeOrders types.TradeOrders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeOrdersKeyPrefix))
	b := k.cdc.MustMarshal(&tradeOrders)
	store.Set(types.TradeOrdersKey(
		tradeOrders.Index,
	), b)
}

// GetTradeOrders returns a tradeOrders from its index
func (k Keeper) GetTradeOrders(
	ctx sdk.Context,
	index string,

) (val types.TradeOrders, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeOrdersKeyPrefix))

	b := store.Get(types.TradeOrdersKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTradeOrders removes a tradeOrders from the store
func (k Keeper) RemoveTradeOrders(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeOrdersKeyPrefix))
	store.Delete(types.TradeOrdersKey(
		index,
	))
}

// GetAllTradeOrders returns all tradeOrders
func (k Keeper) GetAllTradeOrders(ctx sdk.Context) (list []types.TradeOrders) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TradeOrdersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TradeOrders
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
