## 習題預習
{{< ask_chatgpt >}}
給我5題Searching Algorithms的問題
{{< /ask_chatgpt >}}



1. 如何使用深度優先搜索找到一個圖結構中兩個節點之間的最短路徑？

2. 如何在一個已經排序的數組中使用二分查找來查找一個特定的元素？

3. 如何使用廣度優先搜索來找到一個圖結構中所有與給定節點相連的節點？

4. 如何使用雜湊表來實現快速搜索？能否舉例說明一個具體的應用場景？

5. 如何使用迭代加深搜索來找到一個棋盤遊戲的最佳解？   

## 說明知識
{{< ask_chatgpt >}}
說明Searching Algorithms並舉例
{{< /ask_chatgpt >}}



搜尋演算法是一種用來在資料結構（如陣列、列表、樹等）中查找特定值的演算法。這種演算法在許多應用領域中使用，包括數據庫查詢、瀏覽器搜尋、檔案系統搜尋等等。

以下是幾個常見的搜尋演算法：

1. 線性搜尋演算法（Linear Search Algorithm）：逐一搜尋目標值，直到找到為止。這種方法適用於小型數據集。

2. 二分搜尋演算法（Binary Search Algorithm）：這種方法適用於排序後的數據集（如有序陣列）。它通過將數據集逐漸減半進行搜尋，以找到目標值。

3. 哈希搜尋演算法（Hash Search Algorithm）：使用哈希函數將目標值映射到數據集中的特定位置。使用這種方法可以在常數時間內找到目標值。

舉例來說，如果要在一個含有10個元素的陣列中查找值為5的元素，可以使用線性搜尋演算法。開始的時候，將當前索引設為0，然後逐個檢查陣列中的元素，直到找到值為5的元素或遍歷完整個陣列。如果找到了值為5的元素，就返回元素的索引；如果沒有找到，就返回一個錯誤訊息。

若是在一個有序的10個元素的陣列中查找值為5的元素，可以使用二分搜尋演算法。將陣列從中間分為兩部分，如果目標值小於中間的值，就在左半部分繼續尋找，否則在右半部分繼續尋找。這樣每次可以將搜尋範圍減半，以提高效率。如果找到了值為5的元素，就返回元素的索引；如果沒有找到，就返回一個錯誤訊息。

哈希搜尋演算法通常用於數據庫查詢、瀏覽器搜尋等需要快速查找的應用領域。例如，當使用Google搜尋時，Google會在自己的數據庫中使用哈希搜尋演算法來查找相關結果。   

## 彙總重點
{{< ask_chatgpt >}}
條列Searching Algorithms的重點
{{< /ask_chatgpt >}}



1. Linear Search: 一個簡單的搜尋演算法，透過從頭開始，逐一檢查陣列中的元素是否符合搜尋條件；時間複雜度為O(n)。 

2. Binary Search: 將一個已排序的陣列分為兩部分，並透過比較搜尋值大小，判斷目標值可能在哪一個區域，最後從該區域繼續執行二分搜尋；時間複雜度為O(log n)。

3. Jump Search: 在已排序的陣列中，以跳躍的方式搜尋目標值，較為快速，時間複雜度為O(√n)。

4. Interpolation Search：在已排序的數列中，透過數值大小的估計，更快地找到目標值所在的位置；時間複雜度為O(log log n)。

5. Exponential Search: 先進行指數級的搜尋，當有足夠接近目標值的範圍時，再轉為二分搜尋，通常適用於整數搜尋；時間複雜度為O(log n)。

6. Hashing: 透過hash function將目標值轉換為一個索引值，以快速尋找目標值的位置；時間複雜度為O(1)。 

7. Depth-First Search(DFS): 借助stack實現，遍歷整個圖或樹的節點，並依此搜尋目標值所在的節點；時間複雜度為O(V+E)。

8. Breadth-First Search(BFS): 借助queue實現，從起點開始一層一層向外擴展，並依此搜尋目標值所在的節點；時間複雜度為O(V+E)。

9. A* Search: 求解最短路徑問題，以估價函數為基礎評估每個節點的價值，選擇最優的路徑；時間複雜度為O(b^d)。

10. Dijkstra's Algorithm: 求解最短路徑問題，以當前距離為基礎計算每個節點到起點的距離，選擇最短的路徑；時間複雜度為O(E log V)。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Searching Algorithms的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 在一個有序陣列中，查找一個數字的位置，如果不存在則返回-1

答案：二分查找算法（Binary Search）

2. 在一個二叉樹中搜索一個指定值的節點

答案：二叉樹搜索算法（Binary Tree Search）

3. 在一個有向圖中，查找一條從起點到終點的最短路徑

答案：Dijkstra算法

4. 在一個迷宮中查找通往終點的路徑（迷宮由二維陣列表示，0表示可走，1表示障礙物）

答案：深度優先搜索（DFS）或廣度優先搜索（BFS）

5. 在一個字符串中查找一個字串出現的位置

答案：KMP算法（Knuth-Morris-Pratt）   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Searching Algorithms的網路資料
{{< /ask_chatgpt >}}



1. "A Guide to Searching Algorithms" - https://www.geeksforgeeks.org/searching-algorithms/
This article from GeeksforGeeks provides a comprehensive guide to different types of searching algorithms, including linear search, binary search, jump search, interpolation search, exponential search, and more.

2. "Introduction to Searching Algorithms" - https://www.interviewcake.com/article/java/binary-search
Interview Cake offers a detailed overview of binary search, including its advantages and disadvantages, implementation in Java, and examples of its application in real-world scenarios.

3. "Searching Algorithms: A Comprehensive Guide" - https://www.hackerearth.com/practice/algorithms/searching/binary-search/tutorial/
HackerEarth offers a comprehensive tutorial on various searching algorithms, including binary search, ternary search, and linear search. The article also delves into the time complexity of each algorithm and offers examples for each.

4. "Search Algorithms in Artificial Intelligence" - https://www.techopedia.com/definition/16950/search-algorithms-in-artificial-intelligence-ai
Techopedia explores the role of search algorithms in artificial intelligence, examining different search types like depth-first search, breadth-first search and heuristic search. The article delves into the technology's practical application, specifically the use of search algorithms in game development and robotics.

5. "Algorithms for Searching and Sorting" - https://www.brilliant.org/wiki/algorithms-for-searching-and-sorting/
This article from the website Brilliant offers an introduction to basic search algorithms, including linear search, binary search, and interpolation search. The article also provides an introduction to sorting algorithms, such as bubble sort, quicksort, and merge sort, along with an overview of their pros and cons.   

