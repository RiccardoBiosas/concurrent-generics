package concurrent_generic_slice

import (
	"sync"
)

type GenericSlice[T any] []T

type ConcurrentGenericSlice[T any] struct {
	sync.RWMutex
	items GenericSlice[T]
}

type ConcurrentGenericSliceItem[T any] struct {
	Index int
	Value T
}
