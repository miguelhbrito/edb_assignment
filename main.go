package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func main() {
	if os.Args[1] == "" {
		fmt.Println("Invalid paramenter, argument 1 is empty")
		os.Exit(2)
	}

	if os.Args[2] == "" {
		fmt.Println("Invalid paramenter, argument 2 is empty")
		os.Exit(2)
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	fmt.Println("Reading file " + arg1 + ".pbm, to rotate image file with " + arg2 + " degres")

	fileName := fmt.Sprintf("%s.pbm", arg1)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
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

	PBMFormat := data[0][0]
	MatrixParams := data[1]
	fmt.Println(MatrixParams)

	matrix := data[2:]

	var rotatedMatrix [][]string
	var rowsRotated int
	var colsRotated int

	if arg2 == "90" || arg2 == "-270" {
		rotatedMatrix, _, _ = Transpose(matrix)
		rotatedMatrix, rowsRotated, colsRotated = ReverseRows(rotatedMatrix)
	}

	if arg2 == "180" || arg2 == "-180" {
		rotatedMatrix, _, _ = Transpose(matrix)
		rotatedMatrix, _, _ = ReverseRows(rotatedMatrix)
		rotatedMatrix, _, _ = Transpose(rotatedMatrix)
		rotatedMatrix, rowsRotated, colsRotated = ReverseRows(rotatedMatrix)
	}

	if arg2 == "270" || arg2 == "-90" {
		rotatedMatrix, _, _ = ReverseRows(matrix)
		rotatedMatrix, rowsRotated, colsRotated = Transpose(rotatedMatrix)
	}

	fileNameWrite := fmt.Sprintf("%sRotated.pbm", arg1)
	fwrite, err := os.Create(fileNameWrite)
	if err != nil {
		os.Exit(1)
	}

	defer fwrite.Close()

	var dataWrite string
	var dataWriteLine string
	dataWrite += (PBMFormat + "\n")
	dataWriteLine += fmt.Sprintf("%d %d", colsRotated, rowsRotated)
	dataWrite += (dataWriteLine + "\n")

	for _, v := range rotatedMatrix {
		dataWriteLine = strings.Join(v, " ")
		dataWrite += (dataWriteLine + "\n")
	}

	val := fmt.Sprintf("%v", dataWrite)
	toWrite := []byte(val)

	_, err2 := fwrite.Write(toWrite)

	if err2 != nil {
		os.Exit(1)
	}
}
