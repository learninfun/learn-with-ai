+++
title = "N叉树"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题N-ary Tree的问题
{{< /ask_chatgpt >}}



1. 如何計算 N-ary Tree 的深度？
2. 如何實現 N-ary Tree 的前序、中序和後序遍歷？
3. 如何找到 N-ary Tree 中的最大值、最小值以及平均值？
4. 如何找到 N-ary Tree 中某一節點的所有祖先節點？
5. 如何實現 N-ary Tree 的層序遍歷？   

## 说明知识
{{< ask_chatgpt >}}
说明N-ary Tree并举例
{{< /ask_chatgpt >}}



N-ary Tree是一種樹型結構，它由一個根節點和0個或多個子節點組成，但是每個節點可以有多於2個的後代節點。因此，它可以表示多層次的結構，例如組織機構、家族關係等。

舉例如下：

一家人的家譜可以用N-ary Tree表示，以祖先為根節點，每個節點代表一位成員，然後子節點可以是該成員的子女，也可以是孫輩、曾孫輩等後代。如下圖所示：

![N-ary Tree](https://cdn.luogu.com.cn/upload/image_hosting/kyg2o28s.png)

在這個家譜中，每個節點代表一個人，而每個人都有多位子女、孫輩、曾孫輩等後代，這個結構非常適合用N-ary Tree表示。

除家譜外，N-ary Tree還可以用於表示組織架構、網路拓撲等結構，具有很廣泛的應用場景。   

## 汇总重点
{{< ask_chatgpt >}}
条列N-ary Tree的重点
{{< /ask_chatgpt >}}



1. N-ary Tree是一種樹狀結構，每個節點可以有多個子節點。

2. N-ary Tree常常被用來表示有層次結構的數據，例如組織架構或標籤分類。

3. N-ary Tree中，每個節點可以有0到n個子節點，其中n稱為N值。

4. N-ary Tree的節點通常包含一個值，用來表示節點描述的內容。

5. 在N-ary Tree中，每個節點的子節點沒有順序關係，也沒有固定的子節點數目。

6. N-ary Tree的深度可以被定義為從根節點到最深節點的最長路徑長度。

7. N-ary Tree的樹高可以被定義為根節點到葉節點的最長路徑長度。

8. N-ary Tree常常會被用來實現搜尋和排序算法，例如Trie樹和堆積排序。

9. N-ary Tree的遍歷方式包括前序遍歷、中序遍歷、後序遍歷和層序遍歷。

10. 創建N-ary Tree的方法包括遞迴和迭代算法，具體方法取決於數據結構的特點和應用場景的需求。   

## 知识测验
{{< ask_chatgpt >}}
给我5题N-ary Tree的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. Maximum Depth of N-ary Tree
給定一個N-ary Tree，回傳其深度的最大值。

範例輸入:    
       1
     / | \
    3  2  4
   / \
  5   6
範例輸出: 3

答案:
class Solution {
public:
    int maxDepth(Node* root) {
        if(root == nullptr) return 0;
        int maxDepthValue = 0;
        for(auto child : root->children)
        {
            maxDepthValue = max(maxDepthValue, maxDepth(child));
        }
        return 1 + maxDepthValue;
    }
};


2. N-ary Tree Level Order Traversal
給定一個N-ary Tree，回傳其按層次輸出的節點序列。

範例輸入:    
       1
     / | \
    3  2  4
   / \
  5   6
範例輸出: [[1],[3,2,4],[5,6]]

答案:
class Solution {
public:
    vector<vector<int>>  ans;
    void bfs(Node* root){
        if(root==nullptr){
            return;
        }
        queue<Node*> q;
        q.push(root);
        while(!q.empty()){
            vector<int> level;
            int size=q.size();
            for(int i=0;i<size;i++){
                Node* current=q.front();
                q.pop();
                level.push_back(current->val);
                for(auto node:current->children){
                    q.push(node);
                }
            }
            ans.push_back(level);
        }
    }
    vector<vector<int>> levelOrder(Node* root) {
        bfs(root);
        return ans;
    }
};


3. Delete Nodes And Return Forest
給定一個N-ary Tree和一個集合to_delete，請刪除N-ary Tree中所有在集合to_delete中的節點並返回刪除節點後得到的森林(森林即多個樹的集合)。

範例輸入:    
       1
     / | \
    3  2  4
   / \
  5   6
刪除set=[3,5]

範例輸出: [[1,null,4],[6],[2]]

答案:
class Solution {
public:
//返回删除后新生成的树+独立节点
    unordered_set<int> to_delete_set;
    vector<Node*> res;
    vector<Node*> delNodes(Node* root, vector<int>& to_delete) {
        for(auto num:to_delete){
            to_delete_set.insert(num);
        }
        helper(root);
        if(root!=nullptr&&to_delete_set.find(root->val)==to_delete_set.end()){
            res.push_back(root);
        }
        return res;
    }
    Node* helper(Node* root){
        if(root==nullptr) return root;
        root->children[0]=helper(root->children[0]);
        for(int i=1;i<root->children.size();i++){
            root->children[i]=helper(root->children[i]);
            if(to_delete_set.find(root->children[i]->val)!=to_delete_set.end()){
                root->children.erase(root->children.begin()+i,root->children.begin()+i+1);
                i--;
            }
        }
        if(to_delete_set.find(root->val)!=to_delete_set.end()){
            for(int i=0;i<root->children.size();i++){
                if(root->children[i]!=nullptr){
                    res.push_back(root->children[i]);
                }
            }
            return nullptr;
        }
        return root;
    }
};


4. Find Root of N-Ary Tree
給定一個N-ary Tree，其中所有節點除根節點外都有一個父節點。找到並返回這個N-ary Tree的根節點。

範例輸入:
 [[1,null,3,2,4,null,5,6]]
範例輸出: 1

答案:
class Solution {
public:
    Node* findRoot(vector<Node*> tree) {
        unordered_map<int,int> cnt;
        for(auto node:tree){
            if(node->children.size()==0){
                cnt[node->val]++;
            }
            else{
                for(auto c:node->children){
                    cnt[c->val]++;
                }
            }
        }
        for(auto node:tree){
            if(cnt[node->val]==tree.size()){
                return node;
            }
        }
        return nullptr;
    }
};

5. Serialize and Deserialize N-ary Tree
請實現一個序列化和反序列化N-ary Tree的算法。

範例輸入:
      1
    / | \
   3  2  4
  / \
 5   6

範例輸出:
"[1 [3[5 6] 2 4]] "


答案:

class Codec {
public:

    // Encodes a tree to a single string.
    string serialize(Node* root) {
        if(root==nullptr) return "";
        string s=to_string(root->val);
        if(root->children.size()==0) return s;
        s+="[";
        for(int i=0;i<root->children.size();i++){
            s+=serialize(root->children[i]);
            if(i<root->children.size()-1){
                s+=" ";
            }
        }
        s+="]";
        return s;
    }

    // Decodes your encoded data to tree.
    Node* deserialize(string data) {
        if(data=="") return nullptr;
        stack<Node*> st;
        Node* root=nullptr;
        string val;
        for(int i=0;i<data.size();i++){
            if(data[i]=='['){
                st.push(new Node(stoi(val),{}));
                val="";
            }
            else if(data[i]==']'){
                auto pre=st.top();
                st.pop();
                if(st.empty()){
                    root=pre;
                }
                else{
                    st.top()->children.push_back(pre);
                }
            }
            else if(data[i]!=' '){
                val+=data[i];
            }
            else{
                if(val!=""){
                    st.top()->children.push_back(new Node(stoi(val),{}));
                }
                val="";
            }
        }
        if(val!=""){
            root=new Node(stoi(val),{});
        }
        return root;
    }
};   

## 网络数据
{{< ask_chatgpt >}}
给我5篇N-ary Tree的网络数据
{{< /ask_chatgpt >}}



1. "N-ary Trees - GeeksforGeeks"
https://www.geeksforgeeks.org/n-ary-tree/

這篇來自 GeeksforGeeks 網站，是一篇導入式的文章，介紹了 N-ary Tree 的概念以及一些常見的應用。文章中會提到 N-ary Tree 的定義、特性，還有一些基本的操作，例如遍歷、插入節點、刪除節點等。

2. "N-ary Tree Definition and Examples - ThoughtCo."
https://www.thoughtco.com/n-ary-tree-definition-and-examples-373286

這篇文章來自 ThoughtCo.，是一篇介紹 N-ary Tree 的基本概念的文章。文章中介紹了 N-ary Tree 的定義、特性，以及一些基本的操作，例如創建、遍歷、複製等。

3. "N-ary Trees in Data Structures - Programiz"
https://www.programiz.com/dsa/n-ary-tree

這篇文章來自 Programiz 網站，是一份專門介紹 N-ary Tree 的資料結構的文章。文章中介紹了 N-ary Tree 的定義、特性，以及一些基本操作和應用，例如遍歷、插入節點、刪除節點等。

4. "N-ary Tree Tutorial - Tutorialspoint"
https://www.tutorialspoint.com/data_structures_algorithms/n_tree.htm

這篇文章來自 Tutorialspoint 網站，是一份 N-ary Tree 的教學文章。文章介紹了 N-ary Tree 的基礎概念，包括定義、特性，還有複雜度分析。同時，文章也介紹了一些基本的操作和應用，例如插入、刪除、遍歷、搜索等。

5. "N-ary Trees in Python - Real Python"
https://realpython.com/n-ary-trees-python/

這篇文章來自 Real Python 網站，是一篇實戰技巧類型的文章，重點介紹了 N-ary Tree 在 Python 中的實現。文章中會展示如何使用 Python 的內建資料結構和模組來創建、操作和遍歷 N-ary Tree。該篇文章中提供了具體的 Python 代碼示例，並搭配詳細的註釋解說。   

