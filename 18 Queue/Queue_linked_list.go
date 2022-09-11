package main

import "fmt"

// ListNode defines a node of linked-list
type Node struct {
	data string
	next *Node
}

// LinkedListQueue defines a queue based on linked-list
type LinkedListQueue struct {
	head *Node
	tail *Node
	size int
}

func (q *LinkedListQueue) enqueue(element string) {
	node := &Node{data: element, next: nil}
	if q.tail == nil {
		q.tail = node
		q.head = node
	} else {
		q.tail.next = node
		q.tail = node
	}
	q.size++
}

func (q *LinkedListQueue) dequeue() *Node {
	if q.tail == nil {
		return nil
	}
	deqNode := q.head
	q.head = q.head.next
	q.size--
	return deqNode
}

// String returns object string
func (q *LinkedListQueue) output() string {
	if q.head == nil {
		return "empty queue"
	}
	result := "head"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf(" <- %+v ", cur.data)
	}
	result += "<- tail"
	return result
}

func main() {
	queue := LinkedListQueue(LinkedListQueue{})
	queue.enqueue("A")
	queue.enqueue("B")
	queue.enqueue("C")
	queue.enqueue("D")

	fmt.Printf("%s\n", queue.output())

	fmt.Printf("%d\n", queue.size)

	dequeuedNode := queue.dequeue()
	fmt.Printf("Dequed node : %s\n", dequeuedNode.data)

	fmt.Printf("%s\n", queue.output())
}
