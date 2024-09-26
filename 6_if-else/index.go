package main

import "fmt"

func main() {
	age := 20

	if age >= 18 {
		fmt.Println("is adult")
	}else if age >= 12 {
		fmt.Println("is younger")
	}else {
		fmt.Println("is kid")
	}

	if height := 185; height >= 175 {
		fmt.Println("Tall Person")
	}
	// Go does not have ternary operator till now
}

