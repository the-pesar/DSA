# Linked List

A linked list is a linear data structure where elements, called nodes, are connected through pointers. Each node contains two parts: data and a reference (or pointer) to the next node in the sequence. Unlike arrays, linked lists do not require contiguous memory allocation, allowing for dynamic size adjustment and efficient insertion or deletion of elements. There are various types of linked lists, such as singly linked lists (nodes point to the next node only), doubly linked lists (nodes point to both next and previous nodes), and circular linked lists (last node points back to the first).

## Singly Linked List
![alt text](https://cdn.programiz.com/sites/tutorial2program/files/linked-list-concept_0.png)
It is the most common. Each node has data and a pointer to the next node.

```go
type node struct {
  Val  int
  Next *node
}
```

## Doubly Linked List
![alt text](https://cdn.programiz.com/sites/tutorial2program/files/doubly-linked-list-concept.png)
In a Doubly Linked List, each node contains three parts: data, a pointer to the next node, and a pointer to the previous node. This allows bidirectional traversal, making operations like insertion and deletion more flexible compared to a singly linked list. It’s especially useful when you need to frequently navigate both ways through a list.

```go
type node struct {
  Val  int
  Next *node
  Prev *node
}
```

## Circular Linked List
![alt text](https://cdn.programiz.com/sites/tutorial2program/files/circular-linked-list.png)
A Circular Linked List is a type of linked list where the **last node** points back to the **first node**, forming a circular structure. This allows continuous traversal through the list without reaching a null reference, making it ideal for applications like round-robin scheduling. It can be singly or doubly linked.

### Related problems that I have solved:
- [Remove duplicates from sorted list (Leetcode)](https://leetcode.com/problems/remove-duplicates-from-sorted-list/)
- [Linked list cycle (Leetcode)](https://leetcode.com/problems/linked-list-cycle/)
- [Remove linked list elements (Leetcode)](https://leetcode.com/problems/remove-linked-list-elements/)
- [Reverse linked list (Leetcode)](https://leetcode.com/problems/reverse-linked-list/)