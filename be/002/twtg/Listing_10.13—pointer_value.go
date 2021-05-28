package main

import (
	"fmt"
	"math"
)

/*
change() receives a pointer to B, and changes its internal field;
write() only outputs the contents of the B variable and receives its value by copy.

Notice in main() that Go does plumbing work for us,
we ourselves do not have to figure out whether to call the methods on a pointer or not,
Go does that for us.

b1 is a value and b2 is a pointer,
but the methods calls work just fine.
*/

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

type Point3 struct {
	x, y, z float64
}

func (p Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func main() {
	var b1 B //b1 is a value
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) //b2 is a pointer
	b2.change()
	fmt.Println(b2.write())

	var p1 Point3 //p1 is a value
	p1 = Point3{
		x: 2, //4
		y: 4, //16
		z: 8, //64
	}
	fmt.Println(p1.Abs())
}
