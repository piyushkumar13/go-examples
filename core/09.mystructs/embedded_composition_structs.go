package main

import "fmt"

type Vehicle struct {
	name  string
	color string
	model string
}

func NewVehicle(name string, color string, model string) *Vehicle {
	return &Vehicle{name: name, color: color, model: model}
}

func (c *Vehicle) getName() string {
	return c.name
}

func (c *Vehicle) getModel() string {
	return c.model
}

func (c *Vehicle) getColor() string {
	return c.color
}

type Car struct {

	/* anonymous embedded field, allows us to use composition.
	   Methods of Vehicle will be available in Car and can be access using car object directly.*/
	Vehicle
	tyres int
	gates int
}

type Truck struct {

	/* Explicit field, also allows us to use composition.
	   Methods of Vehicle will not be available in Car automatically. To access methods of Vehicle, we need to use truck.V.getName() etc. */
	V     Vehicle
	tyres int
	gates int
}

func main() {

	car := &Car{
		Vehicle: Vehicle{name: "Punch", color: "Grey", model: "2024"},
		tyres:   4,
		gates:   4,
	}

	fmt.Println("Car is ::: ", *car)
	fmt.Println("Car name is ::: ", car.getName()) // methods of Vehicle struct are automatically available in Car struct and can be access using car object.
	fmt.Println("Car model is ::: ", car.getModel())
	fmt.Println("Car color is ::: ", car.getColor())

	truck := &Truck{
		V:     Vehicle{name: "Volvo Truck", color: "White", model: "2024"},
		tyres: 8,
		gates: 2,
	}

	fmt.Println("Truck is ::: ", *truck)

	fmt.Println("Truck name is ::: ", truck.V.getName()) // methods of Vehicle struct are not available in truck struct directly, we need to use V field on truck object to access Vehicle methods.
	fmt.Println("Truck model is ::: ", truck.V.getModel())
	fmt.Println("Truck color is ::: ", truck.V.getColor())

}
