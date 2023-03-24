

1. 實現一個 Trie（字典樹），包含 insert, search, 和 startsWith 這三個操作。
答案：

class Trie {
public:
    /** Initialize your data structure here. */
    Trie() {
        root = new TrieNode();
    }
    
    /** Inserts a word into the trie. */
    void insert(string word) {
        TrieNode* node = root;
        for (char c : word) {
            if (!node->children[c - 'a']) {
                node->children[c - 'a'] = new TrieNode();
            }
            node = node->children[c - 'a'];
        }
        node->isEnd = true;
    }
    
    /** Returns if the word is in the trie. */
    bool search(string word) {
        TrieNode* node = root;
        for (char c : word) {
            if (!node->children[c - 'a']) {
                return false;
            }
            node = node->children[c - 'a'];
        }
        return node->isEnd;
    }
    
    /** Returns if there is any word in the trie that starts with the given prefix. */
    bool startsWith(string prefix) {
        TrieNode* node = root;
        for (char c : prefix) {
            if (!node->children[c - 'a']) {
                return false;
            }
            node = node->children[c - 'a'];
        }
        return true;
    }
    
private:
    struct TrieNode {
        bool isEnd;
        TrieNode* children[26];
        TrieNode() {
            isEnd = false;
            memset(children, 0, sizeof(children));
        }
    };
    
    TrieNode* root;
};


2. 給定一個字符串數組 words，請你實現一個能夠將同構字符串組合在一起的分組函數。
    字符串數組形式為 ["eat", "tea", "tan", "ate", "nat", "bat"]，返回分組結果為:
    [
      ["ate","eat","tea"],
      ["nat","tan"],
      ["bat"]
    ]
答案：

class Solution {
public:
    vector<vector<string>> groupAnagrams(vector<string>& strs) {
        unordered_map<string, vector<string>> hash;
        for (string str : strs) {
            string key = getKey(str);
            hash[key].push_back(str);
        }
        vector<vector<string>> ans;
        for (auto it : hash) {
            ans.push_back(it.second);
        }
        return ans;
    }
    
private:
    string getKey(string str) {
        int count[26] = {0};
        for (char c : str) {
            count[c - 'a']++;
        }
        string key;
        for (int i = 0; i < 26; i++) {
            key += to_string(count[i]) + "#";
        }
        return key;
    }
};


3. 給定一個字符串，請你找出其中不含有重複字符的 最長子串 的長度。
    輸入: "abcabcbb"，輸出: 3 ，解釋: 因為無重複字符的最長子串是 "abc"，所以其長度為 3。
答案：

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        if (s.empty()) {
            return 0;
        }
        int ans = 0;
        unordered_map<char, int> hash;
        for (int i = 0, j = 0; j < s.size(); j++) {
            if (hash.find(s[j]) != hash.end() && hash[s[j]] >= i) {
                i = hash[s[j]] + 1;
            }
            hash[s[j]] = j;
            ans = max(ans, j - i + 1);
        }
        return ans;
    }
};


4. 給定一個字符串 s 和一些長度相同的單詞 words。在 s 中找出可以恰好串聯 words 中所有單詞的子串的起始位置。
    輸入:
      s = "barfoothefoobarman",
      words = ["foo", "bar"]
    輸出: [0,9]
    解釋: 從索引 0 和 9 開始的子串分別是 "barfoor" 和 "foobar"。
答案：

class Solution {
public:
    vector<int> findSubstring(string s, vector<string>& words) {
        vector<int> ans;
        int n = s.size(), m = words.size();
        if (n == 0 || m == 0) {
            return ans;
        }
        unordered_map<string, int> hash;
        for (string word : words) {
            hash[word]++;
        }
        int len = words[0].size();
        for (int i = 0; i < len; i++) {
            int left = i, right = i, count = 0;
            unordered_map<string, int> window;
            while (right + len <= n) {
                string str = s.substr(right, len);
                right += len;
                window[str]++;
                count++;
                while (window[str] > hash[str]) {
                    string temp = s.substr(left, len);
                    left += len;
                    window[temp]--;
                    count--;
                }
                if (count == m) {
                    ans.push_back(left);
                }
            }
        }
        return ans;
    }
};


5. 在一個由大小寫字母組成的字符串中，找到一個最長的子串，要求這個子串中的大小寫字母數量相同。
    輸入: "aAbBABAbBa"
    輸出: 4
    解釋: 最長的子串是 "AbBA"，其中有兩個大寫字母和兩個小寫字母。
答案：

class Solution {
public:
    int findTheLongestSubstring(string s) {
        int ans = 0, state = 0, n = s.size();
        unordered_map<int, int> hash{{0, -1}};
        for (int i = 0; i < n; i++) {
            char c = s[i];
            if (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u') {
                state ^= 1 << (c - 'a');
            }
            if (hash.find(state) != hash.end()) {
                ans = max(ans, i - hash[state]);
            } else {
                hash[state] = i;
            }
        }
        return ans;
    }
};

其中，狀態 state 儲存了當前字符串中每個元音字母是否出現了偶數次，用二進制位來表示，e.g. 0b00000 表示當前字符串中的所有元音字母均出現了偶數次，0b00001 表示當前字符串中 a 出現了奇數次，其餘元音字母出現了偶數次，以此類推。注意到當狀態 state 重複出現時，兩種重複狀態之間的字符必定是符合條件的，因為在兩種狀態之間切換，表示其中一個更少使用的元音字母出現次數變化了一次，並且此時兩種狀態在該元音字母上的出現次數必定有偶奇性正好相反。