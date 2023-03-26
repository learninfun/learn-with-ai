

Bellman-Ford Algorithm是用来解决单源最短路问题的一种算法，可以处理存在负权边的图。该算法基于对每一条边进行V-1次松弛操作(其中V是图中的节点数)，来得到从源点到所有节点的最短路程。

演算法过程中，提供了一个最短路程表，存储所有节点到源点的最小路程。该表一开始会被初始化为至少为无限大。 Bellman-Ford Algorithm 重复进行V-1次的下列回圈：

1. 对图中的每一条边，对其端点进行松弛操作(如果端点的最短路程可以被更新，就更新最短路程表中的路径长度)。

例如，考虑下面的图表示法，它有5个节点。

![Bellman-Ford Algorithm](https://i.imgur.com/ZORymh4.png)

假设起点是节点1。一开始初始化最短路程表为无限大， SOURCE = (0), OTHERS = (+∞)。

对每条边(Tx, Ty, c)进行V-1=4次的松弛操作，一直将节点3、4、5的最短路维护到最新值。

第1次：

    T1 = 0;
    T2 = +∞;
    T3 = +∞;
    T4 = +∞;
    T5 = +∞;

    SOURCE = (0), OTHERS = (T1 T2 T3 T4 T5)

    ...

    T3 = min(T3, T1 + 6) = 6;
    T2 = min(T2, T1 + 3) = 3;

    SOURCE = (0), OTHERS = (T1 3 6 T4 T5)
    

第2次：

    T1 = 0;
    T2 = 3;
    T3 = 6;
    T4 = +∞;
    T5 = +∞;

    SOURCE = (0), OTHERS = (T1 T2 T3 T4 T5)

    ...

    T4 = min(T4, T2 + (-1)) = 2;
    T5 = min(T5, T3 + 1) = 7;

    SOURCE = (0), OTHERS = (T1 3 6 2 7)

第3次：

    T1 = 0;
    T2 = 3;
    T3 = 6;
    T4 = 2;
    T5 = 7;

    SOURCE = (0), OTHERS = (T1 T2 T3 T4 T5)

    ...

    T5 = min(T5, T4 + 3) = 5;

    SOURCE = (0), OTHERS = (T1 3 6 2 5)

第4次：

    T1 = 0;
    T2 = 3;
    T3 = 6;
    T4 = 2;
    T5 = 5;

    SOURCE = (0), OTHERS = (T1 T2 T3 T4 T5)

    ...

Done!