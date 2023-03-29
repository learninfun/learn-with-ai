1. 請問什麼是圖形資料庫，它與傳統關聯式資料庫有何不同之處？
2. 在圖形資料庫中，什麼是節點和邊，它們的作用是什麼？
3. 如何在圖形資料庫中查詢所有已知子節點之間的最短路徑？
4. 如何在圖形資料庫中表示圖論經典問題（如最小生成樹、最長路徑等）？
5. 在圖形資料庫中，如何避免遞迴問題，防止查詢進入無限迴圈？

答案：

1. 圖形資料庫是一種專為圖形資料而設計的資料庫，它使用圖形模型來儲存資料和關係，與傳統關聯式資料庫的最大區別在於可以更好地解決需要分析大量互相連接的資料的應用場景。
2. 節點（Node）在圖形資料庫中代表實體，邊（Edge）則代表實體之間的關係。節點可以有屬性，邊可以有權重以及方向性。節點和邊是圖形資料庫的核心要素，它們的作用是描述資料間的關係。
3. 在圖形資料庫中，有專門的演算法來查詢所有已知子節點之間的最短路徑，例如 Dijkstra 和 A* 等演算法。這些演算法可以計算出圖形資料庫中兩節點之間的最短路徑，從而幫助應用程序解決複雜的路徑問題。
4. 在圖形資料庫中，可以使用圖論基本問題的算法（例如 Kruskal 或 Prim 算法來計算最小生成樹，DAG 算法來計算最長路徑等）來表示圖論經典問題。在圖形資料庫中，密切關注節點和邊的屬性以及權重是非常重要的。
5. 為了避免遞迴問題和無限迴圈，圖形資料庫通常會採用「追溯標記」等技術。當進行遞迴遍歷時，節點會被標記為進入或已經訪問過，如果訪問到已經過訪問過的節點，就可以直接中斷遞迴，從而避免無限迴圈。