

1. Dynamic Programming問題: 最長遞增子序列 (Longest Increasing Subsequence, LIS)
給定一個數列，例如[10, 9, 2, 5, 3, 7, 101, 18]，求其中最長的遞增子序列有多長？

答案: 答案為4，其中最長的遞增子序列為[2, 3, 7, 101]。

2. Greedy Algorithm問題: 貪心算法（Kruskal's Algorithm）
給你一個帶權無向圖，請設計一個算法來找到最小生成樹。

答案: Kruskal's Algorithm

3. Divide and Conquer問題: 在排序數組中查找特定元素
給定一個已排序的整數數組和一個目標值，請在數組中找到目標值的索引。如果目標值不存在於數組中，則返回-1。

答案：針對該問題的解法有很多，可以使用二分搜尋算法、插值搜尋算法等等。

4. Reduction問題: 前置問題 (Traveling Salesman Problem, TSP)
給定一個權重完全圖，請設計一個算法尋找經過所有節點的最短路徑。

答案: 可以將TSP問題針對所有節點求最小生成樹（MST），然後在MST上運行歐拉迴路演算法。

5. Approximation Algorithm問題: 歐拉環路 (Eulerian Circuit)
給定一個無向圖，如果它包含一個權重為奇數的節點，則該圖不可能有一個歐拉迴路。如果該圖中所有節點的權重都是偶數，則該圖可能有歐拉迴路。請編寫一個算法來找到歐拉迴路。

答案: Hierholzer Algorithm

以上問題僅供參考，實際上任何算法設計題目都可以根據不同難度等級進行分類。