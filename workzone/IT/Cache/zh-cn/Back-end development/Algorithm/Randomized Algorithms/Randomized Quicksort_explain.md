

Randomized Quicksort是快速排序算法的一种变体，它与原始快速排序算法的区别在于，它使用随机数来选择主元素（pivot），从而尽可能避免了最坏情况。在原始的快速排序算法中，选择的主元素可能会导致分割的子序列不平衡，进而导致算法的时间复杂度退化为O(n2)；而随机化的选择主元素，则能够尽可能保持分割的子序列平衡，从而保证算法的时间复杂度为O(n log n)。

以下是Randomized Quicksort的实现步骤：

1. 选取一个元素作为主元素（pivot），可以随机选取其中一个元素。
2. 扫瞄整个序列，将比pivot小的元素放到左边，比pivot大的元素放到右边，相等的元素放到任意一边（一般为左边）。
3. 递归排序左右两个子序列。

以下为Randomized Quicksort的示例代码：

```
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void randomized_quicksort(int arr[], int left, int right);

int main(void) {
    int arr[] = {3, 5, 2, 1, 4};
    int n = sizeof(arr) / sizeof(arr[0]);
    randomized_quicksort(arr, 0, n - 1);
    for(int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
    return 0;
}

void randomized_quicksort(int arr[], int left, int right) {
    if(left >= right) {
        return;
    }
    int p = rand() % (right - left + 1) + left;
    int pivot = arr[p];
    int i = left;
    int j = right;
    while(i <= j) {
        while(arr[i] < pivot) {
            i++;
        }
        while(arr[j] > pivot) {
            j--;
        }
        if(i <= j) {
            int temp = arr[i];
            arr[i] = arr[j];
            arr[j] = temp;
            i++;
            j--;
        }
    }
    randomized_quicksort(arr, left, j);
    randomized_quicksort(arr, i, right);
}
```

在上面的示例代码中，我们随机选取一个元素作为主元素（pivot），并通过while循环将比pivot小的元素放到左边，比pivot大的元素放到右边，然后递归的对左右两个子序列进行排序。程序的输出结果为：

```
1 2 3 4 5
```

可以看到，我们使用Randomized Quicksort算法成功地将序列从小到大排序了。