

1. 将所有的边按照权值从小到大排序
2. 开始遍历所有边，如果遇到的边不会形成环路，就将该边加入最小生成树中
3. 判断环路的方法可以是利用Union-Find资料结构
4. 当所有边都遍历完毕或是最小生成树中的边数已达到其顶点数-1时结束
5. Kruskal's Algorithm具有贪心的思想，每次选择权值最小的边，是获得最小生成树的保证
6. 时间复杂度为O(ElogE)，其中E为边的数量，因为排序需要O(ElogE)的时间，接着进行E次寻找，每次寻找需要O(1)的时间