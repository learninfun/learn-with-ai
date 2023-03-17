

1. What is a Min Heap?
Answer: A Min Heap is a binary tree where the value of each parent node is less than or equal to the value of its left and right child nodes.

2. How is a Min Heap different from a Max Heap?
Answer: The main difference between a Min Heap and a Max Heap is that in a Max Heap, the value of each parent node is greater than or equal to the value of its left and right child nodes.

3. What is the time complexity of inserting an element in a Min Heap?
Answer: The time complexity of inserting an element in a Min Heap is O(log n), where n is the number of elements in the heap.

4. How can we extract the minimum element from a Min Heap?
Answer: To extract the minimum element from a Min Heap, we remove the root node and replace it with the last element in the last level of the heap. Then we "heapify" the tree by comparing the new root node with its children and swapping if required.

5. Can a Min Heap have duplicate elements?
Answer: Yes, a Min Heap can have duplicate elements.