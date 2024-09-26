package main

import "fmt"
func main(){
	//in go we have only for loop no while and other loops
	//while loop using for
	// i := 1
	// for i <= 3{
	// 	fmt.Println(i)
	// 	i += 1
	// }
	//infinite loop
	// for {
	// 	fmt.Println(1)
	// }
	//classic for loop
	for i:= 0; i <= 10; i++ {
		if i == 0 {
			continue
		}
		fmt.Println(i)
	}

	for j := range []int{0, 1, 2} {
		fmt.Println(j)
	}
	// for j := 3 {
	// 	fmt.Println(j)
	// }
}

