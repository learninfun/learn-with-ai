## 習題預習
{{< ask_chatgpt >}}
給我5題Online Algorithms的問題
{{< /ask_chatgpt >}}



1. 請你設計一個Online Algorithm，針對一個長度為n的數列，每次新增一個元素，需要在O(1)的時間內找出當前數列的中位數。

2. 請你設計一個Online Algorithm，針對一個動態圖，每次新增一條邊，需要在O(logn)的時間內找出當前圖的最小生成樹。

3. 請你設計一個Online Algorithm，針對一個長度為n的數列，每次新增一個元素，需要在O(1)的時間內找出當前數列的前k大元素。

4. 請你設計一個Online Algorithm，針對一個長度為n的數列，每次新增一個元素，需要在O(1)的時間內找出當前數列的最大值、最小值以及平均值。

5. 請你設計一個Online Algorithm，針對一個動態圖，每次新增一個邊或刪除一個邊，需要在O(nlogn)的時間內計算當前圖的直徑。   

## 說明知識
{{< ask_chatgpt >}}
說明Online Algorithms並舉例
{{< /ask_chatgpt >}}



Online algorithms是一种算法，其可以在数据流到达的同时进行处理，而不需要事先获取所有的数据。它们可以用于许多实时应用程序，例如网络流量和在线广告。

例如，一个经典的在线算法是通过不断取中位数来维护数据流中的中位数。在数据流中，如果我们每次使用一个新的数据来更新当前的中位数，那么我们可以使用一个小根堆和一个大根堆来维护数据流的中位数，其中小根堆存放比当前中位数大的值，大根堆存放比当前中位数小的值。当新的数据流入时，我们选择将其放入合适的堆中，并确保两边的堆数量相同或只差一个。然后我们可以得到中位数，它将是两个堆的顶部元素之一，或者它们的平均数。

另一个例子是在线聚类算法。在处理数据流时，我们可能需要将它们分成一些组或簇。在线聚类算法可以根据数据流中的传入数据自动更新和调整簇，并获得尽可能准确的分类结果。

总的来说，Online algorithms已成为现代计算机科学中的重要研究领域，并可以广泛应用于许多实时应用程序。   

## 彙總重點
{{< ask_chatgpt >}}
條列Online Algorithms的重點
{{< /ask_chatgpt >}}



1. Online Algorithms是一種設計算法的方法，旨在處理動態的資料流或序列，並在每次資料到來時，即時更新最優解。

2. Online Algorithms具有高效執行時間的特性，通常可以在常數時間內完成操作。

3. 為了實現高效率的Online Algorithms，必須將計算功耗減至最低，因此它通常與雲計算等技術相關。

4. Online Algorithms通常使用各種技術來解決各種不同的問題，例如貪婪算法、隨機算法、動態規劃和概率算法等。

5. 在Online Algorithms中，關鍵問題是如何在時間和空間限制下維護良好的解決方案，以及如何處理資料的變化。

6. 通常，Online Algorithms需要在每次資料更新時立即調整運算，以確保有最佳的結果，這需要對資料進行即時分析和調整。

7. Online Algorithms常常用於路徑規劃，最短路徑、搜索、排序、過濾、儲存和匹配等領域，也用於處理在實時數據中隨時間變化的分布律、動態規劃等問題。

8. Online Algorithms是一個非常重要的研究領域，由於它可以讓處理複雜和動態數據變得更加高效。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Online Algorithms的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 在一個長度為n的陣列中，找出一個子陣列的最大和。 (答案: Kadane's Algorithm)
2. 給你一個圖，請找出最小生成樹。 (答案: Prim's Algorithm or Kruskal's Algorithm)
3. 在一個有向無環圖(DAG)中，找出最長路徑。 (答案: Topological Sort and Dynamic Programming)
4. 有 k 個有序陣列，請合併這 k 個陣列成一個單一的有序陣列。 (答案: Merge K Sorted Arrays)
5. 給你一個字符陣列和一個字典，請找出字符陣列中最長的單詞，且該單詞存在字典中。 (答案: Trie Tree and Dynamic Programming)   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Online Algorithms的網路資料
{{< /ask_chatgpt >}}



1. "Introduction to Online Algorithms" - A tutorial from Stanford University that covers the basics of online algorithms, including competitive analysis and decision making with limited information.

2. "Online Algorithms: A Survey" - A comprehensive review of online algorithms that covers a wide range of topics, including optimization, data streaming, and scheduling.

3. "Online Algorithms: The State of the Art" - A research paper that provides an overview of recent developments in online algorithms, highlighting their potential applications and limitations.

4. "Randomized Online Algorithms" - A lecture series from MIT that explores the use of randomized algorithms in online settings, including randomized competitive analysis and the use of randomization in scheduling and resource allocation.

5. "Online Algorithms for Optimization" - A course from the University of California, Berkeley that focuses on online algorithms for optimization problems, such as network flow and primal-dual algorithms.   

