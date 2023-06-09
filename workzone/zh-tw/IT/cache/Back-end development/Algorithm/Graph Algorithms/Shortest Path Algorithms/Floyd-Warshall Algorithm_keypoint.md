

1. Floyd-Warshall Algorithm（F-W算法）是一種動態規划算法，用於解決全源最短路問題。
2. F-W算法利用矩陣來存儲每對節點之間的最短路長度。
3. 矩陣中每個元素的值代表從i到j的最短路長度，其中i和j是矩陣下標。
4. F-W算法每次更新矩陣中保存的最短路長度，以便找出每對節點之間的最短路。
5. F-W算法的核心思想是利用動態規劃思想，將問題分解為子問題依次求解。
6. 在F-W算法中，我們先將圖表示成一個帶權有向圖（即邊帶權），然後利用矩陣來保存最短路長度。
7. F-W算法主要包括初始化矩陣、遍歷圖、更新矩陣等步驟。
8. 儘管F-W算法的時間複雜度較高（為O(N^3)），但它可以處理帶權有向圖中任意兩點之間的最短路問題，並且效果非常穩定。