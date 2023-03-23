+++
title = "最小生成樹演算法"
weight = "3"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Minimum Spanning Tree Algorithms的中文問題
{{< /ask_chatgpt >}}



1. 最小生成樹演算法是什麼？它的定義是什麼？

2. 什麼是Kruskal演算法？它如何找到一個圖的最小生成樹？

3. 什麼是Prim演算法？它如何找到一個圖的最小生成樹？

4. 什麼是Boruvka演算法？它如何找到一個圖的最小生成樹？

5. Dijkstra演算法可以用來找到最短路徑嗎？如果可以，它如何實現？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Minimum Spanning Tree Algorithms並舉例
{{< /ask_chatgpt >}}



Minimum Spanning Tree Algorithms（最小生成樹算法）是用來找尋一個連通無向圖中，權值最小的生成樹（即包含所有節點的樹），其中權值可以是邊的權重或是節點之間的距離。

以下是常被使用的Minimum Spanning Tree Algorithms：

1. Kruskal algorithm（克魯斯卡爾算法）：將所有邊依權值排序，從權值最小的邊開始，依序加入生成樹，直到所有節點都連通。

2. Prim's algorithm（普利姆算法）：以一個起始點開始，找到離該點最近的一個連接節點加入生成樹中，然後繼續找到與生成樹相連的最小的邊，加入生成樹中，直到所有節點都連通。

舉例來說，假設有以下的無向圖： 

若採用Kruskal算法，最小生成樹的邊為： 

若採用Prim算法，最小生成樹的邊為： 

兩種算法得到的最小生成樹都是一樣的，但是在不同的圖形、節點數量上，兩種算法的表現可能會有所不同。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Minimum Spanning Tree Algorithms的重點
{{< /ask_chatgpt >}}



1. 最小生成樹是一個無向圖的生成樹，它的邊權值之和最小。 
2. 最小生成樹算法通常包括普裡姆算法和克魯斯卡爾算法。 
3. 普裡姆算法是從一個起點開始，每次選擇與該點相鄰且權值最小的邊，將該點加入生成樹，並將與該點相鄰的邊加入候選邊集中，繼續選擇下一個節點直到所有點都在生成樹中。 
4. 克魯斯卡爾算法是將邊按權值從小到大排序，然後逐步加入生成樹中，如果連接該邊的兩個點已經在生成樹中，則忽略該邊，否則加入生成樹中。 
5. 最小生成樹算法的時間複雜度通常為O（ElogE），其中E為邊數。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Minimum Spanning Tree Algorithms的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



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

