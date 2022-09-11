package main

import "fmt"

type Node struct {
	data string
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "A"
	head.next = nil
	var nodeB *Node = &Node{data: "B", next: nil}
	head.next = nodeB
	var nodeC *Node = &Node{data: "C", next: nil}
	nodeB.next = nodeC
	tail.data = "D"
	tail.next = head
	nodeC.next = tail
}

func insert(position int, data string) {
	var p = head
	var i = 0
	// Move the node to the insertion position
	for {
		if p.next == nil || i >= position-1 {
			break
		}
		p = p.next
		i++
	}

	var newNode *Node = new(Node)
	newNode.data = data
	newNode.next = p.next // newNode next points to the next node
	p.next = newNode      // current next point to newNode

}

func remove(position int) {
	var p = head
	var i = 0
	for {
		if p.next == nil || i >= position-1 {
			break
		}
		p = p.next
		i++
	}
	var temp = p.next    // Save the node we want to delete
	p.next = p.next.next // Previous node next points to next of delete the node
	temp.next = nil

}

func output(node *Node) {
	var p = node
	for {
		fmt.Printf("%s -> ", p.data)
		p = p.next
		if p == head {
			break
		}
	}
	fmt.Printf("%s \n\n", p.data)
}

func main() {
	initial()
	fmt.Printf("Insert a new node E at index = 2 : \n")
	insert(2, "E")
	output(head)
	fmt.Printf("Delete a new node E at index = 2 : \n")
	remove(2)
	output(head)
}
