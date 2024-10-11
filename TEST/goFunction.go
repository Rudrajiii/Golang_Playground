package main

import "C"
import (
	"fmt"
)

//export GoFunction
func GoFunction() {
	sum := 0
	for i := 0; i < 1000000000; i++ {
		sum += i
	}

	fmt.Println("Go sum:", sum) // Just to ensure the function isn't optimized out
}

func main() {}
