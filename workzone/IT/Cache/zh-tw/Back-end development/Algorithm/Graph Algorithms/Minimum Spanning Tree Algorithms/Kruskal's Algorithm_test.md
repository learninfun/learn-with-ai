

1. 給定一個無向圖，利用Kruskal's Algorithm找出最小生成樹。圖的邊權重為：[(1, 2, 5), (1, 3, 2), (2, 3, 1), (2, 4, 6), (3, 4, 3), (3, 5, 8), (4, 5, 4)]。

答案：最小生成樹的邊為[(2, 3, 1), (1, 3, 2), (4, 5, 4), (3, 4, 3)]，總權重為10。

2. 找出以下有向圖的最小生成樹，使用Kruskal's Algorithm。圖的邊權重為：[(1, 2, 5), (1, 3, 3), (2, 3, 2), (2, 4, 1), (3, 4, 2), (3, 5, 1), (4, 5, 1)]。

答案：這個問題無解。因為圖是有向圖，Kruskal's Algorithm只能處理無向圖。

3. 在下列的圖中，使用Kruskal's Algorithm找出最小生成樹。圖的邊權重為：[(1, 2, 1), (1, 3, 1), (3, 4, 1), (3, 5, 3), (4, 5, 2), (2, 4, 1)]。

答案：最小生成樹的邊為[(1, 2, 1), (1, 3, 1), (2, 4, 1), (4, 5, 2)]，總權重為5。

4. 找到以下圖的最小生成樹，使用Kruskal's Algorithm。圖的邊權重為：[(1, 2, 5), (2, 3, 8), (3, 4, 3), (1, 4, 1), (1, 3, 6), (2, 4, 2)]。

答案：最小生成樹的邊為[(1, 4, 1), (2, 4, 2), (3, 4, 3)]，總權重為6。

5. 給定一個無向圖，使用Kruskal's Algorithm找到最小生成樹。圖的邊權重如下：[(1, 2, 2), (2, 3, 4), (1, 4, 1), (4, 3, 3), (4, 5, 8), (3, 5, 6), (3, 6, 9), (5, 6, 5)]。

答案：最小生成樹的邊為[(1, 4, 1), (1, 2, 2), (2, 3, 4), (4, 5, 8), (5, 6, 5)]，總權重為20。