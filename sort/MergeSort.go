package sort

import "fmt"

func mergeSort(nums []int, lo int, hi int, tmp []int) {
	if lo == hi {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(nums, lo, mid, tmp)
	mergeSort(nums, mid+1, hi, tmp)
	merge(nums, lo, mid, hi, tmp)

}
func merge(nums []int, lo int, mid int, hi int, tmp []int) {
	for i := lo; i <= hi; i++ {
		tmp = append(tmp[:i], nums[i])
	}

	var i, j = lo, mid + 1
	for p := lo; p <= hi; p++ {
		if i == mid+1 {
			// 左半边数组已全部被合并
			nums = append(nums[:p], tmp[j])
			j += 1
		} else if j == hi+1 {
			// 右半边数组已全部被合并
			nums = append(nums[:p], tmp[i])
			i += 1
		} else if tmp[i] > tmp[j] {
			nums = append(nums[:p], tmp[j])
			j += 1
		} else {
			nums = append(nums[:p], tmp[i])
			i += 1
		}

	}
}
func main() {
	//ints := []int{5, 8, 6, 3, 9, 2, 1, 7}
	ints := []int{3, 4, 1, 2}
	hi := len(ints) - 1
	mergeSort(ints, 0, hi, make([]int, hi))
	fmt.Println(ints)
}
