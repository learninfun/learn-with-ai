

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