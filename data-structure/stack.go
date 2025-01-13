package main

import "fmt"

type Stack struct {
	content []int
}

func (s *Stack) Display() {
	fmt.Println(s.content)
}

func (s *Stack) Push(e int) {
	s.content = append(s.content, e)
}

func (s *Stack) Pop() {
	if len(s.content) == 0 {
		fmt.Println("stack is empty")
	}
	lastIndex := len(s.content) - 1
	s.content = s.content[:lastIndex]
}

func main() {
	s := new(Stack)
	s.Display()
	s.Push(1)
	s.Push(2)
	s.Display()
	s.Pop()
	s.Display()
}
