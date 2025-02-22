package main

import (
	"fmt"
	"maps"
)

// map -> hash , dict , object
func main(){
	m := make(map[string]string)
	m["user_name"] = "Rudra"
	m["email"] = "rudra@gmail.com"
	fmt.Println(m , m["user_name"])

	fmt.Println(m["h"] == "") // if key does not exists then it return empty string..

	new_map := make(map[string]int)
	new_map["age"] = 22
	fmt.Println(new_map , new_map["age"])
	delete(m , "user_name")
	fmt.Println(m)

	other_map := map[string]int{"model-no":20393920,"port":5055}
	fmt.Println(other_map)

	_, is_ok := other_map["port"]
	if is_ok {
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}

	//maps.Equal(map1 , map2) func
	map1 := map[string]int{"A":10,"B":20}
	map2 := map[string]int{"A":10,"B":2}
	// fmt.Println(maps.Equal(map1 , map2))
	fmt.Println(maps.Equal(map1 , map2))

}
