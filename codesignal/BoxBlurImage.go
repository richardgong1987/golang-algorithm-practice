package codesignal

func BoxBlurImage(image [][]int) [][]int {
	rowLen := len(image)
	colLen := len(image[0])

	rPoiter := 0
	cPoiter := 0

	var blurImage [][]int

	for rPoiter <= rowLen-3 {
		var rowBlur []int
		for cPoiter <= colLen-3 {
			unitAvg := getUnitSum(image, rPoiter, cPoiter)
			rowBlur = append(rowBlur, unitAvg)
			cPoiter++
		}

		blurImage = append(blurImage, rowBlur)

		rPoiter++
		cPoiter = 0
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
