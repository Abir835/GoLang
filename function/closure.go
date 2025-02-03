package main

import "fmt"

/*
Go closure is a nested function that allows us to access variables of the outer function even after the outer function is closed.

Before we learn about closure, let's first revise the following concepts:

Nested Functions
Returning a function
Closure helps in Data Isolation

*/

func greet(name string) {
	var displayName = func() {
		fmt.Printf("Hello, %s!\n", name)
	}
	displayName()
}

func greet1() func() {
	return func() {
		fmt.Println("Hello, world!")
	}
}

func calculate() func() int {
	num := 1
	return func() int {
		num += num
		return num
	}
}

func main() {
	greet("John Doe")
	obj := greet1()
	obj()
	//IMMF
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2)

	//
	obj1 := calculate()
	fmt.Println(obj1())
	fmt.Println(obj1())
	fmt.Println(obj1())
	fmt.Println(obj1())

	obj2 := calculate()
	fmt.Println(obj2())
}
