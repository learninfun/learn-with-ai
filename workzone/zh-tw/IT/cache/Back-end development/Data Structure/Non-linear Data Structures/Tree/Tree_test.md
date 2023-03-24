

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