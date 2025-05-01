package main

import "fmt"

/* Generic function */
func multiply[length int | float64, breath int | float64](l length, b breath) float64 {

	return float64(l) * float64(b)
}

/* We can also define the parameterized type which can be used in generic function. */

type MyUnit interface { // This is parameterized type
	int | float64
}

func multiplyWithParamType[myunit MyUnit](l myunit) myunit {

	return 2 * 3 * l
}

/* Generic function with any */
func returnAny[length any](l length) length {

	return l

}

/* Use of ~ to allow type definition which is wrapper over primitive type. */
type Int int

func myprint[T ~int](x T) {
	fmt.Println("Tilde example : ", x)
}

type mynumber interface {
	~int | float32 | float64 // here ~ sign I have used which will allow types over primitive type.
}

func myprintNumber[n mynumber](a n) {
	fmt.Println("Printing :: ", a)
}

func main() {

	// we can also call function as multiply[int](2, 3) mentioning the type,
	// however it gets inferred therefore usually not required to be passed.
	resultInt := multiply(2, 3)
	fmt.Println("resultInt: ", resultInt)

	resultFloat := multiply(1.5, 2.5)
	fmt.Println("resultFloat: ", resultFloat)

	resutMix := multiply(2, 2.2)
	fmt.Println("resutMix: ", resutMix)

	fmt.Println("Result multiplyWithParamType with int : ", multiplyWithParamType(2))
	fmt.Println("Result multiplyWithParamType with float : ", multiplyWithParamType(2.2))
	fmt.Println("Result multiplyWithParamType with int : ", multiplyWithParamType(2))

	intSliceReturn := returnAny([]int{1, 2, 3})
	fmt.Println("intSliceReturn: ", intSliceReturn)

	floatReturn := returnAny(3.4)
	fmt.Println("floatReturn: ", floatReturn)

	myprint[Int](Int(1)) // here Int(1) converts int 1 to Int(custom named type)

	// Check here, I am giving Int which is custom named type. This line will give error if I would not have used ~ in type mynumber interface{~int}
	myprintNumber[Int](1)
}
