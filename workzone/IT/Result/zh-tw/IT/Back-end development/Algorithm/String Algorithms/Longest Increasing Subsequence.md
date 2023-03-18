+++
title = "最長上升子序列"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Longest Increasing Subsequence的問題
{{< /ask_chatgpt >}}



1. 給定一個整數序列，求其中最長的遞增子序列的長度是多少？

2. 給定一個由字母組成的字符串，求其中最長的遞增子序列的長度是多少？

3. 給定一個浮點數序列，求其中最長的遞增子序列的和是多少？

4. 給定一個二維整數矩陣，求其中最長的遞增子序列的和是多少？

5. 給定一個由整數構成的有向圖，求其中最長的遞增子序列的長度是多少？   

## 說明知識
{{< ask_chatgpt >}}
說明Longest Increasing Subsequence並舉例
{{< /ask_chatgpt >}}



Longest Increasing Subsequence（最長上升子序列）是指在一個序列中找到一個子序列，使得該子序列中的元素依次遞增，且該子序列的長度最長。例如，序列 [3, 5, 2, 7, 1, 4, 8] 的最長上升子序列為 [3, 5, 7, 8]。

求最長上升子序列的常見方法是使用動態規劃，以下為一個簡單的動態規划算法：

1. 創建一個長度為n的數組dp，其中dp[i]表示以第i個元素為結尾的最長上升子序列的長度。
2. 初始化dp數組為1，因為每個元素本身就是一個最長上升子序列。
3. 遍歷序列中的所有元素，對於每個元素i，從第一個元素到i-1元素中找到比i小的元素j，且dp[j]+1大於dp[i]，則更新dp[i]為dp[j]+1。
4. 最後遍歷dp數組，找到最大的dp[i]，即為該序列的最長上升子序列的長度。

例如，對於序列 [3, 5, 2, 7, 1, 4, 8]，可以使用上述算法求解最長上升子序列。具體步驟如下：

1. 創建一個長度為7的dp數組，初始化所有元素為1。
2. 對於第一個元素3，dp[0]為1。
3. 對於第二個元素5，dp[1]為2。
4. 對於第三個元素2，找不到比它小的元素，dp[2]仍然為1。
5. 對於第四個元素7，dp[3]為3。
6. 對於第五個元素1，找不到比它小的元素，dp[4]仍然為1。
7. 對於第六個元素4，找到比它小的元素3，且dp[3]+1大於dp[5]，因此更新dp[5]為4。
8. 對於第七個元素8，找到比它小的元素7，且dp[3]+1大於dp[6]，因此更新dp[6]為4。
9. 最後遍歷dp數組，找到最大的元素為4，因此該序列的最長上升子序列為 [3, 5, 7, 8]，長度為4。   

## 彙總重點
{{< ask_chatgpt >}}
條列Longest Increasing Subsequence的重點
{{< /ask_chatgpt >}}



1. 定義：Longest Increasing Subsequence（LIS）指數列中最長的遞增子序列。例如，對於數列{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}，其LIS為{1, 2, 5, 6, 9}，長度為5。

2. LIS的求解方法：常用的求解LIS的方法有暴力搜索、動態規劃、貪心、二分查找等。

3. 動態規劃法：將LIS的求解問題分解成子問題，使用動態規劃求解，時間複雜度為O(n^2)。

4. 貪心法：通過選擇適當的元素構造LIS，時間複雜度為O(nlogn)。

5. 二分查找法：通過維護一個升序序列，並不斷更新，來求解LIS，時間複雜度為O(nlogn)。

6. 應用：LIS問題廣泛應用於各種實際問題中，如最長上升子序列、最長下降子序列、最長不下降子序列等。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Longest Increasing Subsequence的中等難度問題，並在後面列出答案
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

## 網路資料
{{< ask_chatgpt >}}
給我5篇Longest Increasing Subsequence的網路資料
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

