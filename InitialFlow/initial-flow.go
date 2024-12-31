package main

import (
	_ "GoLang/AnotherInitFlow"
	"GoLang/NewInitialFlow"
	"fmt"
)

func init() {
	fmt.Println("Initial Flow starting...")
}

func main() {
	fmt.Println("Initial Flow starting... main")
	NewInitialFlow.NewInitialFlow()
}
