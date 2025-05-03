package main

import (
	"fmt"
	"runtime"
)

/*
The init function is a special function that is automatically executed when a package is initialized—before
the main() function runs and before any other code in the package is executed.

* You never call init() manually. Go calls it for you during program startup.

* You can have multiple init() functions in a single package, even in different files within that package.
If multiple exist, Go will execute them in the order they appear per file (lexical file name order).

* You can have multiple init() functions in the same Go file.
They will be executed in the order they appear in the file (top to bottom), before the main() function runs

*/

func init() {

	fmt.Println("First init block")
}

func init() {

	fmt.Println("Second init block")
}

func main() {

	fmt.Println("Hello World to Gophers !!!")

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.NumCPU())
}
