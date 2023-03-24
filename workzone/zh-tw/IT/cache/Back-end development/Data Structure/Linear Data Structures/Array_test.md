

1. 求給定數組中兩個數字的和等於目標值的所有組合。例如，給定數組[2, 7, 11, 15]和目標值9，則返回[[2, 7]]。

答案：

```
def two_sum(arr, target):
    res = []
    for i in range(len(arr)):
        for j in range(i+1, len(arr)):
            if arr[i] + arr[j] == target:
                res.append([arr[i], arr[j]])
    return res

arr = [2, 7, 11, 15]
target = 9
print(two_sum(arr, target))
```

2. 給定一個由非負整數組成的數組，你的任務是將數組中的數字重新排列，使得整個數組的最高位代表數字最大。例如，如果輸入是[3, 30, 34, 5, 9]，則輸出應該是[9, 5, 34, 3, 30]。

答案：

```
def largest_num(arr):
    arr = list(map(str, arr))
    arr.sort(key=lambda x: x*3, reverse=True)
    return ''.join(arr)

arr = [3, 30, 34, 5, 9]
print(largest_num(arr))
```

3. 一個無序整數數組中，找出最長的連續元素序列的長度。例如，給定數組[100, 4, 200, 1, 3, 2]，最長的連續元素序列為[1, 2, 3, 4]，因此輸出4。

答案：

```
def longest_consecutive_sequence(arr):
    if not arr:
        return 0
    nums = set(arr)
    max_len = 0
    for num in nums:
        if num-1 not in nums:
            curr_num = num
            curr_len = 1
            while curr_num+1 in nums:
                curr_num += 1
                curr_len += 1
            max_len = max(max_len, curr_len)
    return max_len

arr = [100, 4, 200, 1, 3, 2]
print(longest_consecutive_sequence(arr))
```

4. 給定一個未排序的整數數組，找到其中未出現過的最小的正整數。例如，輸入[1, 2, 0]，輸出應該是3。

答案：

```
def smallest_missing_positive_num(arr):
    if not arr:
        return 1
    for i in range(len(arr)):
        while 0 < arr[i] <= len(arr) and arr[i] != arr[arr[i]-1]:
            arr[arr[i]-1], arr[i] = arr[i], arr[arr[i]-1]
    for i in range(len(arr)):
        if arr[i] != i+1:
            return i+1
    return len(arr) + 1

arr = [1, 2, 0]
print(smallest_missing_positive_num(arr))
```

5. 給定一個二維數組，表示多個區間段的起始和結束時間，找出這些區間的交集。例如，輸入為[[1, 3], [2, 6], [5, 8], [7, 9]]，則輸出為[[2, 3], [5, 6], [7, 8]]

答案：

```
def interval_intersection(intervals):
    if not intervals:
        return []
    intervals.sort(key=lambda x: x[0])
    res = []
    curr_start, curr_end = intervals[0]
    for interval in intervals[1:]:
        if interval[0] <= curr_end:
            curr_start = max(curr_start, interval[0])
            curr_end = min(curr_end, interval[1])
        else:
            res.append([curr_start, curr_end])
            curr_start, curr_end = interval[0], interval[1]
    res.append([curr_start, curr_end])
    return res

intervals = [[1, 3], [2, 6], [5, 8], [7, 9]]
print(interval_intersection(intervals))
```