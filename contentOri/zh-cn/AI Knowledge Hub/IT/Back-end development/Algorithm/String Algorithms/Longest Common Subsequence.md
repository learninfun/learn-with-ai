+++
title = "最长共同子序列"
weight = "5"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Longest Common Subsequence的问题
{{< /ask_chatgpt >}}



1. 给定两个字串 "abcd" 和 "bdca"，求它们的最长公共子序列。
2. 给定两个字串 "AGGTAB" 和 "GXTXAYB"，求它们的最长公共子序列。
3. 给定两个字串 "ACACGAGT" 和 "AGCACGTGA"，求它们的最长公共子序列。
4. 给定两个字串 "algorithm" 和 "logarithm"，求它们的最长公共子序列。
5. 给定两个字串 "banana" 和 "atana"，求它们的最长公共子序列。   

## 说明知识
{{< ask_chatgpt >}}
说明Longest Common Subsequence并举例
{{< /ask_chatgpt >}}



Longest Common Subsequence（LCS）是指：给定两个序列，找到它们中都出现的最长子序列。其中，子序列是源序列中保持相对顺序的元素子集(不一定相连)，而不需要完全连续。

例如：
序列X：A B C B D A B
序列Y：B D C A B A
它们的最长公共子序列为BCBA，长度为4。

在计算LCS时，通常使用动态规划(Dynamic Programming)的方法解决。可以建立一个矩阵，从高左角开始，往右下角填充。

假设X序列长度为m，Y序列长度为n，LCS的长度为L[m][n]。
则，L[i][j]，代表X序列前i个元素和Y序列前j个元素的LCS长度：

L[i][j] = L[i-1][j-1] + 1 if X[i-1] = Y[j-1]
L[i][j] = max(L[i-1][j], L[i][j-1]) if X[i-1] != Y[j-1]

在填表时，如果X[i-1] = Y[j-1]，即可填L[i-1][j-1] + 1。否则，取L[i-1][j]和L[i][j-1]中的最大值。

最终，L[m][n]就是X与Y序列的最长公共子序列的长度。而其中包含的LCS元素可以通过反向追踪表格而得到。

例如，对于上例中的序列，其LCS长度矩阵如下所示：

    B D C A B A
  0 0 0 0 0 0 0
A 0 0 0 0 0 1 1
B 0 1 1 1 1 1 2
C 0 1 2 2 2 2 2
B 0 1 2 2 3 3 3
D 0 1 2 2 3 3 3
A 0 1 2 2 3 4 4
B 0 1 2 2 3 4 5

最后，经过追踪表格，得到的LCS为BCBA。   

## 汇总重点
{{< ask_chatgpt >}}
条列Longest Common Subsequence的重点
{{< /ask_chatgpt >}}



1. Longest Common Subsequence（LCS）是两个序列中最长的公共子序列。
2. 一个LCS的元素在两个序列中的相对位置必须与原序列中的顺序相同。
3. LCS问题可以使用动态规划来解决，时间复杂度为O（mn）（其中m和n为两个序列的长度）。
4. 通过目前位置的字符如何匹配选择移动（向右、向下、向右下）。
5. 可以使用回溯算法来找到LCS，时间复杂度为O（m+n+k）（其中k是LCS的长度）。
6. LCS问题在文本比较，生物资讯学，图像识别等方面有广泛的应用。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Longest Common Subsequence的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 给定两个字串S1和S2，找出它们之间的最长公共子序列的长度。
S1："abcbdab"，
S2："bdcaba"
答案：4

2. 给定两个长度相等的阵列A和B，找出它们之间的最长公共子序列的长度。其中，一个子序列可以是从阵列中删除某些元素后所得到的序列。
A：[1, 2, 3, 4, 1, 2, 1, 2, 3, 4, 5]
B：[3, 4, 1, 2, 1, 2, 3, 4, 5, 1, 2]
答案：7

3. 给定两个字符串S1和S2，找到它们之间的最长公共 sunsequence 中的字典序最小的子串。
S1： "abcde"
S2："eacbdf"
答案："acb"

4. 给定一个由N个单词构成的序列，找到最长的序列，其中每个单词都是前一个单词的子序列。如果有多个相同长度的最长序列，请选择字典序最小的那个。
["cat", "dog", "god", "at", "hat", "ate"]
答案：["cat", "at", "ate"]

5. 给定三个字串S1、S2和S3，找到它们之间的最长公共子序列。
S1： "AGGTAB"
S2："GXTXAYB"
S3："GTAB"
答案： "GTAB"   

