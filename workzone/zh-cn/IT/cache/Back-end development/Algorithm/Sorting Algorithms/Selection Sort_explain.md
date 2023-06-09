

Selection Sort（选择排序）是一种简单的排序演算法，它会选择最小的元素在整个序列中，将其与序列中的第一个元素交换，然后在剩余的元素中选择最小的元素，将其与序列中的第二个元素交换，以此类推，直到整个序列都被排序过。

下面是一个实现 Selection Sort 的 Python 程序：

```
def selection_sort(arr):
    n = len(arr)
    for i in range(n):
        min_idx = i
        for j in range(i+1, n):
            if arr[j] < arr[min_idx]:
                min_idx = j
        arr[i], arr[min_idx] = arr[min_idx], arr[i]
    return arr
```

该程序使用双重循环，外部循环用于遍历序列中的每一个元素，内部循环用于查找最小的元素。在每次内部循环结束时，我们找到了最小的元素，并且将其移动到序列的开头。

举个例子，假设我们要对以下整数序列进行 Selection Sort 操作：

```
[64, 25, 12, 22, 11]
```

首先，我们找到最小的元素 11，然后将其与序列的第一个元素 64 交换，序列变成：

```
[11, 25, 12, 22, 64]
```

然后，我们从第二个元素 25 到最后一个元素 64 中找到最小的元素 12，然后将其与序列的第二个元素 25 交换，序列变成：

```
[11, 12, 25, 22, 64]
```

然后，继续进行类似的操作，直到整个序列都被排序过：

```
[11, 12, 22, 25, 64]
```

因此，根据上面的算法，我们在 $O(n^2)$ 的时间复杂度内将序列进行了排序。