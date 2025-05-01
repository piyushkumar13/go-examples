package main

import "fmt"

/*
	Go uses duck typing means if there is a struct which has same methods as declared in interface then that struct is the

implementation of interface. Basically, if struct can walk like a duck and quack like a duck, then it's a duck.
*/
type Shape interface {
	Area() int
	Perimeter() int
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) Area() int {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() int {
	return 2*r.width + 2*r.height
}

type Circle struct {
	radius int
}

func (c *Circle) Area() int {
	return c.radius * c.radius // note formula is not correct and thats not important here.
}

func (c *Circle) Perimeter() int {
	return 2*c.radius + 2*c.radius // note formula is not correct and thats not important here.
}

func main() {

	var shape Shape

	// if you dont use & then you will get error Type does not implement Shape as the Area method has a pointer receiver
	// Its because we have specified pointer receiver with methods means pointer receiver have methods which are implementation of
	// the interface. To fix it, either we use & with Rectangle as &Rectangle{} or change the pointer receiver to value receiver.
	// Which would mean value receiver have the methods of interface. BTW, with value receiver methods, both object instantiation works
	// i.e Rectangle{width: 10, height: 5}  as well as &Rectangle{width: 10, height: 5}.
	shape = &Rectangle{width: 10, height: 5}

	fmt.Println("The rectangle area is", shape.Area())
	fmt.Println("The rectangle perimeter is", shape.Perimeter())

	shape = &Circle{radius: 5}

	fmt.Println("The circle area is", shape.Area())
	fmt.Println("The circle perimeter is", shape.Perimeter())
}
