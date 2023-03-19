

1. What is a heap data structure?
Answer: A heap is a binary tree-based data structure that satisfies the Heap Property, which dictates that the parent of each node must be less than or equal to its child nodes.

2. What is the difference between a Min Heap and Max Heap?
Answer: A Min Heap is a heap where the parent node is less than or equal to its child nodes. Conversely, a Max Heap is a heap where the parent node is greater than or equal to its child nodes.

3. What is the time complexity for inserting an element into a heap?
Answer: The time complexity for inserting an element into a heap is O(log n) since the element must be added to the leaf node, and the heap property must be maintained by swapping the node with its parent.

4. How can you remove the root node from a Max Heap data structure?
Answer: To remove the root node from a Max Heap, you can first swap it with the last leaf node, delete the last leaf node, and then ensure the heap property is maintained by repeatedly swapping the root node with its largest child until the heap property is restored.

5. Can a heap be implemented using an array?
Answer: Yes, a heap can be implemented using an array where the parent-child relationship is defined by the array indices and the heap property is maintained by swapping elements to maintain the heap property.