package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in file - My file"
	file, err := os.Create("./myTestGoFile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is :", length)
	file.Close()
	readFile("./myTestGoFile.txt")
}

func readFile(fileName string) {
	dataByte, err := os.ReadFile(fileName)

	checkNilError(err)

	fmt.Println("text data in the file is ", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
