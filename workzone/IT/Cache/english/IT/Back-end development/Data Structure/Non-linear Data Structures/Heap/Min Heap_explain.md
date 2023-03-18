

A Min Heap is a special type of binary heap where the value of each parent node is smaller than or equal to its children. It means that the smallest element is always at the root. This property is called the “min-heap property.”

For example, consider the following binary tree:

        6
      /   \
     7     12
    / \    / \
  10  15  14  20

This is not a valid min heap because the parent node (6) is larger than its children (7 and 12). We can convert this into a min heap by swapping the values of the parent node with its smallest child until the min heap property is satisfied.

After the conversion, the new min heap will look like this:

        6
      /   \
    7      12
   / \    / \
  10 15  14 20

Now, the smallest element (6) is at the root and satisfies the min-heap property. We can use min heap to implement various algorithms such as priority queues, sorting algorithms like heap sort, and graph algorithms like Dijkstra’s shortest path algorithm.