package codesignal

/**

Notice:

j-1 >= 0,j+1 < colLen,i-1 >= 0,i+1 <=rowLen

Mainly to judge whether the subscript of the array is out of bounds, otherwise an exception will be reported

*/
func Minesweeper(matrix [][]bool) [][]int {
	rowLen := len(matrix)
	colLen := len(matrix[0])

	var mineCount = get2DIntSlice(rowLen, colLen)

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			// top
			if j-1 >= 0 && matrix[i][j-1] {
				updateCount(mineCount, i, j)
			}
			// bottom
			if j+1 < colLen && matrix[i][j+1] {
				updateCount(mineCount, i, j)
			}
			// left
			if i-1 >= 0 && matrix[i-1][j] {
				updateCount(mineCount, i, j)
			}
			// right
			if i+1 < rowLen && matrix[i+1][j] {
				updateCount(mineCount, i, j)
			}
			// bottom right
			if i+1 < rowLen && j+1 < colLen && matrix[i+1][j+1] {
				updateCount(mineCount, i, j)
			}
			// bottom left
			if i+1 < rowLen && j-1 >= 0 && matrix[i+1][j-1] {
				updateCount(mineCount, i, j)
			}

			// top right
			if i-1 >= 0 && j+1 < colLen && matrix[i-1][j+1] {
				updateCount(mineCount, i, j)
			}
			// top left
			if i-1 >= 0 && j-1 >= 0 && matrix[i-1][j-1] {
				updateCount(mineCount, i, j)
			}
		}
	}
	return mineCount
}
func updateCount(mineCount [][]int, i, j int) {
	mineCount[i][j]++
}
func get2DIntSlice(rowLen, colLen int) [][]int {
	box := make([][]int, rowLen)
	for i := 0; i < rowLen; i++ {
		box[i] = make([]int, colLen)
	}
	return box
}
