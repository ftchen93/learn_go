package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	numbers := []int{1, -1, 5, 11, -9, 12, 4, 7, 3, -2}

	// === let first goroutine sum first 5 numbers
	go sum(numbers[0:5], c)
	// === let second goroutine sum last 5 numbers
	go sum(numbers[5:], c)

	partA := <-c
	fmt.Printf("Part A: %d\n", partA)
	partB := <-c
	fmt.Printf("Part B: %d\n", partB)

	answer := partA + partB
	fmt.Println(answer)
}

func sum(numbers []int, c chan int) {
	time.Sleep(10 * time.Second)
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	c <- sum
}
