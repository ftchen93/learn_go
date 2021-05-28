package main

import "fmt"

type Employee struct {
	salary float64
}

func (e Employee) giveRaise(percentage float64) (p float64) {
	return e.salary * (1 + percentage/100)
}

func main() {
	Bob := Employee{
		salary: 2000,
	}
	s := Bob.salary
	r := Bob.giveRaise(100)
	fmt.Printf("Bob salary now: %f\n", s)
	fmt.Printf("Bob salary after increase 100 percent : %f\n", r)
}
