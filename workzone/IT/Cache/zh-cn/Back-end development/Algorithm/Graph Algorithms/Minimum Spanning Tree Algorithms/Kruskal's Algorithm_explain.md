

Kruskal's Algorithm是一种用于建立最小生成树（MST）的算法，MST是一种连通加权无向图的树，其权值总和最小。

Kruskal's Algorithm步骤如下：

1. 初始化MST为空。
2. 将图中所有节点分别视为独立的树。
3. 将图中的边按权值从小到大排序。
4. 从权值最小的边开始，依次加入MST中，直到MST包含所有节点或无法再加入边为止。
5. 返回MST。

以下是一个Kruskal's Algorithm的示例：

考虑下图，有5个节点和7条边。

![Kruskal's Algorithm Example](https://i.imgur.com/4Lf6ZCd.png)

按权值从小到大将边排序：(2,3) (2,4) (3,4) (1,2) (1,5) (4,5) (3,5)。

首先添加 (2,3)，MST为：

![Kruskal's Algorithm Example 2](https://i.imgur.com/2Xg20s1.png)

然后添加 (2,4)，MST为：

![Kruskal's Algorithm Example 3](https://i.imgur.com/YKmGv8t.png)

然后添加 (3,4)，MST为：

![Kruskal's Algorithm Example 4](https://i.imgur.com/ZbWjKFF.png)

然后添加 (1,2)，MST为：

![Kruskal's Algorithm Example 5](https://i.imgur.com/zBlEwS1.png)

然后添加 (1,5)，MST为：

![Kruskal's Algorithm Example 6](https://i.imgur.com/UyNZ059.png)

最后添加 (4,5)，MST为：

![Kruskal's Algorithm Example 7](https://i.imgur.com/eR5G177.png)

因此，MST是：

![Kruskal's Algorithm Example 8](https://i.imgur.com/JQuSkKj.png)

权值总和为 12。