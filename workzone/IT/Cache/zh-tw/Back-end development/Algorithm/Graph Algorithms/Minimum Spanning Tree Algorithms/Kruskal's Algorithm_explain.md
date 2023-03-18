

Kruskal's Algorithm是一種用於建立最小生成樹（MST）的算法，MST是一種連通加權無向圖的樹，其權值總和最小。

Kruskal's Algorithm步驟如下：

1. 初始化MST為空。
2. 將圖中所有節點分別視為獨立的樹。
3. 將圖中的邊按權值從小到大排序。
4. 從權值最小的邊開始，依次加入MST中，直到MST包含所有節點或無法再加入邊為止。
5. 返回MST。

以下是一個Kruskal's Algorithm的示例：

考慮下圖，有5個節點和7條邊。

![Kruskal's Algorithm Example](https://i.imgur.com/4Lf6ZCd.png)

按權值從小到大將邊排序：(2,3) (2,4) (3,4) (1,2) (1,5) (4,5) (3,5)。

首先添加 (2,3)，MST為：

![Kruskal's Algorithm Example 2](https://i.imgur.com/2Xg20s1.png)

然後添加 (2,4)，MST為：

![Kruskal's Algorithm Example 3](https://i.imgur.com/YKmGv8t.png)

然後添加 (3,4)，MST為：

![Kruskal's Algorithm Example 4](https://i.imgur.com/ZbWjKFF.png)

然後添加 (1,2)，MST為：

![Kruskal's Algorithm Example 5](https://i.imgur.com/zBlEwS1.png)

然後添加 (1,5)，MST為：

![Kruskal's Algorithm Example 6](https://i.imgur.com/UyNZ059.png)

最後添加 (4,5)，MST為：

![Kruskal's Algorithm Example 7](https://i.imgur.com/eR5G177.png)

因此，MST是：

![Kruskal's Algorithm Example 8](https://i.imgur.com/JQuSkKj.png)

權值總和為 12。