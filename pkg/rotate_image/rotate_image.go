package rotateimage

type (
	Image   [][]string
	Degrees string
)

const (
	ninety                  Degrees = "90"
	oneEighty               Degrees = "180"
	twoHundreadSeventy      Degrees = "270"
	minusNinety             Degrees = "-90"
	minusOneEighty          Degrees = "-180"
	minusTwoHundreadSeventy Degrees = "-270"
)

func (d Degrees) IsDegrees() bool {
	switch d {
	case ninety, oneEighty, twoHundreadSeventy, minusNinety, minusOneEighty, minusTwoHundreadSeventy:
		return true
	default:
		return false
	}
}

func (i Image) transpose(matrix Image) (Image, int, int) {
	cols := len(matrix[0])
	rows := len(matrix)

	result := make(Image, cols)
	for i := range result {
		result[i] = make([]string, rows)
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			result[i][j] = matrix[j][i]
		}
	}

	colsResult := len(result[0])
	rowsResult := len(result)

	return result, rowsResult, colsResult
}

func (i Image) reverseRows(matrix Image) (Image, int, int) {
	cols := len(matrix[0])
	rows := len(matrix)

	result := make(Image, rows)
	for i := range result {
		result[i] = make([]string, cols)
	}

	for i := 0; i < rows; i++ {
		for j, rigthCol := cols-1, 0; j >= 0; j, rigthCol = j-1, rigthCol+1 {
			result[i][rigthCol] = matrix[i][j]
		}
	}

	return result, rows, cols
}

func (i Image) RotateImage(matrix Image, degres Degrees) (Image, int, int) {
	var rotatedMatrix Image
	var rowsRotated int
	var colsRotated int

	switch degres {
	case "90", "-270":
		rotatedMatrix, _, _ = matrix.transpose(matrix)
		rotatedMatrix, rowsRotated, colsRotated = matrix.reverseRows(rotatedMatrix)
	case "180", "-180":
		rotatedMatrix, _, _ = matrix.transpose(matrix)
		rotatedMatrix, _, _ = matrix.reverseRows(rotatedMatrix)
		rotatedMatrix, _, _ = matrix.transpose(rotatedMatrix)
		rotatedMatrix, rowsRotated, colsRotated = matrix.reverseRows(rotatedMatrix)
	case "270", "-90":
		rotatedMatrix, _, _ = matrix.reverseRows(matrix)
		rotatedMatrix, rowsRotated, colsRotated = matrix.transpose(rotatedMatrix)
	case "reverse":
		rotatedMatrix, rowsRotated, colsRotated = matrix.reverseRows(matrix)
	default:
	}

	return rotatedMatrix, rowsRotated, colsRotated
}
