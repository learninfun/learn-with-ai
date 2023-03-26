

以下是Dijkstra's Algorithm的重点：

1. 设定一个起点和终点来解决两点之间的最短路径。
2. 创建一个堆栈来储存已访问的节点和待访问的节点。
3. 初始化起点节点。
4. 对于每一个节点，设置一个到达该节点的“最短距离”和“最短路径”，然后添加到堆栈中。
5. 选择到目前为止最短的节点，设置其为当前节点，从堆栈中移除。
6. 检查该节点的相邻节点，如果通过当前节点到达相邻节点的距离短于之前所保存的“最短距离”，则更新该节点的“最短距离”和“最短路径”。
7. 重复步骤5-6，直到目标节点被访问或者堆栈被清空。

Dijkstra's Algorithm是一种解决图形问题的有效方法，它可以找到两点之间的最短路径。该算法基于“贪心”的策略，通过不断地选择当前最短的路径来搜寻该图形，直到所有节点都被访问过。该算法适用于没有负权重图的情况，并且可以通过最小堆栈实现高效率访问。