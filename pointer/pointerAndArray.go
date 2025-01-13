package main

import "fmt"

func main() {
	//fmt.Println("This is Pointer and Array")
	//
	//var score int = 34
	//var p := &score
	//
	//if p != nill {
	//	fmt.Println("P is having a value: ", *p)
	//} else {
	//	fmt.Println("P is having a nil values")
	//}

	//var lifeBooster float64 = 99.2
	//lifeBoosterRef := &lifeBooster
	//// automatic ref value change
	//lifeBooster = lifeBooster * 2.2
	//
	//fmt.Println(lifeBooster)
	//fmt.Println(*lifeBoosterRef)

	// Array

	var numbers [3]string
	numbers[0] = "Hello"
	numbers[1] = "World"
	numbers[2] = "Golang"
	fmt.Println(numbers)

	var colors = [4]string{"red", "yellow", "green", "blue"}
	fmt.Println(colors)
	fmt.Println(len(colors))
	fmt.Println(colors[1])
}
