package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FinalizingOrdersKeyPrefix is the prefix to retrieve all FinalizingOrders
	FinalizingOrdersKeyPrefix = "FinalizingOrders/value/"
)

// FinalizingOrdersKey returns the store key to retrieve a FinalizingOrders from the index fields
func FinalizingOrdersKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
