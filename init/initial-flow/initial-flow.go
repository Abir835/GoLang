package main

import (
	_ "GoLang/another-init-flow"
	"GoLang/init/new-initial-flow"
	"fmt"
)

func init() {
	fmt.Println("Initial Flow starting...")
}

func main() {
	fmt.Println("Initial Flow starting... main")
	new_initial_flow.NewInitialFlow()
}
