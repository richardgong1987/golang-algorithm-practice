package util

func Swap(arr []int, a int, b int) []int {
	arr[a], arr[b] = arr[b], arr[a]

	return arr
}
