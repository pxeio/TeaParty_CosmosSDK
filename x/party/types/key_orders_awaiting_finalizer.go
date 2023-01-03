package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OrdersAwaitingFinalizerKeyPrefix is the prefix to retrieve all OrdersAwaitingFinalizer
	OrdersAwaitingFinalizerKeyPrefix = "OrdersAwaitingFinalizer/value/"
)

// OrdersAwaitingFinalizerKey returns the store key to retrieve a OrdersAwaitingFinalizer from the index fields
func OrdersAwaitingFinalizerKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
