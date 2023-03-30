1. What is the time complexity of Radix Sort?

Answer: The time complexity of Radix Sort is O(d(n+k)), where d represents the number of digits in the largest number, n represents the number of elements to be sorted, and k represents the range of values.

2. How does Radix Sort work?

Answer: Radix Sort sorts elements by first grouping them by individual digits that share the same significant position within the number and then sorting them in each group based on the next significant position. This process is repeated until all digits have been sorted, resulting in a sorted list.

3. What is the significance of radix in Radix Sort?

Answer: Radix is the base of a number system, such as decimal (base 10) or binary (base 2). In Radix Sort, each digit of a number is treated as a number in a different base system, allowing for sorting based on the value of each digit.

4. What is the best case scenario for Radix Sort?

Answer: The best case scenario for Radix Sort is when all elements have the same number of digits and the range of values is small, resulting in a time complexity of O(n).

5. Can Radix Sort be used to sort negative numbers?

Answer: Radix Sort can only be used to sort non-negative numbers. To sort negative numbers, a modification to the algorithm is needed, such as adding a bias or using two separate lists for positive and negative numbers.