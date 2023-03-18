+++
title = "動態規劃"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Dynamic Programming的問題
{{< /ask_chatgpt >}}



1. 矩陣連乘問題
2. 最長遞增子序列問題
3. 背包問題（分為 01 背包和完全背包）
4. 硬幣找零問題
5. 圖的最短路徑問題（如 Dijkstra 算法）   

## 說明知識
{{< ask_chatgpt >}}
說明Dynamic Programming並舉例
{{< /ask_chatgpt >}}



Dynamic Programming 是一種以分治法為基礎的演算法設計技術，通常應用在需要最佳化（最大值或最小值）的問題上，並且這個問題可以分解成多個子問題，這些子問題的解可以共用之前計算得到的結果，以節省計算時間。

常常使用的技術是，計算出每個子問題的最優解，儲存以便稍後的遞歸或其他處理過程中重複計算的時候使用，透過這樣的優化，可以讓計算總時間從指數級別大幅降低至多項式級別或更低的複雜度。

舉例來說，Fibonacci數列就可以使用 Dynamic Programming 來計算。假如我們想求解 Fibonacci 數列的第 n 項，可以先計算出前面的 n-1 和 n-2 項，然後再利用這兩項的值相加得到第 n 項的值。

以此類推，每次計算一項 Fibonacci 數列的值時，都可以使用前一項和前兩項的值，在遞歸運算時可以透過快取已經計算過的結果，透過這樣的方式，可以將 Fibonacci 數列的計算複雜度從 $O(2^n)$ 優化到 $O(n^2)$。

另一個例子是背包問題（Knapsack Problem）。假設有一個能容納重量為 W 的背包，並且有 n 個物品，每個物品的重量和價值不同。問題是如何從這些物品中選擇一部分放入背包中，使得背包裝入物品的總價值最大。

使用 Dynamic Programming 解決背包問題的一種方法是，先計算出只考慮前 i 個物品而容量為 j 的背包的最優解，再依此方法求出最終問題的最優解。透過剩餘的空間大小和已經考慮了哪些物品，當前背包的最優解會慢慢計算出來，最後得到的就是整個背包問題的最優解。這種方法的時間複雜度是 $O(nW)$，其中 n 是物品數量，W 是背包的容量。   

## 彙總重點
{{< ask_chatgpt >}}
條列Dynamic Programming的重點
{{< /ask_chatgpt >}}



1. 動態規劃是以最佳化問題為基礎的一種演算法。

2. 動態規劃可以用來解決許多複雜的問題，例如最長共同子序列問題、背包問題、最小花費問題等。

3. 動態規劃可以分為兩個階段：建立解空間和尋找最優解。

4. 動態規劃的解決方案通常需要使用到遞迴式，根據問題的具體情況，可能需要使用到不同的遞迴式。

5. 動態規劃的主要特點是利用已知的問題解來解決新的問題，避免了重複計算的工作。

6. 動態規劃的時間複雜度通常是較高的，但可以通過優化遞迴式和使用記憶化搜索等方法來提高效率。

7. 動態規劃的應用非常廣泛，包括圖像處理、自然語言處理、金融建模等方面。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Dynamic Programming的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 最長公共子序列
給定兩個字串s1和s2，找到它們的最長公共子序列的長度。

答案：https://leetcode.com/problems/longest-common-subsequence/

2. 不同的二叉搜索樹
給定整數n，求由1到n形成的所有不同的二叉搜索樹的數量。

答案：https://leetcode.com/problems/unique-binary-search-trees/

3. 最長上升子序列
給定一個整數序列，找到其中的一個最長的子序列，使得它是嚴格上升的。

答案：https://leetcode.com/problems/longest-increasing-subsequence/

4. 最長回文子序列
給定一個字串，找到其中的一個最長的回文子序列。

答案：https://leetcode.com/problems/longest-palindromic-subsequence/

5. 最大子陣列和
給定一個整數數組，找到其中和最大的子陣列。

答案：https://leetcode.com/problems/maximum-subarray/   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Dynamic Programming的網路資料
{{< /ask_chatgpt >}}



1. "Dynamic Programming" by GeeksforGeeks - https://www.geeksforgeeks.org/dynamic-programming/

This comprehensive guide by GeeksforGeeks covers all aspects of dynamic programming, from its definition, types, characteristics, and implementation. It also includes examples of dynamic programming problems and their solutions.

2. "Introduction to Dynamic Programming - Programiz" - https://www.programiz.com/dsa/dynamic-programming

Programiz's beginner-friendly guide to dynamic programming provides a detailed explanation of the concept, its use cases, and how it works. It includes sample problems and step-by-step solutions to help readers understand the process.

3. "Dynamic Programming - Tutorialspoint" - https://www.tutorialspoint.com/dynamic_programming/index.htm

Tutorialspoint's dynamic programming guide explains dynamic programming in simple terms and provides examples of its applications, including the Fibonacci sequence, the Knapsack problem, and the Longest Common Substring problem.

4. "Dynamic Programming Explained: A Comprehensive Introduction" by Towards Data Science - https://towardsdatascience.com/dynamic-programming-explained-8d575c840079

This article on Toward Data Science provides an in-depth explanation of dynamic programming, its common problems, and how it can be applied to machine learning and data science. It also includes sample problems and solutions.

5. "Dynamic Programming: An Optimization Technique for Solving Complex Problems" by UpGrad - https://www.upgrad.com/blog/dynamic-programming/

UpGrad's dynamic programming guide provides a comprehensive overview of the concept, its use cases, and how it can help solve complex problems. It includes examples of dynamic programming problems and their solutions, as well as tips on how to optimize performance.   

