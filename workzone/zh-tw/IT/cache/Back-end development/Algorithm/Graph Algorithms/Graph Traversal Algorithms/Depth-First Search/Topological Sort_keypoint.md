1. 定義：Topological Sort 是一種對有向無環圖（DAG）進行節點排序的演算法。

2. 應用：Topological Sort 常用於尋找項目間的先後關係，例如工作流程中的先後順序、課程選修的先後等。

3. 步驟：Topological Sort 的步驟為尋找 DAG 的一個非循環路徑，並依照該路徑的順序將節點排序。

4. 實現方法：Topological Sort 可以使用 DFS 或 BFS 兩種方式來實現，其中 BFS 的時間複雜度較低。

5. 結果：Topological Sort 的結果並不唯一，可能存在多種排序結果。

6. 特殊情況：如果 DAG 中存在環路則無法進行 Topological Sort，這時需要進行環路檢測或者使用其他方法進行排序。

7. 應用範例：如下圖所示，該 DAG 表示六個項目之間的先後關係，使用 Topological Sort 可以得到的一個排序結果為 B, D, A, C, F, E。

   ![image.png](https://i.imgur.com/XnoZwqa.png)