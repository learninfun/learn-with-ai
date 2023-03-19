

1. What is the difference between O(n) and O(n^2) notation?
Answer: O(n) represents an algorithm that has a time complexity of linear growth, while O(n^2) represents an algorithm with time complexity that grows quadratically with input size.

2. How is Omega notation related to Big O notation?
Answer: Omega notation is just the inverse of Big O notation. While Big O represents the upper bound of an algorithm's runtime, Omega represents the lower bound or best-case scenario.

3. What is the Omega notation of a constant time algorithm?
Answer: The Omega notation of a constant time algorithm is O(1), meaning that the complexity is constant and independent of input size.

4. Can an algorithm have a Big O notation of O(n) and an Omega notation of O(n^2)?
Answer: No, an algorithm cannot have a Big O notation that is less than or equal to its Omega notation. It would mean that the upper and lower bounds of the algorithm's complexity would be different, which is impossible.

5. Suppose an algorithm has a Big O notation of O(n log n) and an Omega notation of O(n). What can we infer about the algorithm's behavior?
Answer: This means that the algorithm performs better than a quadratic function but worse than a linear function. It can also suggest that the algorithm makes use of sorting or merging techniques.