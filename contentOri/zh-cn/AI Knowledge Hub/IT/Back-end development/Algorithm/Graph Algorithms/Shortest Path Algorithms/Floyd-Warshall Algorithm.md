+++
title = "Floyd-Warshall演算法"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Floyd-Warshall Algorithm的问题
{{< /ask_chatgpt >}}



1. 在无向图中，如何找出从给定源节点到所有其他节点的最短路径？
2. 在有向图中，如何找出所有节点彼此之间的最短路径？
3. 在有向图中，如何找出所有节点彼此之间的最长路径？
4. 在一个网格图中，如何找出所有格子之间的最短路径？
5. 在一个带权重的有向图中，如何找出从某个叶子节点到根节点的最短路径？   

## 说明知识
{{< ask_chatgpt >}}
说明Floyd-Warshall Algorithm并举例
{{< /ask_chatgpt >}}



Floyd-Warshall Algorithm（弗洛伊德-沃舍尔算法）是解决图中所有节点之间最短路径的经典算法之一，它是一种动态规划算法，以及解决所有点对之间最短路径的做法，同时也适用于带有负边权的图。

算法的基本思想是维护一个二维数组D，用来记录任意两点之间的最短路径长度。初始时，D[i][j]表示从节点i到节点j的距离，如果i和j不直接相邻，那么D[i][j]=正无穷。算法每进行一轮迭代，便会将D[i][j]和D[i][k]+D[k][j]（相当于以k作为中介）进行比较更新，最终得到任意两点之间的最短路径长度。

以下是弗洛伊德-沃舍尔算法的伪代码：

```
1. for k from 1 to |V|
2.      for i from 1 to |V|
3.          for j from 1 to |V|
4.              if D[i][j] > D[i][k] + D[k][j] then
5.                  D[i][j] = D[i][k] + D[k][j]
```

举例说明：
  考虑下图：

![image.png](attachment:image.png)

构建邻接矩阵如下：

```
    1   2   3   4   5   6
  -------------------------
1 | 0   1   ∞   ∞   5   ∞
2 | ∞   0   1   ∞   ∞   ∞
3 | ∞   ∞   0   2   ∞   ∞
4 | ∞   ∞   ∞   0   ∞   1
5 | ∞   ∞   ∞   1   0   2
6 | ∞   ∞   ∞   ∞   ∞   0
```
其中，∞表示两点之间没有连通路径。

接下来，我们可以运用Floyd-Warshall Algorithm算法来求得任意两点之间的最短路径长度，具体过程如下：

第一轮迭代：k=1
```
D[1][2]=min(∞, ∞+1)=1, 
D[1][5]=min(5, ∞+1)=6,
D[2][3]=min(∞,∞+1)=1,
D[3][4]=min(∞,∞+2)=2,
D[4][6]=min(∞,∞+1)=1,
D[5][6]=min(∞,∞+2)=2
```

```
    1   2   3   4   5   6
  -------------------------
1 | 0   1   ∞   ∞   5   ∞
2 | ∞   0   1   ∞   ∞   ∞
3 | ∞   2   0   2   ∞   ∞
4 | ∞   ∞   ∞   0   ∞   1
5 | ∞   ∞   ∞   1   0   2
6 | ∞   ∞   ∞   ∞   ∞   0
```

第二轮迭代：k=2
```
D[1][3]=min(∞,1+1)=2,
D[1][6]=min(∞,5+2)=7,
D[2][4]=min(1,1+∞)=1,
D[3][6]=min(∞,2+2)=4,
D[4][5]=min(∞,1+2)=3
```
```
    1   2   3   4   5   6
  -------------------------
1 | 0   1   2   ∞   5   7
2 | ∞   0   1   2   ∞   3
3 | ∞   2   0   2   ∞   4
4 | ∞   ∞   ∞   0   3   1
5 | ∞   ∞   ∞   1   0   2
6 | ∞   ∞   ∞   ∞   ∞   0
```

第三轮迭代：k=3
```
D[1][4]=min(∞,2+2)=4,
D[2][5]=min(∞,1+2)=3,
D[3][6]=min(2,2+2)=2,
D[4][6]=min(1,3+1)=1,
D[5][3]=min(∞,2+2)=2
```
```
    1   2   3   4   5   6
  -------------------------
1 | 0   1   2   4   5   6
2 | ∞   0   1   2   3   3
3 | ∞   2   0   2   4   2
4 | ∞   ∞   ∞   0   3   1
5 | ∞   ∞   2   1   0   2
6 | ∞   ∞   ∞   ∞   ∞   0
```

第四轮迭代：k=4
```
D[2][6]=min(3,3+1)=3,
D[3][5]=min(2,2+4)=2,
D[4][2]=min(2,3+1)=2,
D[4][6]=min(1,3+1)=1,
D[5][1]=min(∞,2+5)=7
```
```
    1   2   3   4   5   6
  -------------------------
1 | 0   1   2   4   5   6
2 | ∞   0   1   2   3   3
3 | ∞   2   0   2   4   2
4 | ∞   2   3   0   3   1
5 | 7   6   2   1   0   2
6 | ∞   ∞   ∞   ∞   ∞   0
```

最终的结果是：

```
    1   2   3   4   5   6
  -------------------------
1 | 0   1   2   4   5   6
2 | ∞   0   1   2   3   3
3 | ∞   2   0   2   4   2
4 | ∞   2   3   0   3   1
5 | 7   6   2   1   0   2
6 | ∞   ∞   ∞   ∞   ∞   0
```

因此，我们可以得到所有点对之间的最短距离。例如，点1到点6的最短距离为6，点2到点1的最短距离为∞，点4到点5的最短距离为3 等。

总结：Floyd-Warshall Algorithm算法对有些权值为负的图仍然有效，并且由于操作矩阵，理论上算法时间复杂度为O(n^3)，适用于处理规模不是特别大的图，而且不受图形中负边权的影响，非常适合多源最短路径问题。   

## 汇总重点
{{< ask_chatgpt >}}
条列Floyd-Warshall Algorithm的重点
{{< /ask_chatgpt >}}



1. Floyd-Warshall Algorithm（F-W算法）是一種動態規劃算法，用於解決全源最短路問題。
2. F-W算法利用矩陣來存儲每對節點之間的最短路長度。
3. 矩陣中每個元素的值代表從i到j的最短路長度，其中i和j是矩陣下標。
4. F-W算法每次更新矩陣中保存的最短路長度，以便找出每對節點之間的最短路。
5. F-W算法的核心思想是利用動態規劃思想，將問題分解為子問題依次求解。
6. 在F-W算法中，我們先將圖表示成一個帶權有向圖（即邊帶權），然後利用矩陣來保存最短路長度。
7. F-W算法主要包括初始化矩陣、遍歷圖、更新矩陣等步驟。
8. 儘管F-W算法的時間複雜度較高（為O(N^3)），但它可以處理帶權有向圖中任意兩點之間的最短路問題，並且效果非常穩定。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Floyd-Warshall Algorithm的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 給定一張有向帶權圖，求從任意一點到任意一點的最短路徑。如果兩點之間沒有路徑，輸出inf。
```
程式碼如下：
void Floyd() {
    for (int k = 1; k <= n; k++) {
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                if (d[i][k] != INF && d[k][j] != INF && d[i][j] > d[i][k] + d[k][j]) {
                    d[i][j] = d[i][k] + d[k][j];
                }
            }
        }
    }
}
```

2. 給定一張無向帶權圖，求圖中的最小生成樹。如果圖不連通，則輸出-1。
```
程式碼如下：
void Floyd() {
    for (int k = 1; k <= n; k++) {
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                if (d[i][k] != INF && d[k][j] != INF && d[i][j] > d[i][k] + d[k][j]) {
                    d[i][j] = d[i][k] + d[k][j];
                }
            }
        }
    }
}
```

3. 給定一張有向帶權圖，求從任意一點到任意一點的最短路徑，如果存在負權環，輸出-1。
```
程式碼如下：
int Floyd() {
    for (int k = 1; k <= n; k++) {
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                if (d[i][k] != INF && d[k][j] != INF && d[i][j] > d[i][k] + d[k][j]) {
                    d[i][j] = d[i][k] + d[k][j];
                }
            }
        }
    }
    for (int k = 1; k <= n; k++) {
        if (d[k][k] < 0) { // 存在負權環
            return -1;
        }
    }
    return 0;
}
```

4. 給定一張有向帶權圖，求從每個點出發能到達的所有點之間的最小總權值和。
```
程式碼如下：
void Floyd() {
    for (int k = 1; k <= n; k++) {
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                if (d[i][k] != INF && d[k][j] != INF && d[i][j] > d[i][k] + d[k][j]) {
                    d[i][j] = d[i][k] + d[k][j];
                }
            }
        }
    }
    for (int i = 1; i <= n; i++) {
        int sum = 0;
        for (int j = 1; j <= n; j++) {
            if (d[i][j] != INF) {
                sum += d[i][j];
            }
        }
        printf("%d ", sum);
    }
}

5. 給定一張有向帶權圖，求從任意一點到任意一點經過正好k條邊的最短路徑。
```
程式碼如下：
void Floyd(int k) {
    for (int l = 1; l <= k; l++) { // l表示經過l條邊
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= n; j++) {
                for (int m = 1; m <= n; m++) {
                    if (d[i][m] != INF && d[m][j] != INF && d[i][j] > d[i][m] + d[m][j]) {
                        d[i][j] = d[i][m] + d[m][j];
                    }
                }
            }
        }
    }
}
```   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Floyd-Warshall Algorithm的网络数据
{{< /ask_chatgpt >}}



以下是五篇Floyd-Warshall Algorithm的网路资料：

1. https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm 
这是Floyd-Warshall Algorithm的维基百科页面。它解释了Floyd-Warshall算法的工作原理以及如何实现它。这是一个完整的介绍，从基本概念到具体实践，非常实用。

2. https://www.geeksforgeeks.org/floyd-warshall-algorithm-dp-16/ 
这是GeeksforGeeks的一篇文章，介绍了Floyd-Warshall算法的设计和实现。这篇文章还包括了一些范例和解释，可帮助读者更好地理解和应用该算法。

3. https://www.tutorialspoint.com/Floyd-Warshall-algorithm
这是TutorialsPoint的一篇关于Floyd-Warshall算法的教程。它解释了算法的基本原理，以及如何将其应用于解决实际问题。教程还包括代码示例和练习题，可帮助读者获得更多实战经验。

4. https://www.programiz.com/dsa/floyd-warshall-algorithm
这是Programiz的Floyd-Warshall算法教程，它是针对初学者设计的。教程以简单易懂的方式解释了算法的基本概念和实现方法。而且，您可以在这里找到许多范例和练习，使您可以透过实战获得更多的经验。

5. https://www.hackerearth.com/practice/algorithms/graphs/shortest-path-algorithms/tutorial/
这是HackerEarth的一篇教程，详细介绍了Floyd-Warshall算法的工作原理和应用。这篇文章还包括了一些问题，可以测试您对于该算法的理解程度。如果您想要深入了解Floyd-Warshall算法，这篇文章是一个不错的选择。   

