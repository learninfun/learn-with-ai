+++
title = "字典樹"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Trie的問題
{{< /ask_chatgpt >}}



1. 如何在Trie上實現字典搜索功能，即查找一個字符串是否存在於一個詞典中？
2. 如何在Trie上實現前綴搜索功能，即查找所有以一個特定字符串作為前綴的字串？
3. 如何在Trie中實現自動完成功能，當輸入一部分已知的字符串時，自動推薦可能的後綴字符串？
4. 如何優化Trie節省空間？可以考慮使用壓縮詞典樹等技術。
5. 如何實現在Trie上表示多個字符串？可以使用特殊符號作為字符串的分隔符，也可以將多個字符串合併為一個字符串再建樹。   

## 說明知識
{{< ask_chatgpt >}}
說明Trie並舉例
{{< /ask_chatgpt >}}



Trie，也稱為字典樹或前綴樹，是一種數據結構，用於有效地存儲和檢索字符串。 

Trie樹特別適合於應用場景，如搜尋輸入框中的自動完成，單詞應用程序的拼寫檢查或基因組學中的基因組匹配等。

Trie樹由一組節點構成，其中每個節點代表一個字符，節點之間的連接代表字符之間的關係。每條從根節點到葉節點的路徑代表一個字符串。

下面是一個例子：

假設要存儲以下字典：apple，applaud，application，apply，book，boxing

首先，創建一個空的Trie樹，如下圖所示：

![image.png](attachment:image.png)

然後，將字典中的單詞添加到Trie樹中。每個字母都在Trie樹中表示為一個節點。下面是修改後的Trie樹：

![image-2.png](attachment:image-2.png)

Trie樹允許更快地查找字符串。以「appl」為例，Trie樹可以從根節點開始，按照「a」，「p」，「p」順序遍歷它的三個子節點。當它到達第三個節點時，它會發現它是一個單詞的結尾，因此它可以確定「appl」字符串在字典中存在。

總體來說，Trie樹是一種效率高且易於實現的數據結構，適用於許多應用程序場景。   

## 彙總重點
{{< ask_chatgpt >}}
條列Trie的重點
{{< /ask_chatgpt >}}



1. Trie是一種數據結構，可供高效地存儲和查詢字符串。

2. Trie使用樹形結構來表示所有可能的字符串，每個節點代表一個字符。

3. Trie具有快速查詢複雜度，可以在O(m)的時間內查詢一個長度為m的字符串。

4. Trie也可以用於搜索前綴匹配，通過查詢一個前綴，可以找到所有匹配該前綴的字符串。

5. Trie可以用於字典，拼字檢查和自動完成等應用程序。

6. Trie的缺點是它需要使用大量的空間來存儲所有可能的字符串，尤其是當數據集很大時。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Trie的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 實現一個trie樹，包含addWord, searchWord, searchPrefix三個方法。其中addWord(word)添加單詞word到trie樹中，searchWord(word)返回trie樹是否包含單詞word，searchPrefix(prefix)返回包含前綴prefix的所有單詞列表。

class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_word = False
        
class Trie:
    def __init__(self):
        self.root = TrieNode()

    def addWord(self, word: str) -> None:
        node = self.root
        for char in word:
            if char not in node.children:
                node.children[char] = TrieNode()
            node = node.children[char]
        node.is_word = True

    def searchWord(self, word: str) -> bool:
        node = self.root
        for char in word:
            if char not in node.children:
                return False
            node = node.children[char]
        return node.is_word

    def searchPrefix(self, prefix: str) -> List[str]:
        node = self.root
        for char in prefix:
            if char not in node.children:
                return []
            node = node.children[char]
        words = []
        self.dfs(node, prefix, words)
        return words

    def dfs(self, node, word, words):
        if node.is_word:
            words.append(word)
        for char in node.children:
            self.dfs(node.children[char], word + char, words)


2. 給定一個由一些正整數(以字符串形式給出)組成的數組, 將這些數字連接起來，使之成為一個最小的數。

class Node:
    def __init__(self):
        self.children = {}
        self.is_end = False
        
class Solution:
    def minNumber(self, nums: List[int]) -> str:
        trie = Node()
        for num in nums:
            node = trie
            for char in str(num):
                if char not in node.children:
                    node.children[char] = Node()
                node = node.children[char]
            node.is_end = True
        
        res = []
        self.dfs(trie, '', res)
        return ''.join(res)
    
    def dfs(self, node, path, res):
        if node.is_end:
            res.append(path)
            return 
        for char, child in sorted(node.children.items()):
            self.dfs(child, path + char, res)


3. 給定一個字符集，其包含一些不重複的字符，表示一個字符串集合。你需要使用這個字符集來製作字典，構造出一個能夠識別給定詞語的最小的字典。需要輸出這個最小詞典裡的字符列表。

class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_end = False

class Solution:
    def minimumCharSet(self, words: List[str]) -> str:
        trie = TrieNode()
        for word in words:
            node = trie
            for char in word:
                if char not in node.children:
                    node.children[char] = TrieNode()
                node = node.children[char]
            node.is_end = True
            
        queue = deque([trie])
        res = []
        while queue:
            node = queue.popleft()
            for char in sorted(node.children.keys()):
                child = node.children[char]
                if child.is_end:
                    res.append(char)
                queue.append(child)
        return ''.join(res)

        
4. 給定一個字符串s, 一個單詞字典words, 一個最大寬度maxWidth；把s劃分成若干個單詞子串，每個子串具有相同的長度且不超過maxWidth；一個單詞只能出現在一個子串中。返回劃分方案，每個子串以一個空格隔開且其長度為maxWidth，最後一個子串沒有多餘的空格。子串與子串之間要保證不出現多餘的空格。如果有多個答案，返回字典序最小的答案。

class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_word = False

class Solution:
    def wordBreak(self, s: str, wordDict: List[str], maxWidth: int) -> List[str]:
        trie = TrieNode()
        for word in wordDict:
            node = trie
            for char in word:
                if char not in node.children:
                    node.children[char] = TrieNode()
                node = node.children[char]
            node.is_word = True

        n = len(s)
        dp = [-1] * n
        end = self.dfs(trie, s, 0, dp)
        if end == -1:
            return []

        res = []
        self.dfs2(s, 0, end, maxWidth, [], res, dp)
        return res

    def dfs(self, trie, s, start, dp):
        if start == len(s):
            return start
        if dp[start] != -1:
            return dp[start]
        node = trie
        end = -1
        for i in range(start, len(s)):
            if s[i] not in node.children:
                break
            node = node.children[s[i]]
            if node.is_word:
                end = i
            if i == len(s) - 1:
                end = i + 1
        dp[start] = end
        return end

    def dfs2(self, s, start, end, maxWidth, path, res, dp):
        if start == end:
            res.append(' '.join(path))
            return
        for i in range(start + 1, end + 1):
            if i - start > maxWidth:
                break
            if dp[i] == -1:
                continue
            path.append(s[start:i])
            self.dfs2(s, i, end, maxWidth, path, res, dp)
            path.pop()

        
5. 給定一組單詞，找出其中所有的單詞接龍。例如: ["hot","dot","dog","lot","log","cog"], 所有的單詞接龍如下: 
"hot" -> "dot" -> "dog" -> "cog"
"hot" -> "lot" -> "log" -> "cog"
返回所有的路徑。

class TrieNode:
    def __init__(self):
        self.children = {}
        self.is_word = False
        
class Solution:
    def findLadders(self, beginWord: str, endWord: str, wordList: List[str]) -> List[List[str]]:
        trie = TrieNode()
        for word in wordList:
            node = trie
            for char in word:
                if char not in node.children:
                    node.children[char] = TrieNode()
                node = node.children[char]
            node.is_word = True
        
        if endWord not in wordList:
            return []
        
        paths = []
        path = [beginWord]
        visited = set()
        visited.add(beginWord)
        self.dfs(beginWord, endWord, trie, words, visited, path, paths)
        return paths
    
    def dfs(self, cur, endWord, trie, words, visited, path, paths):
        if cur == endWord:
            paths.append(path)
            return
        
        for i in range(len(cur)):
            for j in range(26):
                c = chr(j + ord('a'))
                if cur[i] == c:
                    continue
                next_word = cur[:i] + c + cur[i+1:]
                if next_word in visited or next_word not in words:
                    continue
                node = trie
                for char in next_word:
                    node = node.children[char]
                visited.add(next_word)
                self.dfs(next_word, endWord, trie, words, visited, path + [next_word], paths)
                visited.remove(next_word)   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Trie的網路資料
{{< /ask_chatgpt >}}



1. Trie - GeeksforGeeks
網址：https://www.geeksforgeeks.org/trie-insert-and-search/
這篇文章介紹了 Trie 的基本操作：插入和搜尋。並且提供了 C++ 的程式碼，也有範例可以測試。

2. Trie Dictionary and Autocomplete - Medium
網址：https://medium.com/swlh/trie-dictionary-and-autocomplete-a71ae077f532
這篇文章介紹了 Trie 樹的應用：字典和自動完成。除了基本的搜索操作，還介紹了整個樹的建構和刪除。也有 Swift 的程式碼示範。

3. Trie - Competitive Programming Algorithms - Topcoder
網址：https://www.topcoder.com/thrive/articles/Trie%20-%20Competitive%20Programming%20Algorithms
這篇文章介紹了 Trie 的一些更進階使用方法，例如文本字符串搜索、點的計數、前綴和等等。還有一個近似匹配的應用例子。也有 Python 3 的程式碼。

4. Implementing a trie in Java - Baeldung
網址：https://www.baeldung.com/trie-java
這篇文章提供了 Java  語言實現 Trie 的方法，同時也介紹了 Trie 的資料結構和性能優化等方面的知識。示範了 Trie 用於單詞搜索和前綴搜尋。

5. Visualization of Trie (Prefix Tree) - VisuAlgo
網址：https://visualgo.net/en/trie
這篇文章提供了一個可視化平台，用來直觀地理解 Trie 樹的建構和搜尋。更加動態地展示 Trie 操作過程，讓讀者可以更加直觀知道其中的細節。此外，平台中也提供了對試驗數據可能性的測試，創造出與實際場景相似的條件。   

