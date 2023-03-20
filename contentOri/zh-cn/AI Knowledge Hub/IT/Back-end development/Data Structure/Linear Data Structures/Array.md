+++
title = "数组"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Array的问题
{{< /ask_chatgpt >}}



1. 如何在数组中查找特定元素？
2. 如何反转数组中的元素顺序？
3. 如何在数组中删除特定元素？
4. 如何按特定条件对数组进行排序？
5. 如何将两个数组合并为一个数组？   

## 说明知识
{{< ask_chatgpt >}}
说明Array并举例
{{< /ask_chatgpt >}}



Array（陣列）是一種儲存多個相同類型值的資料結構。它可以在一個變數中存放多個值，並且每一個值都可以通過索引來訪問。索引通常是整數型別，它從 0 開始，依次增加。

以下是一些 Array 的例子：

1. 儲存一組數值

```
let numbers = [1, 2, 3, 4, 5];
console.log(numbers); // [1, 2, 3, 4, 5]
```

2. 儲存多個字符串

```
let fruits = ["apple", "banana", "orange", "grape"];
console.log(fruits); // ["apple", "banana", "orange", "grape"]
```

3. 儲存布林值

```
let bools = [true, false, true];
console.log(bools); // [true, false, true]
```

4. 儲存對象

```
let users = [
  { name: "Amy", age: 25 },
  { name: "Bob", age: 30 },
  { name: "Carl", age: 28 }
];
console.log(users); // [{name: "Amy", age: 25}, {name: "Bob", age: 30}, {name: "Carl", age: 28}]
```

需要注意的是，Array 有一些屬性和方法，可以對其中的值進行操作，比如 `length` 屬性可以返回陣列的大小，`push` 方法可以在陣列末尾添加一個元素，`splice` 方法可以在指定位置插入或刪除元素等。   

## 汇总重点
{{< ask_chatgpt >}}
条列Array的重点
{{< /ask_chatgpt >}}



以下是Array的重點：

1. 定義：Array是一種數據結構，它由相同類型的元素組成，並存儲在相鄰的內存位置上。

2. 屬性：Array具有下面幾個屬性：

- 長度：Array的長度是指元素的個數。
- 索引：每個元素都有一個索引，從 0 開始編號。
- 元素：Array 的每個元素都具有相同的數據類型。

3. 優點：

- 維護有序數據。
- 可以隨機訪問任何一個元素。
- 可以通過索引快速訪問、修改、刪除和新增元素。

4. 缺點：

- 長度不可變：Array 的長度在創建時已決定且不可變。
- 浪費內存：如果Array 的元素數量沒有達到其容量，那麼可能會造成內存的浪費。
- 需要連續的內存空間：Array 需要連續的內存空間來存儲所有的元素，因此，當需要大量儲存元素時，可能會造成內存不足的問題。

5. 應用：

- 儲存一系列數字或字符串。
- 作為矩陣或向量的儲存方式。
- 當需要訪問元素順序時使用（例如迭代）。
- 用於實現其他數據結構（例如堆棧和佇列）。

總結起來，Array 是一種非常基礎且常用的數據結構，它非常適用於儲存一系列具有相同類型的數據，並提供了高效的隨機訪問、修改、刪除和新增元素的操作。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Array的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Array的网络数据
{{< /ask_chatgpt >}}



1. "Arrays in Java" - Oracle: https://docs.oracle.com/javase/tutorial/java/nutsandbolts/arrays.html 

This article from Oracle covers the basics of arrays in Java, including how to declare and initialize an array, how to access array elements, and how to use multidimensional arrays.

2. "Arrays in JavaScript" - MDN Web Docs: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array 

MDN Web Docs provides a comprehensive guide to arrays in JavaScript, including information on common array methods and how to iterate over arrays using loops.

3. "C++ Arrays" - GeeksforGeeks: https://www.geeksforgeeks.org/arrays-in-c-cpp/ 

This article from GeeksforGeeks provides an introduction to arrays in C++, including syntax and examples of one-dimensional and multidimensional arrays.

4. "Python Arrays" - Programiz: https://www.programiz.com/python-programming/array 

Programiz offers a detailed guide to using arrays in Python, including how to create and manipulate arrays using built-in functions and packages like NumPy.

5. "Swift Arrays" - Hacking with Swift: https://www.hackingwithswift.com/sixty/6/1/arrays 

This tutorial from Hacking with Swift covers arrays in the Swift programming language, including how to create arrays, add and remove elements, and sort arrays.   
