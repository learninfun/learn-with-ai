+++
title = "記憶化搜尋"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Memoization的問題
{{< /ask_chatgpt >}}



1. 利用 Memoization 優化斐波那契數列的計算。

2. 計算兩個字符串之間的編輯距離（編輯距離指的是兩個字符串之間需要進行多少次增刪改操作，才能使它們相等），並利用 Memoization 最小化計算次數。

3. 實現一個函數，該函數輸入一個正整數 n，並返回該整數的所有因數之和。使用 Memoization 避免重複計算。

4. 給定一個正整數 n，求出對於某個 k 值，n 的 k 次方的值（即 n 的 k 次方）。使用 Memoization 減少計算次數。

5. 計算在一個二維矩陣中，從左上角到右下角的最短路徑。矩陣中每個格子的值代表該位置上的代價，行走代價較小的路徑為最終路徑。利用 Memoization 避免冗餘計算。   

## 說明知識
{{< ask_chatgpt >}}
說明Memoization並舉例
{{< /ask_chatgpt >}}



Memoization是一種常用的技術，它可以用來優化重複計算的過程。簡單來說，Memoization就是在計算結果後將其存儲起來，以便在未來需要時可以快速查找。這種方式能夠減少計算的時間複雜度，提高程式效率。

下面舉一個簡單的例子：

假設我們要計算費氏數列中第n個數，費氏數列的定義為：數列中的第一和第二個數都是1，從第三個數開始，每個數都是前面兩個數的和。也就是說，費氏數列的前幾項為：1, 1, 2, 3, 5, 8, 13, ...

我們可以使用遞迴來計算費氏數列中的第n個數，具體實現如下：

```
def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)
```

使用這個函數來計算費氏數列中第n個數的值時，當n較大時會出現計算時間過長的情況。因為計算過程中需要重複計算一些值，比如計算費氏數列中的第10項時，需要先計算出第9項和第8項，計算第9項時又需要先計算第8項和第7項，而計算第8項時還需要計算第7項和第6項，這樣就會浪費大量的計算時間。

使用Memoization可以避免這些重複計算的問題。具體實現方式如下：

```
cache = {}  # 定義一個字典，用來存儲計算結果

def fibonacci(n):
    if n < 2:
        return n
    if n in cache:  # 如果計算結果已經存在於cache中，直接返回該結果
        return cache[n]
    result = fibonacci(n-1) + fibonacci(n-2)
    cache[n] = result  # 將計算結果存儲到cache中
    return result
```

使用這個改進版的函數來計算費氏數列中的第n個數時，計算時間可以大幅降低。因為當需要計算一個已經計算過的值時，可以直接從cache中查找得到，而不需要重新計算。這樣就可以減少重複計算的次數，提高程式效率。   

## 彙總重點
{{< ask_chatgpt >}}
條列Memoization的重點
{{< /ask_chatgpt >}}



1. Memoization是一種常見的優化技術，通常用於優化計算過程相對昂貴的函數。 

2. 該技術通過記憶已經計算過的結果，避免重複執行相同的計算，從而提高程序的效率。 

3. Memoization的實現方式包括使用散列表、緩存、緩存注入等方法。 

4. 適用情況通常包括需要重複執行相同計算的場景，例如遞迴運算和複雜的算法問題等。 

5. 注意事項包括緩存大小的設計、緩存溢出的處理、並發環境下的緩存同步等。 

6. 需要注意的是，Memoization在某些情況下可能會導致錯誤的結果，因此需要根據具體的場景進行適當的評估和測試。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Memoization的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 爬樓梯問題
假設有一階樓梯，你可以一次爬一格或兩格，求爬上n階樓梯有幾種不同的方式。

答案：
```python
def climbStairs(n: int) -> int:
    memo = [0] * (n + 1)
    memo[0], memo[1] = 1, 1
    for i in range(2, n + 1):
        memo[i] = memo[i - 1] + memo[i - 2]
    return memo[n]
```

2. 切繩子問題
有一條長度為n的繩子，你可以把它剪成任意多段，但每段長度必須是整數，並且至少剪一刀，求剪後各段乘積的最大值。

答案：
```python
def max_product_after_cutting(n: int) -> int:
    memo = [0] * (n + 1)
    memo[2] = 1
    for i in range(3, n + 1):
        for j in range(1, i):
            memo[i] = max(memo[i], max(j * memo[i - j], j * (i - j)))
    return memo[n]
```

3. Edit Distance問題
有兩個字符串s1和s2，可以進行三種操作：插入一個字符、刪除一個字符、替換一個字符，求s1轉化成s2所需的最少操作次數。

答案：
```python
def min_distance(s1: str, s2: str) -> int:
    memo = [[0 for _ in range(len(s2) + 1)] for _ in range(len(s1) + 1)]
    for i in range(len(s1) + 1):
        memo[i][0] = i
    for j in range(len(s2) + 1):
        memo[0][j] = j
    for i in range(1, len(s1) + 1):
        for j in range(1, len(s2) + 1):
            if s1[i - 1] == s2[j - 1]:
                memo[i][j] = memo[i - 1][j - 1]
            else:
                memo[i][j] = 1 + min(memo[i - 1][j], memo[i][j - 1], memo[i - 1][j - 1])
    return memo[-1][-1]
```

4. 跳躍遊戲問題
有一個非負整數的數列，每個數表示在該位置上最多可以向前跳躍該數字個單位，請求是否可以到達最後一個位置。

答案：
```python
def can_jump(nums: List[int]) -> bool:
    memo = [0] * len(nums)
    memo[0] = nums[0]
    for i in range(1, len(nums)):
        if memo[i - 1] < i:
            return False
        memo[i] = max(memo[i - 1], i + nums[i])
    return True
```

5. 費波那契數列問題
請列出費波那契數列的前n項。

答案：
```python
def fib(n: int) -> int:
    if n == 0:
        return 0
    memo = [0] * (n + 1)
    memo[1] = 1
    for i in range(2, n + 1):
        memo[i] = memo[i - 1] + memo[i - 2]
    return memo[n]
```   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Memoization的網路資料
{{< /ask_chatgpt >}}



1. "Memoization in Python: How to Implement and Use Memoization"
Link: https://www.freecodecamp.org/news/python-memoization-how-to-implement-and-use-memoization/

This article provides a beginner-friendly explanation of memoization and, more specifically, how to use memoization in Python. It covers the concept of recursion, which can be made more efficient with memoization, and includes a code example to help readers understand how to implement memoization in their own projects.

2. "Understanding Memoization in JavaScript"
Link: https://medium.com/@siddharthac6/understanding-memoization-in-javascript-context-usage-and-pitfalls-a45b5b5116cd

This article provides a detailed explanation of memoization, including its benefits and limitations, and how to use it in JavaScript. It covers how memoization helps to optimize slow functions and includes a code example to demonstrate its implementation.

3. "Memoization in Java"
Link: https://www.geeksforgeeks.org/memoization-in-java/

This article explains how to implement memoization in Java. It covers the concept of recursion and how it can be made more efficient with memoization. The article also includes a code example to make it easier for readers to implement memoization in their own Java projects.

4. "Memoization in C++"
Link: https://www.geeksforgeeks.org/memoization-in-c-template/

This article provides a detailed explanation of memoization in C++. It covers the basic concept of memoization, including how it works and its benefits. The article also provides a C++ code example for readers to follow and understand how to implement memoization in their own C++ projects.

5. "An Introduction to Memoization and Dynamic Programming in Go"
Link: https://kylewbanks.com/blog/memoization-and-dynamic-programming-in-golang

This article provides an introduction to memoization and dynamic programming in Go. It covers the concept of memoization, why it is useful, and includes a code example to help readers understand how to implement memoization in their own Go projects. The article also covers the basics of dynamic programming and how memoization can be used as part of dynamic programming.   

