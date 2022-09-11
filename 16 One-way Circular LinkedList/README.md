# One-way Circular LinkedList

**One-way Circular List:**: It is a chain storage structure of a linear table, which is connected to form a ring and each node is composed of data and a pointer to next.

![example linked list](example%20circular%20linked%20list.png)

UML

![uml representation](uml.png)

```go
type Node struct {
    data string
    next *Node
}
```

1. One-way Circular Linked List initialization and traversal ouput

![initialization and insert](initialization%20and%20traversal.png)


2. Insert a node E in position 2

![](insert.png)

![](insert%202.png)

3. Delete the index=2 node

![](delete%201.png)

![](delete%202.png)