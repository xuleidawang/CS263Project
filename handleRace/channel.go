package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(getNumber())
	elapsed := time.Since(start)
    
	fmt.Println(elapsed)
}

func getNumber() int {
	var i int
	// Create a channel to push an empty struct to once we're done
	done := make(chan struct{})
	go func() {
		i = 5
		// Push an empty struct once we're done
		done <- struct{}{}
	}()
	// This statement blocks until something gets pushed into the `done` channel
	<-done
	return i
}