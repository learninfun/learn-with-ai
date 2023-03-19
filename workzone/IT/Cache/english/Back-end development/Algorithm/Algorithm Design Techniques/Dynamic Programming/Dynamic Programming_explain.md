Dynamic programming (DP) is an algorithmic approach that involves solving complex problems by breaking them down into smaller subproblems, solving them independently, and storing the results for future reference. It is used when the same subproblems are encountered multiple times, and solving them repeatedly is time-consuming. Dynamic programming seeks to optimize solutions by reducing the number of repeated computations.

An example of dynamic programming is the Fibonacci sequence. The sequence can be defined as follows: 

Fib(0) = 0
Fib(1) = 1
Fib(n) = Fib(n-1) + Fib(n-2) for n>1

To find the nth Fibonacci number, one can use the recursive formula in the definition above. However, this approach is inefficient because it requires the same calculations to be repeated multiple times. For example, to calculate Fib(n), one needs to calculate Fib(n-1) and Fib(n-2), which in turn require the calculation of other Fibonacci numbers.

With dynamic programming, we can solve this problem more efficiently by storing the results of the previous calculations and using them when necessary. One common approach is to use a table to store the values of Fib(n) for all n from 0 to n. Starting from Fib(0) and Fib(1), we can calculate Fib(2) by adding Fib(0) and Fib(1), Fib(3) by adding Fib(2) and Fib(1), and so on. By storing the values in a table, we can avoid calculating the same Fibonacci numbers multiple times, leading to a more efficient algorithm.