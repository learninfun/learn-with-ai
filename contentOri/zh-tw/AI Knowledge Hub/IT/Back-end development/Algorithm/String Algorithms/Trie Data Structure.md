## 習題預習
{{< ask_chatgpt >}}
給我5題Trie Data Structure的問題
{{< /ask_chatgpt >}}



1. 如何在Trie中查找一个字符串？

2. 如何向Trie中添加一个字符串？

3. 如何删除Trie中的一个字符串？

4. 如何统计Trie中所有字符串出现的次数？

5. 如何在Trie中查找所有以给定字符串为前缀的字符串？   

## 說明知識
{{< ask_chatgpt >}}
說明Trie Data Structure並舉例
{{< /ask_chatgpt >}}



Trie (也稱為"字典樹"或"前綴樹")是一種能夠有效儲存字串的數據結構。它的主要優勢在於快速查找字串，而且其查詢的時間複雜度只和被查詢字串的長度有關無論數據集中有多少字串。

Trie的結構是一個根節點，往下分支出多個子節點，每個子節點代表某個字母，而節點與節點之間的連線代表字母之間的關係。換句話說，Trie是一種多層次的數據結構，每一層代表一個字母，而每個節點可以存儲任意數量的子節點。

例如，在下圖中，是一個包含一些單詞的Trie。

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

在這個Trie中，我們可以看到單詞“and”、“ape”和“any”都被儲存在根節點的第一層子節點之下，以此類推。單詞的結尾可以特別註明，例如在“and”單詞的最後一個字母節點上加上一個標記。這樣我們就可以簡單地識別單詞的結尾，而不必將整個單詞儲存在每個節點中。

Trie的查詢速度很快，因為每次查詢只需要從Trie的根節點開始往下遍歷，直到找到目標字串的最後一個字母節點。如果目標字串不存在，就無法找到對應的節點，這樣就可以快速得出結論。

使用Trie可以輕鬆地實現自動完成、拼寫檢查、搜索引擎等功能。

關於Trie的應用：

- 單詞查找
- 自動完成
- 拼寫檢查
- 搜索引擎
- 統計單詞的出現頻率等等。   

## 彙總重點
{{< ask_chatgpt >}}
條列Trie Data Structure的重點
{{< /ask_chatgpt >}}



1. Trie是一種數據結構，它以樹形結構存儲字符串，並且能夠快速查詢和插入字符串。

2. Trie的根節點代表一個空字符串，每個節點都包含一個字符和對應子節點的指針。

3. Trie結構中每個節點上的字元皆不相同，即對於同一個字符串，沒有兩個節點包含相同的字符。

4. Trie結構能夠高效的搜尋、插入、刪除字符串，時間複雜度與字符串長度呈線性關係。

5. Trie結構的應用包括：單詞查詢、自動補全、字符串匹配、字符串壓縮等。

6. Trie結構優化方式包括：壓縮型Trie、詞頻統計型Trie等。

7. Trie中的變種結構包括：可壓縮的Trie、可關鍵字覆蓋的Trie、有權重的Trie等。

8. Trie的缺點是占用空間較大，並且對於含有大量相同前缀的字符串，Trie的效率不如其他數據結構。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Trie Data Structure的中等難度問題，並在後面列出答案
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

## 網路資料
{{< ask_chatgpt >}}
給我5篇Trie Data Structure的網路資料
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

