

Bellman-Ford Algorithm是用來解決單源最短路問題的一種算法，可以處理存在負權邊的圖。該算法基於對每一條邊進行V-1次松弛操作(其中V是圖中的節點數)，來得到從源點到所有節點的最短路程。

演算法過程中，提供了一個最短路程表，存儲所有節點到源點的最小路程。該表一開始會被初始化為至少為無限大。 Bellman-Ford Algorithm 重複進行V-1次的下列迴圈：

1. 對圖中的每一條邊，對其端點進行松弛操作(如果端點的最短路程可以被更新，就更新最短路程表中的路徑長度)。

例如，考慮下面的圖表示法，它有5個節點。

![Bellman-Ford Algorithm](https://i.imgur.com/ZORymh4.png)

假設起點是節點1。一開始初始化最短路程表為無限大， SOURCE = (0), OTHERS = (+∞)。

對每條邊(Tx, Ty, c)進行V-1=4次的松弛操作，一直將節點3、4、5的最短路維護到最新值。

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