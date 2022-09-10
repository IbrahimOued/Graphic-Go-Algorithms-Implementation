# Doubly Linked List

**Doubly Linked List**:
It is a chained storge structure of a linear table. It is connected by nodes in two directions. Each node consists of data, pointing to the previous node and pointing to the next node.

![](exanple%20doubly%20linked%20list.png)

UML

![](uml.png)

```go
type Node struct{
    data string
    prev *Node
    next *Node
}

1. Doubly linked list initialization

Example : Construct a SF subway doubly linked list
```

2. Traversal output

3. Add a node Walnut at the end of Fremont

![](insert.png)

4. Delete the index=2 node

![](delete.png)