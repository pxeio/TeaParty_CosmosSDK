package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TradeOrdersKeyPrefix is the prefix to retrieve all TradeOrders
	TradeOrdersKeyPrefix = "TradeOrders/value/"
)

// TradeOrdersKey returns the store key to retrieve a TradeOrders from the index fields
func TradeOrdersKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
