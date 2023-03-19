

The longest increasing subsequence (LIS) is a fundamental problem in computer science and mathematics that entails finding the longest subsequence in a given array of numbers that is strictly increasing. 

For example, consider the array [3, 4, -1, 0, 6, 2, 3] The LIS in the given array of numbers is [3, 4, 6], where the subsequence is strictly increasing in terms of index values. There are other increasing subsequences in the array like [3, 4, 6] and [0, 2, 3], but the length of these subsequences is not the longest possible length of increasing subsequence.

The problem is solved using dynamic programming techniques that find the length of the LIS at each index of the array, which is based on the previously calculated LIS lengths. Finally, the subsequence is then constructed by backtracking through the array.