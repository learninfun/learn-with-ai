+++
title = "t-SNE"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題t-SNE的中文問題
{{< /ask_chatgpt >}}

1. 什麼是t-SNE？它是如何運作的？
2. t-SNE常用於哪些領域？它有哪些應用？
3. 如何選擇t-SNE的參數？參數的設定對結果有哪些影響？
4. 什麼是“局部最優”問題？t-SNE是如何解決這個問題的？
5. t-SNE的優缺點是什麼？與PCA相比，它有哪些優勢和不足？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明t-SNE並舉例
{{< /ask_chatgpt >}}

t-SNE是一種降維技術，能夠將高維度的數據轉換為低維度的數據，同時保持原始數據之間的關係和相似度。它可以繪製出具有可視化效果的數據圖形，這些圖形可以幫助我們更好地理解數據之間的關係。

例如，假設有個數據集，其中包含了很多關於動物的信息，例如種類、體型、食性等等。透過t-SNE降維，可以將這些數據轉換為二維空間中的點，並將相似的數據點放在附近，不相似的數據點放在遠處，從而形成一個具有可視化效果的動物分類圖形。這樣我們就可以更清晰地看到各種不同動物之間的相似度和關係了。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列t-SNE的重點
{{< /ask_chatgpt >}}

1. t-SNE是一種非監督式機器學習算法，用於可視化高維數據。

2. t-SNE通過在高維空間和低維空間之間建立一個映射，將高維數據轉換為低維數據。通常，t-SNE將高維數據映射到2D或3D空間。

3. t-SNE將計算每對數據點之間的相似度，並輸出一個數據矩陣。

4. t-SNE利用機率分布來表示高維數據和低維數據之間的關係。

5. t-SNE主要關注如何在低維空間中保持數據點之間的局部和全局關係。即，緊密相鄰的數據點在低維空間中仍要靠近，但相互遠離的數據點在低維空間中應該更加分散。

6. t-SNE的效能受到多個超參數的影響，如學習率、正則化費用、嵌入維度等等。這些超參數的不同取值可能會導致不同的結果。

7. t-SNE的主要應用領域是在視覺化高維數據方面，特別是像圖像、音訊、自然語言處理等領域中。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題t-SNE的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}

1. 什麼是t-SNE？為什麼用它進行資料降維？

2. t-SNE的優點和缺點是什麼？

3. 如何調整t-SNE的參數來得到最佳的降維結果？

4. 在t-SNE中，如何處理缺失值和離群值？

5. 在t-SNE中，如何判斷降維後的資料是否仍保持原始資料的特徵？

答案：
1. t-SNE是一種基於機器學習的降維技術，它可以將高維度的資料降低到二維或三維的空間，使得人們可以更好地理解和可視化資料。透過嵌入學習和維度的轉換，t-SNE能夠在保留高維空間中的資料特徵的同時，讓其在低維空間中的分佈更加均勻、相似。

2. 優點：t-SNE在降維的同時，能更好地呈現資料之間的相似度；對於複雜的資料集有很好的效果；具有較強的可視化效果。缺點：計算時間較長，需要較大的計算資源；降維後的資料可能會失去原始資料的某些特徵。

3. 主要參數有兩個，分別是「perplexity」和「learning rate」。perplexity指出了附近的點應作為考慮相對於每個點的數量。learning rate則是直接關係到和當前梯度的距離。調整這些參數可以更好地控制t-SNE的效果。

4. 缺失值的處理：可以透過填補缺失值的方法，如平均值替換、插值法等。離群值的處理：可以使用一些現成的離群值檢測方法，如局部敏感哈希、基於密度的局部離群值等。或者，可以使用更高的perplexity來減輕離群點的影響。

5. 可以透過定性和定量的方​​式來幫助評估t-SNE的降維效果。定量的方法包括單調性和穩定性的分析；定性的方法包括視覺比較和對比原始和降維資料的分佈等。   
