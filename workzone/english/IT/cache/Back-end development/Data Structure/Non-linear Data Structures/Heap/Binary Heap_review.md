1. What is a binary heap and how is it structured?
Answer: A binary heap is a data structure that is represented as a binary tree, where each parent node has at most two child nodes. The heap follows the property that the value of the parent node is always greater than or equal to its children (in a max heap) or less than or equal to its children (in a min heap). 

2. How does a binary heap differ from a binary search tree?
Answer: A binary heap is different from a binary search tree in that it is not sorted. The value of any parent node in a binary heap is only greater than or less than its immediate children, whereas in a binary search tree, the value of any node is greater than or equal to its left child and less than or equal to its right child. 

3. How do you insert an element into a binary heap?
Answer: To insert an element into a binary heap, it must first be added as the last element in the last level of the heap. Then, the heap must be adjusted so that the parent nodes maintain their heap property, either by swapping the new element with its parent or recursively swapping with its parent until it satisfies the heap property. 

4. What is the time complexity of finding the minimum or maximum element in a binary heap?
Answer: The time complexity of finding the minimum or maximum element in a binary heap is O(1), as the root node always contains the minimum or maximum element depending on whether it is a min heap or max heap. 

5. How do you remove the minimum or maximum element from a binary heap?
Answer: To remove the minimum or maximum element from a binary heap, the root node is removed and replaced with the last element in the heap. The heap property is then restored by recursively swapping the new root node with its children until the heap property is satisfied.