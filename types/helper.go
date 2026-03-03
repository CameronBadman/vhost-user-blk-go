package types

import "unsafe"

func castToBytes[T any](v *T) []byte {
	size := unsafe.Sizeof(*v)
	return unsafe.Slice((*byte)(unsafe.Pointer(v)), size)
}
