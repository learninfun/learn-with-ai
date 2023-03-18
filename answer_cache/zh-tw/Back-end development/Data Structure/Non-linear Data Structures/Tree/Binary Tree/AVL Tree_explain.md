

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