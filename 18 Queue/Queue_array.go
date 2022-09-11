package main

import "fmt"

// ArrayQueue defines a queue based on array
type ArrayQueue struct {
	data     []string
	capacity int
	head     int
	tail     int
}

// NewArrayQueue creates a queue
func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{data: make([]string, capacity), capacity: capacity, head: 0, tail: 0}
}

// Enqueue puts an element in the tail queue
func (q *ArrayQueue) enqueue(element string) bool {
	if q.tail == q.capacity {
		if q.head == 0 {
			return false
		}
		// Move data
		for i := q.head; i < q.tail; i++ {
			q.data[i-q.head] = q.data[i]
		}
		q.tail -= q.head
		q.head = 0
	}
	q.data[q.tail] = element
	q.tail++
	return true
}

// Dequeue fetches a element from queue
func (q *ArrayQueue) dequeue() string {
	if q.head == q.tail {
		return ""
	}
	deqElement := q.data[q.head]
	q.head++
	return deqElement
}

// Output prints the queue
func (q *ArrayQueue) output() string {
	if q.head == q.tail {
		return "Queue empty"
	}
	result := "head"
	for i := q.head; i <= q.tail-1; i++ {
		result += fmt.Sprint(" <- ", q.data[i])
	}
	result += " <- tail"
	return result
}

func main() {
	queue := NewArrayQueue(10)
	queue.enqueue("A")
	queue.enqueue("B")
	queue.enqueue("C")
	queue.enqueue("D")

	fmt.Printf("%s\n", queue.output())

	fmt.Printf("%d\n", len(queue.data))

	dequeuedNode := queue.dequeue()
	fmt.Printf("Dequed node : %s\n", dequeuedNode)

	fmt.Printf("%s\n", queue.output())
}
