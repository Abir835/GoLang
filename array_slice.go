package main

import "fmt"

func main() {

	var arr [][]int
	arr = append(arr, []int{1})
	arr = append(arr, []int{2})
	arr = append(arr, []int{3})
	arr = append(arr, []int{4})
	arr = append(arr, []int{5})
	//arr = append(arr, 2)
	//arr = append(arr, 3)
	//arr = append(arr, 4)
	//arr = append(arr, 5)
	fmt.Println(arr[1][0])

}
