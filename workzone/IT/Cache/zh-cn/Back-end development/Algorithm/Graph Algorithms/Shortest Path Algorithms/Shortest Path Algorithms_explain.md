

Shortest Path Algorithms 是一种常用的图论算法，用于在图论中，找出一个特定起点与终点之间的最短路径。这种算法主要应用于路径导航、交通运输等方面。以下是三种常用的 Shortest Path Algorithms:

1. Dijkstra 算法
Dijkstra 算法是一种贪心算法，通过选择当前节点到起始节点距离最短的节点，来逐步构造最短路径。它的时间复杂度与图中边的数量有关，通常为 $O(|E| + |V|\log|V|)$。

举例：在一张地图中，我们需要从 A 点出发，前往 B 点，请问最短路径是哪条？

<img src="https://i.imgur.com/Z7v1zIC.png" width="200" height="150">

通过 Dijkstra 算法，我们可以得到 A - C - D - B 这条路径，其长度为 10。

2. Bellman-Ford 算法
Bellman-Ford 算法也是一种动态规划算法，通过对图中所有边进行 V - 1 次松弛操作，来找到最短路径。它的时间复杂度为 $O(|V||E|)$。

举例：在一张地图中，我们需要从 A 点出发，前往 B 点，请问最短路径是哪条？

<img src="https://i.imgur.com/IjOojgn.png" width="200" height="150">

通过 Bellman-Ford 算法，我们可以得到 A - C - B 这条路径，其长度为 6。

3. Floyd-Warshall 算法
Floyd-Warshall 算法是一种动态规划算法，用于解决所有节点之间的最短路径问题。它的时间复杂度为 $O(|V|^3)$。

举例：在一张地图中，我们需要找到所有节点之间的最短路径。

<img src="https://i.imgur.com/sy0W8Sc.png" width="200" height="150">

通过 Floyd-Warshall 算法，我们可以得到下面的最短路径矩阵：

```
0 3 8 1
INF 0 INF 2
INF 4 0 INF
INF INF 5 0
```

其中，INF 表示无限大。例如，第一列中的数字 0，3，8，1 分别表示从 A 点到 A 点的距离为 0，A 点到 B 点的距离为 3，以此类推。