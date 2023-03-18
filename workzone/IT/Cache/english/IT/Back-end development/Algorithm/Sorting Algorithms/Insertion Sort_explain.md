

Insertion sort is an algorithm in computer science used to sort an array or list of elements into ascending or descending order. It works by sorting one element at a time and placing them in their correct position.

The idea behind insertion sort is that for each element in the unsorted list, it is compared to the elements that are already sorted, and then placed in its correct position relative to them.

Here is how it works:

1. We start with an unsorted list of integers: [5, 3, 1, 8, 6, 9]

2. We take the first element and compare it to the second element. If the first element is greater than the second element, we swap them.

[5, 3, 1, 8, 6, 9] -> [3, 5, 1, 8, 6, 9]

3. We take the third element and compare it to the first two elements. We swap it with the second element if it is smaller, and swap it with the first element if it is smaller than both the first and the second element.

[3, 5, 1, 8, 6, 9] -> [1, 3, 5, 8, 6, 9]

4. We repeat this process for all the remaining elements in the array until the entire array is sorted.

[1, 3, 5, 8, 6, 9] -> [1, 3, 5, 6, 8, 9]

5. The resulting sorted array is [1, 3, 5, 6, 8, 9].

Insertion sort is a simple and efficient sorting algorithm for small arrays, but its performance deteriorates quickly for large arrays.