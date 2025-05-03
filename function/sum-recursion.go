package main

import "fmt"

func sum(number int) int {
	if number == 0 {
		return 0
	} else {
		return number + sum(number-1)
	}
}

func main() {
	sum_ := sum(10)
	fmt.Println(sum_)
}
