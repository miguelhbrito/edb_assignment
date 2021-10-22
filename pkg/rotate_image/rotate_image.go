package rotateimage

import (
	"fmt"
)

func Transpose(matrix [][]string) ([][]string, int, int) {
	cols := len(matrix[0])
	rows := len(matrix)

	result := make([][]string, cols)
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

func ReverseRows(matrix [][]string) ([][]string, int, int) {
	cols := len(matrix[0])
	rows := len(matrix)

	result := make([][]string, rows)
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

func RotateImage(matrix [][]string, degres string) ([][]string, int, int) {
	var rotatedMatrix [][]string
	var rowsRotated int
	var colsRotated int

	switch degres {
	case "90", "-270":
		rotatedMatrix, _, _ = Transpose(matrix)
		rotatedMatrix, rowsRotated, colsRotated = ReverseRows(rotatedMatrix)
	case "180", "-180":
		rotatedMatrix, _, _ = Transpose(matrix)
		rotatedMatrix, _, _ = ReverseRows(rotatedMatrix)
		rotatedMatrix, _, _ = Transpose(rotatedMatrix)
		rotatedMatrix, rowsRotated, colsRotated = ReverseRows(rotatedMatrix)
	case "270", "-90":
		rotatedMatrix, _, _ = ReverseRows(matrix)
		rotatedMatrix, rowsRotated, colsRotated = Transpose(rotatedMatrix)
	case "reverse":
		rotatedMatrix, rowsRotated, colsRotated = ReverseRows(matrix)
	default:
		fmt.Printf("Please, type the right input to degres, examples:\n" +
			"90, -90, 180, -180, 270, -270 or reverse")
	}

	return rotatedMatrix, rowsRotated, colsRotated
}
