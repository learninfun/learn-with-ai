

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