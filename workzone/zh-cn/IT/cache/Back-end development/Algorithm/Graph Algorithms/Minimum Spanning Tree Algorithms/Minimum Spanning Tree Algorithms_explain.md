

Minimum Spanning Tree Algorithms（最小生成树算法）是用来找寻一个连通无向图中，权值最小的生成树（即包含所有节点的树），其中权值可以是边的权重或是节点之间的距离。

以下是常被使用的Minimum Spanning Tree Algorithms：

1. Kruskal algorithm（克鲁斯卡尔算法）：将所有边依权值排序，从权值最小的边开始，依序加入生成树，直到所有节点都连通。

2. Prim's algorithm（普利姆算法）：以一个起始点开始，找到离该点最近的一个连接节点加入生成树中，然后继续找到与生成树相连的最小的边，加入生成树中，直到所有节点都连通。

举例来说，假设有以下的无向图： 

若采用Kruskal算法，最小生成树的边为： 

若采用Prim算法，最小生成树的边为： 

两种算法得到的最小生成树都是一样的，但是在不同的图形、节点数量上，两种算法的表现可能会有所不同。