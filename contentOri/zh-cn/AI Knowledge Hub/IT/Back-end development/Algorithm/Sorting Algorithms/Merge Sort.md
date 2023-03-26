+++
title = "合并排序"
weight = "4"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Merge Sort的中文问题
{{< /ask_chatgpt >}}



1. 什么是 Merge Sort？它是如何工作的？
2. Merge Sort 是一种时间和空间效率如何？有何优缺点？
3. Merge Sort 在哪些场景下被广泛使用？能否举几个例子？
4. Merge Sort 的最坏时间复杂度是多少？它是如何达到的？
5. Merge Sort 和 Quick Sort 的异同有哪些？它们在哪些情况下能起到更好的作用？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Merge Sort并举例
{{< /ask_chatgpt >}}



Merge Sort是一种分治法排序演算法，它的基本概念是将待排序数列不断拆分为较小的子序列，直到每个子序列里面只有一个元素为止，再将相邻的子序列合并为更大的有序子序列，直到最后只剩下一个有序的序列。

以下是Merge Sort的步骤：

1. 分割：把待排序的序列拆分成两个子序列，每个子序列包含约一半的元素。 

2. 对每个子序列递归排序，直到每个子序列里面只剩下一个元素。 

3. 合并：把排序好的子序列合并成一个大的有序序列。 

举例来说，如果我们要对以下数列进行排序：

{38, 27, 43, 3, 9, 82, 10}

分割后会得到两个子序列：

{38, 27, 43, 3}

{9, 82, 10}

对每个子序列递归排序，最终得到排序好的子序列：

{3, 27, 38, 43}

{9, 10, 82}

最后将排序好的子序列合并，得到最终的有序序列：

{3, 9, 10, 27, 38, 43, 82}   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Merge Sort的重点
{{< /ask_chatgpt >}}



1. Divide and conquer: Merge Sort使用“分而治之”的策略，将一个大问题分解成较小的子问题，并且解决每个子问题，最后将它们合并成一个有序的问题。

2. Recursion: Merge Sort使用递归的形式实现分治法，每递归一层都会分解原问题并且排序子数组。

3. Merge: Merge Sort将两个已排序的子数组合并为一个已排序的数组，使得最终的数组保持正确的顺序。

4. Time complexity: Merge Sort的时间复杂度为nlogn，是一种高效的排序算法。

5. Space complexity: Merge Sort的空间复杂度是O(n)，因为必须在排序过程中创建一个与输入数组大小相同的临时数组。

6. Stability: Merge Sort是稳定的，因为它使用了相等的情况的规则，以保持同等大小的数字的相对位置。

7. Adaptive: Merge Sort也是自适应的，可以在原地排序的情况下使用。

8. Parallel: Merge Sort还有一个显著的特性，就是它很容易实现并发排序，因为它将大数组分成许多小的子数组进行排序，并且这些子数组可以并发处理   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Merge Sort的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 实现Merge Sort的一个变体，将所有小于等于pivot的元素放在左边，大于pivot的元素放在右边。
答案：以下是代码示例
```python
def partition(arr, low, high):
    pivot = arr[high]
    i = low - 1
    for j in range(low, high):
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i+1

def mergeSort(arr, low, high):
    if low < high:
        pi = partition(arr, low, high)
        mergeSort(arr, low, pi-1)
        mergeSort(arr, pi+1, high)
```

2. 给定一个整数数组，找到其中第K大的元素。
答案：以下是代码示例
```python
def mergeSort(arr):
    if len(arr) > 1:
        mid = len(arr) // 2
        L = arr[:mid]
        R = arr[mid:]

        mergeSort(L)
        mergeSort(R)

        i = j = k = 0

        while i < len(L) and j < len(R):
            if L[i] < R[j]:
                arr[k] = L[i]
                i += 1
            else:
                arr[k] = R[j]
                j += 1
            k += 1

        while i < len(L):
            arr[k] = L[i]
            i += 1
            k += 1

        while j < len(R):
            arr[k] = R[j]
            j += 1
            k += 1


def findKthLargest(nums, k):
    mergeSort(nums)
    return nums[-k]
```

3. 设计一个算法来对一个整数数组进行排序，其中一个数是重复的，且该重复数可能有多个位置。
答案：以下是代码示例
```python
def mergeSort(arr, l, r):
    if l < r:
        m = (l + r) // 2

        mergeSort(arr, l, m)
        mergeSort(arr, m+1, r)

        # merge the two sorted halves
        i = l
        j = m + 1
        k = 0
        temp = [0] * (r-l+1)

        while i <= m and j <= r:
            if arr[i] < arr[j]:
                temp[k] = arr[i]
                i += 1
            else:
                temp[k] = arr[j]
                j += 1
            k += 1

        while i <= m:
            temp[k] = arr[i]
            i += 1
            k += 1

        while j <= r:
            temp[k] = arr[j]
            j += 1
            k += 1

        for p in range(len(temp)):
            arr[l+p] = temp[p]

def findDuplicate(nums):
    mergeSort(nums, 0, len(nums)-1)
    for i in range(1, len(nums)):
        if nums[i] == nums[i-1]:
            return nums[i]
```

4. 给定一个整数数组和一个目标值，找到数组中三个数的和最接近目标值的和。
答案：以下是代码示例
```python
def mergeSort(nums, l, r):
    if l < r:
        m = (l + r) // 2

        mergeSort(nums, l, m)
        mergeSort(nums, m+1, r)

        i = l
        j = m + 1
        k = 0
        temp = [0] * (r-l+1)

        while i <= m and j <= r:
            if nums[i] < nums[j]:
                temp[k] = nums[i]
                i += 1
            else:
                temp[k] = nums[j]
                j += 1
            k += 1

        while i <= m:
            temp[k] = nums[i]
            i += 1
            k += 1

        while j <= r:
            temp[k] = nums[j]
            j += 1
            k += 1

        for p in range(len(temp)):
            nums[l+p] = temp[p]

def threeSumClosest(nums, target):
    mergeSort(nums, 0, len(nums)-1)
    closest_sum = nums[0] + nums[1] + nums[2]
    for i in range(len(nums)-2):
        j = i + 1
        k = len(nums) - 1
        while j < k:
            current_sum = nums[i] + nums[j] + nums[k]
            if abs(current_sum - target) < abs(closest_sum - target):
                closest_sum = current_sum
            if current_sum < target:
                j += 1
            else:
                k -= 1
    return closest_sum
```

5. 给定一个整数数组和一个目标值，找到所有三个数的和为目标值。
答案：以下是代码示例
```python
def mergeSort(nums, l, r):
    if l < r:
        m = (l + r) // 2

        mergeSort(nums, l, m)
        mergeSort(nums, m+1, r)

        i = l
        j = m + 1
        k = 0
        temp = [0] * (r-l+1)

        while i <= m and j <= r:
            if nums[i] < nums[j]:
                temp[k] = nums[i]
                i += 1
            else:
                temp[k] = nums[j]
                j += 1
            k += 1

        while i <= m:
            temp[k] = nums[i]
            i += 1
            k += 1

        while j <= r:
            temp[k] = nums[j]
            j += 1
            k += 1

        for p in range(len(temp)):
            nums[l+p] = temp[p]

def threeSum(nums, target):
    mergeSort(nums, 0, len(nums)-1)
    result = []
    for i in range(len(nums)-2):
        j = i + 1
        k = len(nums) - 1
        while j < k:
            current_sum = nums[i] + nums[j] + nums[k]
            if current_sum == target:
                result.append([nums[i], nums[j], nums[k]])
                j += 1
                k -= 1
            elif current_sum < target:
                j += 1
            else:
                k -= 1
    return result
```   

