package main

import "fmt"

/*
type A = string creates an alias for string. Whenever you use A in your code,
it works just like string. So for example, you can't define methods on it.

type A string defines a new type, which has the same representation as string.
One of the use case of defining type definition on primitive type is when using context to store values.
*/

type mystringalias = string // this is type alias.
type mystringdef string     // this is type definition, you can also define method for this type.

func main() {

	var strAlias mystringalias = "hello"
	fmt.Printf("type of strAlias : %T, value : %v", strAlias, strAlias) // type will be printed as string

	fmt.Println()
	var strDef mystringdef = "hello"
	fmt.Printf("type of strDef : %T, value : %v", strDef, strDef) // type will be printed as main.mystringdef

}
