+++
title = "紅黑樹"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Red-Black Tree的問題
{{< /ask_chatgpt >}}



1. Red-Black Tree是什麼?
2. Red-Black Tree的優點是什麼?
3. 為什麼Red-Black Tree可以保證平衡?
4. 如何實現紅黑樹的插入操作，保持平衡?
5. 如何實現紅黑樹的查找操作，保持平衡?   

## 說明知識
{{< ask_chatgpt >}}
說明Red-Black Tree並舉例
{{< /ask_chatgpt >}}



Red-Black Tree是一種平衡二叉樹，它在BST（二叉搜索樹）的基礎上，增加了顏色的概念，使得在插入、刪除操作時，能夠保持樹的平衡性，同時也提高了查詢操作的效率。

在Red-Black Tree中，每個節點都有顏色屬性，紅色或黑色，並且樹中的所有空白節點都視為黑色節點。樹中有以下幾個重要特性：

1. 根節點與葉節點都視為黑節點
2. 紅色節點的子節點都是黑色節點
3. 任何一個節點到它的葉子節點（空節點）的所有路徑上，經過的黑節點數量是相同的

插入操作會對紅黑樹造成變化，為了維持平衡性，插入後需要將新增節點標記為紅色。如果插入的位置破壞了紅黑樹的特性，例如紅節點有紅色子節點，這時需要進行調整，使得紅黑樹重新滿足所有特性。

舉例：

以下是一棵紅黑樹，其中紅色表示節點為紅色，黑色表示節點為黑色。

![red-black tree example](https://i.imgur.com/ZKPoJPy.png)

將13插入上圖紅黑樹：

1. 顏色為紅色，插入到節點6的右邊，得到下面的樹：

![red-black tree example 2](https://i.imgur.com/D8VwWvc.png)

此時插入的節點13沒有破壞紅黑樹的特性，因此可以保持不變。

2. 顏色為紅色，插入到節點1的右邊，得到下面的樹：

![red-black tree example 3](https://i.imgur.com/TQdRl6U.png)

此時13的父節點1為黑色節點，但是它的兄弟節點2也為黑色節點，這破壞了特性2。由於13、1、2三個節點都在同一條路徑上，因此需要進行旋轉操作。

3. 將1作為旋轉軸向左旋轉，得到樹：

![red-black tree example 4](https://i.imgur.com/Di9Sohu.png)

4. 由於旋轉之後，13的祖父節點6變成了原本的13的父節點，因此需要重新檢查樹的平衡性，經過修正之後，得到最終的紅黑樹：

![red-black tree example 5](https://i.imgur.com/VcmDZyl.png)   

## 彙總重點
{{< ask_chatgpt >}}
條列Red-Black Tree的重點
{{< /ask_chatgpt >}}



1. 每個節點都是紅色或黑色的。
2. 根節點是黑色的。
3. 每個葉節點（NIL節點）都是黑色的。
4. 如果一個節點是紅色的，那麼它的兩個子節點都是黑色的（不能有兩個相鄰的紅色節點）。
5. 從任一節點到其每個子節點的所有路徑都包含相同數量的黑色節點。
6. 插入和刪除節點時，要保持樹的上述性質。
7. 紅黑樹是一種自平衡二叉搜索樹，它保證了最壞情況下基本操作的時間複雜度為O(log N)。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Red-Black Tree的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 設計一個紅黑樹類，並實現以下方法：
   - insert(value)：插入一個值為value的節點，要求插入後仍符合紅黑樹的性質。
   - delete(value)：刪除一個值為value的節點，要求刪除後仍符合紅黑樹的性質。
   - search(value)：查找一個值為value的節點。

2. 實現一個紅黑樹的遍歷方法，輸出節點的值集合。要求按照先序、中序以及後序遍歷方式分別實現。

3. 在紅黑樹中查找一個節點的祖先節點中與它最接近的那個紅色節點。

4. 給定一個紅黑樹節點，求出這個節點的排名（即排序後它在樹中所有節點中的位置編號）。

5. 設計一個算法，判斷兩個紅黑樹是否完全相同（即樹的形狀和節點值完全相同）。


答案：

1. 略。

2. 先序遍歷：
   1. 訪問根節點。
   2. 遞歸地先序遍歷左子樹。
   3. 遞歸地先序遍歷右子樹。
   
   中序遍歷：
   1. 遞歸地中序遍歷左子樹。
   2. 訪問根節點。
   3. 遞歸地中序遍歷右子樹。
   
   後序遍歷：
   1. 遞歸地後序遍歷左子樹。
   2. 遞歸地後序遍歷右子樹。
   3. 訪問根節點。

3. 可以按照以下步驟進行：
   1. 從給定的節點開始向上遍歷，直到達到根節點或者找到一個紅色節點。
   2. 如果找到的是一個紅色節點，則結果就是這個節點；否則，結果就是根節點。

4. 可以按照以下步驟進行：
   1. 如果要查詢排名的節點有左子樹，則排名等於左子樹大小加1。
   2. 如果要查詢排名的節點沒有左子樹，則排名等於祖先節點中第一個比它右子樹節點數目多的節點的排名加上右子樹大小加1。

5. 可以按照以下步驟進行：
   1. 如果兩個紅黑樹的根節點都不存在，則它們相同。
   2. 如果兩個紅黑樹的根節點的值不相等，則它們不相同。
   3. 遞歸地判斷兩個紅黑樹的左、右子樹是否相同，只有當左右子樹都相同時，兩個紅黑樹才相同。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Red-Black Tree的網路資料
{{< /ask_chatgpt >}}



1. Red-Black Tree - GeeksforGeeks
https://www.geeksforgeeks.org/red-black-tree-set-1-introduction-2/

這是一篇介紹 Red-Black Tree 的基本概念和原理的文章，也提供了紅黑樹的節點結構、插入和刪除操作的實現等，是一個相當詳細的資料。

2. Red-Black Tree - Wikipedia
https://en.wikipedia.org/wiki/Red%E2%80%93black_tree

這是一篇從理論和歷史的角度介紹紅黑樹的文章。文章中說明了紅黑樹的性質和特點，比較了紅黑樹相對於其他平衡樹的優缺點，以及歷史上紅黑樹的發展和演變。

3. Algorithms and Data Structures - Red-Black Tree - Tutorialspoint
https://www.tutorialspoint.com/data_structures_algorithms/red_black_tree_algorithm.htm

這是一篇在紅黑樹的基礎上深入介紹了 AVL 樹的文章。原文中講解了紅黑樹的插入和刪除操作的實現以及紅黑樹中節點旋轉的基本概念，同時也介紹了 AVL 樹和紅黑樹的比較和區別。

4. Red-Black Tree - Visualgo.net
https://visualgo.net/en/bst

這是一個用動畫的形式展示紅黑樹的特點和操作過程的網站。它展示了紅黑樹的遞迴算法實現過程，可以讓讀者更好地理解紅黑樹的運作原理。

5. Binary Trees: Red-Black Trees - Princeton University
https://www.cs.princeton.edu/~rs/talks/LLRB/RedBlack.pdf

這是一篇公開課程的講義，其中介紹了紅黑樹和 LLRB（Left-Leaning Red Black Tree）的介紹和實現。講義中介紹了紅黑樹的性質和特點，以及用雜技演示了紅黑樹的節點旋轉操作，是一份相當好的資料。   
