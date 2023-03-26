+++
title = "N叉树"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题N-ary Tree的中文问题
{{< /ask_chatgpt >}}



1. 如何计算 N-ary Tree 的深度？
2. 如何实现 N-ary Tree 的前序、中序和后序遍历？
3. 如何找到 N-ary Tree 中的最大值、最小值以及平均值？
4. 如何找到 N-ary Tree 中某一节点的所有祖先节点？
5. 如何实现 N-ary Tree 的层序遍历？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明N-ary Tree并举例
{{< /ask_chatgpt >}}



N-ary Tree是一种树型结构，它由一个根节点和0个或多个子节点组成，但是每个节点可以有多于2个的后代节点。因此，它可以表示多层次的结构，例如组织机构、家族关系等。

举例如下：

一家人的家谱可以用N-ary Tree表示，以祖先为根节点，每个节点代表一位成员，然后子节点可以是该成员的子女，也可以是孙辈、曾孙辈等后代。如下图所示：

![N-ary Tree](https://cdn.luogu.com.cn/upload/image_hosting/kyg2o28s.png)

在这个家谱中，每个节点代表一个人，而每个人都有多位子女、孙辈、曾孙辈等后代，这个结构非常适合用N-ary Tree表示。

除家谱外，N-ary Tree还可以用于表示组织架构、网路拓扑等结构，具有很广泛的应用场景。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列N-ary Tree的重点
{{< /ask_chatgpt >}}



1. N-ary Tree是一种树状结构，每个节点可以有多个子节点。

2. N-ary Tree常常被用来表示有层次结构的数据，例如组织架构或标签分类。

3. N-ary Tree中，每个节点可以有0到n个子节点，其中n称为N值。

4. N-ary Tree的节点通常包含一个值，用来表示节点描述的内容。

5. 在N-ary Tree中，每个节点的子节点没有顺序关系，也没有固定的子节点数目。

6. N-ary Tree的深度可以被定义为从根节点到最深节点的最长路径长度。

7. N-ary Tree的树高可以被定义为根节点到叶节点的最长路径长度。

8. N-ary Tree常常会被用来实现搜寻和排序算法，例如Trie树和堆积排序。

9. N-ary Tree的遍历方式包括前序遍历、中序遍历、后序遍历和层序遍历。

10. 创建N-ary Tree的方法包括递回和迭代算法，具体方法取决于数据结构的特点和应用场景的需求。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题N-ary Tree的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. Maximum Depth of N-ary Tree
给定一个N-ary Tree，回传其深度的最大值。

范例输入:    
       1
     / | \
    3  2  4
   / \
  5   6
范例输出: 3

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
给定一个N-ary Tree，回传其按层次输出的节点序列。

范例输入:    
       1
     / | \
    3  2  4
   / \
  5   6
范例输出: [[1],[3,2,4],[5,6]]

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
给定一个N-ary Tree和一个集合to_delete，请删除N-ary Tree中所有在集合to_delete中的节点并返回删除节点后得到的森林(森林即多个树的集合)。

范例输入:    
       1
     / | \
    3  2  4
   / \
  5   6
删除set=[3,5]

范例输出: [[1,null,4],[6],[2]]

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
给定一个N-ary Tree，其中所有节点除根节点外都有一个父节点。找到并返回这个N-ary Tree的根节点。

范例输入:
 [[1,null,3,2,4,null,5,6]]
范例输出: 1

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
请实现一个序列化和反序列化N-ary Tree的算法。

范例输入:
      1
    / | \
   3  2  4
  / \
 5   6

范例输出:
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

