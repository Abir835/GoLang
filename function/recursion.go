package main

import "fmt"

func countDown(num int) {
	fmt.Println(num)
	if num == 5 {
		return
	}
	countDown(num + 1)
}

func main() {
	countDown(1)
}
