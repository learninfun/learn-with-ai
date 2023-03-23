+++
title = "克鲁斯卡尔演算法"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Kruskal's Algorithm的问题
{{< /ask_chatgpt >}}



1. 使用Kruskal的演算法，在以下边权重图形中找到最小生成树:
      
      **图片:**

2. 在一个Galaxy中，有N个星球，M条航线，每条航线有不同的费用，请使用Kruskal演算法找出一条最小生成树。

3. 一家公司需要将他们的厂房连接起来，厂房数量为N，他们想用最小的代价去建立连接，请透过Kruskal演算法解决此问题。

4. 假设您正在建立一个网路，其中您需要用光纤线连接数百个建筑物。每个建筑物需要使用不同的专用光纤线进行连接，并且每条连接的费用不同，请利用Kruskal演算法找出最小成本的光纤线之连接方式。

5. 在一个城市里，有许多地点需要连接，建立通路的费用也不同，请使用Kruskal演算法找出连接全部地点的最小成本。   

## 说明知识
{{< ask_chatgpt >}}
说明Kruskal's Algorithm并举例
{{< /ask_chatgpt >}}



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

## 汇总重点
{{< ask_chatgpt >}}
条列Kruskal's Algorithm的重点
{{< /ask_chatgpt >}}



1. 将所有的边按照权值从小到大排序
2. 开始遍历所有边，如果遇到的边不会形成环路，就将该边加入最小生成树中
3. 判断环路的方法可以是利用Union-Find资料结构
4. 当所有边都遍历完毕或是最小生成树中的边数已达到其顶点数-1时结束
5. Kruskal's Algorithm具有贪心的思想，每次选择权值最小的边，是获得最小生成树的保证
6. 时间复杂度为O(ElogE)，其中E为边的数量，因为排序需要O(ElogE)的时间，接着进行E次寻找，每次寻找需要O(1)的时间   

## 知识测验
{{< ask_chatgpt >}}
给我5题Kruskal's Algorithm的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 给定一个无向图，利用Kruskal's Algorithm找出最小生成树。图的边权重为：[(1, 2, 5), (1, 3, 2), (2, 3, 1), (2, 4, 6), (3, 4, 3), (3, 5, 8), (4, 5, 4)]。

答案：最小生成树的边为[(2, 3, 1), (1, 3, 2), (4, 5, 4), (3, 4, 3)]，总权重为10。

2. 找出以下有向图的最小生成树，使用Kruskal's Algorithm。图的边权重为：[(1, 2, 5), (1, 3, 3), (2, 3, 2), (2, 4, 1), (3, 4, 2), (3, 5, 1), (4, 5, 1)]。

答案：这个问题无解。因为图是有向图，Kruskal's Algorithm只能处理无向图。

3. 在下列的图中，使用Kruskal's Algorithm找出最小生成树。图的边权重为：[(1, 2, 1), (1, 3, 1), (3, 4, 1), (3, 5, 3), (4, 5, 2), (2, 4, 1)]。

答案：最小生成树的边为[(1, 2, 1), (1, 3, 1), (2, 4, 1), (4, 5, 2)]，总权重为5。

4. 找到以下图的最小生成树，使用Kruskal's Algorithm。图的边权重为：[(1, 2, 5), (2, 3, 8), (3, 4, 3), (1, 4, 1), (1, 3, 6), (2, 4, 2)]。

答案：最小生成树的边为[(1, 4, 1), (2, 4, 2), (3, 4, 3)]，总权重为6。

5. 给定一个无向图，使用Kruskal's Algorithm找到最小生成树。图的边权重如下：[(1, 2, 2), (2, 3, 4), (1, 4, 1), (4, 3, 3), (4, 5, 8), (3, 5, 6), (3, 6, 9), (5, 6, 5)]。

答案：最小生成树的边为[(1, 4, 1), (1, 2, 2), (2, 3, 4), (4, 5, 8), (5, 6, 5)]，总权重为20。   

