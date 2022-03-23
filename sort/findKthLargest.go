package sort

import "fmt"

func main() {
	largest := findKthLargest(&[]int{3, 2, 1, 5, 6, 4}, 2)
	fmt.Println(largest == 5)
}
func findKthLargest(nums *[]int, k int) int {
	hi := len(*nums) - 1
	lo := 0

	k = hi - k + 1
	for lo <= hi {
		pivot := partition2(nums, lo, hi)
		if k > pivot {
			lo = pivot + 1
		} else if k < pivot {
			hi = pivot - 1
		} else {
			return (*nums)[k]
		}
	}

	return -1
}
func partition2(nums *[]int, lo int, hi int) int {
	pivot := (*nums)[lo]
	left := lo
	right := hi

	for left != right {
		for left < right && (*nums)[right] > pivot {
			right--
		}
		for left < right && (*nums)[left] <= pivot {
			left++
		}
		if left < right {
			swap(nums, left, right)
		}
	}
	swap(nums, lo, left)
	return left
}

func swap(nums *[]int, i int, j int) {
	tmp := (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = tmp
}
