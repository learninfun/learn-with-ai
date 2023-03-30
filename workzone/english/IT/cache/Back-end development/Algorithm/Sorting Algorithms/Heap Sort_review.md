1. What is the time complexity of Heap Sort in the worst case scenario?
Answer: O(nlogn)

2. How does Heap Sort compare to Quick Sort in terms of average and worst-case time complexity?
Answer: Quick Sort has a better average-case time complexity, but Heap Sort is more consistent with its worst-case time complexity of O(nlogn).

3. What is a Heap data structure, and how does it relate to Heap Sort?
Answer: A Heap is a binary tree data structure in which every parent node has a value greater than or equal to its children. In Heap Sort, the input data is organized into a heap before being sorted.

4. How does Heap Sort differ from other sorting algorithms in terms of memory usage?
Answer: Heap Sort is an in-place sorting algorithm, meaning it sorts the input data without requiring any additional memory space beyond the original array. Therefore, it has a relatively low memory footprint compared to other sorting algorithms.

5. How does Heap Sort sort elements in decreasing order?
Answer: To sort elements in decreasing order, a max heap is used instead of a min heap. This means that the largest element is always placed at the root of the heap, and each successive element is placed in smaller sub-heaps until the entire input array is sorted in descending order.