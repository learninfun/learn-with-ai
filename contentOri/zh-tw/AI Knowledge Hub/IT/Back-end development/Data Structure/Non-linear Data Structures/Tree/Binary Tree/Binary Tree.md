+++
title = "二叉樹"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Binary Tree的中文問題
{{< /ask_chatgpt >}}



1. 如何計算一個Binary Tree的高度？
2. 如何求出一個Binary Tree的最大路徑總和？
3. 如何找到一個Binary Tree中的最小公共祖先？
4. 如何判斷一個Binary Tree是不是平衡二叉樹？
5. 如何將一個Binary Tree轉換為其鏈結表遍歷的序列？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Binary Tree並舉例
{{< /ask_chatgpt >}}



二叉樹(Binary Tree)是由節點(node)和邊(edge)所組成的樹狀結構，每個節點最多只有兩個子節點，左子節點與右子節點之分別代表了左右分支。其中，最頂端的節點稱為根節點(root)，沒有子節點的節點稱為葉節點(leaf)。

舉例來說，以下為一個Binary Tree:

```
             A
           /   \
          B     C
         / \   / \
        D   E F   G
```

其中，根節點為A，左子節點為B，右子節點為C，B的左子節點為D，右子節點為E，C的左子節點為F，右子節點為G。A、B、C、D、E、F、G都是節點，由線條所組成的部分則為邊。 

在此Binary Tree中，每個節點最多只有兩個子節點，因此符合Binary Tree的定義。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Binary Tree的重點
{{< /ask_chatgpt >}}



1. Binary Tree是一種基於節點的數據結構。
2. 每個節點最多只能有兩個子節點，這些節點被稱為左子節點和右子節點。
3. Binary Tree可以是空的，或者只有一個root node，或者包含多個節點組成的樹。
4. Binary Tree擁有三種遍歷方式：前序遍歷(pre-order traversal)、中序遍歷(in-order traversal)、後序遍歷(post-order traversal)。
5. Binary Tree可以用遞歸或迭代的方式進行遍歷，具體方法取決於節點和子節點之間的關係。
6. Binary Tree的高度(hight)等於節點最大深度(depth)+1。
7. 在Binary Tree中，從root node到任何一個節點的路徑都是唯一的。
8. Binary Tree可以通過不同的方法進行優化，例如平衡二叉樹(AVL Tree)、紅黑樹(Red-Black Tree)等。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Binary Tree的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 給定一個二叉搜索樹，找到其中第k小的元素。
答案：https://leetcode.com/problems/kth-smallest-element-in-a-bst/

2. 給定一個二叉樹，找到該樹中兩個指定節點的最近公共祖先。
答案：https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/

3. 給定一個二叉樹，找到從根節點到葉子節點的所有路徑，這些路徑需要滿足以下條件：路徑的和等於給定的 sum。
答案：https://leetcode.com/problems/path-sum-ii/

4. 給定一個二叉樹，找到該樹的最大深度。
答案：https://leetcode.com/problems/maximum-depth-of-binary-tree/

5. 給定一個二叉樹，找到該樹中每一棵子樹中的最大值與最小值之差的最大值。
答案：https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/   

