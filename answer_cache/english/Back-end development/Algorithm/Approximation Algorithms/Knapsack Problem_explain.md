

Knapsack problem is a classical optimization problem that involves finding the maximum value that can be obtained by selecting a subset of items with given weights and values, such that their total weight does not exceed a maximum weight limit. It is a well-known NP-hard problem, and many real-world applications can be modeled as knapsack problems.

For example, imagine that you are a thief who just broke into a jewelry store, and your goal is to maximize your profit by stealing the most valuable items you can carry in your backpack. The items have different weights and values, and your backpack can only hold a maximum weight of $W$ kilograms. Your task is to pick a subset of items with maximum total value that does not exceed the backpack's weight limit.

More formally, the knapsack problem can be defined as follows: Given a set of $n$ items, where each item $i$ has a weight $w_i$ and a value $v_i$, find a subset $S$ of items that maximizes the sum of the values of the items in $S$ subject to the constraint that the total weight of the items in $S$ does not exceed a given limit $W$.