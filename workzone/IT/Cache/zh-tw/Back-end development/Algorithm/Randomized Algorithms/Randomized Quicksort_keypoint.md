

以下是Randomized Quicksort的重點：

1. 快速排序是一種具有高效率的排序演算法，而 Randomized Quicksort 是在快速排序的基礎上改進的演算法。

2. Randomized Quicksort 將 pivot 選擇從固定位置改變成隨機位置，透過亂數產生 pivot，讓一開始選取的 pivot 位置不會對其效率有太大的影響。

3. 一開始先將數列分成三個部分，分別為小於 pivot、等於 pivot 和大於 pivot 的部分。然後將小於 pivot 和大於 pivot 的部分再各自選定一個 pivot，繼續進行分區。

4. Randomized Quicksort 的時間複雜度平均為 O(nlogn)，最壞為 O(n^2)。

5. 在數列已經排好序或者都是相同元素的情況下，有可能會發生效率低落的情況，因此需要定期暫停採用 Randomized Quicksort 來處理排序。

6. Randomized Quicksort 的實現方式與普通快速排序大同小異，只是在選擇 pivot 時需要多加一些隨機的因素。