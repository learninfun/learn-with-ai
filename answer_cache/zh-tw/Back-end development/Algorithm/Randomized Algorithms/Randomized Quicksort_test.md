

1. 如何避免Randomized Quicksort的最差時間複雜度O(n^2)？請說明原因及解決方法。
答案：避免最差情況的方法之一是使用新的選擇pivot的方法，例如Median of 3或Random Median。此外，也可以使用尾遞迴或迭代版本，或者實現Randomized Quicksort的隨機樣本版本，以減少最差情況的概率。

2. 如何實現Randomized Quicksort的平均時間複雜度？請給出算法的時間複雜度和空間複雜度。
答案：Randomized Quicksort的平均時間複雜度為O(nlogn)，空間複雜度為O(logn)。算法的關鍵在於隨機選擇pivot，並使用partition方法將數組分為兩個部分。對左側和右側的子數組遞歸地應用相同的算法，直到排序完成。

3. 如何在Randomized Quicksort中實現重複元素的處理？請說明原因及解決方法。
答案：在Randomized Quicksort中，如果數組中存在大量重複元素，則可能導致一些情況下的最壞時間複雜度。解決此問題的一種方法是使用三向切分快速排序，將數組分為三個部分（小於、等於和大於pivot）。這樣可以避免重複元素被反覆交換，在許多情況下可以提高算法的效率。

4. 如何測試Randomized Quicksort的性能？請單獨列出利用實驗數據進行性能分析的步驟。
答案：測試Randomized Quicksort的性能可以進行以下步驟：
（1）選擇數據集：從不同的數據集中選擇數據，包括已排序、部分排序或完全隨機的數據。
（2）選擇公平的比較方式：比較Randomized Quicksort和其他排序算法，如Merge Sort和Heap Sort。
（3）記錄運行時間：使用編程語言的時間記錄功能，記錄每次運行算法花費的時間。
（4）重複運行多次：運行算法多次以獲得平均值。
（5）分析數據：比較不同數據集之間的性能差異，並觀察對快速排序的優化效果。

5. 如何在Randomized Quicksort中實現非遞迴版？請給出代碼實現。
答案：
非遞迴版Randomized Quicksort的主要思路是使用一個棧來模擬遞迴的運行過程。以下是Python代碼：

def quickSort(arr):

    # Create an empty stack
    stack = []

    # Push initial values of l and h to stack
    l = 0
    h = len(arr) - 1
    stack.append((l, h))

    # Keep popping from stack while is not empty
    while stack:

        # Pop values from stack
        (low, high) = stack.pop()

        # Set pivot element at its correct position
        p = partition(arr, low, high)

        # If there are elements on left side of pivot,
        # push the left side to stack
        if p - 1 > low:
            stack.append((low, p - 1))

        # If there are elements on right side of pivot,
        # push the right side to stack
        if p + 1 < high:
            stack.append((p + 1, high))

    # Check the sorted array
    return arr
  
def partition(arr, low, high):
    i = low - 1
    pivot = arr[high]
    for j in range(low, high):
        if arr[j] < pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i + 1], arr[high] = arr[high], arr[i + 1]
    return i + 1

# Test
arr = [10, 80, 30, 90, 40, 50, 70]
print('Original array:', arr)

quickSort(arr)
print('Sorted array:', arr)