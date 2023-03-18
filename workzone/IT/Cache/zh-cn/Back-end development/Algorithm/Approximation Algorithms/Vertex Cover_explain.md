

Vertex Cover是一种图论中的问题，目的是找到一个最小的集合，可以覆盖一张无向图中所有的边。换句话说，就是找到一些点，让这些点所相连的边涵盖了整个图。

例如，下图中，有一个六个节点的无向图，其中的所有边都用虚线标记。如果要找到一个Vertex Cover，可以选择以下三个点：1、3和5。这三个点所连接的边（用实线表示）可以涵盖整个图中的所有边。

![Vertex Cover example](https://miro.medium.com/max/3248/1*FovPhPTOG64NKGpbv1zExQ.png)

在这个例子中，这个Vertex Cover的大小为3，因为我们只需要三个节点就可以完全涵盖整个图了。Vertex Cover问题是一个NP完全问题，因此通常需要使用近似算法进行求解。