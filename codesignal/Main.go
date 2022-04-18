package codesignal

import "fmt"

func Main() {
	image := Minesweeper([][]bool{
		{false, false, false},
		{false, false, false},
	})
	fmt.Println(image)
}
