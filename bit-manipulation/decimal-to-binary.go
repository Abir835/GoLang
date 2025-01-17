package main

import "fmt"

func main() {
	n := 13
	result := ""

	for n > 0 {
		if n%2 == 1 {
			result = "1" + result
		} else {
			result = "0" + result
		}
		n = n / 2
	}

	fmt.Println(result)
}
