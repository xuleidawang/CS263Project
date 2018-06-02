package main

import (
	"fmt"
	"math/rand"
	"matrix"
)



func main() {
	PrintMemUsage()
	
	var a,b MatrixRO
	Println(a.ParallelProduct(b))


	fmt.Println(Product)

	runtime.GC()
    PrintMemUsage()
}

