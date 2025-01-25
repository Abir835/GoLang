package main

import "fmt"

/*
A struct is used to store variables of different data types. For example,

Suppose we want to store the name and age of a person. We can create two variables: name and age and store value.

However, suppose we want to store the same information of multiple people.

In this case, creating variables for a person might be a tedious task. We can create a struct that stores the name and age to overcome this.

And, we can use this same struct for every person.
*/

type Rec func(int, int) int

type Rectangle struct {
	Height int
	Width  int
	rect   Rec
}

func main() {
	rect := Rectangle{Height: 5, Width: 10, rect: func(Height, Weight int) int { return Height + Weight }}
	fmt.Println(rect)
}
