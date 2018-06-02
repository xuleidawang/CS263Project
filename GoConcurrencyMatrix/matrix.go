package main

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
)
var size int =6
// A Matrix is a square 6*6 of type int.
type Matrix [1500][1500]int

func (m *Matrix) String() string {
	result := ""

	for j := 0; j < 1500; j++ {
		result = result + fmt.Sprintf("%v\n", m[j])
	}
	return result
}

// Random generates a random matrix.
func Random() *Matrix {
	content := [1500][1500]int{}
	for i := 0; i < 1500*1500; i++ {
		rand.Seed(time.Now().UnixNano())
		content[i/1500][i%1500] = rand.Int()
	}
	result := Matrix(content)
	return &result
}


// Product is the result of matrix multiplications.
var Product = Random()

// func actor(transforms <-chan *Matrix, rowNum int) {
// 	for multiplier := range transforms {
// 		row := [6]int{}

// 		for j := 0; j < 6; j++ {
// 			running := 0
// 			for k := 0; k < 6; k++ {
// 				running = running + Product[rowNum][k]*multiplier[k][j]
// 			}
// 			row[j] = running
// 		}
// 		Product[rowNum] = row
// 	}
// }

// func runActors(transforms <-chan *Matrix) {
// 	inputs := [6]chan *Matrix{}
// 	for i := 0; i < 6; i++ {
// 		inputs[i] = make(chan *Matrix)
// 		go actor(inputs[i], i)
// 		defer close(inputs[i])
// 	}

// 	for x := range transforms {
// 		for _, in := range inputs {
// 			in <- x
// 		}
// 	}
// }

func main() {
	PrintMemUsage()

	// transforms := make(chan *Matrix)
	// go runActors(transforms)
	// for i := 0; i < 1000; i++ {
	// 	transforms <- Random()
	// }
	// close(transforms)
	// fmt.Println(Product)

	// runtime.GC()
	start := time.Now()

    
    var A = Random()
	var B = Random()
	var C = Random()

	in := make(chan int)
	quit := make(chan bool)

	dotRowCol := func() {
		for {
			select {
			case i := <-in:
				sums := make([]int , 1500)
				for k := 0; k < 1500; k++ {
					for j := 0; j < 1500; j++ {
						sums[j] += A[i][k] * B[k][j]
					}
				}
				for j := 0; j < 1500; j++ {
					C[i][j]=sums[j]
				}
			case <-quit:
				return
			}
		}
	}

	threads := 36

	for i := 0; i < threads; i++ {
		go dotRowCol()
	}

	for i := 0; i < 1500; i++ {
		in <- i
	}

	for i := 0; i < threads; i++ {
		quit <- true
	}

	
	elapsed := time.Since(start)
    
	
	//fmt.Println(C)
	PrintMemUsage()
	runtime.GC()
    PrintMemUsage()
    fmt.Println(elapsed)
}

func PrintMemUsage() {
        var m runtime.MemStats
        runtime.ReadMemStats(&m)
        // For info on each, see: https://golang.org/pkg/runtime/#MemStats
        fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
        fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
        fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
        fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}