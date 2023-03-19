

1. A binary heap is a binary tree that satisfies the heap property.
2. The heap property states that the key of any node is greater than or equal to the key of its children (if they exist).
3. Binary heaps can be represented as an array, where the root is at index 1, and the children of node i are at indices 2i and 2i+1.
4. Binary heaps are commonly used to implement priority queues.
5. The two main operations on a binary heap are insert and extract-min (or extract-max in a max heap).
6. Inserting a new element involves adding it to the bottom level of the heap and then swapping it up until the heap property is satisfied.
7. Extracting the minimum element involves removing the root and then swapping the last element in the heap to the root and sifting it down until the heap property is satisfied.
8. Binary heaps can be built from a given array of elements in O(n) time, using a bottom-up approach.
9. The height of a binary heap with n elements is logarithmic in n, making its operations very efficient.