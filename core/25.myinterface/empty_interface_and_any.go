package main

import "fmt"

func main() {

	/* Empty interface */
	var x interface{} // interface{} is same as any. Any is short for interface{}

	x = 10
	x = 30.5

	x = "Hi"

	fmt.Println("Value of x is ::: ", x)

	/* any type */

	var y any //any was introduced in Go 1.18. It's just a type alias for interface{} i.e type any = interface{}

	y = 15
	y = 3.5

	y = "Hello"

	fmt.Println("Value of y is ::: ", y)
}
