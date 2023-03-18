+++
title = "字典树"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Trie的问题
{{< /ask_chatgpt >}}



1. 如何在Trie上实现字典搜索功能，即查找一个字符串是否存在于一个词典中？
2. 如何在Trie上实现前缀搜索功能，即查找所有以一个特定字符串作为前缀的字串？
3. 如何在Trie中实现自动完成功能，当输入一部分已知的字符串时，自动推荐可能的后缀字符串？
4. 如何优化Trie节省空间？可以考虑使用压缩词典树等技术。
5. 如何实现在Trie上表示多个字符串？可以使用特殊符号作为字符串的分隔符，也可以将多个字符串合并为一个字符串再建树。   

## 说明知识
{{< ask_chatgpt >}}
说明Trie并举例
{{< /ask_chatgpt >}}



Trie，也称为字典树或前缀树，是一种数据结构，用于有效地存储和检索字符串。 

Trie树特别适合于应用场景，如搜寻输入框中的自动完成，单词应用程序的拼写检查或基因组学中的基因组匹配等。

Trie树由一组节点构成，其中每个节点代表一个字符，节点之间的连接代表字符之间的关系。每条从根节点到叶节点的路径代表一个字符串。

下面是一个例子：

假设要存储以下字典：apple，applaud，application，apply，book，boxing

首先，创建一个空的Trie树，如下图所示：

![image.png](attachment:image.png)

然后，将字典中的单词添加到Trie树中。每个字母都在Trie树中表示为一个节点。下面是修改后的Trie树：

![image-2.png](attachment:image-2.png)

Trie树允许更快地查找字符串。以“appl”为例，Trie树可以从根节点开始，按照“a”，“p”，“p”顺序遍历它的三个子节点。当它到达第三个节点时，它会发现它是一个单词的结尾，因此它可以确定“appl”字符串在字典中存在。

总体来说，Trie树是一种效率高且易于实现的数据结构，适用于许多应用程序场景。   

## 汇总重点
{{< ask_chatgpt >}}
条列Trie的重点
{{< /ask_chatgpt >}}



1. Trie是一種數據結構，可供高效地存儲和查詢字符串。

2. Trie使用樹形結構來表示所有可能的字符串，每個節點代表一個字符。

3. Trie具有快速查詢複雜度，可以在O(m)的時間內查詢一個長度為m的字符串。

4. Trie也可以用於搜索前綴匹配，通過查詢一個前綴，可以找到所有匹配該前綴的字符串。

5. Trie可以用於字典，拼字檢查和自動完成等應用程序。

6. Trie的缺點是它需要使用大量的空間來存儲所有可能的字符串，尤其是當數據集很大時。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Trie的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 实现一个trie树，包含addWord, searchWord, searchPrefix三个方法。其中addWord(word)添加单词word到trie树中，searchWord(word)返回trie树是否包含单词word，searchPrefix(prefix)返回包含前缀prefix的所有单词列表。

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


2. 给定一个由一些正整数(以字符串形式给出)组成的数组, 将这些数字连接起来，使之成为一个最小的数。

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


3. 给定一个字符集，其包含一些不重复的字符，表示一个字符串集合。你需要使用这个字符集来制作字典，构造出一个能够识别给定词语的最小的字典。需要输出这个最小词典里的字符列表。

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

        
4. 给定一个字符串s, 一个单词字典words, 一个最大宽度maxWidth；把s划分成若干个单词子串，每个子串具有相同的长度且不超过maxWidth；一个单词只能出现在一个子串中。返回划分方案，每个子串以一个空格隔开且其长度为maxWidth，最后一个子串没有多余的空格。子串与子串之间要保证不出现多余的空格。如果有多个答案，返回字典序最小的答案。

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

        
5. 给定一组单词，找出其中所有的单词接龙。例如: ["hot","dot","dog","lot","log","cog"], 所有的单词接龙如下: 
"hot" -> "dot" -> "dog" -> "cog"
"hot" -> "lot" -> "log" -> "cog"
返回所有的路径。

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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Trie的网络数据
{{< /ask_chatgpt >}}



1. Trie - GeeksforGeeks
网址：https://www.geeksforgeeks.org/trie-insert-and-search/
这篇文章介绍了 Trie 的基本操作：插入和搜寻。并且提供了 C++ 的程式码，也有范例可以测试。

2. Trie Dictionary and Autocomplete - Medium
网址：https://medium.com/swlh/trie-dictionary-and-autocomplete-a71ae077f532
这篇文章介绍了 Trie 树的应用：字典和自动完成。除了基本的搜索操作，还介绍了整个树的建构和删除。也有 Swift 的程式码示范。

3. Trie - Competitive Programming Algorithms - Topcoder
网址：https://www.topcoder.com/thrive/articles/Trie%20-%20Competitive%20Programming%20Algorithms
这篇文章介绍了 Trie 的一些更进阶使用方法，例如文本字符串搜索、点的计数、前缀和等等。还有一个近似匹配的应用例子。也有 Python 3 的程式码。

4. Implementing a trie in Java - Baeldung
网址：https://www.baeldung.com/trie-java
这篇文章提供了 Java  语言实现 Trie 的方法，同时也介绍了 Trie 的资料结构和性能优化等方面的知识。示范了 Trie 用于单词搜索和前缀搜寻。

5. Visualization of Trie (Prefix Tree) - VisuAlgo
网址：https://visualgo.net/en/trie
这篇文章提供了一个可视化平台，用来直观地理解 Trie 树的建构和搜寻。更加动态地展示 Trie 操作过程，让读者可以更加直观知道其中的细节。此外，平台中也提供了对试验数据可能性的测试，创造出与实际场景相似的条件。   

