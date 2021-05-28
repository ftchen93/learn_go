package main

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
	wheelCount int
}

func (c *Car) Start() {
	fmt.Println("car start")
}

func (c *Car) Stop() {
	fmt.Println("car stop")
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	//drive to work
	c.Stop()
	// get out car
}

func (c *Car) numberOfWheels() int {
	return c.wheelCount
}

type Mercedes struct {
	Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi Merkel")

}

func main() {
	c := &Mercedes{Car{nil, 4}}

	fmt.Println("I have wheel of ", c.Car.numberOfWheels())
	c.GoToWorkIn()
	c.sayHiToMerkel()
}
