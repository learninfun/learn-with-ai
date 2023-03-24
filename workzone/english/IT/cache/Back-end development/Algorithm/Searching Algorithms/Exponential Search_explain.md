

Exponential Search is a searching algorithm that works on sorted arrays. It tries to find the range where the target element may lie, and then performs a binary search within that range to locate the element. 

The algorithm works as follows:

1. Set a minimum value (start) to 0 and a maximum value (end) to the length of the array minus 1
2. Check if the target value is equal to the first element of the array. If it is, return 0 as the index.
3. Set an exponent value (exp) to 1 and keep increasing it until the index of the element at the exponent position is greater than the target value or the end of the array is reached.
4. Once an exponent value that meets the criteria is found, perform a binary search within the range of the previous exponent and the current exponent to find the target value. 

Here's an example: 

Suppose we have a sorted array [2, 7, 13, 18, 22, 25, 32, 35, 40]. We want to search for the element 22 using the exponential search algorithm.

1. Start = 0, End = 8, Target = 22
2. Check if target = array[0] = 2, it's not.
3. The first exponent is 1, so we check array[1] = 7. As 7 < target, we continue to the next exponent.
4. The second exponent is 2, so we check array[2] = 13. As 13 < target, we continue to the next exponent.
5. The third exponent is 4, so we check array[4] = 22. As 22 = target, we perform a binary search within the range of previous exponent (2) and current exponent (4).
6. Binary search finds the target element at index 4.

In this case, exponential search took 3 iterations to find the range where the target element lies, and binary search took 2 iterations to locate the element. The total number of iterations is 5 in this case, which is less than the 9 iterations required for a traditional binary search.