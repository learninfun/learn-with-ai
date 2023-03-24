

Memoization is an optimization technique used in computer programming to speed up the performance of a program by caching the result of expensive function calls and returning them when the same inputs occur again. This can be seen as a way to store and reuse the computed values of a function, avoiding the need to recalculate them every time the function is called.

Here is an example of Memoization in Python:

```
# regular recursive function
def fibonacci(n):
    if (n < 2):
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)

# memoized version of the same function
def fibonacci_mem(n, memo = {}):
    if (n < 2):
        return n
    if (n in memo):
        return memo[n]
    memo[n] = fibonacci_mem(n - 1, memo) + fibonacci_mem(n - 2, memo)
    return memo[n]

print(fibonacci(10)) # output: 55
print(fibonacci_mem(10)) # output: 55
```

In this example, the `fibonacci` function calculates the nth number in the Fibonacci sequence using recursion. However, this approach can be computationally expensive for larger values of n. To speed up the function, we can use memoization to store the results of previously calculated values and reuse them when the function is called again with the same input.

The `fibonacci_mem` function is a memoized version of the same function, which takes an additional `memo` argument. When the function is called, it first checks if the value of `n` is already present in the memo dictionary. If it is, the function returns the cached value instead of computing it again. If the value is not present in the dictionary, the function calculates it and stores it in the memo dictionary for future use. This greatly reduces the time complexity of the function, making it faster and more efficient.