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

	degrees := rotateimage.Degrees(arg2)
	if !degrees.IsDegrees() {
		fmt.Printf("Please, type the right input to degres, examples:\n" +
			"90, -90, 180, -180, 270, -270 or reverse")
		return
	}

	rotatedMatrix, rowsRotated, colsRotated := matrix.RotateImage(matrix, degrees)

	err = utils.WriteOnFile(arg1, rotatedMatrix, rowsRotated, colsRotated)
	if err != nil {
		return
	}

	fmt.Println("File successfully generated !")
}
