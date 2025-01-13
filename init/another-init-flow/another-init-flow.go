package another_init_flow

import "fmt"

func init() {
	fmt.Println("Another init flow")
}

func NewAnotherInitFlow() {
	fmt.Println("New another init flow")
}
