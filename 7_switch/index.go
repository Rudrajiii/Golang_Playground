package main

import (
	"fmt"
	"time"
)

func main() {
	//Normal type of switch
	i := 3
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	//Multiple condition wala switch
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend")
	default:
		fmt.Println("It's Workday")
	}

	//type switch
	whoAmI := func(i interface{}){
		switch t := i.(type){
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("boolean")
		default:
			fmt.Println("other" , t)
		}
	}
	whoAmI(true)
}

