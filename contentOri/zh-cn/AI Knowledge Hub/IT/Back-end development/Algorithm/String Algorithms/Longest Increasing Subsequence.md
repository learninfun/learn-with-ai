+++
title = "最长上升子序列"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Longest Increasing Subsequence的问题
{{< /ask_chatgpt >}}



1. 给定一个整数序列，求其中最长的递增子序列的长度是多少？

2. 给定一个由字母组成的字符串，求其中最长的递增子序列的长度是多少？

3. 给定一个浮点数序列，求其中最长的递增子序列的和是多少？

4. 给定一个二维整数矩阵，求其中最长的递增子序列的和是多少？

5. 给定一个由整数构成的有向图，求其中最长的递增子序列的长度是多少？   

## 说明知识
{{< ask_chatgpt >}}
说明Longest Increasing Subsequence并举例
{{< /ask_chatgpt >}}



Longest Increasing Subsequence（最长上升子序列）是指在一个序列中找到一个子序列，使得该子序列中的元素依次递增，且该子序列的长度最长。例如，序列 [3, 5, 2, 7, 1, 4, 8] 的最长上升子序列为 [3, 5, 7, 8]。

求最长上升子序列的常见方法是使用动态规划，以下为一个简单的动态规划算法：

1. 创建一个长度为n的数组dp，其中dp[i]表示以第i个元素为结尾的最长上升子序列的长度。
2. 初始化dp数组为1，因为每个元素本身就是一个最长上升子序列。
3. 遍历序列中的所有元素，对于每个元素i，从第一个元素到i-1元素中找到比i小的元素j，且dp[j]+1大于dp[i]，则更新dp[i]为dp[j]+1。
4. 最后遍历dp数组，找到最大的dp[i]，即为该序列的最长上升子序列的长度。

例如，对于序列 [3, 5, 2, 7, 1, 4, 8]，可以使用上述算法求解最长上升子序列。具体步骤如下：

1. 创建一个长度为7的dp数组，初始化所有元素为1。
2. 对于第一个元素3，dp[0]为1。
3. 对于第二个元素5，dp[1]为2。
4. 对于第三个元素2，找不到比它小的元素，dp[2]仍然为1。
5. 对于第四个元素7，dp[3]为3。
6. 对于第五个元素1，找不到比它小的元素，dp[4]仍然为1。
7. 对于第六个元素4，找到比它小的元素3，且dp[3]+1大于dp[5]，因此更新dp[5]为4。
8. 对于第七个元素8，找到比它小的元素7，且dp[3]+1大于dp[6]，因此更新dp[6]为4。
9. 最后遍历dp数组，找到最大的元素为4，因此该序列的最长上升子序列为 [3, 5, 7, 8]，长度为4。   

## 汇总重点
{{< ask_chatgpt >}}
条列Longest Increasing Subsequence的重点
{{< /ask_chatgpt >}}



1. 定义：Longest Increasing Subsequence（LIS）指数列中最长的递增子序列。例如，对于数列{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}，其LIS为{1, 2, 5, 6, 9}，长度为5。

2. LIS的求解方法：常用的求解LIS的方法有暴力搜索、动态规划、贪心、二分查找等。

3. 动态规划法：将LIS的求解问题分解成子问题，使用动态规划求解，时间复杂度为O(n^2)。

4. 贪心法：通过选择适当的元素构造LIS，时间复杂度为O(nlogn)。

5. 二分查找法：通过维护一个升序序列，并不断更新，来求解LIS，时间复杂度为O(nlogn)。

6. 应用：LIS问题广泛应用于各种实际问题中，如最长上升子序列、最长下降子序列、最长不下降子序列等。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Longest Increasing Subsequence的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. LeetCode 300 - Longest Increasing Subsequence
給定一個未排序的整數數組，找到最長遞增子序列的長度。
例如：
Input：[10,9,2,5,3,7,101,18]
Output：4
Explanation：
最長的遞增子序列是 [2,3,7,101]，因此長度為 4。

2. GeeksforGeeks - Longest Increasing Subsequence
給定一個整數數列，找到其中的最長遞增子序列。
例如：
Input：[3, 10, 2, 1, 20]
Output：3
Explanation：
最長的遞增子序列是 [3, 10, 20]，因此長度為 3。

3. SPOJ - Increasing Subsequences
給定一個整數數列，找出有多少個子序列是遞增的。
例如：
Input：[5,2,8,6,3,6,9,7]
Output：8
Explanation：
遞增子序列分別為 [5], [2], [8], [6], [3], [6], [9], [7]。

4. CodeChef - LISAND - Longest Increasing Subsequence and Non-decreasing Array
給定一個整數數列，找到其中的最長遞增子序列的長度，同時也需要找到能夠使數列變成非遞減的最少操作次數。
例如：
Input：[2, 2, 2, 1, 1, 1, 3, 3, 3]
Output：2 3 
Explanation：
最長的遞增子序列是 [2, 3]，長度為 2。最少需要執行 3 次操作才能使數列變成非遞減的：將第 4, 5, 6 個數變成 2。

5. Hackerrank - The Longest Increasing Subsequence of an Array
給定一個整數數列，找到其中的最長遞增子序列，同時需要找到所有符合該要求的最長子序列中，字典序最小的那個。
例如：
Input：[0, 8, 4, 12, 2, 10, 6, 14, 1, 9]
Output：[0, 4, 6, 9, 14]
Explanation：
最長的遞增子序列是 [0, 2, 6, 9, 14] 和 [0, 4, 6, 9, 14]，但是前者字典序比後者大，因此輸出後者。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Longest Increasing Subsequence的网络数据
{{< /ask_chatgpt >}}



1. "Longest Increasing Subsequence" by GeeksForGeeks
Link: https://www.geeksforgeeks.org/longest-increasing-subsequence-dp-3/

This article explains the concept of Longest Increasing Subsequence (LIS) and provides a dynamic programming solution to find it. It includes a detailed explanation of the algorithm and also provides time and space complexity analysis with examples.

2. "Longest Increasing Subsequence" by TopCoder
Link: https://www.topcoder.com/thrive/articles/Longest%20Increasing%20Subsequence

This article provides an introduction to the problem of Longest Increasing Subsequence and various approaches to solve it, including brute force, dynamic programming, and binary search. It also discusses the time complexity of each approach and their advantages and disadvantages.

3. "Optimal Subsequence Algorithm for Longest Increasing Subsequence" by Codeforces
Link: https://codeforces.com/blog/entry/13225

This article introduces a new algorithm for finding the Longest Increasing Subsequence, which has a faster execution time than other well-known algorithms. It includes a detailed explanation of the algorithm and also discusses time complexity analysis with examples.

4. "An O(n log n) Algorithm for Longest Increasing Subsequence" by National Academy of Sciences
Link: https://www.pnas.org/content/98/5/2286

This article presents an O(n log n) algorithm for finding the Longest Increasing Subsequence. It discusses the algorithm in detail, including its time complexity analysis and optimization techniques used in the implementation.

5. "Maximum Increasing Subsequence" by University of California, Berkeley
Link: https://www2.cs.duke.edu/courses/cps100e/fall09/notes/longest_increasing_subsequence.pdf

This article presents an approach to find the Maximum Increasing Subsequence, which is a variant of the classic Longest Increasing Subsequence problem. It describes the algorithm, provides time complexity analysis, and includes examples to illustrate the algorithm.   

