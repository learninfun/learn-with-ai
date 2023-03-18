

Greedy algorithms are problem-solving techniques that make the best possible choices at each step in hopes of finding the optimal solution. The approach is to choose the local optimum at each step without looking for the big picture. 

One example of a greedy algorithm is the famous "Coin Change" problem. Suppose you have a total of n cents, and you need to make change for n cents in the minimum number of coins. Knowing the available denominations of coins, the objective is to select the minimum number of coins required to make that amount of change.

The algorithm starts with the largest denomination of coin and keeps deducting it from the total until it can't be deducted completely, then it moves to the next highest denomination coin till it meets the required amount. The algorithm continues this process until the change is made. 

For instance, if the denominations of coins are 1, 5, 10, 25, and 50, and we need to make change for 75 cents, the greedy algorithm would choose three 25-cent coins, as it is the largest denomination, to make up the 75 cents. Finally, the algorithm would return the solution 3, as the minimum number of coins required. 

While this method offers a simple solution to the coin change problem, it may not always lead to the optimal solution. Therefore, the greedy algorithm's feasibility strongly depends on selecting the right problem, analyzing the problem carefully, and defining a strategy that guarantees the optimal solution.