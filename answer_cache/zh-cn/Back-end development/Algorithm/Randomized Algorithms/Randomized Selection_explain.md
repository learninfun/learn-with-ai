

Randomized Selection是一种选择性能比QuickSort还要优异的算法，可以在O(n)的时间复杂度内选择一个阵列中第k小（大）的元素。

该算法的基本思想是选择一个随机数作为pivot（基准点），将阵列分为两个部分，一部分比pivot小，一部分比pivot大。然后判断第k小（大）的元素在哪个部分中，进行分治递归操作，直到找到第k小（大）的元素为止。

以下是一个说明Randomized Selection的Python代码范例：

```
import random

def randomized_selection(arr, k):
    if len(arr) == 1:
        return arr[0]
    pivot = random.choice(arr)
    left = [x for x in arr if x < pivot]
    right = [x for x in arr if x > pivot]
    mid = [x for x in arr if x == pivot]
    if k <= len(left):
        return randomized_selection(left, k)
    elif k > len(left) + len(mid):
        return randomized_selection(right, k-len(left)-len(mid))
    else:
        return mid[0]
```

在这个例子中，我们首先选择随机pivot，然后将阵列分成三部分：比pivot小的元素、比pivot大的元素、和等于pivot的元素。接着，我们判断第k小（大）的元素在哪个部分中，进行分治递归操作。如果k小于等于左边部分的元素个数，那么第k小的元素必然在左边，我们就递归对左边的阵列进行操作。如果k大于左边部分的元素个数加上等于pivot的元素个数，那么第k小的元素必然在右边，我们递归对右边的阵列进行操作。如果k介于左边部分的元素和等于pivot的元素的个数之间，那么第k小的元素就是pivot本身。

以上就是Randomized Selection的基本概念和使用范例。该算法通常比其他排序算法效率更高，因为它采用随机pivot，可以避免最坏情况下的时间复杂度。