+++
title = "记忆化搜寻"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Memoization的问题
{{< /ask_chatgpt >}}



1. 利用 Memoization 優化斐波那契數列的計算。

2. 計算兩個字符串之間的編輯距離（編輯距離指的是兩個字符串之間需要進行多少次增刪改操作，才能使它們相等），並利用 Memoization 最小化計算次數。

3. 實現一個函數，該函數輸入一個正整數 n，並返回該整數的所有因數之和。使用 Memoization 避免重複計算。

4. 給定一個正整數 n，求出對於某個 k 值，n 的 k 次方的值（即 n 的 k 次方）。使用 Memoization 減少計算次數。

5. 計算在一個二維矩陣中，從左上角到右下角的最短路徑。矩陣中每個格子的值代表該位置上的代價，行走代價較小的路徑為最終路徑。利用 Memoization 避免冗餘計算。   

## 说明知识
{{< ask_chatgpt >}}
说明Memoization并举例
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

## 汇总重点
{{< ask_chatgpt >}}
条列Memoization的重点
{{< /ask_chatgpt >}}



1. Memoization是一种常见的优化技术，通常用于优化计算过程相对昂贵的函数。 

2. 该技术通过记忆已经计算过的结果，避免重复执行相同的计算，从而提高程序的效率。 

3. Memoization的实现方式包括使用散列表、缓存、缓存注入等方法。 

4. 适用情况通常包括需要重复执行相同计算的场景，例如递回运算和复杂的算法问题等。 

5. 注意事项包括缓存大小的设计、缓存溢出的处理、并发环境下的缓存同步等。 

6. 需要注意的是，Memoization在某些情况下可能会导致错误的结果，因此需要根据具体的场景进行适当的评估和测试。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Memoization的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 爬楼梯问题
假设有一阶楼梯，你可以一次爬一格或两格，求爬上n阶楼梯有几种不同的方式。

答案：
```python
def climbStairs(n: int) -> int:
    memo = [0] * (n + 1)
    memo[0], memo[1] = 1, 1
    for i in range(2, n + 1):
        memo[i] = memo[i - 1] + memo[i - 2]
    return memo[n]
```

2. 切绳子问题
有一条长度为n的绳子，你可以把它剪成任意多段，但每段长度必须是整数，并且至少剪一刀，求剪后各段乘积的最大值。

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

3. Edit Distance问题
有两个字符串s1和s2，可以进行三种操作：插入一个字符、删除一个字符、替换一个字符，求s1转化成s2所需的最少操作次数。

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

4. 跳跃游戏问题
有一个非负整数的数列，每个数表示在该位置上最多可以向前跳跃该数字个单位，请求是否可以到达最后一个位置。

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

5. 费波那契数列问题
请列出费波那契数列的前n项。

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

