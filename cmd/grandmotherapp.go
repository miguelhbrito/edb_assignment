package main

import (
	"fmt"
	"os"

	rotateimage "github.com/edb_test/pkg/rotate_image"
	"github.com/edb_test/pkg/utils"
)

func main() {

	var arg1 string
	var arg2 string

	if len(os.Args) != 3 {
		fmt.Println("Invalid paramenter," +
			" please type the file name and degres you desire to rotate, example:\n" +
			"grandmotherapp fileName 90")
		return
	} else {
		arg1 = os.Args[1]
		arg2 = os.Args[2]
	}

	fmt.Println("Reading file " + arg1 + ".pbm, to rotate image file with " + arg2 + " degres")

	matrix, err := utils.ReadFile(arg1)
	if err != nil {
		return
	}

	rotatedMatrix, rowsRotated, colsRotated := rotateimage.RotateImage(matrix, arg2)

	err = utils.WriteOnFile(arg1, rotatedMatrix, rowsRotated, colsRotated)
	if err != nil {
		return
	}

	fmt.Println("File successfully generated !")
}
