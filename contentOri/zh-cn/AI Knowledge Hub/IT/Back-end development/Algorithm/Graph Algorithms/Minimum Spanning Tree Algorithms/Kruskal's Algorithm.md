+++
title = "克魯斯卡爾演算法"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Kruskal's Algorithm的中文問題
{{< /ask_chatgpt >}}



1. 使用Kruskal的演算法，在以下邊權重圖形中找到最小生成樹:
      
      **圖片:**

2. 在一個Galaxy中，有N個星球，M條航線，每條航線有不同的費用，請使用Kruskal演算法找出一條最小生成樹。

3. 一家公司需要將他們的廠房連接起來，廠房數量為N，他們想用最小的代價去建立連接，請透過Kruskal演算法解決此問題。

4. 假設您正在建立一個網路，其中您需要用光纖線連接數百個建築物。每個建築物需要使用不同的專用光纖線進行連接，並且每條連接的費用不同，請利用Kruskal演算法找出最小成本的光纖線之連接方式。

5. 在一個城市裡，有許多地點需要連接，建立通路的費用也不同，請使用Kruskal演算法找出連接全部地點的最小成本。   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Kruskal's Algorithm並舉例
{{< /ask_chatgpt >}}



Kruskal's Algorithm是一種用於建立最小生成樹（MST）的算法，MST是一種連通加權無向圖的樹，其權值總和最小。

Kruskal's Algorithm步驟如下：

1. 初始化MST為空。
2. 將圖中所有節點分別視為獨立的樹。
3. 將圖中的邊按權值從小到大排序。
4. 從權值最小的邊開始，依次加入MST中，直到MST包含所有節點或無法再加入邊為止。
5. 返回MST。

以下是一個Kruskal's Algorithm的示例：

考慮下圖，有5個節點和7條邊。

![Kruskal's Algorithm Example](https://i.imgur.com/4Lf6ZCd.png)

按權值從小到大將邊排序：(2,3) (2,4) (3,4) (1,2) (1,5) (4,5) (3,5)。

首先添加 (2,3)，MST為：

![Kruskal's Algorithm Example 2](https://i.imgur.com/2Xg20s1.png)

然後添加 (2,4)，MST為：

![Kruskal's Algorithm Example 3](https://i.imgur.com/YKmGv8t.png)

然後添加 (3,4)，MST為：

![Kruskal's Algorithm Example 4](https://i.imgur.com/ZbWjKFF.png)

然後添加 (1,2)，MST為：

![Kruskal's Algorithm Example 5](https://i.imgur.com/zBlEwS1.png)

然後添加 (1,5)，MST為：

![Kruskal's Algorithm Example 6](https://i.imgur.com/UyNZ059.png)

最後添加 (4,5)，MST為：

![Kruskal's Algorithm Example 7](https://i.imgur.com/eR5G177.png)

因此，MST是：

![Kruskal's Algorithm Example 8](https://i.imgur.com/JQuSkKj.png)

權值總和為 12。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Kruskal's Algorithm的重點
{{< /ask_chatgpt >}}



1. 將所有的邊按照權值從小到大排序
2. 開始遍歷所有邊，如果遇到的邊不會形成環路，就將該邊加入最小生成樹中
3. 判斷環路的方法可以是利用Union-Find資料結構
4. 當所有邊都遍歷完畢或是最小生成樹中的邊數已達到其頂點數-1時結束
5. Kruskal's Algorithm具有貪心的思想，每次選擇權值最小的邊，是獲得最小生成樹的保證
6. 時間複雜度為O(ElogE)，其中E為邊的數量，因為排序需要O(ElogE)的時間，接著進行E次尋找，每次尋找需要O(1)的時間   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Kruskal's Algorithm的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



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

