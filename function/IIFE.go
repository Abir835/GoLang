package main

import "fmt"

// Immediately invoked function expression

func main() {
	func(a, b int) {
		c := a + b
		fmt.Println(c)
	}(1, 2)
}

func init() {
	fmt.Println("<UNK>")
}
