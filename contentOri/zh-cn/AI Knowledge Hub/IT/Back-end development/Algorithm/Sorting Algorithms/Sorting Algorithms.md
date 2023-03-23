+++
title = "排序演算法"
weight = "4"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Sorting Algorithms的问题
{{< /ask_chatgpt >}}



1. 什么是Sorting Algorithms？

2. Sorting Algorithms的时间复杂度计算方式是什么？

3. 哪些Sorting Algorithms可以实现稳定排序？

4. Sorting Algorithms在什么情况下可能会导致排序失败？

5. 什么是“in-place sorting”？那些Sorting Algorithms可以实现in-place sorting？   

## 说明知识
{{< ask_chatgpt >}}
说明Sorting Algorithms并举例
{{< /ask_chatgpt >}}



排序算法是计算机科学中处理数据的常用算法之一，它可以将一组数据按照顺序排列。常用的排序算法包括冒泡排序、插入排序、选择排序、快速排序、合并排序、堆排序等。

下面举例来说明不同的排序算法：

1. 冒泡排序（Bubble Sort）：从序列的开始位置开始比较相邻的两个数，如果前一个数比后一个数大，则交换两个数的位置，一直比较到序列的最后一位，再从头开始重复以上步骤，直至整个序列有序。

例如，对于以下数列进行冒泡排序：[5, 2, 8, 6, 3, 9, 1]

第一轮排序：[2, 5, 6, 3, 8, 1, 9]
第二轮排序：[2, 5, 3, 6, 1, 8, 9]
第三轮排序：[2, 3, 5, 1, 6, 8, 9]
第四轮排序：[2, 3, 1, 5, 6, 8, 9]
第五轮排序：[2, 1, 3, 5, 6, 8, 9]
第六轮排序：[1, 2, 3, 5, 6, 8, 9]

2. 插入排序（Insertion Sort）：从未排序的数列中，按照顺序遍历每一个元素，将每一个元素插入到已排序的序列中，直到所有元素都插入完成。

例如，对于以下数列进行插入排序：[5, 2, 8, 6, 3, 9, 1]

第一轮排序：[2, 5, 8, 6, 3, 9, 1]
第二轮排序：[2, 5, 8, 6, 3, 9, 1]
第三轮排序：[2, 5, 6, 8, 3, 9, 1]
第四轮排序：[2, 3, 5, 6, 8, 9, 1]
第五轮排序：[1, 2, 3, 5, 6, 8, 9]

3. 选择排序（Selection Sort）：从未排序的数列中，每一次找到未排序部分的最小值，将其排到已排序部分的末尾。

例如，对于以下数列进行选择排序：[5, 2, 8, 6, 3, 9, 1]

第一轮排序：[1, 2, 8, 6, 3, 9, 5]
第二轮排序：[1, 2, 3, 6, 8, 9, 5]
第三轮排序：[1, 2, 3, 5, 8, 9, 6]
第四轮排序：[1, 2, 3, 5, 6, 9, 8]
第五轮排序：[1, 2, 3, 5, 6, 8, 9]

4. 快速排序（Quick Sort）：通过将一个数列划分为较小和较大的两部分，不断对较小部分和较大部分重复上述步骤，直到所有元素都有序。

例如，对于以下数列进行快速排序：[5, 2, 8, 6, 3, 9, 1]

第一轮排序：[1, 2, 3, 5, 8, 9, 6]
第二轮排序：[1, 2, 3, 5, 6, 8, 9]

5. 合并排序（Merge Sort）：将一个数列划分成两个子数列，对每个子数列进行排序，然后将两个子数列合并成一个有序数列。

例如，对于以下数列进行合并排序：[5, 2, 8, 6, 3, 9, 1]

第一轮排序：[2, 5, 8, 6, 1, 3, 9]
第二轮排序：[1, 2, 3, 5, 6, 8, 9]   

## 汇总重点
{{< ask_chatgpt >}}
条列Sorting Algorithms的重点
{{< /ask_chatgpt >}}



1. Bubble Sort: 透过不断比较相邻的元素，将最大值或最小值移到数列的最后面或最前面。
2. Quick Sort: 以一个基准值为标准分成左右两部分，将小于基准值的数值移到左边，大于基准值的数值移到右边，再以同样的方式处理左右两个区间，递回进行直到整个数列有序。
3. Merge Sort: 将数列分成若干个子序列进行排序，再将这些子序列合并成一个大序列进行排序。
4. Insertion Sort: 逐一取出元素并按照适当的顺序插入已排好序的数列中。
5. Selection Sort: 依次找出数列中最小的元素，放到数列起始位置，再从剩下的元素中找到最小值，将其放到已排序部分的结尾，直到整个数列有序。
6. Shell Sort: 基于插入排序的一种算法，会先进行一次分组排序，然后逐渐缩小分组区间，最终当区间缩小成1时，整个数列也就有序了。
7. Heap Sort: 利用二叉堆进行排序，先将数列变成一个最小堆或最大堆，然后将堆顶元素取出加入结果集，再把剩余堆调整成新的堆，重复以上步骤直到整个数列有序。
8. Counting Sort: 用于小范围整数排序，其核心理念在于将输入的数据值转化为键存储在额外开辟的数组空间中，然后依次把计数大于 1 的填回原数组。

以上就是Sorting Algorithms的重点，各算法在时间复杂度和空间复杂度等方面都有不同优点和限制，需要根据应用场景选择适合的算法来进行排序。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Sorting Algorithms的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 如何實現QuickSort的In-Place版本？
答案：由於QuickSort是將一個數列切割成較小和較大的元素，因此可以使用一個指針來追踪分割元素的位置，將小於分割元素的元素放在左側，大於分割元素的元素放在右側。請參考以下代碼：

```python
def quickSortInPlace(arr, low, high):
    if low < high:
        pivot = partition(arr, low, high)
        quickSortInPlace(arr, low, pivot - 1)
        quickSortInPlace(arr, pivot + 1, high)
 
def partition(arr, low, high):
    pivot_value = arr[high]
    i = low - 1
    for j in range(low, high):
        if arr[j] <= pivot_value:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1
```

2. 如何使用MergeSort來對超過1TB的大型數據進行排序？
答案：將數據分成小的塊，分別使用MergeSort進行排序，然後依次合併這些已排序的塊，直到所有數據都被排序。這種方法稱為外部排序（external sorting）。例如，可以將數據分成100GB的塊，排序每個塊，然後使用儲存器並行合併這些已排序的塊。

3. 如何實現HeapSort算法？
答案：HeapSort使用最大堆（Max Heap）來實現排序的過程。首先將數據構建成最大堆，然後依次取出堆頂元素（最大元素），放到數列最後，再進行最大堆重建操作。請參考以下代碼：

```python
def heapSort(arr):
    n = len(arr)
    for i in range(n//2 - 1, -1, -1):
        heapify(arr, n, i)
    for i in range(n-1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        heapify(arr, i, 0)
 
def heapify(arr, n, i):
    largest = i
    l = 2*i + 1
    r = 2*i + 2
    if l < n and arr[l] > arr[largest]:
        largest = l
    if r < n and arr[r] > arr[largest]:
        largest = r
    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)
```

4. 如何在O(n)時間複雜度下找到數列中第k小的元素？
答案：可以使用QuickSelect算法，類似於QuickSort的思想，只需繼續對分割後的一個子數列進行遞迴，直到找到第k小的元素。平均時間複雜度為 O(n)。以下是Python示例代碼：

```python
import random
 
def quickSelect(arr, left, right, k):
    if left == right:
        return arr[left]
    pivotIndex = random.randint(left, right)
    pivotIndex = partition(arr, left, right, pivotIndex)
    if k == pivotIndex:
        return arr[k]
    elif k < pivotIndex:
        return quickSelect(arr, left, pivotIndex - 1, k)
    else:
        return quickSelect(arr, pivotIndex + 1, right, k)
 
def partition(arr, left, right, pivotIndex):
    pivotValue = arr[pivotIndex]
    arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
    storeIndex = left
    for i in range(left, right):
        if arr[i] < pivotValue:
            arr[i], arr[storeIndex] =  arr[storeIndex], arr[i]
            storeIndex += 1
    arr[storeIndex], arr[right] = arr[right], arr[storeIndex]
    return storeIndex
```

5. 如何使用BubbleSort來對鏈表進行排序？
答案：在鏈表上進行交換操作較為困難，因此可以使用一個標誌來標記最後一個已排序節點的位置，然後將下一個節點與其比較，如果需要排序則進行交換，直到標誌移動到鏈表末尾為止。以下是示例代碼：

```python
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
 
def bubbleSortList(head):
    if not head:
        return None
    flag = True
    while flag:
        flag = False
        curr = head
        while curr.next:
            if curr.val > curr.next.val:
                curr.val, curr.next.val = curr.next.val, curr.val
                flag = True
            curr = curr.next
    return head
```

以上所有代碼均為Python 3.6+。   

