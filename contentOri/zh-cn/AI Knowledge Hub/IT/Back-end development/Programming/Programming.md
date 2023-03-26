+++
title = "程式設計"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Programming的中文問題
{{< /ask_chatgpt >}}

1. 實作一個簡單的To-Do List應用程式，能夠新增、編輯、刪除待辦事項，以及將完成的事項標記為已完成。
2. 實作一個程式，計算一個文字檔中出現頻率最高的單字，並顯示該單字出現的次數。
3. 實作一個簡單的購物車應用程式，能夠新增、編輯、刪除商品，以及計算總價格。
4. 實作一個程式，能夠查詢指定月份的天數，例如輸入2月，回傳28或29。
5. 實作一個程式，能夠將一個整數轉換為字串，不得使用內建函數。   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Programming並舉例
{{< /ask_chatgpt >}}

Programming是指使用特定的程序語言來創建、設計和實現軟件和應用程序的過程。編程的目的是讓電腦能夠執行特定的任務或操作。 

例如，當我們使用編程來創建一個網站，我們需要使用HTML、CSS和JavaScript等程序語言來設計和編寫網站的視覺效果和交互功能。在這個過程中，我們需要創建關於網站內容和結構的代碼。最終，當一個用戶訪問這個網站時，他們能夠看到和互動我們設計好的網站頁面。

其他的程式設計例如：建立行動應用程式、數據科學、網絡安全、人工智能和機器人等等。無論在哪個領域，編程都是講一個問題分解成一組數字計算，並以程式設計的方式實現解決方案的過程。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Programming的重點
{{< /ask_chatgpt >}}



1. 編程語言和工具：學習並掌握主要的編程語言和開發工具，例如Java、Python、C++、Visual Studio等。

2. 算法和數據結構：瞭解各種基本算法和數據結構，例如排序、搜索、鏈表、數組等等。

3. 面向對像程序設計（OOP）：實現OOP的基本概念，例如類、對像、繼承、多態等等。

4. 效能優化：學習如何優化代碼，並掌握對效能進行測量和分析的技能。

5. 版本控制：瞭解版本控制的基本概念和工具，例如Git和SVN，以便更好地管理代碼和合作項目。

6. 軟件設計和架構：熟悉如何設計和實現大型軟件項目，包括分層設計、資料庫架構、API設計等等。

7. Web開發：掌握Web開發的相關技能，包括HTML、CSS、JavaScript、PHP等等。

8. 資料庫：瞭解SQL和NoSQL資料庫的基本原理，以及如何在自己的代碼中使用它們。

9. 測試：學習如何實施單元測試、集成測試和自動化測試，以確保代碼的正確性和可重用性。

10. 問題解決和Debugging：掌握解決問題和Debugging的技巧，並熟悉各種工具和技術，以便快速定位和修復錯誤。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Programming的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



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

