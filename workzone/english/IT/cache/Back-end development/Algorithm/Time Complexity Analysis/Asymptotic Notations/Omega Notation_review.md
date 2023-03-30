1. What is the range of values for the set of functions that can be expressed using Omega notation? 
Answer: The range of values is the set of functions that have a lower bound on their growth rate.

2. What is the difference between Omega notation and Big-O notation?
Answer: Omega notation provides a lower bound on the growth rate of a function, whereas Big-O notation provides an upper bound.

3. How do you determine whether a function is in Omega notation of another function?
Answer: A function f(n) is in Omega notation of another function g(n) if there exist constants c and n0 such that f(n) >= c * g(n) for all n >= n0.

4. Can a function be in both Big-O and Omega notation of another function?
Answer: Yes, if a function f(n) is in Big-O notation of another function g(n), and also in Omega notation of another function h(n), then f(n) is said to be in Theta notation of g(n) and h(n).

5. How does the Omega notation of a function relate to its worst-case time complexity?
Answer: The Omega notation provides a lower bound on the growth rate of the function, which is useful in determining the minimum amount of time that a program using the function would take on its worst-case input.