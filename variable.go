package main

import "fmt"

func main() {
	var score int = 4
	thor := "I am Abir"
	var defaultValue int

	var (
		age  = 10
		name = "Abir"
	)

	fmt.Println("hello world")
	fmt.Println(score)
	fmt.Println(thor)
	fmt.Printf("%v, %T", thor, thor)
	fmt.Println(defaultValue)
	fmt.Println(age, " ", name)
}
