

1. What is a doubly linked list?
Answer: A doubly linked list is a type of data structure in which each node contains two pointers - one that points to the previous node in the list and one that points to the next node in the list.

2. How is a node added to the end of a doubly linked list?
Answer: To add a node to the end of a doubly linked list, you need to traverse the list until you reach the last node. then create a new node and set its previous pointer to point to the previous last node, and the next pointer of the previous last node to point to the new node.

3. What is the time complexity of searching for an element in a doubly linked list?
Answer: The time complexity of searching for an element in a doubly linked list is O(n) in the worst case, as you may need to traverse through the entire list to find the element.

4. What is the advantage of a doubly linked list over a singly linked list?
Answer: The advantage of a doubly linked list over a singly linked list is that you can traverse the list in both directions, which makes it easier to navigate and perform modifications, especially when dealing with large lists.

5. How is a node removed from a doubly linked list?
Answer: To remove a node from a doubly linked list, you need to update the previous node's next pointer to point to the next node, and the next node's previous pointer to point to the previous node. Then, you can free the memory of the node that you want to remove.