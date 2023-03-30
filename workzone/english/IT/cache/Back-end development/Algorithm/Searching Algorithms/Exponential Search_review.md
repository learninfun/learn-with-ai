1. What is the time complexity of exponential search?
Answer: The time complexity of exponential search is O(log n), where n is the number of elements in the array.

2. What is the basic idea behind exponential search?
Answer: The basic idea behind exponential search is to search for an element in a sorted array by repeatedly doubling the search range until the element is found or the search range exceeds the size of the array.

3. Why is exponential search more efficient than linear search?
Answer: Exponential search is more efficient than linear search because it takes advantage of the fact that the array is sorted, and it reduces the search range by half with each iteration. This makes the search time logarithmic rather than linear.

4. Can exponential search be used on an unsorted array?
Answer: No, exponential search can only be used on a sorted array, since it relies on the assumption that the array is sorted in order to reduce the search range.

5. How does exponential search compare to binary search in terms of time complexity?
Answer: Exponential search and binary search both have a time complexity of O(log n), but exponential search may take more time than binary search for smaller values of n, since it may need to perform more comparisons in its initial search range. However, for larger values of n, exponential search may be faster than binary search, since it can reduce the search range more quickly.