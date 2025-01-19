package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter number: ")
	fmt.Scan(&number)
	fmt.Printf("Number is %d\n", number)

	var name string
	var age int

	fmt.Println("Enter name and age: ")
	fmt.Scanln(&name, &age)
	fmt.Printf("Name is %s, age is %d\n", name, age)

	// take name and age input using format specifier
	fmt.Println("Enter your name and age:")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Name is %s, age is %d\n", name, age)

}
