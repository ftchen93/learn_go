package main

import "fmt"

type Celcius struct {
	degree float64
}

func main() {
	d := new(Celcius)
	d.degree = 35
	fmt.Println(d)
}

func (c *Celcius) String() string {
	d := fmt.Sprintf("%1f", c.degree)
	return d + " °C "
}
