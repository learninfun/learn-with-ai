## 習題預習
{{< ask_chatgpt >}}
給我5題Selection Sort的問題
{{< /ask_chatgpt >}}



1. Selection Sort是什麼排序演算法？它是如何運作的？
2. 在Selection Sort中，對於n個元素的未排序子數列，需要執行多少次比較操作和交換操作？
3. Selection Sort和Bubble Sort之間有什麼差異？哪個演算法的效率更高？
4. 在Selection Sort中，當要將一個元素插入到已排序子數列中時，是從哪個方向開始執行比較操作？
5. 如果要對一個數組進行升序排序，什麼情況下會選擇使用Selection Sort？   

## 說明知識
{{< ask_chatgpt >}}
說明Selection Sort並舉例
{{< /ask_chatgpt >}}



Selection Sort（選擇排序）是一種簡單的排序演算法，它會選擇最小的元素在整個序列中，將其與序列中的第一個元素交換，然後在剩餘的元素中選擇最小的元素，將其與序列中的第二個元素交換，以此類推，直到整個序列都被排序過。

下面是一個實現 Selection Sort 的 Python 程序：

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

該程序使用雙重循環，外部循環用於遍歷序列中的每一個元素，內部循環用於查找最小的元素。在每次內部循環結束時，我們找到了最小的元素，並且將其移動到序列的開頭。

舉個例子，假設我們要對以下整數序列進行 Selection Sort 操作：

```
[64, 25, 12, 22, 11]
```

首先，我們找到最小的元素 11，然後將其與序列的第一個元素 64 交換，序列變成：

```
[11, 25, 12, 22, 64]
```

然後，我們從第二個元素 25 到最後一個元素 64 中找到最小的元素 12，然後將其與序列的第二個元素 25 交換，序列變成：

```
[11, 12, 25, 22, 64]
```

然後，繼續進行類似的操作，直到整個序列都被排序過：

```
[11, 12, 22, 25, 64]
```

因此，根據上面的算法，我們在 $O(n^2)$ 的時間複雜度內將序列進行了排序。   

