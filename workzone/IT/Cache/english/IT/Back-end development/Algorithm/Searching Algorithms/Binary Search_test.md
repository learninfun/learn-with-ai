

1. What is the time complexity of Binary Search when searching for an element in a sorted array of size n?
Answer: O(log(n))

2. What is the main advantage of using Binary Search over Linear Search for finding elements in an array?
Answer: Binary Search has a faster average search time, which is O(log(n)) instead of O(n) for Linear Search.

3. Can Binary Search be used on unsorted arrays?
Answer: No, Binary Search requires the array to be sorted in order to work correctly.

4. What is the worst-case scenario for Binary Search when searching for an element that does not exist in the array?
Answer: The worst-case scenario is when the target element is not in the array at all, and Binary Search must check every element. This takes O(log(n)) time, just like when the element is found.

5. How does Binary Search work internally to find a target element in a sorted array?
Answer: Binary Search continually narrows down the search range by comparing the target element to the middle element of the current range. If the target is less than the middle element, the search range is narrowed to the lower half of the current range. If the target is greater, the search range is narrowed to the upper half. This process is repeated until either the target element is found or the search range is empty.