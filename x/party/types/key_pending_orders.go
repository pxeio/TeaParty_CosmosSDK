package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// PendingOrdersKeyPrefix is the prefix to retrieve all PendingOrders
	PendingOrdersKeyPrefix = "PendingOrders/value/"
)

// PendingOrdersKey returns the store key to retrieve a PendingOrders from the index fields
func PendingOrdersKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
