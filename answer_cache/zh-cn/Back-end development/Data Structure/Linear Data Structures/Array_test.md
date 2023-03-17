

1. 求给定数组中两个数字的和等于目标值的所有组合。例如，给定数组[2, 7, 11, 15]和目标值9，则返回[[2, 7]]。

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

2. 给定一个由非负整数组成的数组，你的任务是将数组中的数字重新排列，使得整个数组的最高位代表数字最大。例如，如果输入是[3, 30, 34, 5, 9]，则输出应该是[9, 5, 34, 3, 30]。

答案：

```
def largest_num(arr):
    arr = list(map(str, arr))
    arr.sort(key=lambda x: x*3, reverse=True)
    return ''.join(arr)

arr = [3, 30, 34, 5, 9]
print(largest_num(arr))
```

3. 一个无序整数数组中，找出最长的连续元素序列的长度。例如，给定数组[100, 4, 200, 1, 3, 2]，最长的连续元素序列为[1, 2, 3, 4]，因此输出4。

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

4. 给定一个未排序的整数数组，找到其中未出现过的最小的正整数。例如，输入[1, 2, 0]，输出应该是3。

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

5. 给定一个二维数组，表示多个区间段的起始和结束时间，找出这些区间的交集。例如，输入为[[1, 3], [2, 6], [5, 8], [7, 9]]，则输出为[[2, 3], [5, 6], [7, 8]]

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