

1. What is a linked list and how is it different from other data structures? 
Answer: A linked list is a linear data structure that consists of a head node and a set of nodes that each has a reference to the next node in the list. It differs from arrays and other data structures in that it allows for dynamic memory allocation and is well-suited for inserting and deleting elements in the middle of the list. 

2. What is the time complexity of searching for an element in a singly linked list? 
Answer: The time complexity of searching for an element in a singly linked list is O(n), where n is the number of elements in the list. This is because we must traverse the list from the head node until we find the element we are searching for or reach the end of the list. 

3. How do you insert an element at the beginning of a linked list? 
Answer: To insert an element at the beginning of a linked list, we create a new node with the value we want to insert and set its next node to the current head node. Then, we set the head of the list to the new node. 

4. What is the difference between a singly linked list and a doubly linked list? 
Answer: The difference between a singly linked list and a doubly linked list is that in a singly linked list each node only has a reference to the next node in the list, while in a doubly linked list each node has a reference to both the next and previous nodes in the list. This allows for more efficient traversal and operations such as inserting or deleting elements in the middle of the list. 

5. How do you delete an element from a linked list? 
Answer: To delete an element from a linked list, we first traverse the list to find the node that contains the value we want to delete. Once we find the node, we set the next reference of the previous node to the next node after the one we want to delete, effectively skipping over it. Then, we deallocate the memory used by the deleted node.