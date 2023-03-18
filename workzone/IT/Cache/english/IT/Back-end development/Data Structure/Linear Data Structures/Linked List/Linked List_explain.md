

A linked list is a common data structure that stores a sequence of values, called nodes. Each node contains a value and a reference to the next node in the list. The reference may be a pointer or a memory address. This way, the nodes are not stored in contiguous memory locations like arrays, but rather linked together using pointers or references.

For example, let's consider a singly linked list of integers 1, 2, 3, 4. It would consist of four nodes:
- Node 1: value=1, next=Node 2
- Node 2: value=2, next=Node 3
- Node 3: value=3, next=Node 4
- Node 4: value=4, next=NULL

In this example, each node stores an integer value, and the reference or pointer "next" points to the next node in the list. The last node has a NULL reference for "next," indicating the end of the list. 

Linked lists can be used to implement various data structures like stacks, queues, and hash tables. They are efficient for inserting and deleting elements from the list, but accessing individual nodes can be slower than in arrays because it requires sequential traversal.