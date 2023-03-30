1. What is Memoization and how does it work?

Memoization is a technique used to speed up programs by caching function results so that they can be reused later. When a function is called with the same arguments, the cached result is returned instead of calculating it again, saving time.

2. What are the benefits of using Memoization in a program?

The benefits of using Memoization in a program are that it can improve the program's performance by reducing time complexity, as well as simplify code and reduce duplication by caching function results.

3. What are the potential drawbacks of using Memoization in a program?

The potential drawbacks of using Memoization in a program are increased memory usage due to caching results, as well as the need to ensure that cached results are still valid and up-to-date.

4. When should Memoization be used in a program?

Memoization should be used in a program when a function is called multiple times with the same arguments, and the function's computation is expensive or time-consuming. It is especially useful for recursive functions where the same calculation is repeated multiple times.

5. Can Memoization be applied to all types of functions?

Memoization can be applied to most functions, but it is most effective for functions that are deterministic, meaning that they always return the same output for the same input. Memoization is not as effective for non-deterministic functions, such as those that rely on external factors or produce random outputs.