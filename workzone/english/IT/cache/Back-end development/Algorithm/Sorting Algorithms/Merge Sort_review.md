1. What is the worst-case time complexity of Merge Sort and why?
Answer: The worst-case time complexity of Merge Sort is O(n log n) because the algorithm divides the input array into two halves recursively until a single element is left, and then merges the sorted halves back together.

2. What is the space complexity of Merge Sort and why?
Answer: The space complexity of Merge Sort is O(n) because the algorithm creates temporary arrays to store the divided halves of the input array during the sorting process.

3. How does Merge Sort differ from other comparison-based sorting algorithms like Quick Sort and Heap Sort?
Answer: Merge Sort is a stable sorting algorithm, which means that equal elements in the input array retain their relative order in the output. Quick Sort and Heap Sort are not necessarily stable. Additionally, Merge Sort has a guaranteed worst-case time complexity of O(n log n), while Quick Sort and Heap Sort have a worst-case time complexity of O(n^2).

4. What is the basic idea behind the divide and conquer approach used by Merge Sort?
Answer: The divide and conquer approach used by Merge Sort involves dividing the input array into smaller and smaller halves recursively until a single element is left. Each of these halves is then sorted and merged back together with its neighbor halves until the original input array is completely sorted.

5. Can Merge Sort be implemented as an in-place sorting algorithm?
Answer: No, Merge Sort requires additional memory in the form of temporary arrays to store the divided halves of the input array during the sorting process. Therefore, it cannot be implemented as an in-place sorting algorithm.