package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const PBMFormat = "P1"

func WriteOnFile(fileName string, matrix [][]string, rows, cols int) error {
	fileNameWrite := fmt.Sprintf("%s.pbm", fileName)
	fwrite, err := os.Create(fileNameWrite)
	if err != nil {
		fmt.Printf("Error to create response file: %v\nPlease try again.", err)
		return err
	}
	defer fwrite.Close()

	var dataReponse string
	var dataReponseLine string
	dataReponse += (PBMFormat + "\n")
	dataReponseLine += fmt.Sprintf("%d %d", cols, rows)
	dataReponse += (dataReponseLine + "\n")

	for _, v := range matrix {
		dataReponseLine = strings.Join(v, " ")
		dataReponse += (dataReponseLine + "\n")
	}

	dataValue := fmt.Sprintf("%v", dataReponse)
	dataResult := []byte(dataValue)

	_, err = fwrite.Write(dataResult)

	if err != nil {
		fmt.Printf("Error to write data to the file: %v\nPlease try again.", err)
		return err
	}

	return nil
}

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

func ReadFile(fName string) ([][]string, error) {
	fileName := fmt.Sprintf("%s.pbm", fName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\nPlease enter the right file name.", err)
		return nil, err
	}

	data := make([][]string, 0)
	reader := bufio.NewReader(file)
	stringLine, err := Readln(reader)
	for err == nil {
		if strings.Contains(string(stringLine[0]), "#") {
			fmt.Printf(stringLine)
		} else {
			splitData := strings.Split(stringLine, " ")
			data = append(data, splitData)
		}
		stringLine, err = Readln(reader)
	}

	PBMFormatFile := data[0][0]
	if PBMFormat != PBMFormatFile {
		fmt.Printf("PBM file is different thand P1 type. Please try again with right file.")
		return nil, err
	}

	matrix := data[2:]

	return matrix, nil
}

func Readln(reader *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = reader.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func main() {

	var arg1 string
	var arg2 string

	if len(os.Args) != 3 {
		fmt.Println("Invalid paramenter," +
			" please type the file name and degres you desire to rotate, example:\n" +
			"appName fileName 90")
		return
	} else {
		arg1 = os.Args[1]
		arg2 = os.Args[2]
	}

	fmt.Println("Reading file " + arg1 + ".pbm, to rotate image file with " + arg2 + " degres")

	matrix, err := ReadFile(arg1)
	if err != nil {
		fmt.Printf("Error reading file")
		return
	}

	rotatedMatrix, rowsRotated, colsRotated := RotateImage(matrix, arg2)

	err = WriteOnFile(arg1, rotatedMatrix, rowsRotated, colsRotated)
	if err != nil {
		return
	}

	fmt.Println("Success ! Your file is done !")
}
