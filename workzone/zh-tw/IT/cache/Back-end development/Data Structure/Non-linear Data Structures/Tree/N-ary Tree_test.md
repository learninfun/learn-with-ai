

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