package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("::::: Reading input from console :::: ")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n') // here ignoring the error. This is a comma ok syntax

	fmt.Println("The input read from console is :::: ", input)
	fmt.Printf("The input read from console type is %T \n", input)

	i, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		panic(err)
	}

	numRating := i + 1

	fmt.Println("The rating is ::: ", numRating)

}
