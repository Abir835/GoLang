package main

import "fmt"

// array is a collection of similar types of data like exm: [1,2,3]

func main() {
	// creating array
	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	// creating undefined array
	var undefinedArray = [...]int{1}
	fmt.Println(undefinedArray)

	// shorthand notation use for creating array
	shorthandArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(shorthandArray)

	// accessing array
	fmt.Println(len(shorthandArray))
	fmt.Println(shorthandArray[0])

}
