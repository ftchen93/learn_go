package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//goroutineTimeSleep()
	goroutineWaitGroup()
}

func goroutineTimeSleep() {
	for i := 0; i < 10; i++ {
		fmt.Printf("call: %d\n", i)
		go printNumber(i)
	}

	// *** what happens if we comment out the following line?
	// go routine printNumber wnot able to print out
	time.Sleep(3 * time.Second)
}

func goroutineWaitGroup() {
	waitgroup := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		fmt.Printf("call: %d\n", i)
		waitgroup.Add(1)
		go printNumberWithWG(i, &waitgroup)
	}
	waitgroup.Wait()
}

func printNumber(call int) {
	fmt.Printf("execute: %d\n", call)
}

func printNumberWithWG(call int, wg *sync.WaitGroup) {
	fmt.Printf("execute: %d\n", call)
	wg.Done()
}
