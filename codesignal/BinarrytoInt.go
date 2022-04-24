package codesignal

import (
	"math"
	"strconv"
)

func BinarrytoInt(str string) int32 {
	l := len(str)
	result := 0.0
	for i, n := range str {
		num, _ := strconv.ParseFloat(string(n), -1)
		result += math.Exp2(float64(l-i-1)) * num
	}
	return int32(result)
}
