

Average-case time complexity refers to the amount of time it takes to execute an algorithm assuming input that has average properties. In other words, average-case time complexity is defined in terms of the statistical distribution of the input data.

For example, consider the problem of sorting a large array of numbers. The average-case time complexity of an algorithm to sort an array is usually O(n log n). This means that the time it takes to sort an array of n elements is proportional to n log n on average. However, this assumes that the input data is randomly ordered, as it would be if the data was generated from a uniformly distributed probability distribution. 

If the input array is already sorted or contains many duplicates, this could result in a worst-case time complexity that is higher than O(n log n). Conversely, if the input array contains a lot of randomness, the actual running time of the algorithm may be lower than O(n log n) on average. Therefore, average-case time complexity is a more realistic measure of an algorithm's performance than worst-case time complexity.