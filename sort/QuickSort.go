package sort

import "fmt"

func QuickSort(arr *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}
	pivot := partition(arr, lo, hi)
	QuickSort(arr, lo, pivot-1)
	QuickSort(arr, pivot+1, hi)
}

func partition(arr *[]int, lo int, hi int) int {
	pivot := (*arr)[lo]
	var left = lo + 1
	var right = hi

	for left != right {
		// right
		for left < right && (*arr)[right] > pivot {
			right--
		}
		// left
		for left < right && (*arr)[left] <= pivot {
			left++
		}

		if left < right {
			tmp := (*arr)[left]
			(*arr)[left] = (*arr)[right]
			(*arr)[right] = tmp
		}
	}

	(*arr)[lo] = (*arr)[left]
	(*arr)[left] = pivot
	return left

}

func main() {
	arr := []int{4, 4, 6, 5, 3, 2, 8, 1}
	QuickSort(&arr, 0, len(arr)-1)
	fmt.Println(arr)

}
