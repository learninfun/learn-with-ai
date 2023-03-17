

1. Heap Sort is a comparison-based sorting algorithm that works by sorting an unsorted array of elements by creating a binary heap.
2. A binary heap is a complete binary tree where all the nodes have either a greater or smaller value than their children nodes.
3. The two types of binary heaps are the max heap, where the parent node has a greater value than its children nodes, and the min heap, where the parent node has a smaller value than its children nodes.
4. The heap sort algorithm starts by creating a max heap from the unsorted array of elements.
5. The first step in creating a max heap is to build a complete binary tree.
6. Once the complete binary tree has been built, the algorithm sifts down the nodes in the tree until it is a valid max heap.
7. After the max heap has been constructed, the algorithm performs a swap operation between the root node and the last node in the list.
8. The next step is to reduce the size of the array by one, and then re-heapify the binary tree.
9. The re-heapification process is achieved by sifting down the nodes in the binary tree.
10. The above steps are repeated until the entire list is sorted in ascending order.