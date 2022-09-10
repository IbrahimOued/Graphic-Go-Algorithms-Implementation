package main

import "fmt"

type Node struct {
	data string
	next *Node
}

/*
The new built-in function allocates memory.
The first argument is a type, not a value,
and the value returned is a pointer to a newly
allocated zero value of that type.
*/
var head *Node = new(Node) // The first node called head node

func initial() {
	head.data = "San Franciso"
	head.next = nil

	var nodeOakland *Node = &Node{data: "Oakland", next: nil}
	head.next = nodeOakland

	var nodeBerkeley *Node = &Node{data: "Berkeley", next: nil}
	nodeOakland.next = nodeBerkeley

	var tail *Node = &Node{data: "Fremont", next: nil}
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
	p.next = newNode
	/*
		or
		var newNode *Node = new(Node)
		newNode.next = p.next	// newNode next point to next node
		p.next = newNode		// CUrrent next point to newNode
	*/
}

func removeNode(position int) {
	var p = head
	var i = 0
	// Move the node to the previous position that was deleted
	for {
		if p.next == nil || i >= position-1 {
			break
		}
		p = p.next
		i++
	}
	var temp = p.next    // Save thee node you want to delete
	p.next = p.next.next // Previous node next points to next of delete the node
	temp.next = nil
}

func output(node *Node) {
	var p = node
	for {
		if p == nil {
			break
		}
		fmt.Printf("%s -> ", p.data)
		p = p.next
	}
	fmt.Printf("End\n\n")
}

func main() {
	initial()

	fmt.Printf("Insert a new node Walnut at index = 2 : \n")
	insert(2, "Walnut")
	fmt.Println("=====================Node inserted========================")
	output(head)
	fmt.Printf("Delete a new node Berkeley at index = 2 : \n")
	removeNode(2)
	fmt.Println("=====================Node deleted========================")
	output(head)
}
