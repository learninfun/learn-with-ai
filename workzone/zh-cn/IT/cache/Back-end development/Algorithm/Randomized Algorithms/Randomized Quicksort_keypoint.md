

以下是Randomized Quicksort的重点：

1. 快速排序是一种具有高效率的排序演算法，而 Randomized Quicksort 是在快速排序的基础上改进的演算法。

2. Randomized Quicksort 将 pivot 选择从固定位置改变成随机位置，透过乱数产生 pivot，让一开始选取的 pivot 位置不会对其效率有太大的影响。

3. 一开始先将数列分成三个部分，分别为小于 pivot、等于 pivot 和大于 pivot 的部分。然后将小于 pivot 和大于 pivot 的部分再各自选定一个 pivot，继续进行分区。

4. Randomized Quicksort 的时间复杂度平均为 O(nlogn)，最坏为 O(n^2)。

5. 在数列已经排好序或者都是相同元素的情况下，有可能会发生效率低落的情况，因此需要定期暂停采用 Randomized Quicksort 来处理排序。

6. Randomized Quicksort 的实现方式与普通快速排序大同小异，只是在选择 pivot 时需要多加一些随机的因素。