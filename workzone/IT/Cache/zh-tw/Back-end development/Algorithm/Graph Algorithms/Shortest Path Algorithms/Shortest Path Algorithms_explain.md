

Shortest Path Algorithms 是一種常用的圖論算法，用於在圖論中，找出一個特定起點與終點之間的最短路徑。這種算法主要應用於路徑導航、交通運輸等方面。以下是三種常用的 Shortest Path Algorithms:

1. Dijkstra 算法
Dijkstra 算法是一種貪心算法，通過選擇當前節點到起始節點距離最短的節點，來逐步構造最短路徑。它的時間複雜度與圖中邊的數量有關，通常為 $O(|E| + |V|\log|V|)$。

舉例：在一張地圖中，我們需要從 A 點出發，前往 B 點，請問最短路徑是哪條？

<img src="https://i.imgur.com/Z7v1zIC.png" width="200" height="150">

通過 Dijkstra 算法，我們可以得到 A - C - D - B 這條路徑，其長度為 10。

2. Bellman-Ford 算法
Bellman-Ford 算法也是一種動態規划算法，通過對圖中所有邊進行 V - 1 次鬆弛操作，來找到最短路徑。它的時間複雜度為 $O(|V||E|)$。

舉例：在一張地圖中，我們需要從 A 點出發，前往 B 點，請問最短路徑是哪條？

<img src="https://i.imgur.com/IjOojgn.png" width="200" height="150">

通過 Bellman-Ford 算法，我們可以得到 A - C - B 這條路徑，其長度為 6。

3. Floyd-Warshall 算法
Floyd-Warshall 算法是一種動態規划算法，用於解決所有節點之間的最短路徑問題。它的時間複雜度為 $O(|V|^3)$。

舉例：在一張地圖中，我們需要找到所有節點之間的最短路徑。

<img src="https://i.imgur.com/sy0W8Sc.png" width="200" height="150">

通過 Floyd-Warshall 算法，我們可以得到下面的最短路徑矩陣：

```
0 3 8 1
INF 0 INF 2
INF 4 0 INF
INF INF 5 0
```

其中，INF 表示無限大。例如，第一列中的數字 0，3，8，1 分別表示從 A 點到 A 點的距離為 0，A 點到 B 點的距離為 3，以此類推。