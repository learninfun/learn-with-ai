1. What is Set Cover problem?
- Set Cover is a computational problem in which given a set of elements and a collection of subsets, we want to find the smallest number of subsets that covers all the elements in the original set.

2. What is the difference between the Set Cover and the Subset Sum problem?
- The Set Cover problem involves selecting a minimum number of subsets that cover all the elements in a given set. Whereas, Subset Sum problem involves selecting a subset of elements from a given set that add up to a particular sum.

3. Can the Set Cover problem be solved in polynomial time?
- No, the Set Cover problem is known to be NP-complete.

4. What is the Greedy Algorithm for solving the Set Cover problem?
- The Greedy Algorithm for Set Cover involves selecting the subset that covers the most uncovered elements at each iteration, until all elements are covered.

5. Can the Set Cover problem be approximated within a constant factor of the optimal solution?
- Yes, the Set Cover problem can be approximated within a constant factor of the optimal solution using the greedy algorithm. Specifically, the greedy algorithm has an approximation ratio of ln(n), where n is the number of elements in the original set.