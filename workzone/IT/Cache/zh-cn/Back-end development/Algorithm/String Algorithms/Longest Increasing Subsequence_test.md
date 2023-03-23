

1. LeetCode 300 - Longest Increasing Subsequence
给定一个未排序的整数数组，找到最长递增子序列的长度。
例如：
Input：[10,9,2,5,3,7,101,18]
Output：4
Explanation：
最长的递增子序列是 [2,3,7,101]，因此长度为 4。

2. GeeksforGeeks - Longest Increasing Subsequence
给定一个整数数列，找到其中的最长递增子序列。
例如：
Input：[3, 10, 2, 1, 20]
Output：3
Explanation：
最长的递增子序列是 [3, 10, 20]，因此长度为 3。

3. SPOJ - Increasing Subsequences
给定一个整数数列，找出有多少个子序列是递增的。
例如：
Input：[5,2,8,6,3,6,9,7]
Output：8
Explanation：
递增子序列分别为 [5], [2], [8], [6], [3], [6], [9], [7]。

4. CodeChef - LISAND - Longest Increasing Subsequence and Non-decreasing Array
给定一个整数数列，找到其中的最长递增子序列的长度，同时也需要找到能够使数列变成非递减的最少操作次数。
例如：
Input：[2, 2, 2, 1, 1, 1, 3, 3, 3]
Output：2 3 
Explanation：
最长的递增子序列是 [2, 3]，长度为 2。最少需要执行 3 次操作才能使数列变成非递减的：将第 4, 5, 6 个数变成 2。

5. Hackerrank - The Longest Increasing Subsequence of an Array
给定一个整数数列，找到其中的最长递增子序列，同时需要找到所有符合该要求的最长子序列中，字典序最小的那个。
例如：
Input：[0, 8, 4, 12, 2, 10, 6, 14, 1, 9]
Output：[0, 4, 6, 9, 14]
Explanation：
最长的递增子序列是 [0, 2, 6, 9, 14] 和 [0, 4, 6, 9, 14]，但是前者字典序比后者大，因此输出后者。