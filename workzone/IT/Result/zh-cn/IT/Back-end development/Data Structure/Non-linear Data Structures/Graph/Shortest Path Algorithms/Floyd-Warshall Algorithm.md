+++
title = "Floyd-Warshall演算法"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Floyd-Warshall Algorithm的问题
{{< /ask_chatgpt >}}



1. 在一个图中，如何利用 Floyd-Warshall Algorithm 找出从每一个节点到其他节点的最短路径？

2. 如果一个图中存在负权重边，如何使用 Floyd-Warshall Algorithm 求最短路径？

3. 如果一个图中存在环，可以使用 Floyd-Warshall Algorithm 求最短路吗？为什么？

4. 如果一个图中存在多个源顶点（即，多个节点到其他节点的距离都要求最小），该如何使用 Floyd-Warshall Algorithm 处理？

5. 如何使用 Floyd-Warshall Algorithm 找出两个节点之间的最短路径？   

## 说明知识
{{< ask_chatgpt >}}
说明Floyd-Warshall Algorithm并举例
{{< /ask_chatgpt >}}



Floyd-Warshall算法是一种动态规划算法，用于在带权图中查找所有节点的最短路径。该算法的时间复杂度为O(n^3)，其中n是图中节点的数量。

具体实现步骤如下：

1. 创建一个n x n的二维数组D，并将所有元素初始化为无穷大（表示i和j之间没有边）。

2. 对于图中每个有权边（i，j）和边权w，将D[i][j]设置为w。

3. 对于每个节点i，将D[i][i]设置为0（表示节点到自身的距离为0）。

4. 使用三重循环迭代计算每个节点之间的最短路径，并更新D数组。具体来说，对于每个节点k，遍历所有节点i和j，如果从节点i到节点j的路径经过节点k会更短，则将D[i][j]更新为D[i][k] + D[k][j]。

5. 最终，数组D中包含所有节点之间的最短路径。

下面是一个简单的示例：

假设有以下带权图：

```
     12
(0)------>(1)
 |         /|\
10|          |
 |          | 1
\|/         |
(2)------>(3)
     3
```

根据上述步骤，我们可以进行如下的计算：

1. 初始化2维数组D：
```
   0 ∞ ∞ ∞
 ∞   0 ∞ ∞
 ∞ ∞   0 ∞
 ∞ ∞ ∞   0
```

2. 将边加入D：
```
   0 12 ∞ ∞
12   0 10  1
 ∞  10  0  3
 ∞   1  3  0
```

3. 对角线元素为0：
```
0 12 ∞ ∞
12 0 10 1
∞ 10 0 3
∞ 1 3 0
```

4. Floyd-Warshall算法计算最短路径：
```
0 12 22 13
12 0 10 1
22 10 0 3
13 1 3 0
```

从上述结果可以看出，节点0到节点1的最短路径长度为12，从节点0到节点3的最短路径为13，等等。   

## 汇总重点
{{< ask_chatgpt >}}
条列Floyd-Warshall Algorithm的重点
{{< /ask_chatgpt >}}



1. 使用DP（動態規劃）思想
2. 用矩陣表示每個節點之間的距離，可以初始化矩陣中不存在的路徑為無窮大或撇除
3. 使用三個嵌套循環來遍歷所有節點同時更新其距離
4. 循環時先比較走直接路徑跟轉移路徑哪一個更短，若轉移路徑更短則更新矩陣
5. 循環的順序很重要，需要先遍歷所有點，再考慮該點作為轉移節點時能夠到達的其他節點
6. 算法的時間複雜度是O(n^3)，空間複雜度是O(n^2)，其中n是節點的數量   

## 知识测验
{{< ask_chatgpt >}}
给我5题Floyd-Warshall Algorithm的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 給定一個有向加權圖，求任意兩點之間的最短距離，並找出這些最短距離中的最大值。

答案：

求出任意兩點之間的最短距離後，遍歷所有距離的值，找出其中的最大值即可。

2. 給定一個有向加權圖，求圖的直徑，即圖中最遠兩點之間的距離。

答案：

運用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，然後遍歷所有距離的值，找出其中的最大值即可。

3. 給定一個有向加權圖，該圖邊權值為正數，求圖中所有負環。

答案：

運用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，如果任意一個點到自己的距離是負數，則該點在負環上。

4. 給定一個有向加權圖，求任意兩點之間恰好經過一次某條特定邊的最短路徑。

答案：

用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，然後將該特定邊設為一個很大的值，再求一次最短路徑，即可得到恰好經過該特定邊的最短路徑。

5. 給定一個無向加權圖，求任意兩點之間的最短路徑，但要求經過的邊權值的和不能超過一個給定值。

答案：

將邊權值視為花費，運用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，然後在求最短路徑時限制經過的邊權值之和不超過給定值，可以用 Dijkstra Algorithm 實現。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Floyd-Warshall Algorithm的网络数据
{{< /ask_chatgpt >}}



1. https://www.geeksforgeeks.org/floyd-warshall-algorithm-dp-16/

这是一篇GeeksforGeeks网站上的Floyd-Warshall Algorithm教学文章，详细介绍了此演算法的原理和示范程式码，并以实际的图形表示法解释演算法的操作过程。

2. https://www.tutorialspoint.com/Floyd-Warshall-Algorithm

这是一篇Tutorials Point网站上的Floyd-Warshall Algorithm教学文章，同样提供了演算法的原理和示范程式码，并以实例示范如何应用在不同种类的图形问题中。

3. https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm

这是维基百科上的Floyd-Warshall Algorithm介绍页面，其中介绍了演算法的历史、使用情况、算法流程和时间复杂度等。此外，也有图形和数学示例解释演算法中所涉及的基本概念和处理方法。

4. https://www.cs.usfca.edu/~galles/visualization/Floyd.html

这是一个互动式网站演示Floyd-Warshall Algorithm操作的过程，可自由设定输入图形的大小和权重，让使用者更好地了解演算法中每个步骤的细节。

5. https://www.youtube.com/watch?v=4BUDmYf5eFA

这是一个由YouTube上的MyCodesSchool频道上传的Floyd-Warshall Algorithm介绍影片，有详细的文字解说和画面演示，用生动的方式让观众理解演算法的核心思想和应用场景。   

