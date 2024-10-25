package main

import "fmt"

func main(){
	//* -> fixed size , that is predictable
	//* -> memory optimization
	//* -> constant time access
	var array [5]int
	fmt.Println(len(array))
	fmt.Println(array)

	var float32_array [5]float32
	fmt.Println(float32_array)
	//* [0 0 0 0]

	var float64_array [5]float64
	fmt.Println(float64_array)
	//* [0 0 0 0]

	num := [3]int{1,2,3}
	fmt.Println(num)

	num_2D := [3][3]float64{{7.9,9.1,5.5},{1.2,9.3,7.8} ,{8.3,6.5,4.9}}
	fmt.Println(num_2D)
}
