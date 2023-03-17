

Heap Sort is a sorting algorithm that works by creating a binary heap and repeatedly removing the maximum element from it until the heap is empty. The binary heap can be either min-heap or max-heap, but in the context of Heap Sort, it is typically a max-heap.

The steps involved in Heap Sort are:

1. Convert the given array into a max-heap by using the heapify procedure, which is a way of converting a binary tree into a heap by swapping nodes until the ordering property is maintained.

2. Remove the maximum element from the root node, and swap it with the last element in the heap.

3. Reduce the size of the heap by one, and repeat steps 1 and 2 until the heap is empty.

4. The resulting array will be sorted in ascending order.

Example:

Let's say we have an array of integers as follows:

[23, 15, 34, 55, 47, 12, 56, 72]

To sort this array using Heap Sort, we first create a max-heap by using the heapify procedure as follows:

[72, 55, 56, 23, 47, 12, 34, 15]

We then remove the maximum element from the root node, which is 72, and swap it with the last element in the heap, which is 15, resulting in the following heap:

[55, 47, 56, 23, 15, 12, 34, 72]

We repeat this process until the heap is empty, resulting in the sorted array:

[12, 15, 23, 34, 47, 55, 56, 72]