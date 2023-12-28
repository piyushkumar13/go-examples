package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println(":::: Files with defer example ::::")

	content := "This is the example of files"

	file, err := os.Create("./samplefile.txt")

	if err != nil {
		panic("Error occurred while creating file")
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic("Error occurred while writing string to file")
	}
	fmt.Println("The written file length is ::: ", length)

	defer file.Close()

	readFile("./samplefile.txt")
}

func readFile(fileName string) {

	dataBytes, err := os.ReadFile(fileName)

	if err != nil {
		panic("Error occurred while reading file")
	}

	fmt.Println("The read string is ::: ", string(dataBytes))
}
