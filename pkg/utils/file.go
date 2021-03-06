package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	rotateimage "github.com/edb_test/pkg/rotate_image"
)

const PBMFormat = "P1"

func ReadFile(fName string) (rotateimage.Image, error) {
	fileName := fmt.Sprintf("%s.pbm", fName)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\nPlease enter the right file name.", err)
		return nil, err
	}

	data := make(rotateimage.Image, 0)
	reader := bufio.NewReader(file)
	stringLine, err := readln(reader)
	for err == nil {
		if strings.Contains(string(stringLine[0]), "#") {
			fmt.Printf(stringLine)
		} else {
			splitData := strings.Split(stringLine, " ")
			data = append(data, splitData)
		}
		stringLine, err = readln(reader)
	}

	PBMFormatFile := data[0][0]
	if PBMFormat != PBMFormatFile {
		fmt.Printf("PBM file is different than P1 type. Please try again with right file.")
		return nil, err
	}

	matrix := data[2:]

	return matrix, nil
}

func WriteOnFile(fileName string, matrix rotateimage.Image, rows, cols int) error {
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

func readln(reader *bufio.Reader) (string, error) {
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
