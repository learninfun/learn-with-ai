

Bellman-Ford Algorithm是一种求解单源最短路径问题的动态规划算法，它可以处理含有负权值的图形。这种算法经常用于路由演算法，它可以在 O（V·E）的时间复杂度内完成任务，其中 V 和 E 分别是图形中的节点和边的数量。

该算法通过将节点的最短路径设置为无限大，然后遍历图形并使用动态规划更新最短路径。如果更新某个节点的最短路径时发现该节点到其相邻节点的距离比原先的距离更短，则更新该节点的最短路径。

以下是一个简单的例子，展示Bellman-Ford算法的运作过程：

考虑以下图形，其中顶点编号为1，2，3，4，5，6：

![image.png](attachment:image.png)

1. 设置每个节点的最短路径为∞，并设置起点节点（1）的最短路径为0。
2. 遍历图形中的每个边，并使用动态规划更新最短路径。
3. 现在，我们已经开始更新最短路径了。将节点1的最短路径分别设置为0，其他节点设置为∞。
4. 将节点1与节点2之间的边长更新为4，因为4比∞更小。同理，将节点1与节点3之间的边更新为3，节点1与节点4之间的边更新为5，节点1与节点5之间的边更新为2，节点1与节点6之间的边更新为1。
5. 接下来，我们遍历图形的所有边，并再次更新最短路径。现在，我们可以看到最短路径已经更新了，因为节点2到节点5的距离为7，而节点1到节点5的距离等于3 + 2 = 5。
6. 再次遍历所有边，更新最短路径。现在已经无法更新任何节点的最短路径了，因为所有节点的距离已经达到最小值。

因此，Bellman-Ford算法可以找到从起点到每个节点的最短路径，即运行此算法将可以找到上图中，节点1到其他节点的最短路径。