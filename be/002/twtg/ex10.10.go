package main

import "fmt"

type Base struct {
	id int
}

func (b *Base) Id() int {
	return b.id
}

func (b *Base) SetId(newid int) {
	b.id = newid
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	Salary float64
}

func main() {
	e := new(Employee)
	e.SetId(1)
	fmt.Println("Employee ID: ", e.Id())
}
