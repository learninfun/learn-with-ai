

1. What is a Max Heap?
Answer: A Max Heap is a binary tree where the parent node is larger than its child nodes.

2. How is a Max Heap represented in an array?
Answer: In an array representation of a Max Heap, each node is assigned an index, and the parent node is assigned the index value (i-1)/2, the left child node is assigned the index value 2i+1, and the right child node is assigned the index value 2i+2.

3. What is the time complexity of inserting a new element in a Max Heap?
Answer: The time complexity of inserting a new element in a Max Heap is O(log n), where n is the number of nodes in the heap.

4. How is the maximum element extracted from a Max Heap?
Answer: To extract the maximum element from a Max Heap, the root node (maximum element) is swapped with the last node in the heap. The last node is then removed, and the heap is re-arranged to maintain the Max Heap property.

5. What is the time complexity of extracting the maximum element from a Max Heap?
Answer: The time complexity of extracting the maximum element from a Max Heap is O(log n), where n is the number of nodes in the heap.