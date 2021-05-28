package main

import (
	"fmt"
)

func main() {
	createArray()
	// createArrayLiteral()
}

func createArray() {
	// create array
	var studentIds [20]uint32
	fmt.Println()

	// check length of array
	fmt.Println("--- get length of array")
	fmt.Printf("studentIds array length: %d\n", len(studentIds))
	fmt.Println()

	// set values at index
	fmt.Println("-- set values of items in array")
	studentIds[0] = 1000
	studentIds[19] = 1019
	printArray(studentIds)
	fmt.Println()

	// get values at index
	fmt.Println("--- get values of items in array")
	lastStudentId := studentIds[19]
	fmt.Printf("Value at index 19: %d\n", lastStudentId)
	fmt.Println()

	// ** what happens if we uncomment the following line?
	// index out of bounds
	// studentIds[99] = 1099
}

func createArrayLiteral() {
	studentIds := [20]uint32{88, 91, 12, 12, 12}
	printArray(studentIds)
}

func printArray(array [20]uint32) {
	for _, value := range array {
		fmt.Print(value)
		fmt.Print(", ")
	}
	fmt.Println()
}
