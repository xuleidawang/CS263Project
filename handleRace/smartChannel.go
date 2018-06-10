package main

import (
	"fmt"
	
	"time"
)

func main() {
	// The code is blocked until something gets pushed into the returned channel
	// As opposed to the previous method, we block in the main function, instead
	// of the function itself
	start := time.Now()


	i := <-getNumberChan()
	fmt.Println(i)

	
	elapsed := time.Since(start)
    
	fmt.Println(elapsed)
}

// return an integer channel instead of an integer
func getNumberChan() <-chan int {
	// create the channel
	c := make(chan int)
	go func() {
		// push the result into the channel
		c <- 5
	}()
	// immediately return the channel
	return c
}