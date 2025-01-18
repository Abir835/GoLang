package main

import "fmt"

func main() {
	n := 40
	res := n & (n - 1)
	fmt.Println(res)
}
