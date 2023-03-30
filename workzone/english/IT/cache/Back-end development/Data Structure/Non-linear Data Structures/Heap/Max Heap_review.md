1. What is a Max Heap? 
Answer: A Max Heap is a data structure where the parent nodes hold a higher value than their child nodes, where the highest value is the root node.

2. How is a Max Heap different from a Min Heap? 
Answer: A Max Heap is a similar data structure as a Min Heap, but in a Max Heap, the parent node holds a higher value than the child nodes, while in a Min Heap, the parent node holds a smaller value than the child nodes.

3. Describe the process of inserting a new element into a Max Heap. 
Answer: To insert a new element into a Max Heap, the new element is added as the last element in the heap, and then the heap property is maintained by swapping the new element with its parent if it is larger than its parent until the heap property is satisfied.

4. What is the time complexity of finding the maximum element in a Max Heap? 
Answer: The maximum element in a Max Heap is always the root node; therefore, it can be found in constant time O(1).

5. Can a Max Heap have duplicate elements? 
Answer: Yes, a Max Heap can have duplicate elements, but the priority of the elements is determined by their index in the heap. The first element that was inserted becomes the root node, and the nodes with the same value will be inserted to the right of it.