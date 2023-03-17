## 習題預習
{{< ask_chatgpt >}}
給我5題Merge Sort的問題
{{< /ask_chatgpt >}}



1. 什麼是 Merge Sort？它是如何工作的？
2. Merge Sort 是一種時間和空間效率如何？有何優缺點？
3. Merge Sort 在哪些場景下被廣泛使用？能否舉幾個例子？
4. Merge Sort 的最壞時間複雜度是多少？它是如何達到的？
5. Merge Sort 和 Quick Sort 的異同有哪些？它們在哪些情況下能起到更好的作用？   

## 說明知識
{{< ask_chatgpt >}}
說明Merge Sort並舉例
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
條列Merge Sort的重點
{{< /ask_chatgpt >}}



1. Divide and conquer: Merge Sort使用「分而治之」的策略，將一個大問題分解成較小的子問題，並且解決每個子問題，最後將它們合併成一個有序的問題。

2. Recursion: Merge Sort使用遞歸的形式實現分治法，每遞歸一層都會分解原問題並且排序子數組。

3. Merge: Merge Sort將兩個已排序的子數組合併為一個已排序的數組，使得最終的數組保持正確的順序。

4. Time complexity: Merge Sort的時間複雜度為nlogn，是一種高效的排序算法。

5. Space complexity: Merge Sort的空間複雜度是O(n)，因為必須在排序過程中創建一個與輸入數組大小相同的臨時數組。

6. Stability: Merge Sort是穩定的，因為它使用了相等的情況的規則，以保持同等大小的數字的相對位置。

7. Adaptive: Merge Sort也是自適應的，可以在原地排序的情況下使用。

8. Parallel: Merge Sort還有一個顯著的特性，就是它很容易實現並發排序，因為它將大數組分成許多小的子數組進行排序，並且這些子數組可以並發處理   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Merge Sort的中等難度問題，並在後面列出答案
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

## 網路資料
{{< ask_chatgpt >}}
給我5篇Merge Sort的網路資料
{{< /ask_chatgpt >}}



1. Merge Sort in Java: A Complete Guide

Link: https://www.baeldung.com/java-merge-sort

This article on Baeldung explains the concept of Merge Sort in Java in a clear and concise manner. It starts with an overview of Merge Sort, followed by a detailed explanation of how to implement it in Java. It also provides sample code and a step-by-step guide on how to execute it.

2. Merge Sort Algorithm

Link: https://www.geeksforgeeks.org/merge-sort/

GeeksforGeeks is a well-known website among programmers, and this article on Merge Sort is one of the many helpful resources they have. It dives deep into the algorithm of Merge Sort, discussing the working principle and various steps involved. Additionally, the article covers a few examples and provides code snippets in various programming languages.

3. Merge Sort: Sorting Algorithms

Link: https://www.tutorialspoint.com/data_structures_algorithms/merge_sort_algorithm.htm

Tutorialspoint is another popular website for programmers, and this article on Merge Sort is another helpful resource. It covers the basic concepts of Merge Sort, such as the working principle, time complexity, and comparison with other sorting algorithms. It also provides Java code for implementation.

4. Why is Merge Sort Better Than Quick Sort?

Link: https://blog.hyperiondev.com/index.php/2019/02/04/why-is-merge-sort-better-than-quick-sort/

Hyperion Dev Blog is a platform that focuses on coding, and this article explores the comparison between Merge Sort and Quick Sort. It explains why Merge Sort is better in certain scenarios, such as when dealing with large datasets. It also provides code snippets in both Java and Python.

5. Understanding Merge Sort Algorithm in Depth

Link: https://www.softwaretestinghelp.com/merge-sort-algorithm/

Software Testing Help covers a wide range of topics related to software development, and this article focuses on a detailed understanding of Merge Sort. It discusses the working of Merge Sort in detail, along with examples and illustrations. It also provides a visual representation of the algorithm's execution.   

