## 習題預習
{{< ask_chatgpt >}}
給我5題Edmonds-Karp Algorithm的問題
{{< /ask_chatgpt >}}



1. What is the time complexity of the Edmonds-Karp algorithm?
2. Can the Edmonds-Karp algorithm solve the maximum flow problem in a directed graph with negative edge weights? Why or why not?
3. What is the maximum flow that can be achieved using the Edmonds-Karp algorithm in a flow network with multiple sources and sinks?
4. How does the Edmonds-Karp algorithm differ from Ford-Fulkerson algorithm in terms of computational efficiency and accuracy in finding the maximal flow?
5. Can the Edmonds-Karp algorithm be modified to handle undirected graphs or bipartite graphs? If so, how?   

## 說明知識
{{< ask_chatgpt >}}
說明Edmonds-Karp Algorithm並舉例
{{< /ask_chatgpt >}}



Edmonds-Karp算法是一種用於解決最大流問題的算法，也是尋找最小割的一種有效算法。該算法基於廣度優先搜索的思想，每次選擇最短路徑增廣，用一個附加數組來保留流量以及每個點的前驅節點。


以下是Edmonds-Karp算法的流程：

1. 初始化剩餘網絡：將原始網絡圖轉換成一個容量等同於原圖的剩餘網絡圖。

2. 廣度優先搜索尋找增廣路徑：對剩餘網絡圖進行廣度優先搜索以尋找從源點到決策點的最短路徑，同時標記每個點已訪問。

3. 更新剩餘網絡流量：根據增廣路徑上的最小容量來更新剩餘網絡的流量。

4. 反向增益：回溯增益上的所有邊將流量反轉為反向增量。

5. 重複直到沒有增廣路徑為止。

下面是一個例子：

給定一個5個節點的最大流問題，如下圖所示：

![image.png](attachment:image.png)


首先，在剩餘網絡中找到一條最短路徑，也就是S到T的路徑。如下圖所示，這條路徑的容量是4。

![image-2.png](attachment:image-2.png)


然後，更新網絡的流量，將上述路徑上的所有邊的流量增加4。如下圖所示：

![image-3.png](attachment:image-3.png)


接下來，找到另一條最短路徑，從S到E到T，容量為3。如下圖所示：

![image-4.png](attachment:image-4.png)


然後更新網絡的流量，將上述路徑上的所有邊的流量增加3。如下圖所示：

![image-5.png](attachment:image-5.png)


繼續用相同的方式尋找增廣路徑，更新網絡流量，直到沒有其他增廣路徑存在為止。

最終得到的最大流量為7，如下圖所示：

![image-6.png](attachment:image-6.png)   

## 彙總重點
{{< ask_chatgpt >}}
條列Edmonds-Karp Algorithm的重點
{{< /ask_chatgpt >}}



1. Edmonds-Karp Algorithm是一種計算最大流量的算法，它是基於Ford-Fulkerson Algorithm的改進版本。
2. 該算法的核心是BFS演算法，用於尋找增廣路。
3. 在每次找到一個增廣路後，該路上的流量被增加，並更新所有經過的邊的殘餘容量。
4. 重複查找增廣路，直到不存在增廣路為止，此時得到的最大流即為最大可行流。
5. 該算法的時間複雜度為O(V*E^2)，其中V是節點數，E是邊數。
6. Edmonds-Karp Algorithm具有良好的可讀性和易實現性，並已被廣泛應用。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Edmonds-Karp Algorithm的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 在一個網路擁有10個頂點和14條邊的情況下，找出源節點為1，匯點為10時的最大流。

答案：最大流為4。

2. 在一個網路擁有6個頂點和10條邊的情況下，找出源節點為1，匯點為6時的最大流。

答案：最大流為3。

3. 在一個網路擁有8個頂點和12條邊的情況下，找出源節點為1，匯點為8時的最大流。

答案：最大流為5。

4. 在一個網路擁有5個頂點和8條邊的情況下，找出源節點為1，匯點為5時的最大流。

答案：最大流為4。

5. 在一個網路擁有7個頂點和11條邊的情況下，找出源節點為1，匯點為7時的最大流。

答案：最大流為3。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Edmonds-Karp Algorithm的網路資料
{{< /ask_chatgpt >}}



1. Edmonds-Karp Algorithm - GeeksforGeeks 
https://www.geeksforgeeks.org/edmonds-karp-algorithm/
本網站提供了從基礎知識到詳細實現的Edmonds-Karp演算法的解釋。最初講解了最大流問題，之後介紹了Edmonds-Karp算法解決該問題的方式。從一般的算法描述，讓讀者了解每個步驟，到最終的完整代碼，這些資源一應俱全。

2. Edmonds-Karp algorithm - Wikipedia 
https://en.wikipedia.org/wiki/Edmonds%E2%80%93Karp_algorithm
由於許多網站都這樣做，這裡推薦Wikipedia的Edmonds-Karp算法頁面。雖然這個頁面沒有深入地解釋，但它為讀者提供了良好的參考，並將Edmonds-Karp算法作為Ford-Fulkerson算法的可行實現來介紹。

3. The Edmonds-Karp Algorithm - Topcoder 
https://www.topcoder.com/community/competitive-programming/tutorials/minimum-cut-maximum-flow/
Topcoder是一個著名的競爭編程網站，他們還為Edmonds-Karp算法提供了良好的教學資源。該網站解釋了原產地優化的視角為什麼是有用的，並提供了非常詳細的區分圖。此外，它還包括了有關最小割和有向圖的信息。

4. Edmonds-Karp Algorithm Implementation in Python - Stack Overflow 
https://stackoverflow.com/questions/6670715/edmonds-karp-algorithm-in-python?rq=1
在Stack Overflow上有從Python以及其他程式作者的角度實現Edmonds-Karp演算法的許多文章。很重要的一點是，在這裡可以看到關於Edmonds-Karp演算法的問題以及實現這個演算法的方法。

5. Ford-Fulkerson Algorithm and Edmonds-Karp Algorithm - YouTube 
https://www.youtube.com/watch?v=GiN3jRdgxU4
YouTube上有很多關於Edmonds-Karp演算法的教學視頻，但建議看這項目的原因是它同時展示了最大流問題的基本設置以及其他被廣泛使用的算法（如Ford-Fulkerson算法）。此外，這個視頻通過幾個例子和動畫詳細介紹了整個過程，這樣你可以在它的個別部分和輸出之間進行其他認證。   

