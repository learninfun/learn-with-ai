

1. 實作Heap Sort的函數，將一個陣列從小到大排序。
範例輸入：[7, 5, 1, 8, 3]
範例輸出：[1, 3, 5, 7, 8]

2. 如何將Heap Sort的執行時間降至O(nlogn)?
答案：使用heapify函數來建立heap，此函數的時間複雜度為O(n)。

3. 請解釋Heap Sort不穩定的排序特性。
答案：當兩個元素的值相等時，排序後它們的相對位置可能會改變。

4. 請列出執行Heap Sort時使用的空間複雜度。
答案：Heap Sort的空間複雜度為O(1)，因為排序是直接在原始陣列上進行的。

5. 請解釋Heap Sort與Merge Sort的差異。
答案：Heap Sort使用heap來實現排序，時間複雜度為O(nlogn)，空間複雜度為O(1)。Merge Sort使用分治法來實現排序，時間複雜度為O(nlogn)，空間複雜度為O(n)。