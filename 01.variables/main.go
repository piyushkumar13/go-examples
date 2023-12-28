package main

import "fmt"

/* Global variables and constants */

var myGlobalPrivateStr string = "My global private string" // as first letter is small so it is non-exportable, non-visible, private variable.
var MyGlobalPublicStr string = "My global public string"   // as first letter is capital so it is exportable, visible, public variable.

var myGlobalIntWithImplicitType = 123 // private
var MyGlobalIntWithImplicitType = 123 // public

const myGlobalPrivateConst string = "My global private const"
const MyGlobalPublicConst string = "My global public const"

const myGlobalPrivateConstWithImplicitType string = "My global private const with implicit type"
const MyGlobalPublicConstWithImplicitType = "My global private const with implicit type"

//myStrGlobalVar := "Walrus operator is not allowed here" // walrus operator is only allowed within a method. Not in global scope.

func main() {

	/* Variables in go */

	fmt.Println(":::::::::::::::::::::::::::::::::: Variables in go :::::::::::::::::::::::::::::::::::::::::::: ")
	fmt.Println("::::::::::: Variables in go with explicit type declaration :::::::::: ")

	var myStrVar string = "Hello to variable class"
	fmt.Println(myStrVar)
	fmt.Printf("The myStrVar type is %T \n", myStrVar)

	var myUnsignedIntVar uint8 = 255
	fmt.Println(myUnsignedIntVar)
	fmt.Printf("The myUnsignedIntVar type is %T \n", myUnsignedIntVar)

	var myFloatVar float32 = 255.1234566778655
	fmt.Println(myFloatVar)
	fmt.Printf("The myFloatVar type is %T \n", myFloatVar)

	var myBigFloatVar float64 = 255.1234566778655
	fmt.Println(myBigFloatVar)
	fmt.Printf("The myBigFloatVar type is %T \n", myBigFloatVar)

	/* Default values */

	fmt.Println("::::::::::: Variables in go with default values :::::::::: ")

	var strDefault string
	fmt.Println("The default value for string is :::", strDefault) // empty default value for string
	fmt.Printf("The strDefault type is %T \n", strDefault)

	var intDefault int8
	fmt.Println("The default value for int8 is :::", intDefault) // 0 default value
	fmt.Printf("The intDefault type is %T \n", intDefault)

	var floatDefault float32
	fmt.Println("The default value for float is :::", floatDefault) // 0 default value
	fmt.Printf("The floatDefault type is %T \n", floatDefault)

	/* Implicit types */
	// We can define variables without the type, lexer will interpret the type of the variable.

	fmt.Println("::::::::::: Variables in go with implicit types :::::::::: ")

	var myImplicitStrVar = "Hello to variable class"
	fmt.Println(myImplicitStrVar)
	fmt.Printf("The myImplicitStrVar type is %T \n", myImplicitStrVar)
	//myImplicitStrVar = 23 // Cannot do this once myImplicitStrVar type is inferred by lexer

	var myImplicitUnsignedIntVar = 255
	fmt.Println(myImplicitUnsignedIntVar)
	fmt.Printf("The myImplicitUnsignedIntVar type is %T \n", myImplicitUnsignedIntVar)

	var myImplicitFloatVar = 255.1234566778655
	fmt.Println(myImplicitFloatVar)
	fmt.Printf("The myImplicitFloatVar type is %T \n", myImplicitFloatVar)

	var myImplicitBigFloatVar = 255.1234566778655
	fmt.Println(myImplicitBigFloatVar)
	fmt.Printf("The myImplicitBigFloatVar type is %T \n", myImplicitBigFloatVar)

	/* Int type alias */

	fmt.Println("::::::::::: Variables in go with int alias :::::::::: ")

	var myIntVarWithAliasType int = 2555555
	fmt.Println(myIntVarWithAliasType)
	fmt.Printf("The myIntVarWithAliasType type is %T \n", myIntVarWithAliasType)

	/* Walrus operator */

	fmt.Println("::::::::::: Walrus operator in go :::::::::: ")

	myIntVarWithoutType := 2555555 // lexer will automatically infer the type
	fmt.Println(myIntVarWithoutType)
	fmt.Printf("The myIntVarWithoutType type is %T \n", myIntVarWithoutType)

	myStrVarWithoutType := "example of walrus operator" // lexer will automatically infer the type
	fmt.Println(myStrVarWithoutType)
	fmt.Printf("The myStrVarWithoutType type is %T \n", myStrVarWithoutType)

	/* Constants in go */

	fmt.Println("::::::::::: Constants in go :::::::::: ")

	const myConstVar int = 123
	fmt.Println(myConstVar)
	fmt.Printf("The myConstVar type is %T \n", myConstVar)
	//myConstVar = 56 // Cannot do this as myConstVar is constant

	const myConstVarWithImplicitType = "Something"
	fmt.Println(myConstVarWithImplicitType)
	fmt.Printf("The myConstVarWithImplicitType type is %T \n", myConstVarWithImplicitType)
	//myConstVarWithImplicitType = "something else" // Cannot do this as myConstVarWithImplicitType is constant

	/* Print global vars */

	fmt.Println("::::::::::: Constants in go :::::::::: ")

	fmt.Println(myGlobalPrivateStr)
	fmt.Printf("The myGlobalPrivateStr type is %T \n", myGlobalPrivateStr)

	fmt.Println(MyGlobalPublicStr)
	fmt.Printf("The MyGlobalPublicStr type is %T \n", MyGlobalPublicStr)

	fmt.Println(myGlobalIntWithImplicitType)
	fmt.Printf("The myGlobalIntWithImplicitType type is %T \n", myGlobalIntWithImplicitType)

	fmt.Println(MyGlobalIntWithImplicitType)
	fmt.Printf("The MyGlobalIntWithImplicitType type is %T \n", MyGlobalIntWithImplicitType)

	fmt.Println(myGlobalPrivateConst)
	fmt.Printf("The myGlobalPrivateConst type is %T \n", myGlobalPrivateConst)

	fmt.Println(MyGlobalPublicConst)
	fmt.Printf("The MyGlobalPublicConst type is %T \n", MyGlobalPublicConst)

	fmt.Println(myGlobalPrivateConstWithImplicitType)
	fmt.Printf("The myGlobalPrivateConstWithImplicitType type is %T \n", myGlobalPrivateConstWithImplicitType)

	fmt.Println(MyGlobalPublicConstWithImplicitType)
	fmt.Printf("The MyGlobalPublicConstWithImplicitType type is %T \n", MyGlobalPublicConstWithImplicitType)
}
