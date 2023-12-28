package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("::::: Reading input from console :::: ")

	reader := bufio.NewReader(os.Stdin)

	readString, _ := reader.ReadString('\n') // here ignoring the error. This is a comma ok syntax

	fmt.Println("The input read from console is :::: ", readString)
	fmt.Printf("The input read from console type is %T", readString)
}
