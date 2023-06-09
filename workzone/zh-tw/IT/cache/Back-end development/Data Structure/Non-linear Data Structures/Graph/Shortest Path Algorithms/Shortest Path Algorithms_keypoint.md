

最短路徑演算法（Shortest Path Algorithms）是尋找兩個點之間最短路徑的方法，以下是其重點：

1. Dijkstra演算法

- Dijkstra演算法是最常用的尋找單源最短路徑的演算法。
- 它基於貪心法的思想，每次選擇未處理節點中距離起點最近的節點。
- 該演算法只能處理無負權重的圖。

2. Bellman-Ford演算法

- Bellman-Ford演算法是一種廣泛應用的尋找單源最短路徑的演算法，也可處理有負權重的圖。
- 該演算法通過遍歷所有邊來找到最短路徑，並通過逐步縮短範圍的策略避免了死循環的問題。

3. Floyd-Warshall演算法

- Floyd-Warshall演算法是一種尋找所有節點之間最短路徑的演算法。
- 該演算法使用動態規劃的思想，從中介節點尋找路徑，並使用矩陣來表示路徑和距離。
- 該演算法能夠處理有負權重的圖。

4. A*演算法

- A*演算法是一種啟髮式搜索演算法，通常用於找到兩個點之間最短路徑。
- 該演算法通過評估每個節點到目標節點的距離來決定搜索路徑，並使用該距離和到起點的距離來計算路徑成本。
- 該演算法通常比Dijkstra演算法更快，但需要一個合適的啟發函數來評估距離。

5. Johnson演算法

- Johnson演算法是一種先使用Bellman-Ford演算法求解圖中任意兩個節點之間的最短路徑，再使用Dijkstra演算法進行快速查詢的方法。
- 該演算法避免了負權重的計算，並且能夠快速查詢出所有節點之間的最短路徑。

以上是最短路徑演算法的重點，瞭解這些演算法有助於選擇適用的演算法並加快計算速度。