package word

import (
	"fmt"
	"strings"
)

func Main() {
	//beginWord := "hit"
	//endWord := "cog"
	//wordList := []string{"hot", "dot", "dog", "lot", "log"}
	//
	//length := ladderLength(beginWord, endWord, wordList)
	//fmt.Println(length)
	fmt.Println(strings.Repeat("*", 4))

}
func solution(inputString string) string {
	return ""
}
func compareTriplets(a []int32, b []int32) []int32 {
	// Write your code here
	arr := []int32{0, 0}
	for i := range a {
		if a[i] > b[i] {
			arr[0] = arr[0] + 1
		} else if a[i] < b[i] {
			arr[1] = arr[1] + 1
		}
	}
	return arr
}
func strReverse(str string) string {
	ret := ""
	for _, v := range str {
		ret = string(v) + ret
	}
	return ret
}
