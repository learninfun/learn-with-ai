## 習題預習
{{< ask_chatgpt >}}
給我5題Floyd-Warshall Algorithm的問題
{{< /ask_chatgpt >}}



1. 在一個圖中，如何利用 Floyd-Warshall Algorithm 找出從每一個節點到其他節點的最短路徑？

2. 如果一個圖中存在負權重邊，如何使用 Floyd-Warshall Algorithm 求最短路徑？

3. 如果一個圖中存在環，可以使用 Floyd-Warshall Algorithm 求最短路嗎？為什麼？

4. 如果一個圖中存在多個源頂點（即，多個節點到其他節點的距離都要求最小），該如何使用 Floyd-Warshall Algorithm 處理？

5. 如何使用 Floyd-Warshall Algorithm 找出兩個節點之間的最短路徑？   

## 說明知識
{{< ask_chatgpt >}}
說明Floyd-Warshall Algorithm並舉例
{{< /ask_chatgpt >}}



Floyd-Warshall算法是一種動態規划算法，用於在帶權圖中查找所有節點的最短路徑。該算法的時間複雜度為O(n^3)，其中n是圖中節點的數量。

具體實現步驟如下：

1. 創建一個n x n的二維數組D，並將所有元素初始化為無窮大（表示i和j之間沒有邊）。

2. 對於圖中每個有權邊（i，j）和邊權w，將D[i][j]設置為w。

3. 對於每個節點i，將D[i][i]設置為0（表示節點到自身的距離為0）。

4. 使用三重循環迭代計算每個節點之間的最短路徑，並更新D數組。具體來說，對於每個節點k，遍歷所有節點i和j，如果從節點i到節點j的路徑經過節點k會更短，則將D[i][j]更新為D[i][k] + D[k][j]。

5. 最終，數組D中包含所有節點之間的最短路徑。

下面是一個簡單的示例：

假設有以下帶權圖：

```
     12
(0)------>(1)
 |         /|\
10|          |
 |          | 1
\|/         |
(2)------>(3)
     3
```

根據上述步驟，我們可以進行如下的計算：

1. 初始化2維數組D：
```
   0 ∞ ∞ ∞
 ∞   0 ∞ ∞
 ∞ ∞   0 ∞
 ∞ ∞ ∞   0
```

2. 將邊加入D：
```
   0 12 ∞ ∞
12   0 10  1
 ∞  10  0  3
 ∞   1  3  0
```

3. 對角線元素為0：
```
0 12 ∞ ∞
12 0 10 1
∞ 10 0 3
∞ 1 3 0
```

4. Floyd-Warshall算法計算最短路徑：
```
0 12 22 13
12 0 10 1
22 10 0 3
13 1 3 0
```

從上述結果可以看出，節點0到節點1的最短路徑長度為12，從節點0到節點3的最短路徑為13，等等。   

## 彙總重點
{{< ask_chatgpt >}}
條列Floyd-Warshall Algorithm的重點
{{< /ask_chatgpt >}}



1. 使用DP（動態規劃）思想
2. 用矩陣表示每個節點之間的距離，可以初始化矩陣中不存在的路徑為無窮大或撇除
3. 使用三個嵌套循環來遍歷所有節點同時更新其距離
4. 循環時先比較走直接路徑跟轉移路徑哪一個更短，若轉移路徑更短則更新矩陣
5. 循環的順序很重要，需要先遍歷所有點，再考慮該點作為轉移節點時能夠到達的其他節點
6. 算法的時間複雜度是O(n^3)，空間複雜度是O(n^2)，其中n是節點的數量   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Floyd-Warshall Algorithm的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 給定一個有向加權圖，求任意兩點之間的最短距離，並找出這些最短距離中的最大值。

答案：

求出任意兩點之間的最短距離後，遍歷所有距離的值，找出其中的最大值即可。

2. 給定一個有向加權圖，求圖的直徑，即圖中最遠兩點之間的距離。

答案：

運用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，然後遍歷所有距離的值，找出其中的最大值即可。

3. 給定一個有向加權圖，該圖邊權值為正數，求圖中所有負環。

答案：

運用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，如果任意一個點到自己的距離是負數，則該點在負環上。

4. 給定一個有向加權圖，求任意兩點之間恰好經過一次某條特定邊的最短路徑。

答案：

用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，然後將該特定邊設為一個很大的值，再求一次最短路徑，即可得到恰好經過該特定邊的最短路徑。

5. 給定一個無向加權圖，求任意兩點之間的最短路徑，但要求經過的邊權值的和不能超過一個給定值。

答案：

將邊權值視為花費，運用 Floyd-Warshall Algorithm 求出任意兩點之間的最短距離，然後在求最短路徑時限制經過的邊權值之和不超過給定值，可以用 Dijkstra Algorithm 實現。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Floyd-Warshall Algorithm的網路資料
{{< /ask_chatgpt >}}



1. https://www.geeksforgeeks.org/floyd-warshall-algorithm-dp-16/

這是一篇GeeksforGeeks網站上的Floyd-Warshall Algorithm教學文章，詳細介紹了此演算法的原理和示範程式碼，並以實際的圖形表示法解釋演算法的操作過程。

2. https://www.tutorialspoint.com/Floyd-Warshall-Algorithm

這是一篇Tutorials Point網站上的Floyd-Warshall Algorithm教學文章，同樣提供了演算法的原理和示範程式碼，並以實例示範如何應用在不同種類的圖形問題中。

3. https://en.wikipedia.org/wiki/Floyd%E2%80%93Warshall_algorithm

這是維基百科上的Floyd-Warshall Algorithm介紹頁面，其中介紹了演算法的歷史、使用情況、算法流程和時間複雜度等。此外，也有圖形和數學示例解釋演算法中所涉及的基本概念和處理方法。

4. https://www.cs.usfca.edu/~galles/visualization/Floyd.html

這是一個互動式網站演示Floyd-Warshall Algorithm操作的過程，可自由設定輸入圖形的大小和權重，讓使用者更好地瞭解演算法中每個步驟的細節。

5. https://www.youtube.com/watch?v=4BUDmYf5eFA

這是一個由YouTube上的MyCodesSchool頻道上傳的Floyd-Warshall Algorithm介紹影片，有詳細的文字解說和畫面演示，用生動的方式讓觀眾理解演算法的核心思想和應用場景。   

