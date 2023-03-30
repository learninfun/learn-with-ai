1. What is a Singly Linked List?
Answer: A Singly Linked List is a linear data structure in which each element is a node that contains a data value and a pointer to the next node.

2. How is insertion of a new node in a Singly Linked List performed?
Answer: Insertion of a new node in a Singly Linked List is performed by creating a new node and setting its pointer to the next node in the list. Then, the previous node's pointer is updated to point to the new node, effectively inserting the new node into the list.

3. What is the time complexity of searching for a particular element in a Singly Linked List?
Answer: The time complexity of searching for a particular element in a Singly Linked List is O(n), where n is the number of nodes in the list. This is because we need to traverse the entire list to find the desired element.

4. How can we remove a node from a Singly Linked List?
Answer: To remove a node from a Singly Linked List, we need to update the pointer of the previous node to point to the next node in the list, effectively skipping over the node we want to remove. Then, we can delete the node we want to remove from memory.

5. What is the difference between a Singly Linked List and a Doubly Linked List?
Answer: The main difference between a Singly Linked List and a Doubly Linked List is that in the latter, each node contains a pointer to both the next node and the previous node. This allows for traversal of the list in both directions, making certain operations more efficient, but it also requires more memory than a Singly Linked List.