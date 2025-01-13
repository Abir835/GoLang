package main

import "fmt"

type User1 struct {
	Name string
	Age  int
}

func main() {
	var userName *User1

	userName = new(User1)
	userName.Age = 18
	userName.Name = "aaaaa"

	fmt.Println(userName)

	//user := User{"Abir Hasan", 21}
	//user.Age = 10
	//user.Name = "Momin Bhai"
	//
	//p := &user
	//fmt.Println(p)
	//
	//fmt.Println(user)

}
