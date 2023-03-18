

Randomized Quicksort是快速排序算法的一種變體，它與原始快速排序算法的區別在於，它使用隨機數來選擇主元素（pivot），從而盡可能避免了最壞情況。在原始的快速排序算法中，選擇的主元素可能會導致分割的子序列不平衡，進而導致算法的時間複雜度退化為O(n2)；而隨機化的選擇主元素，則能夠盡可能保持分割的子序列平衡，從而保證算法的時間複雜度為O(n log n)。

以下是Randomized Quicksort的實現步驟：

1. 選取一個元素作為主元素（pivot），可以隨機選取其中一個元素。
2. 掃瞄整個序列，將比pivot小的元素放到左邊，比pivot大的元素放到右邊，相等的元素放到任意一邊（一般為左邊）。
3. 遞歸排序左右兩個子序列。

以下為Randomized Quicksort的示例代碼：

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

在上面的示例代碼中，我們隨機選取一個元素作為主元素（pivot），並通過while循環將比pivot小的元素放到左邊，比pivot大的元素放到右邊，然後遞歸的對左右兩個子序列進行排序。程序的輸出結果為：

```
1 2 3 4 5
```

可以看到，我們使用Randomized Quicksort算法成功地將序列從小到大排序了。