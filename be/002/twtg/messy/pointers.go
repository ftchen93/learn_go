package main

import "fmt"

// A pointer holds the memory address of a value

func main() {
	i, j := 42, 2701

	p := &i         // p point to i
	fmt.Println(*p) //read value of i through the pointer

	*p = 1         //set i through the pointer
	fmt.Println(i) //see new value of i

	p = &j         //p point to j
	*p = *p / 37   //divide j through the pointer ( j = j /37 )
	fmt.Println(j) //see new value of j
}
