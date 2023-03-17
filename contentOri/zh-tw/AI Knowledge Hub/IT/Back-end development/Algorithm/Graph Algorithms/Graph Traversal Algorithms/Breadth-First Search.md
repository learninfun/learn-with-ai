## 習題預習
{{< ask_chatgpt >}}
給我5題Breadth-First Search的問題
{{< /ask_chatgpt >}}



1. 在一張無向圖中，如何找出兩個給定節點之間的最短路徑？
2. 在一個網格中，如何找到從起點到終點的最短路徑，其中障礙物可能會阻擋移動？
3. 如何在一個社交網絡中找到關於某個人的直接或間接聯繫，並計算出與他們之間的距離？
4. 如何從起點到盡可能多的目標節點，而不必找到最短路徑？
5. 如何在一個巨大的圖形資料庫中找到與搜索條件最相符的所有圖形，並以最短路徑顯示它們之間的關係？   

## 說明知識
{{< ask_chatgpt >}}
說明Breadth-First Search並舉例
{{< /ask_chatgpt >}}



Breadth-First Search（BFS）是一種搜尋演算法，以廣度優先的方式來遍歷一個圖形或樹狀結構。

其遍歷的順序是逐層往下，也就是先遍歷所有的同一深度節點，再遍歷下一深度的節點。在BFS遍歷中，使用一個隊列來維護已經被廣度遍歷的節點，以便按層訪問下一階段節點時使用。

舉個例子：假設我們有一個有向圖如下圖所示：

<img src="https://i.imgur.com/Vhh9XfO.png" width="200"/>

我們從節點1開始進行BFS遍歷，首先將節點1加入隊列中。接下來，按照節點編號的大小順序，先遍歷節點2和節點3。

<img src="https://i.imgur.com/s8BDGdc.png" width="200"/>

然後，把節點2的相鄰節點4，7加入隊列中，把節點3的相鄰節點5，6加入隊列中。

<img src="https://i.imgur.com/gPXs4L4.png" width="200"/>

再遍歷節點4和節點7，因為它們沒有相鄰節點可加入隊列中，所以直接跳過。

<img src="https://i.imgur.com/yujfiFB.png" width="200"/>

最後，遍歷節點5和節點6，發現節點5有一個相鄰節點8，所以把節點8加入隊列中。遍歷完節點5和節點6，隊列已經空了，此時遍歷結束。

<img src="https://i.imgur.com/iCCvSTN.png" width="200"/>

這個例子中，我們首先訪問距離起點節點最近的節點，然後是次近的節點，然後是更遠的節點。這種BFS的訪問順序可以確保我們在最短時間內訪問到所有節點。   

## 彙總重點
{{< ask_chatgpt >}}
條列Breadth-First Search的重點
{{< /ask_chatgpt >}}



1. 廣度優先搜尋演算法（BFS）是一種用來搜索圖形或樹形結構的技術。

2. BFS從起始節點開始搜尋，然後逐層擴展，直到達到終點或所有節點都被訪問為止。

3. BFS適用於找到最短路徑和最少操作的問題，因為它保證了先找到的路徑長度最短或者操作最少。

4. BFS使用FIFO（先進先出）佇列來保存待處理的節點，這有助於記錄搜尋順序和計算層次。

5. BFS通常需要使用標記訪問過的節點，以避免重複訪問和死循環。

6. BFS可以用來應對未知的圖形和樹形結構，並且可以與其他搜尋演算法結合使用。

7. BFS的時間複雜度為O(V+E)，其中V是圖形的節點數，E是圖形的邊數。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Breadth-First Search的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. Word Ladder
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list. 
Note that beginWord is not a transformed word.

Example:
Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5

Answer:
https://leetcode.com/problems/word-ladder/

2. Minimum Genetic Mutation
A gene string can be represented by an 8-character long string, with choices from "A", "C", "G", "T".

Suppose we need to investigate about a mutation (mutation from "start" to "end"), where ONE mutation is defined as ONE single character changed in the gene string.

For example, "AACCGGTT" -> "AACCGGTA" is 1 mutation.

Also, there is a given gene "bank", which records all the valid gene mutations. A gene must be in the bank to make it a valid gene string.

Now, given 3 things - start, end, bank, your task is to determine what is the minimum number of mutations needed to mutate from "start" to "end". If there is no such a mutation, return -1.

Note:

Starting point is assumed to be valid, so it might not be included in the bank.
If multiple mutations are needed, all mutations during in the sequence must be valid.
You may assume start and end string is not the same.

Example:
start: "AACCGGTT"
end:   "AACCGGTA"
bank: ["AACCGGTA"]

return: 1
 

Answer:
https://leetcode.com/problems/minimum-genetic-mutation/

3. Bus Routes
We have a list of bus routes. Each routes[i] is a bus route that the i-th bus repeats forever. For example if routes[0] = [1, 5, 7], this means that the first bus (0-th indexed) travels in the sequence 1->5->7->1->5->7->1->... forever.

We start at bus stop S (initially not on a bus), and we want to go to bus stop T. Travelling by buses only, what is the least number of buses we must take to reach our destination? Return -1 if it is not possible.

Example:
Input: 
routes = [[1, 2, 7], [3, 6, 7]]
S = 1
T = 6
Output: 2

Answer:
https://leetcode.com/problems/bus-routes/

4. Cut Off Trees for Golf Event
You are asked to cut off trees in a forest for a golf event. The forest is represented as a non-negative 2D map, in this map:

0 represents the obstacle can't be reached.
1 represents the ground can be walked through.
The place with number bigger than 1 represents a tree can be walked through, and this positive number represents the tree's height.

You are asked to cut off all the trees in this forest in the order of tree's height - always cut off the tree with lowest height first. And after cutting, the original place has the tree will become a grass (value 1).

You will start from the point (0, 0) and you should output the minimum steps you need to walk to cut off all the trees. If you can't cut off all the trees, output -1 in that situation.

You are guaranteed that no two trees have the same height and there is at least one tree needs to be cut off.

Example:
Input: 
[
 [1,2,3],
 [0,0,4],
 [7,6,5]
]
Output: 6

Answer:
https://leetcode.com/problems/cut-off-trees-for-golf-event/

5. Pacific Atlantic Water Flow
Given an m x n matrix of non-negative integers representing the height of each unit cell in a continent, the "Pacific ocean" touches the left and top edges of the matrix and the "Atlantic ocean" touches the right and bottom edges.

Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.

Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.

Note:

The order of returned grid coordinates does not matter.
Both m and n are less than 150.

Example:
Given the following 5x5 matrix:

  Pacific ~   ~   ~   ~ 
       ~  1   2   2   3  (5) 
       ~  3   2   3  (4) (4) 
       ~  2   4  (5)  3   1  
       ~ (6) (7)  1   4   5  
           Atlantic   ~   ~   ~   ~
           
Answer:
https://leetcode.com/problems/pacific-atlantic-water-flow/   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Breadth-First Search的網路資料
{{< /ask_chatgpt >}}



1. "Introduction to Breadth-First Search (BFS) Algorithm" by GeeksforGeeks - This article provides a detailed introduction to the BFS algorithm, explaining how it works and how it is implemented. It includes code examples and illustrations to help readers understand the concept.

2. "Breadth First Search (BFS) algorithm" by Tutorials Point - This tutorial offers a step-by-step guide to implementing BFS, including pseudocode and a sample implementation. It also explains how BFS can be used for different applications, such as finding the shortest path or distributed computing.

3. "Breadth-First Search in Artificial Intelligence" by Analytics India Magazine - This article explores how BFS is used in AI, particularly in solving problems related to search and optimization. It explains the advantages and limitations of BFS, as well as its variations and applications in other domains.

4. "Breadth-First Search Algorithm" by Computer Science Refined - This article provides a comprehensive overview of BFS, including examples of how it can be applied in various scenarios. It also explains the differences between BFS and other search algorithms, such as depth-first search (DFS).

5. "Breadth First Search (BFS): A Graph Traversal Algorithm with Examples in Python" by DataCamp - This tutorial introduces BFS using Python code samples and visualizations. It explains the basic concepts of graph theory and traversal, and shows how BFS can be applied to solve problems such as finding the shortest path in a graph.   

