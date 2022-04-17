package util

func SliceShift[T any](slice []T) []T {
	if len(slice) > 0 {
		return append(slice[:0], slice[1:]...)
	}
	return slice
}
