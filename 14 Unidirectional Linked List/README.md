# Unidirectional Linked List

**Unidirectional Linked List Single Link**:
Is a chained storage structure of a linear table, which is connected by a node. Each node consists of data and next pointer to the next node.

![](example%20linked%20list.png)

UML diagram\
![](uml%20linked%20list.png)

```go
type Node struct {
    data string
    next *Node
}
```

## 1. Unidirectional Linked List Initialization

Example: Construcut a SF subway Unidirectional Linked List

![hint](SF%20subway.png)

## 2. Traversal output

![traversal](traversal.png)

## 3 Insert a node Walnut in position 2

![](insert.png)

## 4 Delete the index = 2 node

![](delete.png)