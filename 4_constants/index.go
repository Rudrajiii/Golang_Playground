package main

import "fmt"

const age int = 20
var name string = "Rudra"

func main() {
	const name string = "golang"
	fmt.Println(name) //golang
	fmt.Println(age) // 20

	const (
		PORT = 5000
		HOST = "\nlocalhost://"
	)
	fmt.Println(PORT , HOST)
	var (
		username="vicky"
		email="\nvicky@gmail.com"
	)
	fmt.Println(username , email)
}

