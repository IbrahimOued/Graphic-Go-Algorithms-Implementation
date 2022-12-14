package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

var root *Node = nil

func createNewNode(newData int) *Node {
	var newNode *Node = new(Node)
	newNode.data = newData
	newNode.left = nil
	newNode.right = nil
	return newNode
}

// In order traversal binary search tree
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left) // Traversing the left subtree
	fmt.Printf("%d, ", root.data)
	inOrder(root.right) // Traversing the right subtree
}

func insert(node *Node, newData int) {
	if root == nil {
		root = &Node{data: newData, left: nil, right: nil}
		return
	}
	var compareValue = newData - node.data
	// Recursive left subtree, continue to find the insertion position
	if compareValue < 0 {
		if node.left == nil {
			node.left = createNewNode((newData))

		} else {
			insert(node.left, newData)
		}
	} else if compareValue > 0 {
		// Recursive right subtree, continue to find the insertion position
		if node.right == nil {
			node.right = createNewNode(newData)
		} else {
			insert(node.right, newData)
		}
	}
}

func main() {
	//Constructing a binary search tree
	insert(root, 60)
	insert(root, 40)
	insert(root, 20)
	insert(root, 10)
	insert(root, 30)
	insert(root, 50)
	insert(root, 80)
	insert(root, 70)
	insert(root, 90)
	fmt.Printf("In-order traversal binary search tree \n")
	inOrder(root)
}
