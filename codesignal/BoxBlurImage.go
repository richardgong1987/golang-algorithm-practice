package codesignal

func BoxBlurImage(image [][]int) [][]int {
	rowLen := len(image)
	colLen := len(image[0])

	var blurImage [][]int

	for rPoiter := 0; rPoiter <= rowLen-3; rPoiter++ {
		var rowBlur []int
		for cPoiter := 0; cPoiter <= colLen-3; cPoiter++ {
			unitAvg := getUnitSum(image, rPoiter, cPoiter)
			rowBlur = append(rowBlur, unitAvg)
		}
		blurImage = append(blurImage, rowBlur)
	}

	return blurImage
}

func getUnitSum(image [][]int, rPoiter int, cPoiter int) int {
	var unitItem [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			unitItem[i][j] = image[rPoiter+i][cPoiter+j]
		}
	}
	return getAVG(unitItem)
}

func getAVG(square [3][3]int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += square[i][j]
		}
	}
	return sum / 9
}
