+++
title = "程式设计"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Programming的问题
{{< /ask_chatgpt >}}

1. 實作一個簡單的To-Do List應用程式，能夠新增、編輯、刪除待辦事項，以及將完成的事項標記為已完成。
2. 實作一個程式，計算一個文字檔中出現頻率最高的單字，並顯示該單字出現的次數。
3. 實作一個簡單的購物車應用程式，能夠新增、編輯、刪除商品，以及計算總價格。
4. 實作一個程式，能夠查詢指定月份的天數，例如輸入2月，回傳28或29。
5. 實作一個程式，能夠將一個整數轉換為字串，不得使用內建函數。   

## 说明知识
{{< ask_chatgpt >}}
说明Programming并举例
{{< /ask_chatgpt >}}

Programming是指使用特定的程序語言來創建、設計和實現軟件和應用程序的過程。編程的目的是讓電腦能夠執行特定的任務或操作。 

例如，當我們使用編程來創建一個網站，我們需要使用HTML、CSS和JavaScript等程序語言來設計和編寫網站的視覺效果和交互功能。在這個過程中，我們需要創建關於網站內容和結構的代碼。最終，當一個用戶訪問這個網站時，他們能夠看到和互動我們設計好的網站頁面。

其他的程式設計例如：建立行動應用程式、數據科學、網絡安全、人工智能和機器人等等。無論在哪個領域，編程都是講一個問題分解成一組數字計算，並以程式設計的方式實現解決方案的過程。   

## 汇总重点
{{< ask_chatgpt >}}
条列Programming的重点
{{< /ask_chatgpt >}}



1. 编程语言和工具：学习并掌握主要的编程语言和开发工具，例如Java、Python、C++、Visual Studio等。

2. 算法和数据结构：了解各种基本算法和数据结构，例如排序、搜索、链表、数组等等。

3. 面向对象程序设计（OOP）：实现OOP的基本概念，例如类、对象、继承、多态等等。

4. 效能优化：学习如何优化代码，并掌握对效能进行测量和分析的技能。

5. 版本控制：了解版本控制的基本概念和工具，例如Git和SVN，以便更好地管理代码和合作项目。

6. 软件设计和架构：熟悉如何设计和实现大型软件项目，包括分层设计、资料库架构、API设计等等。

7. Web开发：掌握Web开发的相关技能，包括HTML、CSS、JavaScript、PHP等等。

8. 资料库：了解SQL和NoSQL资料库的基本原理，以及如何在自己的代码中使用它们。

9. 测试：学习如何实施单元测试、集成测试和自动化测试，以确保代码的正确性和可重用性。

10. 问题解决和Debugging：掌握解决问题和Debugging的技巧，并熟悉各种工具和技术，以便快速定位和修复错误。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Programming的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Programming的网络数据
{{< /ask_chatgpt >}}



1. Programming Language Basics - https://www.geeksforgeeks.org/programming-language-basics/
This tutorial from GeeksforGeeks covers the basics of programming languages, including compilers, interpreters, syntax, variables, data types, operators, loops, and conditional statements.

2. Learn Python - https://www.learnpython.org/
This website offers a free, interactive introduction to the Python programming language. Topics covered include syntax, data types, loops, functions, and modules.

3. Java Programming Tutorial - https://www.javatpoint.com/java-tutorial
This tutorial from javatpoint.com provides an indepth tutorial of the Java programming language. It covers topics such as object-oriented programming, exceptions handling, collections, and generics.

4. C++ Language - https://www.tutorialspoint.com/cplusplus/index.htm
This tutorial from TutorialsPoint covers the C++ programming language. Topics covered include data types, operators, arrays, functions, and classes.

5. Web development tutorial - https://www.w3schools.com/
This website covers the basics of web development, including HTML, CSS, JavaScript, SQL, PHP, and ASP. It also offers examples and quizzes to test your knowledge.   

