package main

import (
	"fmt"
)

// an array cannot be resize(size is fixed)
// a slice, is a dynamically-sized, flexible

func main() {
	// createSlice()
	appendSlice()
	// getPartOfSlice()
	// copySlice()
}

func createSlice() {
	fmt.Println("--- creating slices")
	// create with var
	var sliceA []string

	// create with :=
	sliceB := []string{"a", "b", "c"}

	// create with make()
	sliceC := make([]string, 5)

	fmt.Println("slice A:")
	printSlice(sliceA)
	fmt.Println(len(sliceA))
	fmt.Println(cap(sliceA))
	fmt.Println()

	fmt.Println("slice B:")
	printSlice(sliceB)
	fmt.Println(len(sliceB))
	fmt.Println(cap(sliceB))
	fmt.Println()

	fmt.Println("slice C:")
	printSlice(sliceC)
	fmt.Println(len(sliceC))
	fmt.Println(cap(sliceC))

	fmt.Println()
}

func appendSlice() {
	fmt.Println("--- appending to slices")
	sliceX := []string{}
	sliceX = append(sliceX, "apple", "banana", "orange", "mango", "durian", "pear")
	fmt.Println("after appending to slice X:")
	printSlice(sliceX)

	/** what happens if appending to a slice exceeds its capacity? **/
	sliceY := make([]string, 2)
	fmt.Println(len(sliceY))
	fmt.Println(cap(sliceY))
	sliceY = []string{"a", "b", "c"}
	//sliceY = append(sliceY, "a", "b", "c")
	fmt.Println("after appending to slice Y:")
	printSlice(sliceY)
	fmt.Println(len(sliceY))
	/** what if we want to append a slice to a slice? **/

	fmt.Println()
}

func getPartOfSlice() {
	fmt.Println("--- getting part of slice")
	mainSlice := []string{"apple", "banana", "orange", "mango", "durian", "pear"}
	fmt.Println("mainSlice:")
	printSlice(mainSlice)

	subSliceA := mainSlice[:3]
	fmt.Println("subSliceA:")
	printSlice(subSliceA)

	subSliceB := mainSlice[3:]
	fmt.Println("subSliceB:")
	printSlice(subSliceB)

	subSliceC := mainSlice[1:4]
	fmt.Println("subSliceC:")
	printSlice(subSliceC)

	fmt.Println()

	fmt.Println("--- changing value in subslice")
	subSliceC[0] = "jackfruit"
	fmt.Println("mainSlice after changing subSlice:")

	/** what are the values in mainSlice after modifying subSliceC ? **/
	//printSlice(mainSlice)

	fmt.Println()
}

func copySlice() {
	originalSlice := []string{"cookie", "bread", "noodle", "rice"}
	fmt.Println("originalSlice:")
	printSlice(originalSlice)
	fmt.Println()

	// copy a part original slice
	copySlice := make([]string, 3)
	copy(copySlice, originalSlice[1:])
	fmt.Println("copySlice:")
	printSlice(copySlice)
	fmt.Println()

	// change a value in copy slice
	copySlice[0] = "vegetable"
	fmt.Println("new values in copySlice:")
	printSlice(copySlice)
	fmt.Println()

	fmt.Println("originalSlice after changing copySlice:")
	printSlice(originalSlice)
}

func printSlice(slice []string) {
	for _, value := range slice {
		fmt.Print(value)
		fmt.Print(", ")
	}
	fmt.Println()
}
