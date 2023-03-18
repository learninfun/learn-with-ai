+++
title = "字典树资料结构"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Trie Data Structure的问题
{{< /ask_chatgpt >}}



1. 如何在Trie中查找一个字符串？

2. 如何向Trie中添加一个字符串？

3. 如何删除Trie中的一个字符串？

4. 如何统计Trie中所有字符串出现的次数？

5. 如何在Trie中查找所有以给定字符串为前缀的字符串？   

## 说明知识
{{< ask_chatgpt >}}
说明Trie Data Structure并举例
{{< /ask_chatgpt >}}



Trie (也称为"字典树"或"前缀树")是一种能够有效储存字串的数据结构。它的主要优势在于快速查找字串，而且其查询的时间复杂度只和被查询字串的长度有关无论数据集中有多少字串。

Trie的结构是一个根节点，往下分支出多个子节点，每个子节点代表某个字母，而节点与节点之间的连线代表字母之间的关系。换句话说，Trie是一种多层次的数据结构，每一层代表一个字母，而每个节点可以存储任意数量的子节点。

例如，在下图中，是一个包含一些单词的Trie。

```
         (root)
          /  \
         a    b
        / \    \
       n   p    y
      /     \    \
     d       t    e
    / \     / \    \
   e   y   o   a    s
  /     \         / 
 l       e       t
```

在这个Trie中，我们可以看到单词“and”、“ape”和“any”都被储存在根节点的第一层子节点之下，以此类推。单词的结尾可以特别注明，例如在“and”单词的最后一个字母节点上加上一个标记。这样我们就可以简单地识别单词的结尾，而不必将整个单词储存在每个节点中。

Trie的查询速度很快，因为每次查询只需要从Trie的根节点开始往下遍历，直到找到目标字串的最后一个字母节点。如果目标字串不存在，就无法找到对应的节点，这样就可以快速得出结论。

使用Trie可以轻松地实现自动完成、拼写检查、搜索引擎等功能。

关于Trie的应用：

- 单词查找
- 自动完成
- 拼写检查
- 搜索引擎
- 统计单词的出现频率等等。   

## 汇总重点
{{< ask_chatgpt >}}
条列Trie Data Structure的重点
{{< /ask_chatgpt >}}



1. Trie是一种数据结构，它以树形结构存储字符串，并且能够快速查询和插入字符串。

2. Trie的根节点代表一个空字符串，每个节点都包含一个字符和对应子节点的指针。

3. Trie结构中每个节点上的字元皆不相同，即对于同一个字符串，没有两个节点包含相同的字符。

4. Trie结构能够高效的搜寻、插入、删除字符串，时间复杂度与字符串长度呈线性关系。

5. Trie结构的应用包括：单词查询、自动补全、字符串匹配、字符串压缩等。

6. Trie结构优化方式包括：压缩型Trie、词频统计型Trie等。

7. Trie中的变种结构包括：可压缩的Trie、可关键字覆盖的Trie、有权重的Trie等。

8. Trie的缺点是占用空间较大，并且对于含有大量相同前缀的字符串，Trie的效率不如其他数据结构。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Trie Data Structure的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Trie Data Structure的网络数据
{{< /ask_chatgpt >}}



1. "Trie Data Structure" - GeeksforGeeks ：https://www.geeksforgeeks.org/trie-data-structure/

這篇文章提供了詳盡的 Trie Data Structure 的介紹，也包含了 Trie 的基本原理、插入、刪除、搜尋等操作，並且透過範例讓讀者更易理解 Trie 的使用方法。

2. "Trie (Prefix Tree) & its implementation" - TutorialsPoint ：https://www.tutorialspoint.com/data_structures_algorithms/trie_algorithm.htm

TutorialsPoint 提供了與 Trie 有關的技術用語與相關概念的講解，並以 Python 為範例，解析了 Trie 的程式實作方法，讓讀者可輕鬆學會使用 Trie。

3. "Trie：字典查詢，完美解法" - 程序員垃圾場：https://mp.weixin.qq.com/s/5ZszCn2XfsXLtxVBmB2QWw

這篇文章以中文呈現，專注於 Trie 資料結構的應用，透過插入操作與字典查詢的範例，成功解釋了 Trie 的語法與使用方法。

4. "A Complete Guide to Trie Data Structure" - Towards Data Science：https://towardsdatascience.com/a-complete-guide-to-trie-data-structure-2626db71bc98

這篇文章使用圖片來輔助說明 Trie 的使用，並且透過比較不同使用 Trie 的情境，進一步的解說 Trie 的效能與優勢。

5. "Trie Data Structure" - Programiz：https://www.programiz.com/dsa/trie-data-structure

Programiz 透過專業的方式，提供了 Trie 的基礎知識、實作、示例和小結等內容，並以簡潔明瞭的方式呈現給讀者，讓讀者能快速了解 Trie 的物件結構特征和使用方式。   

