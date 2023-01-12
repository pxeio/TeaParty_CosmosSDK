package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// OrdersUnderWatchKeyPrefix is the prefix to retrieve all OrdersUnderWatch
	OrdersUnderWatchKeyPrefix = "OrdersUnderWatch/value/"
)

// OrdersUnderWatchKey returns the store key to retrieve a OrdersUnderWatch from the index fields
func OrdersUnderWatchKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
