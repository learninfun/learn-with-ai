## 習題預習
{{< ask_chatgpt >}}
給我5題AVL Tree的問題
{{< /ask_chatgpt >}}



1. AVL Tree是什麼？ 
2. AVL Tree與Binary Tree之間的區別是什麼？ 
3. AVL Tree如何保持平衡？ 
4. AVL Tree是如何進行插入和刪除操作？ 
5. AVL Tree的時間複雜度是多少，並舉例說明？   

## 說明知識
{{< ask_chatgpt >}}
說明AVL Tree並舉例
{{< /ask_chatgpt >}}



AVL Tree是一種自平衡二元搜尋樹，也就是說在插入或刪除節點時，會自動調整其結構，使樹保持平衡狀態，使搜索操作的時間複雜度保持在O(log n)級別。

AVL Tree的平衡是在節點的左右子樹高度之差不超過1的情況下進行的。當發現某一個節點的左右子樹高度差超過1時，就需要進行平衡操作。

常見的平衡操作有四種：左旋、右旋、先左旋再右旋和先右旋再左旋。旋轉操作會改變節點的位置，但不會改變節點的子節點，因此旋轉操作不會導致整棵樹的搜索順序改變。

以下是一個AVL Tree的例子：

![AVL Tree Example](https://i.imgur.com/ndKjJuX.png)

在這個例子中，插入節點3時，發現節點7和節點9的高度差為2，需要進行平衡操作，先進行右旋操作，得到以下樹：

![AVL Tree Rotated](https://i.imgur.com/ZJvGU23.png)

接著再插入節點5，這次發現節點6和節點9的高度差為2，需要進行平衡操作，先進行左旋操作，得到以下樹：

![AVL Tree Rotated 2](https://i.imgur.com/eMu6M8f.png)

這樣的AVL Tree就是一個平衡樹，可以保證搜索操作的時間複雜度為O(log n)級別。   

## 彙總重點
{{< ask_chatgpt >}}
條列AVL Tree的重點
{{< /ask_chatgpt >}}



AVL Tree 的重點如下：

1. 平衡因子：每個節點的左子樹和右子樹的高度差稱為平衡因子，AVL Tree中平衡因子的絕對值不超過1。

2. 左旋和右旋：為了維護AVL Tree的平衡性，需要通過左旋或右旋操作對樹進行調整。

3. 插入和刪除：插入和刪除節點時，要注意平衡因子的變化，需要將不平衡的節點通過旋轉調整為平衡狀態。

4. 查找和遍歷：AVL Tree 支持快速查找和各種遍歷操作，包括前序遍歷、中序遍歷和後序遍歷。

5. 時間複雜度：AVL Tree 的操作時間複雜度為O(log n)，其中n為樹中節點的數量。因此，AVL Tree在實際應用中具有很高的效率和可靠性。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題AVL Tree的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 在AVL Tree中插入元素10、15、12、20、25，請列出插入後AVL Tree的前序遍歷結果。
答案：15 12 10 20 25

2. 在一個AVL Tree中，刪除元素3，請列出刪除後AVL Tree的後序遍歷結果。
答案：4 2 5 1

3. 在AVL Tree中查找元素8的祖先，請列出查找結果。
答案：10 9

4. 在AVL Tree中查找元素11的後繼，請列出查找結果。
答案：12

5. 給定一組數列{8, 12, 10, 6, 11, 9, 14, 17, 20}，請構建出一棵AVL Tree。
答案：12 8 9 6 10 11 17 14 20   

## 網路資料
{{< ask_chatgpt >}}
給我5篇AVL Tree的網路資料
{{< /ask_chatgpt >}}



1. 資料結構：平衡二元樹（AVL Tree）
https://www.itread01.com/content/1545181177.html
本篇介紹 AVL Tree 平衡二元樹，解決二元樹退化成鏈表的問題。

2. AVL Tree：實現及其性質解釋
https://www.jianshu.com/p/e136ec79235c
本篇介紹了 AVL Tree 的實現及其性質解釋。

3. AVL Tree - 魯迅有句話
https://www.youtube.com/watch?v=1TmVnUy8n88
本篇是一個 AVl Tree 的講解影片，輕鬆有趣易懂。

4. AVL Tree (平衡二叉樹)
https://zhuanlan.zhihu.com/p/34858054
本篇介紹AVL Tree的三步驟、旋轉、非遞迴、以及優缺點。

5. 數據結構——平衡二叉樹（AVL Tree）詳解
https://www.cnblogs.com/ider/p/avlTree.html
本篇介紹 AVL Tree 的基本定義、被動更新、主動更新、查找、添加、刪除等詳細內容。   

