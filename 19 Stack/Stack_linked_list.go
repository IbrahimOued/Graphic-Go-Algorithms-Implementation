package main

import "fmt"

// ListNode defines a node of linked-list
type Node struct {
	data string
	next *Node
}

// LinkedListStack defines a stack based on linked-list
type LinkedListStack struct {
	head *Node
	size int
}

func (s *LinkedListStack) isEmpty() bool {
	return s.head == nil
}

func (s *LinkedListStack) push(element string) {
	s.head = &Node{data: element, next: s.head}
	s.size++
}

func (s *LinkedListStack) pop() *Node {
	if s.isEmpty() {
		return nil
	}
	data := s.head
	s.head = s.head.next
	return data
}

// String returns object string
func (s *LinkedListStack) output() string {
	if s.head == nil {
		return "empty stack"
	}
	result := "head\n"
	for cur := s.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("%+v\n", cur.data)
	}
	result += "tail\n"
	return result
}

// Get the front
func (s *LinkedListStack) getFront() *Node {
	if s.isEmpty() {
		return nil
	}
	return s.head
}

func main() {
	stack := LinkedListStack(LinkedListStack{})
	stack.push("A")
	stack.push("B")
	stack.push("C")
	stack.push("D")

	fmt.Printf("%s\n", stack.output())

	fmt.Printf("%d\n", stack.size)

	popedNode := stack.pop()
	fmt.Printf("Popped node : %s\n", popedNode.data)

	fmt.Printf("%s\n", stack.output())
}
