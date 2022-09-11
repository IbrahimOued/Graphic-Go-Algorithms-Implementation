package main

import (
	"fmt"
)

// ArrayStack defines the stack based on array
type ArrayStack struct {
	data     []string
	capacity int
	head     int //The stack top index in data
}

func (s *ArrayStack) isEmpty() bool {
	return s.head < 0
}

// NewArrayStack creates a empty stack
func NewArrayStack(capacity int) *ArrayStack {
	return &ArrayStack{
		data: make([]string, capacity), capacity: capacity, head: -1}
}

func (s *ArrayStack) push(element string) {
	if s.head < 0 {
		s.head = 0
	} else {
		s.head++
	}
	if s.head > len(s.data)-1 {
		s.data = append(s.data, element)
	} else {
		s.data[s.head] = element
	}
}

// Dequeue fetches a element from queue
func (s *ArrayStack) pop() string {
	if s.isEmpty() {
		return ""
	}
	topEle := s.data[s.head]
	s.head--
	return topEle
}

// Output prints the queue
func (s *ArrayStack) output() {
	if s.isEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := s.head; i >= 0; i-- {
			fmt.Println(s.data[i])
		}
	}
}

func main() {
	stack := NewArrayStack(10)
	stack.push("A")
	stack.push("B")
	stack.push("C")
	stack.push("D")

	stack.output()

	fmt.Printf("%d\n", len(stack.data))

	popedElement := stack.pop()
	fmt.Printf("Popped node : %s\n", popedElement)

	stack.output()
}
