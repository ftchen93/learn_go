package main

import "fmt"

type Days uint

const (
	MON = iota
	TUE
	WED
	THU
	FRI
	SAT
	SUN
)

var daynames = [7]string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

func main() {
	var today Days = SUN
	fmt.Printf("Today is: %s.\n", today)
}

func (d Days) String() string {
	return daynames[d]
}
