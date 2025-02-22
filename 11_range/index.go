package main

import "fmt"

func main(){

	arr := []int{4,8,9}
	for i,num := range arr {
		fmt.Println(num , i)
	}
	m := map[string]int{"A":10,"B":20,"C":30}
	for k , v := range m {
		fmt.Println(k , v)
	}

	for i , c := range "python" {
		fmt.Println(c , i , string(c))
	}
}

