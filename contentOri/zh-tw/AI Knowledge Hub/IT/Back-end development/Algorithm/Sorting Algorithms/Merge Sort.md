+++
title = "合併排序"
weight = "4"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Merge Sort的中文問題
{{< /ask_chatgpt >}}



1. 什麼是 Merge Sort？它是如何工作的？
2. Merge Sort 是一種時間和空間效率如何？有何優缺點？
3. Merge Sort 在哪些場景下被廣泛使用？能否舉幾個例子？
4. Merge Sort 的最壞時間複雜度是多少？它是如何達到的？
5. Merge Sort 和 Quick Sort 的異同有哪些？它們在哪些情況下能起到更好的作用？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Merge Sort並舉例
{{< /ask_chatgpt >}}



Merge Sort是一種分治法排序演算法，它的基本概念是將待排序數列不斷拆分為較小的子序列，直到每個子序列裡面只有一個元素為止，再將相鄰的子序列合併為更大的有序子序列，直到最後只剩下一個有序的序列。

以下是Merge Sort的步驟：

1. 分割：把待排序的序列拆分成兩個子序列，每個子序列包含約一半的元素。 

2. 對每個子序列遞歸排序，直到每個子序列裡面只剩下一個元素。 

3. 合併：把排序好的子序列合併成一個大的有序序列。 

舉例來說，如果我們要對以下數列進行排序：

{38, 27, 43, 3, 9, 82, 10}

分割後會得到兩個子序列：

{38, 27, 43, 3}

{9, 82, 10}

對每個子序列遞歸排序，最終得到排序好的子序列：

{3, 27, 38, 43}

{9, 10, 82}

最後將排序好的子序列合併，得到最終的有序序列：

{3, 9, 10, 27, 38, 43, 82}   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Merge Sort的重點
{{< /ask_chatgpt >}}



1. Divide and conquer: Merge Sort使用“分而治之”的策略，将一个大问题分解成较小的子问题，并且解决每个子问题，最后将它们合并成一个有序的问题。

2. Recursion: Merge Sort使用递归的形式实现分治法，每递归一层都会分解原问题并且排序子数组。

3. Merge: Merge Sort将两个已排序的子数组合并为一个已排序的数组，使得最终的数组保持正确的顺序。

4. Time complexity: Merge Sort的时间复杂度为nlogn，是一种高效的排序算法。

5. Space complexity: Merge Sort的空间复杂度是O(n)，因为必须在排序过程中创建一个与输入数组大小相同的临时数组。

6. Stability: Merge Sort是稳定的，因为它使用了相等的情况的规则，以保持同等大小的数字的相对位置。

7. Adaptive: Merge Sort也是自适应的，可以在原地排序的情况下使用。

8. Parallel: Merge Sort还有一个显著的特性，就是它很容易实现并发排序，因为它将大数组分成许多小的子数组进行排序，并且这些子数组可以并发处理   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Merge Sort的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 實現Merge Sort的一個變體，將所有小於等於pivot的元素放在左邊，大於pivot的元素放在右邊。
答案：以下是代碼示例
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

2. 給定一個整數數組，找到其中第K大的元素。
答案：以下是代碼示例
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

3. 設計一個算法來對一個整數數組進行排序，其中一個數是重複的，且該重複數可能有多個位置。
答案：以下是代碼示例
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

4. 給定一個整數數組和一個目標值，找到數組中三個數的和最接近目標值的和。
答案：以下是代碼示例
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

5. 給定一個整數數組和一個目標值，找到所有三個數的和為目標值。
答案：以下是代碼示例
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

