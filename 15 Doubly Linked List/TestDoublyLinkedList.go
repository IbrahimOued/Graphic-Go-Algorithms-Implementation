package main

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}

/*
The new built-in function allocates memory.
The first argument is a type, not a value,
and the value returned is a pointer to a newly
allocated zero value of that type.
*/
var head *Node = new(Node) // The first node called head node
var tail *Node = new(Node)

func initial() {
	head.data = "San Franciso"
	head.prev = nil
	head.next = nil

	var nodeOakland *Node = &Node{data: "Oakland", prev: head, next: nil}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley", prev: nodeOakland, next: nil}
	nodeOakland.next = nodeBerkeley

	tail.data = "Fremont"
	tail.prev = nodeBerkeley
	tail.next = nil
	nodeBerkeley.next = tail
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
	var newNode *Node = &Node{data: data, next: p.next}
	newNode.data = data
	newNode.next = p.next
	p.next = newNode
	newNode.prev = p
	newNode.next.prev = newNode

}

func removeNode(position int) {
	var p = head
	var i = 0
	// Move the node to the previous node position that was deleted
	for {
		if p.next == nil || i >= position-1 {
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

func output(node *Node) {
	var p = node
	var end *Node = nil
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.data)
		end = p
		p = p.next
	}
	fmt.Printf("End\n")
	p = end
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.data)
		p = p.prev
	}
	fmt.Printf("Start\n\n")
}

func main() {
	initial()
	output(head)
	// fmt.Printf("Insert a new node Walnut at index 2 : \n")
	// insert(2, "Walnut")
	// output(head)
	fmt.Printf("Delete a node Walnut at index 2 : \n")
	removeNode(2)
	output(head)
}
