package main

import "fmt"

//func opProcess(a int, b int, op func(x int, y int)) {
//	op(a, b)
//}

func call() func(x, y int) int {
	return sum_
}

func sum_(a int, b int) int {
	//fmt.Println(a + b)
	return a + b
}

func main() {
	//opProcess(10, 20, sum_)
	sum := call()
	ans := sum(1, 2)
	fmt.Println(ans)
}
