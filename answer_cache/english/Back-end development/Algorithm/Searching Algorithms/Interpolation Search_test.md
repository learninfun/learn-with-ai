

1. What is the time complexity of Interpolation Search algorithm in worst case?
Answer: The worst case time complexity for Interpolation search is O(n).

2. Explain the basic concept of Interpolation search.
Answer: Interpolation search is a search algorithm which works on a sorted array of elements. It uses the value of the target key and the estimation of its position in the array to determine where to search next. It is more efficient than binary search if the values in the array are uniformly distributed.

3. What is the difference between Interpolation search and Binary search?
Answer: Interpolation search is a more efficient search algorithm than binary search because it uses the value of the target key and the estimation of its position in the array to determine where to search next. Binary search uses the midpoint of the array as the index for each comparison.

4. What are the advantages of using Interpolation search?
Answer: Interpolation Search has the following advantages:
a) It is more efficient than binary search in interpolation.
b) It performs better when values in the array are uniformly distributed.
c) It improves the complexity of the algorithm.

5. What is the formula used in Interpolation search algorithm to estimate the position of the target key?
Answer: The formula used in Interpolation search algorithm to estimate the position of the target key is:
mid = low + [(target - arr[low]) / (arr[high] - arr[low]) * (high - low)]