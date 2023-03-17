

1. 題目: 在一個長度為n的整數陣列中，查找是否存在三個元素a, b, c，其值滿足a + b + c = 0？如果存在，找到所有滿足該條件的三元組，並返回。

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

2. 題目: 給定n個非負整數，其中每個數代表一個垂直線上的點。將n個點連接起來，找出可以容納最多水的兩個線段，並返回其面積。

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

3. 題目: 給定n個硬幣，每個硬幣都有一個數字值。將硬幣分成兩組，使得每組硬幣的數字值之和相等。找到這樣的兩組硬幣，並返回它們的數組下標。

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

4. 題目: 給定一個整數數組nums和一個整數k，找出是否存在兩個不同的索引i和j，使得nums [i] = nums [j]和i和j之間的絕對差最大至多為k。

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

5. 題目: 給定一個完美散列表，該散列表允許O(1)時間查找特定元素的存在性。該散列表使用的碰撞解決方案是鏈式解決方案，即如果兩個元素在該位置發生碰撞，則將它們放在同一鏈表中。現在，您要根據存在於兩個完美散列表中的相同的值來決定是否為共享索引對的一部分。

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