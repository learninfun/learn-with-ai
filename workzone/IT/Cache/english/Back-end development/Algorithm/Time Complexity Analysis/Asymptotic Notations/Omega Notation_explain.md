

Omega notation is a mathematical notation used to describe the lower bound of a function or algorithm. It defines the minimum growth rate of a function, so the actual running time of an algorithm cannot exceed this minimum bound.

For example, if a function f(n) is said to be in Omega(g(n)), then there exist constants c and n0 such that f(n) >= c * g(n) for all n >= n0. This means that the function f(n) grows at least as fast as g(n) for large enough values of n.

One example of Omega notation is: 

f(n) = n^2 + 3n + 1 is in Omega(n^2)

This means that the running time of this algorithm is at least proportional to n^2. No matter how efficient the algorithm is, it can't have a running time that is better than this lower bound.