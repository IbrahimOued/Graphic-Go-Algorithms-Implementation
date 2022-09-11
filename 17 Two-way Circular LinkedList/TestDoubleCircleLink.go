package main

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "A"
	head.prev = nil
	head.next = nil

	var nodeB *Node = &Node{data: "B", prev: head, next: nil}
	head.next = nodeB

	var nodeC *Node = &Node{data: "C", prev: nodeB, next: nil}
	nodeB.next = nodeC

	tail.data = "D"
	tail.prev = nodeC
	tail.next = head
	nodeC.next = tail
	head.prev = tail
}

func insert(insertPosition int, data string) {
	var p = head
	var i = 0
	for {
		if p.next == nil || i >= insertPosition-1 {
			break
		}
		p = p.next
		i++
	}
	var newNode *Node = new(Node)
	newNode.data = data
	newNode.next = p.next // newNode next point to next node
	p.next = newNode      // current next point to newNode
	newNode.prev = p
	newNode.next.prev = newNode
}

func remove(removePosition int) {
	var p = head
	var i = 0
	for {
		if p.next == nil || i >= removePosition-1 {
			break
		}
		p = p.next
		i++
	}
	var temp = p.next    // Save the node you want to delete
	p.next = p.next.next // Previous node next points to next of delete the node
	p.next.prev = p
	temp.next = nil // Set the delete node next to null
	temp.prev = nil // Set the delete node prev to null
}

func output() {
	var p = head
	for {
		fmt.Printf("%s -> ", p.data)
		p = p.next
		if p == head {
			break
		}
	}
	fmt.Printf("%s ", p.data)
	fmt.Printf("End\n")
	p = tail
	for {
		fmt.Printf("%s -> ", p.data)
		p = p.prev
		if p == tail {
			break
		}
	}
	fmt.Printf("%s ", p.data)
	fmt.Printf("Start\n\n")
}
func main() {
	initial()
	fmt.Printf("Insert a new node E at index 2 : \n")
	insert(2, "E")
	output()
	fmt.Printf("Delete a new node C at index = 2 : \n")
	remove(2)
	output()
}
