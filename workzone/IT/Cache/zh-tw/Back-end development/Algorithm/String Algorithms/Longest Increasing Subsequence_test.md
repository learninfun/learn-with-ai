

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