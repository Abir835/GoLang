package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {
	u := User{}
	u.Name = "Bob"
	u.Email = "bob@gmail.com"
	u.Age = 18
	fmt.Printf("%+v\n", u)

	var sam = new(User)
	sam.Name = "Sam"
	sam.Email = "sam@gmail.com"
	sam.Age = 18

	fmt.Printf("%+v\n", sam)

	var toddy = &User{"toddy", "toddy@gmail.com", 18}
	fmt.Printf("%+v\n", toddy)
}
