1) What is a linked list, and how does it differ from an array?

Answer: A linked list is a data structure consisting of a series of nodes, each containing a value and a reference or pointer to the next node. Unlike an array, which stores elements in a contiguous block of memory, a linked list can be created dynamically and the nodes can be scattered throughout memory.

2) What is the time complexity for inserting a node at the beginning of a singly linked list?

Answer: The time complexity for inserting a node at the beginning of a singly linked list is O(1), as it only requires updating the reference of the new node to point to the previous head node, and updating the head reference to point to the new node.

3) What is a doubly linked list, and how does it differ from a singly linked list?

Answer: A doubly linked list is a data structure consisting of a series of nodes, each containing a value and references or pointers to the previous and next nodes. Unlike a singly linked list, which only has a reference to the next node, a doubly linked list allows for traversal in both directions.

4) What is the time complexity for searching for a specific value in a linked list?

Answer: The time complexity for searching for a specific value in a linked list depends on the implementation. In a singly linked list, the worst-case time complexity is O(n), as all nodes must be traversed until the value is found or the end of the list is reached. In a doubly linked list, the worst-case time complexity can be reduced to O(n/2) by searching from both ends simultaneously.

5) What is a circular linked list, and what are some use cases for it?

Answer: A circular linked list is a variation of a linked list in which the last node points back to the first node, creating a circular structure. This can be useful for implementing circular data structures such as buffers or queues, in which elements are added and removed in a cyclic manner. It can also be used for circular navigation in a game or interactive application.