

Bubble Sort is a simple and commonly used algorithm for sorting an array or list of elements. It works by comparing adjacent elements in the array and swapping them if they are out of order. The algorithm continues to pass through the list until no more swaps are required.

For example, consider the following array: 

[5, 1, 4, 2, 8] 

In the first pass, the algorithm compares the first two elements (5 and 1) and swaps them because 1 is smaller than 5. Then it compares 5 and 4 and swaps them to get [1, 5, 4, 2, 8]. 

In the second pass, it compares 5 and 4 and swaps them again to get [1, 4, 5, 2, 8]. It also compares 5 and 2 and swaps them to get [1, 4, 2, 5, 8]. 

The third pass compares 5 and 8 but no swap is necessary because 5 is not smaller than 8. The final pass compares 4 and 2 and swaps them to get the final sorted array [1, 2, 4, 5, 8]. 

The time complexity of bubble sort is O(n^2) in the worst case, where n is the number of elements in the array.