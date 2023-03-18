

Randomized Quicksort is a sorting algorithm that partitions an array by picking a random pivot element and sorting the smaller elements on the left and the larger elements on the right. It is a randomized version of the Quicksort algorithm, which picks the first or the last element as the pivot.

Steps of Randomized Quicksort Algorithm:

1. Select a random element from the array as the pivot.

2. Partition the array around the pivot element by comparing every other element to it. All the smaller elements go to the left of the pivot and all the larger elements go to the right.

3. Recursively apply the algorithm to the left and the right partition.

4. The base case for the recursion is when there is only one element in the subarray.

Example:

Suppose we have an array [7, 2, 8, 1, 4, 6, 9, 3, 5]. We can apply the Randomized Quicksort algorithm to sort the array as follows:

1. Pick a random pivot. Let's say the pivot is 6.

2. Partition the array around the pivot: [2, 1, 4, 3, 5], [7, 8, 9]

3. Recursively apply the algorithm to the left and the right partition.

- On the left partition, pick a new pivot randomly. Let's say the pivot is 4.
- Partition the left partition: [2, 1, 3], [4, 5]
- Recursively apply the algorithm to the left subpartition and the right subpartition.
- Continue this process until the base case is reached.

- On the right partition, pick a new pivot randomly. Let's say the pivot is 8.
- Partition the right partition: [7], [8, 9]
- Recursively apply the algorithm to the left subpartition and the right subpartition.
- Continue this process until the base case is reached.

4. Finally, combine the sorted subarrays to get the sorted array. [1, 2, 3, 4, 5, 6, 7, 8, 9]