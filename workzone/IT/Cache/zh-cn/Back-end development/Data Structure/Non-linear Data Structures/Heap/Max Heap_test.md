

1. 找到第 k 大的元素
給定一個大小爲 n 的整數數組，請找到第 k 大的元素。可以假設 1 ≤ k ≤ n。

答案：使用最大堆維護數組的前 k 大元素，然後返回堆頂元素即可。

2. 合併 k 個有序數組
給定 k 個已經按照升序排列的數組，請將它們合併成一個新的有序數組。

答案：建立一個大小爲 k 的最小堆，每次將 k 個數組中的最小元素加入堆中，然後彈出堆頂元素並添加到結果數組中，直到堆為空。

3. 求解中位數
給定一個數組，求解其中位數（如果數組大小為偶數，則返回中間兩個數的平均值）。

答案：使用兩個堆，一個最大堆維護數組的前半部分，一個最小堆維護數組的後半部分。當數組大小爲奇數時，中位數就是最大堆的堆頂，當數組大小爲偶數時，中位數就是最大堆的堆頂和最小堆的堆頂的平均值。

4. 找到 k 個最接近的元素
給定一個排序後的數組和一個數 k，請找到 k 個和指定數最接近的元素。可以假設給定數字一定存在於數組中。

答案：使用大小爲 k 的最小堆維護 k 個距離最近的元素。每次加入一個新元素時，如果堆大小小於 k，就直接加入，否則就判斷新元素是否比堆頂元素更接近，如果是則彈出堆頂元素並加入新元素。

5. 找到多數元素
給定一個大小爲 n 的整數數組，請找到出現次數超過 ⌊ n/2 ⌋ 的元素。

答案：使用大小爲 n/2+1 的最小堆（或最大堆），將數組中的元素加入堆中。由於最多只有一個元素的出現次數超過 ⌊ n/2 ⌋，所以堆頂元素一定是答案。