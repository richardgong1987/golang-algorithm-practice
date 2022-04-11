package word

import "fmt"

func Main() {
	//beginWord := "hit"
	//endWord := "cog"
	//wordList := []string{"hot", "dot", "dog", "lot", "log"}
	//
	//length := ladderLength(beginWord, endWord, wordList)
	//fmt.Println(length)

	i := solution([]int{8, 1, 0, 4, 7})
	fmt.Println(i)

}

func solution(statues []int) int {
	quickSort(statues, 0, len(statues)-1)
	arrLen := len(statues)
	return (statues[arrLen-1] - statues[0] + 1) - arrLen
}

func quickSort(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	pivot := partition(arr, lo, hi)
	quickSort(arr, lo, pivot-1)
	quickSort(arr, pivot+1, hi)
}
func partition(arr []int, lo int, hi int) int {
	pivot := arr[lo]
	left := lo
	right := hi
	for left != right {
		for left < right && arr[right] > pivot {
			right--
		}
		for left < right && arr[left] <= pivot {
			left++
		}

		if left < right {
			swap(arr, left, right)
		}
	}
	swap(arr, left, lo)
	return left
}
func swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
