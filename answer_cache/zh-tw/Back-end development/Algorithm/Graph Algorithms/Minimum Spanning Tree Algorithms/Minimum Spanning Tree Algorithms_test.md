

1. 請實作 Kruskal 演算法找出下列圖的最小生成樹。
![image](https://i.imgur.com/9TD7b8e.png)
答案：{(A, C), (C, D), (D, E), (E, F), (A, B)}

2. 請實作 Prim 演算法找出下列圖的最小生成樹。
![image](https://i.imgur.com/5tKjJtH.png)
答案：{(A, D), (A, C), (C, B), (C, E), (E, F)}

3. 假設你有一個無向圖，但是其中有些邊是有向的，也就是說，你無法順利的去走過那些只有單向的路段。請實現 Kruskal 演算法過濾掉所有無法雙向通行的邊，只找出仍可構成 MST 的邊。
答案： {(A, B), (B, C), (B, D), (D, F), (D, E)}

4. 請實作 Boruvka 演算法找出下列圖的最小生成樹。
![image](https://i.imgur.com/yhp5Z5I.png)
答案：{(A, B), (B, D), (C, D), (D, E), (E, F)}

5. 請將下列圖例圖使用 Kruskal 演算法找出最小生成樹，但是加入特別的限制：所有的最短邊都只能夠使用一次，而非兩次。
![image](https://i.imgur.com/OY6oPHU.png)
答案：{(A, B), (B, C), (C, E), (C, F), (A, D)}