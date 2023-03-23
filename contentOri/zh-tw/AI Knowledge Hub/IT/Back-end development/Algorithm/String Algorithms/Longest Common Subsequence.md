+++
title = "最長共同子序列"
weight = "5"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Longest Common Subsequence的中文問題
{{< /ask_chatgpt >}}



1. 給定兩個字串 "abcd" 和 "bdca"，求它們的最長公共子序列。
2. 給定兩個字串 "AGGTAB" 和 "GXTXAYB"，求它們的最長公共子序列。
3. 給定兩個字串 "ACACGAGT" 和 "AGCACGTGA"，求它們的最長公共子序列。
4. 給定兩個字串 "algorithm" 和 "logarithm"，求它們的最長公共子序列。
5. 給定兩個字串 "banana" 和 "atana"，求它們的最長公共子序列。   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Longest Common Subsequence並舉例
{{< /ask_chatgpt >}}



Longest Common Subsequence（LCS）是指：給定兩個序列，找到它們中都出現的最長子序列。其中，子序列是源序列中保持相對順序的元素子集(不一定相連)，而不需要完全連續。

例如：
序列X：A B C B D A B
序列Y：B D C A B A
它們的最長公共子序列為BCBA，長度為4。

在計算LCS時，通常使用動態規劃(Dynamic Programming)的方法解決。可以建立一個矩陣，從高左角開始，往右下角填充。

假設X序列長度為m，Y序列長度為n，LCS的長度為L[m][n]。
則，L[i][j]，代表X序列前i個元素和Y序列前j個元素的LCS長度：

L[i][j] = L[i-1][j-1] + 1 if X[i-1] = Y[j-1]
L[i][j] = max(L[i-1][j], L[i][j-1]) if X[i-1] != Y[j-1]

在填表時，如果X[i-1] = Y[j-1]，即可填L[i-1][j-1] + 1。否則，取L[i-1][j]和L[i][j-1]中的最大值。

最終，L[m][n]就是X與Y序列的最長公共子序列的長度。而其中包含的LCS元素可以通過反向追蹤表格而得到。

例如，對於上例中的序列，其LCS長度矩陣如下所示：

    B D C A B A
  0 0 0 0 0 0 0
A 0 0 0 0 0 1 1
B 0 1 1 1 1 1 2
C 0 1 2 2 2 2 2
B 0 1 2 2 3 3 3
D 0 1 2 2 3 3 3
A 0 1 2 2 3 4 4
B 0 1 2 2 3 4 5

最後，經過追蹤表格，得到的LCS為BCBA。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Longest Common Subsequence的重點
{{< /ask_chatgpt >}}



1. Longest Common Subsequence（LCS）是兩個序列中最長的公共子序列。
2. 一個LCS的元素在兩個序列中的相對位置必須與原序列中的順序相同。
3. LCS問題可以使用動態規劃來解決，時間複雜度為O（mn）（其中m和n為兩個序列的長度）。
4. 通過目前位置的字符如何匹配選擇移動（向右、向下、向右下）。
5. 可以使用回溯算法來找到LCS，時間複雜度為O（m+n+k）（其中k是LCS的長度）。
6. LCS問題在文本比較，生物資訊學，圖像識別等方面有廣泛的應用。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Longest Common Subsequence的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 給定兩個字串S1和S2，找出它們之間的最長公共子序列的長度。
S1："abcbdab"，
S2："bdcaba"
答案：4

2. 給定兩個長度相等的陣列A和B，找出它們之間的最長公共子序列的長度。其中，一個子序列可以是從陣列中刪除某些元素後所得到的序列。
A：[1, 2, 3, 4, 1, 2, 1, 2, 3, 4, 5]
B：[3, 4, 1, 2, 1, 2, 3, 4, 5, 1, 2]
答案：7

3. 給定兩個字符串S1和S2，找到它們之間的最長公共 sunsequence 中的字典序最小的子串。
S1： "abcde"
S2："eacbdf"
答案："acb"

4. 給定一個由N個單詞構成的序列，找到最長的序列，其中每個單詞都是前一個單詞的子序列。如果有多個相同長度的最長序列，請選擇字典序最小的那個。
["cat", "dog", "god", "at", "hat", "ate"]
答案：["cat", "at", "ate"]

5. 給定三個字串S1、S2和S3，找到它們之間的最長公共子序列。
S1： "AGGTAB"
S2："GXTXAYB"
S3："GTAB"
答案： "GTAB"   

