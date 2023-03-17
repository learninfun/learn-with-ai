

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