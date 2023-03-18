

1. 实现一个 Trie（字典树），包含 insert, search, 和 startsWith 这三个操作。
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


2. 给定一个字符串数组 words，请你实现一个能够将同构字符串组合在一起的分组函数。
    字符串数组形式为 ["eat", "tea", "tan", "ate", "nat", "bat"]，返回分组结果为:
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


3. 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
    输入: "abcabcbb"，输出: 3 ，解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
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


4. 给定一个字符串 s 和一些长度相同的单词 words。在 s 中找出可以恰好串联 words 中所有单词的子串的起始位置。
    输入:
      s = "barfoothefoobarman",
      words = ["foo", "bar"]
    输出: [0,9]
    解释: 从索引 0 和 9 开始的子串分别是 "barfoor" 和 "foobar"。
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


5. 在一个由大小写字母组成的字符串中，找到一个最长的子串，要求这个子串中的大小写字母数量相同。
    输入: "aAbBABAbBa"
    输出: 4
    解释: 最长的子串是 "AbBA"，其中有两个大写字母和两个小写字母。
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

其中，状态 state 储存了当前字符串中每个元音字母是否出现了偶数次，用二进制位来表示，e.g. 0b00000 表示当前字符串中的所有元音字母均出现了偶数次，0b00001 表示当前字符串中 a 出现了奇数次，其余元音字母出现了偶数次，以此类推。注意到当状态 state 重复出现时，两种重复状态之间的字符必定是符合条件的，因为在两种状态之间切换，表示其中一个更少使用的元音字母出现次数变化了一次，并且此时两种状态在该元音字母上的出现次数必定有偶奇性正好相反。