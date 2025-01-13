package main

import "fmt"

func main() {
	// make new
	// new - only allocate - no init of memory
	// make- - allocate and init - non zeored storage
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	//delete(m, "a")
	//fmt.Println(m)

	//var score map[string]int
	//score = map[string]int{"a": 1, "b": 2}
	//fmt.Println(score)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
