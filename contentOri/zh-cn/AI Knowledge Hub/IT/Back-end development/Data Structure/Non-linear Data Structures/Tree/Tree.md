+++
title = "树"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Tree的中文问题
{{< /ask_chatgpt >}}



1. 了解Tree资料结构的特点及应用场景有哪些？
2. 如何在Tree中搜索某一个元素，并且应用在什么场景中？
3. 如何从一个有序列表中建立一棵二叉搜索树？
4. 如何求解一个二叉树的深度？
5. 如何在一个二叉搜索树中插入一个新的节点，并且保证它仍是一个有效的树结构？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Tree并举例
{{< /ask_chatgpt >}}



Tree是一种常见的资料结构，它由许多组成称为节点，节点之间可以有连接称为边的二元关系，因此Tree也称为二元树。Tree以分层的方式呈现，每一层节点与它上层的节点之间，都有唯一的连接，而与同一层的其他节点之间则没有连接。树状结构的节点具有父节点和子节点之分，根节点位于顶层，没有父节点，而叶节点则没有子节点。

树的节点上可以附加任意的资料，一棵树可以用这些附加的资料来表达一个超级多的信息，例如家族树，以个人为节点，附加的资料可以是姓名，性别，年龄，电话等信息，通过树的分层结构，实现每个人之间的关系。

下面举一个二元树的例子：

     1
    / \
   2   3
  / \   \
 4   5   6

上面的示例是一个深度为3的树，含有6个节点和5个边，其中节点1是根节点，它有两个子节点2和3，节点2又有两个子节点4和5，节点3有一个子节点6。节点1到节点2之间有一条边，节点2到节点4之间有一条边，节点2到节点5之间有一条边，节点1到节点3之间有一条边，节点3到节点6之间有一条边。

透过这些节点之间的连接，可以实现很多不同的应用，例如档案系统、编译器、协议分析等。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Tree的重点
{{< /ask_chatgpt >}}



1. Tree是一种资料结构，以树状的方式组织资料。

2. 每个Tree都有一个根节点，并由节点和边构成。

3. 节点可分为内部节点与叶节点，内部节点为有子节点的节点，叶节点则为没有子节点的节点。

4. 每个节点可以有任意数量的子节点，而每个子节点的值都必须小于其父节点的值。

5. Tree有许多不同形式的变体，如Binary Tree、Binary Search Tree、AVL Trees和B-Trees等。

6. Tree可以用来储存大量的数据，以及在搜索和排序等问题上提供快速和高效的解决方案。

7. Tree的概念是许多高级演算法和数据结构的基础，包括优先伫列、哈希表和图形算法等。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Tree的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 路径最大和 

给定一个二叉树，找到所有路径中的最大路径和。

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

2. 二叉树的锯齿形层序遍历 

给定一棵二叉树，按照锯齿形的顺序返回其节点值。

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

3. 递增顺序查找树 

给定一棵二叉搜索树，将其转换为一棵只有右子树的递增顺序查找树。

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

4. 寻找树中第 k 小的元素 

给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法寻找其中第 k 个最小元素。

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

5. 二叉搜索树中的众数 

给定一个有相同值的二叉搜索树，找出 BST 中出现次数最多的值。

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

