+++
title = "分支限界法"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Branch and Bound的問題
{{< /ask_chatgpt >}}



以下是5個Branch and Bound的問題:

1. 在路徑規劃中，你需要找到通過一系列城市的最短路徑。使用Branch and Bound算法可以幫助你找到最短的路徑。

2. 在資源分配問題中，你需要將可用資源分配給多個項目，以最大化全局效益。Branch and Bound算法可以幫助你找到最優解。

3. 在排程問題中，你需要在多個工作之間分配時間，以最小化完成時間。Branch and Bound算法可以幫助你減少總體完成時間。

4. 在0-1背包問題中，你需要填充一個有限大小的背包，以最大化價值。Branch and Bound算法可以幫助你找到最理想的裝載方案。

5. 在數據壓縮中，你需要找到最小的壓縮方式，以減少數據量。Branch and Bound算法可以幫助你找到最佳的壓縮方法。   

## 說明知識
{{< ask_chatgpt >}}
說明Branch and Bound並舉例
{{< /ask_chatgpt >}}



Branch and Bound（分支定界法）是一種求解最優解問題的算法，通常用於組合優化，如旅行商問題、背包問題等。該算法利用了一個重要原則：通過一系列分支和限制條件的不斷縮小搜索空間，同時利用上界和下界對每個搜索分支進行評估和剪枝，以找到最優解。 

具體來說，Branch和Bound的搜索過程分為以下幾步：

1. 分支：從初始狀態出發，不斷將問題拆分為若干個子問題，即各種可能的解決途徑。對於背包問題，一個可能的拆分方法是選或不選某個物品。

2. 限制條件：給每個子問題限制條件，從而使得搜索空間逐漸地收斂。對於背包問題，限制條件是背包容量必須小於等於總容量。

3. 上界估計：對每個子問題計算一個最大可行解的上界，用於和問題的實際解進行比較，確定是否需要進一步搜索。對於背包問題，上界可能是已經被選中的物品總價值加上剩餘物品中價值最大的物品的價值。

4. 下界估計：對每個子問題計算一個最小可行解的下界，用於對搜索樹進行剪枝，即丟棄那些估值低於當前已知最優解的搜索分支。對於背包問題，下界可能是已選擇物品的總價值加上還能放入背包的最大價值。

5. 搜索回溯：在搜索分支全部完成後，輸出最優解。如果搜索到某個分支的上界小於已知最優值，則剪枝退出。 

舉例來說，假設我們要用Branch and Bound算法，解決一個旅行商問題，即從一個城市出發，依次經過其他城市回到原城市。已知每個城市之間的距離，求最短路程。

第一步是分支，我們從任意一個城市出發，可以分別選擇下一個要訪問的城市。例如，從城市A出發，我們可以選擇去B、C、D等城市。

第二步是限制條件，我們需要給每個分支設置約束條件，例如當前已經經過的城市不能再次經過。

第三步是上界估計，對於每個分支，我們需要計算已經選擇的路徑長度與未選擇的城市到起點的距離之和，同時加上一個下界估計。

第四步是下界估計，我們需要計算剩餘城市之間的最小距離，假設路徑已訪問過的城市均為i1，i2， ...，ik，剩下未訪問的城市為j1， j2， ...，jm，則下界估計為：

LB = Sum(D i j ) + min(D i1 k + D k j1 , D i2 k + D k j2 , ... , D ik k + D k jm )

其中，D i j 表示城市i到城市j的距離。我們選擇未訪問城市中到目前為止距離最近的一個城市，加入已知路徑。以這種方式對每個分支進行下界估計。

第五步是回溯搜索，在已經遍歷整個搜索樹後，輸出最優的路徑及其長度。

總之，Branch and Bound是一種高效的求解組合優化問題的算法。通過不斷縮小搜索空間並且對每個搜索分支進行剪枝，算法可以在較短的時間內找到最優解。   

## 彙總重點
{{< ask_chatgpt >}}
條列Branch and Bound的重點
{{< /ask_chatgpt >}}



1. 分支限界法是一種優化問題的求解方法，可以用於求解許多困難的問題，如最小化或最大化目標函數的線性規劃問題、TSP問題等。
2. 分支限界法的基本思想是逐步構造一個搜索樹，並使用啟發函數將可行域切割成多個子域，以加快搜索速度。
3. 在搜索過程中，需要保存每個還未被擴展的節點的狀態以及當前的最優解，並根據啟發函數選擇合適的節點進行擴展。
4. 分支限界法的剪枝策略是在搜索樹的每個節點上進行，通常包括擇優策略、可行性剪枝和界限剪枝等。
5. 分支限界法的時間複雜度較高，但通常可以得到最優解或近似最優解。為了減少時間消耗，可以使用一些技巧，如啟髮式搜索、並行搜索等。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Branch and Bound的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 在一個完全連通的無向圖中，找到最小權重生成樹。 

答案: Kruskal algorithm 和 Prim's algorithm 

2. 將$n$個項目分成$m$個集合，使得每個集合的總和盡可能接近一個給定的值$K$。 

答案: Knapsack問題的分支限界解法 

3. 在二元搜索樹中查找第$k$小的元素。 

答案: 計算每個節點的大小並使用分支限制法 

4. 旅行商問題：一個商業旅行者必須訪問$n$個城市，每個城市只能訪問一次，並在所有訪問城市的路徑中找到最短路徑。 

答案: 已知經過點集將事件分解為小問題的方法 

5. 在$0-1$背包中，找出總重量不超過$W$，但總價值最大的物品子集。 

答案: Branch and Bound解法   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Branch and Bound的網路資料
{{< /ask_chatgpt >}}



1. Branch and bound algorithm - GeeksforGeeks 
https://www.geeksforgeeks.org/branch-and-bound-algorithm/

這篇文章介紹了branch and bound algorithm (B&B)的概要和步驟，同時附有python程式碼示範。文章中也提到一些B&B解決問題的例子，例如Traveling Salesman Problem (TSP)和Knapsack Problem等等。

2. Branch and Bound Algorithm - Brilliant.org 
https://brilliant.org/wiki/branch-and-bound/

這篇文章介紹了B&B的原理和實作方式。這篇的重點在於說明B&B算法如何運作以及如何應用它來求解最優解。文章中列舉了一些範例，包括TSP和Maximum Clique Problem等等。

3. The Branch and Bound Algorithm - Princeton University Press 
https://press.princeton.edu/books/hardcover/9780691122218/the-branch-and-bound-algorithm

這個網站介紹了一本書，書名為The Branch and Bound Algorithm。該書是由Prof. Barnard Korte和Prof. Jens Vygen撰寫，並且主要講授了B&B算法的原理和在解決各種組合優化問題上的應用。

4. Branch and Bound - Tutorialspoint 
https://www.tutorialspoint.com/branch_and_bound_algorithm/index.htm

這篇文章介紹了B&B算法的結構和基本概念，以及如何透過這種算法求解最優解。文章中還提到如何使用B&B算法解決一些組合優化問題，例如Traveling Salesman Problem和Knapsack Problem等問題。

5. Branch and Bound Algorithm - Georgia Tech 
https://www.isye.gatech.edu/sites/default/files/documents/ipy2/Lecture6.pdf

這個網站是由Georgia Tech的師生撰寫，介紹了B&B的運作方式，並且同時列舉了一些關於BP問題的範例，以及如何使用B&B算法求解。簡而言之，這個資源將B&B算法的重點視為如何在數學上解決組合優化問題。   

