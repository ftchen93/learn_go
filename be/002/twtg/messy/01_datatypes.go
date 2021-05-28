package main

import (
	"fmt"
	"strconv"
	// "strconv"
)

/** Basic Data Type **/

func main() {
	// declareVariables()
	wrapAround()
	// typeConversion()
}

func declareVariables() {
	// using 'var'
	fmt.Println("--- using var")
	var a int32
	fmt.Printf("Default value of int32 is: %d\n", a)
	a = 47
	fmt.Printf("We have set the value of variable 'a' to: %d\n", a)
	fmt.Println()

	// multiple declaration
	fmt.Println("--- multiple declaration")
	var b, c, d int32
	b = 57
	c = 67
	d = 77
	fmt.Printf("value of b is: %d\n", b)
	fmt.Printf("value of c is: %d\n", c)
	fmt.Printf("value of d is: %d\n", d)
	fmt.Println()

	j, k, t := 34, 44, 54
	fmt.Printf("value of j is: %d\n", j)
	fmt.Printf("value of k is: %d\n", k)
	fmt.Printf("value of t is: %d\n", t)
	fmt.Println()

	// ** we are not specifying the size of integer type. Is 'e' an 32bits or 64bits? **
	// ** read: "Sizes of Numeric Types" section of https://www.digitalocean.com/community/tutorials/understanding-data-types-in-go
	fmt.Println("--- using :=")
	e := 37
	fmt.Println(strconv.IntSize)
	fmt.Printf("Value of e is: %d\n", e)
	fmt.Println()
}

func wrapAround() {
	// wrap around
	fmt.Println("--- wrap around")
	var x uint32
	x = 4294967295
	fmt.Printf("Current value of x is: %d\n", x)
	x = x + 1
	fmt.Printf("New value of x is: %d\n", x)
	fmt.Println()

	// overflow
	// var z uint32
	// z = 4294967296
}

func typeConversion() {
	var j uint32
	var k uint64

	j = 5
	k = 10000000000000000000

	// badSum := j + k

	fmt.Println("--- convert uint32 to uint64")
	var sum64 uint64
	sum64 = uint64(j) + k
	fmt.Printf("Sum of j and k is: %d\n", sum64)
	fmt.Println()

	// ** what happens if we convert uint64 to uint32? **
	fmt.Println("--- convert uint64 to uint32")
	var sum32 uint32
	sum32 = j + uint32(k)
	fmt.Printf("Sum of j and k is: %d\n", sum32)
	fmt.Println()
}
