

1. 题目: 在一个长度为n的整数阵列中，查找是否存在三个元素a, b, c，其值满足a + b + c = 0？如果存在，找到所有满足该条件的三元组，并返回。

答案:

```python
def threeSum(nums):
    nums.sort()
    res = []
    for i in range(len(nums)-2):
        if i > 0 and nums[i] == nums[i-1]:
            continue
        l, r = i+1, len(nums)-1
        while l < r:
            s = nums[i] + nums[l] + nums[r]
            if s > 0:
                r -= 1
            elif s < 0:
                l += 1
            else:
                res.append([nums[i], nums[l], nums[r]])
                while l < r and nums[l] == nums[l+1]:
                    l += 1
                while l < r and nums[r] == nums[r-1]:
                    r -= 1
                l += 1
                r -= 1
    return res
```

2. 题目: 给定n个非负整数，其中每个数代表一个垂直线上的点。将n个点连接起来，找出可以容纳最多水的两个线段，并返回其面积。

答案:

```python
def maxArea(height):
    l, r =0, len(height)-1
    ans=0
    while (l < r):
        area= min(height[l], height[r]) * (r - l)
        ans = max(ans, area)
        if (height[l] < height[r]):
            l += 1
        else:
            r -= 1
    return ans
```

3. 题目: 给定n个硬币，每个硬币都有一个数字值。将硬币分成两组，使得每组硬币的数字值之和相等。找到这样的两组硬币，并返回它们的数组下标。

答案:

```python
def coinPartition(coins):
    total_sum = sum(coins)
    if total_sum % 2 != 0:
        return []
    target_sum = total_sum // 2
    dp = [[0] * (target_sum+1) for _ in range(len(coins)+1)]
    for i in range(1, len(coins)+1):
        for j in range(1, target_sum+1):
            if coins[i-1] > j:
                dp[i][j] = dp[i-1][j]
            else:
                dp[i][j] = max(dp[i-1][j], dp[i-1][j-coins[i-1]]+coins[i-1])
    if dp[-1][-1] != target_sum:
        return []
    res = []
    i, j = len(coins), target_sum
    while i > 0 and j >= 0:
        if dp[i][j] == dp[i-1][j]:
            i -= 1
        else:
            res.append(i-1)
            j -= coins[i-1]
            i -= 1
    return res[::-1]
```

4. 题目: 给定一个整数数组nums和一个整数k，找出是否存在两个不同的索引i和j，使得nums [i] = nums [j]和i和j之间的绝对差最大至多为k。

答案:

```python
def containsNearbyDuplicate(nums, k):
    d = {}
    for i in range(len(nums)):
        if nums[i] in d and i - d[nums[i]] <= k:
            return True
        d[nums[i]] = i
    return False
```

5. 题目: 给定一个完美散列表，该散列表允许O(1)时间查找特定元素的存在性。该散列表使用的碰撞解决方案是链式解决方案，即如果两个元素在该位置发生碰撞，则将它们放在同一链表中。现在，您要根据存在于两个完美散列表中的相同的值来决定是否为共享索引对的一部分。

答案:

```python
def sharedIndexes(h1, h2):
    res = []
    for i in range(len(h1)):
        if h1[i]:
            node = h1[i]
            while node:
                if h2[i] and node.val in h2[i]:
                    res.append(i)
                    break
                node = node.next
    return res
```