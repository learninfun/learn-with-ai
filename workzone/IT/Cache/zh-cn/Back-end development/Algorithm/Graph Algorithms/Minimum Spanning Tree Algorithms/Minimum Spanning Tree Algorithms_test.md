

1. 请实作 Kruskal 演算法找出下列图的最小生成树。
![image](https://i.imgur.com/9TD7b8e.png)
答案：{(A, C), (C, D), (D, E), (E, F), (A, B)}

2. 请实作 Prim 演算法找出下列图的最小生成树。
![image](https://i.imgur.com/5tKjJtH.png)
答案：{(A, D), (A, C), (C, B), (C, E), (E, F)}

3. 假设你有一个无向图，但是其中有些边是有向的，也就是说，你无法顺利的去走过那些只有单向的路段。请实现 Kruskal 演算法过滤掉所有无法双向通行的边，只找出仍可构成 MST 的边。
答案： {(A, B), (B, C), (B, D), (D, F), (D, E)}

4. 请实作 Boruvka 演算法找出下列图的最小生成树。
![image](https://i.imgur.com/yhp5Z5I.png)
答案：{(A, B), (B, D), (C, D), (D, E), (E, F)}

5. 请将下列图例图使用 Kruskal 演算法找出最小生成树，但是加入特别的限制：所有的最短边都只能够使用一次，而非两次。
![image](https://i.imgur.com/OY6oPHU.png)
答案：{(A, B), (B, C), (C, E), (C, F), (A, D)}