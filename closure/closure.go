package main

import "fmt"

const p = 10

var a = 100

func outer() func() {
	money := 100
	age := 30

	fmt.Println("Age: ", age)

	show := func() {
		money = p + a + money
		fmt.Println("Money: ", money)
	}
	return show
}

func call() {
	out1 := outer()
	out1()
	out1()

	//out2 := outer()
	//out2()
}

func main() {
	call()
}
