package main

import "fmt"

func add(a int , b int) int {
	return a + b
}

func langs() (string , string , string , string ,int) {
	return "Java" , "C" , "Go" , "Python" , 1
}

func processIt(fn func(a int) int) int {
	return fn(1)
}

func complexFuncReturn() func(a int) int {
	return func(a int) int {
		return a
	}
}

func main(){
	r := add(1,2)
	fmt.Println(r) 
	fmt.Println(langs())

	fn := func(a int) int {
		return a
	}
	ans := processIt(fn)
	fmt.Println(ans)

	cf := complexFuncReturn()
	fmt.Println(cf(0))
	print_result := cf(3)
	fmt.Println(print_result)
}
