+++
title = "深度優先搜尋"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Depth-First Search的問題
{{< /ask_chatgpt >}}



1. 如何使用深度優先搜索算法（DFS）在無向圖中尋找所有連通份量？

2. 使用DFS尋找無向圖中的最長路徑是如何實現的？

3. 如何使用DFS在有向無環圖中查找最長路徑？

4. 如果我想要使用DFS在迷宮中找到出口，應該如何實現？

5. 如何使用DFS找到一張圖中所有的割點（關鍵節點）？   

## 說明知識
{{< ask_chatgpt >}}
說明Depth-First Search並舉例
{{< /ask_chatgpt >}}



Depth-First Search(DFS)是一種遍歷圖（Graph）或樹（Tree）的演算法，以從起點（Start）開始，往下搜索並幾乎盡可能地走到底部，然後返回搜索路徑，直到發現無法再繼續為止。

DFS遍歷圖的步驟如下：

1. 選擇一個未被訪問過的節點當作起點。
2. 訪問該節點，將其標記為已訪問。
3. 對該節點的相鄰節點進行遍歷，如果相鄰節點未訪問過，就按照相同的方式進行訪問，如果已經訪問過，則返回上一個節點繼續遍歷。
4. 重複步驟3，直到所有節點都被訪問過。

以下是一個DFS的示例，設有一個圖如下所示：

```
A —— B —— C
|     |
D —— E
```

從節點A開始，先把它標記為已訪問，然後從它的相鄰節點中選擇B進行訪問。對節點B進行訪問，在其相鄰節點中選擇C進行訪問，當C訪問完畢，返回B，然後選擇E進行訪問，接下來訪問D，再返回A，最後訪問完畢。

在此遍歷過程中，每個節點只被訪問一次，所以時間複雜度為O（V+E），其中V是節點數，E是邊數。   

## 彙總重點
{{< ask_chatgpt >}}
條列Depth-First Search的重點
{{< /ask_chatgpt >}}



1. DFS 是一種搜索算法，它是以深度為優先考慮，在搜索過程中盡可能深入地訪問每個可訪問節點。
2. DFS 是一種遞歸算法，使用棧（stack）或者遞歸來實現。
3. DFS 可以用於求解圖或樹的遍歷、連通性、最短路徑等問題。
4. DFS 算法的時間複雜度為 O(V+E)，其中 V 表示節點數，E 表示邊數。
5. 在 DFS 中，可以通過標記已訪問節點或使用 visited 數組來避免無限遞歸的問題。
6. DFS 算法分為前序遍歷、中序遍歷和後序遍歷三種方式。
7. DFS 可以應用到許多問題中，例如迷宮求解、拓撲排序、詞法分析、圖像處理等。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Depth-First Search的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 岛屿数量：给定一个由‘0’和‘1’组成的二维网格地图，其中‘1’表示陆地，‘0’表示海水。计算这个地图上岛屿的数量。（LeetCode 200）

答案：https://leetcode.com/problems/number-of-islands/solution/

2. 从键盘输入一个有向图的邻接矩阵，输出其深度优先遍历序列。

答案：https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/

3. 朋友圈的数量：在一个社交圈子里，有 N 个人。每个人都有若干个朋友，这些关系用一个二维矩阵表示，矩阵中的 1 表示两个人互为朋友，0 表示不是。判断这个圈子里有多少个朋友圈。 （LeetCode 547）

答案：https://leetcode.com/problems/friend-circles/solution/

4. 课程安排：将课程表表示成一个名为prerequisites的二维数组，其中 prerequisites[i] = [ai, bi] 表示要想学习课程 i，在修完课程 ai 后必须先修完课程 bi。如果一条直接的修课路线可以完成所有课程程度，则返回 true；否则，返回 false 。（LeetCode 207）

答案：https://leetcode.com/problems/course-schedule/solution/

5. 机器人运动范围：地上有一个 m 行和 n 列的方格，机器人从坐标 (0, 0) 的格子开始移动，每次可以向上、下、左、右四个方向移动一格，但不能进入坐标数位之和大于 k 的格子，求机器人能到达多少个格子。（剑指Offer 13）

答案：https://leetcode.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/solution/   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Depth-First Search的網路資料
{{< /ask_chatgpt >}}



1. Depth-First Search Algorithm and Applications
https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/

This article from GeeksforGeeks provides a comprehensive overview of Depth-First Search (DFS) algorithm and its applications. It explains the concept of DFS, how it works, and its various implementations. It also includes code examples and detailed explanations of how DFS can be used in different scenarios, such as finding connected components in a graph, detecting cycles in a graph, and solving maze problems.

2. Depth First Search (DFS) Algorithm
https://www.tutorialspoint.com/data_structures_algorithms/depth_first_traversal.htm

This tutorial from Tutorials Point explains how DFS is used to traverse a graph or a tree data structure. The article includes detailed explanations of DFS steps, DFS tree, stack-based DFS algorithm, and recursive DFS algorithm. It also provides a sample code snippet to implement DFS in C++.

3. Depth-First Search and Breadth-First Search in Python
https://eddmann.com/posts/depth-first-search-and-breadth-first-search-in-python/

This tutorial from Edward Mann's programming blog provides an overview of DFS and Breadth-First Search (BFS) algorithms and how they can be implemented in Python. The article has code examples that demonstrate how DFS and BFS are used to traverse a graph, including the use of recursion and a queue data structure to implement the algorithm.

4. Graph Traversal – Depth-First Search and Breadth-First Search
https://www.cdn.geeksforgeeks.org/graph-traversal-depth-first-search-and-breadth-first-search/

This article from GeeksforGeeks provides an in-depth overview of graph traversal algorithms, specifically DFS and BFS. It covers both recursive and iterative implementations of DFS, as well as explaining how BFS works. It also includes code examples in C++ and Java, along with a detailed explanation of the algorithm's time complexity.

5. Depth-First Search (DFS) and Applications in Python
https://towardsdatascience.com/depth-first-search-dfs-and-applications-in-python-50d8da7997e1

This article from Towards Data Science provides an overview of DFS and its applications in Python. The author explains how DFS can be used to solve problems such as graph traversal, topological sorting, and cycle detection. The article also includes code examples that demonstrate how DFS can be implemented in Python.   

