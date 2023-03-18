+++
title = "樹"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Tree的問題
{{< /ask_chatgpt >}}



1. 瞭解Tree資料結構的特點及應用場景有哪些？
2. 如何在Tree中搜索某一個元素，並且應用在什麼場景中？
3. 如何從一個有序列表中建立一棵二叉搜索樹？
4. 如何求解一個二叉樹的深度？
5. 如何在一個二叉搜索樹中插入一個新的節點，並且保證它仍是一個有效的樹結構？   

## 說明知識
{{< ask_chatgpt >}}
說明Tree並舉例
{{< /ask_chatgpt >}}



Tree是一種常見的資料結構，它由許多組成稱為節點，節點之間可以有連接稱為邊的二元關係，因此Tree也稱為二元樹。Tree以分層的方式呈現，每一層節點與它上層的節點之間，都有唯一的連接，而與同一層的其他節點之間則沒有連接。樹狀結構的節點具有父節點和子節點之分，根節點位於頂層，沒有父節點，而葉節點則沒有子節點。

樹的節點上可以附加任意的資料，一棵樹可以用這些附加的資料來表達一個超級多的信息，例如家族樹，以個人為節點，附加的資料可以是姓名，性別，年齡，電話等信息，通過樹的分層結構，實現每個人之間的關係。

下面舉一個二元樹的例子：

     1
    / \
   2   3
  / \   \
 4   5   6

上面的示例是一個深度為3的樹，含有6個節點和5個邊，其中節點1是根節點，它有兩個子節點2和3，節點2又有兩個子節點4和5，節點3有一個子節點6。節點1到節點2之間有一條邊，節點2到節點4之間有一條邊，節點2到節點5之間有一條邊，節點1到節點3之間有一條邊，節點3到節點6之間有一條邊。

透過這些節點之間的連接，可以實現很多不同的應用，例如檔案系統、編譯器、協議分析等。   

## 彙總重點
{{< ask_chatgpt >}}
條列Tree的重點
{{< /ask_chatgpt >}}



1. Tree是一種資料結構，以樹狀的方式組織資料。

2. 每個Tree都有一個根節點，並由節點和邊構成。

3. 節點可分為內部節點與葉節點，內部節點為有子節點的節點，葉節點則為沒有子節點的節點。

4. 每個節點可以有任意數量的子節點，而每個子節點的值都必須小於其父節點的值。

5. Tree有許多不同形式的變體，如Binary Tree、Binary Search Tree、AVL Trees和B-Trees等。

6. Tree可以用來儲存大量的數據，以及在搜索和排序等問題上提供快速和高效的解決方案。

7. Tree的概念是許多高級演算法和數據結構的基礎，包括優先佇列、哈希表和圖形算法等。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Tree的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 路徑最大和 

給定一個二叉樹，找到所有路徑中的最大路徑和。

答案：

```
class Solution {
public:
    int maxPathSum(TreeNode* root) {
        int res = INT_MIN;
        dfs(root, res);
        return res;
    }
    
    int dfs(TreeNode* root, int& res) {
        if (!root) return 0;
        int left = max(0, dfs(root->left, res));
        int right = max(0, dfs(root->right, res));
        res = max(res, left + right + root->val);
        return max(left, right) + root->val;
    }
};
```

2. 二叉樹的鋸齒形層序遍歷 

給定一棵二叉樹，按照鋸齒形的順序返回其節點值。

答案：

```
class Solution {
public:
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>> res;
        if (!root) return res;
        queue<TreeNode*> q{{root}};
        bool zigzag = false;
        while (!q.empty()) {
            int size = q.size();
            vector<int> level(size);
            for (int i = 0; i < size; ++i) {
                auto node = q.front(); q.pop();
                int idx = zigzag ? size - i - 1 : i;
                level[idx] = node->val;
                if (node->left) q.push(node->left);
                if (node->right) q.push(node->right);
            }
            zigzag = !zigzag;
            res.push_back(level);
        }
        return res;
    }
};
```

3. 遞增順序查找樹 

給定一棵二叉搜索樹，將其轉換為一棵只有右子樹的遞增順序查找樹。

答案：

```
class Solution {
public:
    TreeNode* increasingBST(TreeNode* root) {
        return dfs(root, nullptr);
    }
    
    TreeNode* dfs(TreeNode* root, TreeNode* tail) {
        if (!root) return tail;
        auto res = dfs(root->left, root);
        root->left = nullptr;
        root->right = dfs(root->right, tail);
        return res;
    }
};
```

4. 尋找樹中第 k 小的元素 

給定一個二叉搜索樹的根節點 root ，和一個整數 k ，請你設計一個算法尋找其中第 k 個最小元素。

答案：

```
class Solution {
public:
    int kthSmallest(TreeNode* root, int k) {
        vector<int> inorder;
        inorderTraversal(root, inorder);
        return inorder[k - 1];
    }
    
    void inorderTraversal(TreeNode* root, vector<int>& inorder) {
        if (!root) return;
        inorderTraversal(root->left, inorder);
        inorder.push_back(root->val);
        inorderTraversal(root->right, inorder);
    }
};
```

5. 二叉搜索樹中的眾數 

給定一個有相同值的二叉搜索樹，找出 BST 中出現次數最多的值。

答案：

```
class Solution {
public:
    vector<int> findMode(TreeNode* root) {
        vector<int> res;
        int cnt = 1, max_cnt = 0;
        TreeNode* pre = nullptr;
        inorderTraversal(root, pre, cnt, max_cnt, res);
        return res;
    }
    
    void inorderTraversal(TreeNode* root, TreeNode*& pre, int& cnt, int& max_cnt, vector<int>& res) {
        if (!root) return;
        inorderTraversal(root->left, pre, cnt, max_cnt, res);
        if (pre) {
            cnt = root->val == pre->val ? cnt + 1 : 1;
        }
        if (cnt == max_cnt) {
            res.push_back(root->val);
        } else if (cnt > max_cnt) {
            max_cnt = cnt;
            res = {root->val};
        }
        pre = root;
        inorderTraversal(root->right, pre, cnt, max_cnt, res);
    }
};
```   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Tree的網路資料
{{< /ask_chatgpt >}}



1. "The Importance of Trees in our Environment" by The Environmental Literacy Council
Link: https://enviroliteracy.org/environment-society/trees-forests/importance-of-trees-in-our-environment/

This article discusses the crucial role that trees play in our environment, including their impact on air quality, water conservation, and climate change.

2. "Types of Trees: Common Tree Types in the United States" by Arbor Day Foundation
Link: https://www.arborday.org/trees/treeguide/

This resource provides information about the different types of trees commonly found in the United States, including descriptions of their physical characteristics and growth habits.

3. "10 Surprising Benefits of Trees for Health and Wellness" by Clean Air Gardening
Link: https://www.cleanairgardening.com/benefits-of-trees/

This article highlights some of the unexpected health and wellness benefits that trees can provide, such as reducing stress levels and improving mental clarity.

4. "The Incredible Ways Trees Help Humans and the Planet" by Forbes
Link: https://www.forbes.com/sites/lauratenenbaum/2018/08/08/the-incredible-ways-trees-help-humans-and-the-planet/?sh=602a797d175a

This Forbes article explores some of the diverse ways that trees contribute to human and environmental health, from reducing air pollution to supporting local economies.

5. "Why Planting and Growing Trees is Key to Saving Our Planet" by World Wildlife Fund
Link: https://www.worldwildlife.org/stories/why-planting-and-growing-trees-is-key-to-saving-our-planet

The World Wildlife Fund provides a comprehensive overview of the importance of trees in combatting climate change and sustaining biodiversity, and outlines ways that individuals and communities can get involved in tree planting and preservation efforts.   

