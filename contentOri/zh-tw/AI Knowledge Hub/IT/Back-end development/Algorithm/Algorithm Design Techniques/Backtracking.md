## 習題預習
{{< ask_chatgpt >}}
給我5題Backtracking的問題
{{< /ask_chatgpt >}}



1. N皇后問題 - 在一個N × N的棋盤上放置N個皇后，使得它們不能相互攻擊（即：不同行，不同列和不同對角線上不能有兩個皇后）。
2. 0/1背包問題 - 在有限的容量下，給出一個物品列表和它們的重量和價值，找出一個最佳方案，使得總重量不超過容量限制，而總價值最大化。
3. 字符串排列 - 給定一個字符串，返回它的所有可能的排列集合。
4. 圖形著色問題 - 在一個圖形中，著色每個點，使得相鄰的兩個點都有不同的顏色，最少使用多少種顏色。
5. 八皇后問題 - 在一個8×8的棋盤上放置8個皇后，使得它們不能相互攻擊。   

## 說明知識
{{< ask_chatgpt >}}
說明Backtracking並舉例
{{< /ask_chatgpt >}}



背朝追蹤是解決問題的一種演算法，其中通過在解決方案的過程中回溯並反悔，以避免訪問無用的狀態和降低演算法的複雜性。這個演算法通常在求解組合問題（如最短路徑、旅行商問題、數獨）和搜索問題（如填字、八皇后）時使用。

在背朝追蹤中，我們開始尋找問題的解決方案，並假定解決方案是正確的。如果發現某個部分無法通過限制條件，則回溯到先前考慮的狀態並嘗試其他解決方案。在這個過程中，我們保留了先前的計算和結果，以節省計算成本。

舉一個背朝追蹤演算法的例子，考慮以下問題：從一個起點開始，找到一條到達終點的最短路徑。我們可以使用深度優先搜索來找到路徑，並通過回溯來避免重複搜索相同的狀態。以下是一個簡單的背向搜索算法，在網格上搜索從起點到終點的最短路徑：

```python
def backtrack(start, end, grid, visited):
    # Base case: we have reached the end
    if start == end:
        return [end]
    
    # Check if we already visited this state
    if start in visited:
        return []
    
    visited.add(start)
    
    # Try moving in all directions
    rows, cols = len(grid), len(grid[0])
    x, y = start
    candidates = []
    
    for dx, dy in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
        nx, ny = x + dx, y + dy
        if 0 <= nx < rows and 0 <= ny < cols and grid[nx][ny]:
            candidates.append((nx, ny))
    
    # Try to backtrack using each candidate
    for next_pos in candidates:
        path = backtrack(next_pos, end, grid, visited)
        if path:
            return [start] + path
    
    # No valid path found, so backtrack
    return []
```

使用背向搜索算法，我們可以以線性時間和空間複雜度找到從一個起點到另一個終點的最短路徑。   

## 彙總重點
{{< ask_chatgpt >}}
條列Backtracking的重點
{{< /ask_chatgpt >}}



1. 確定問題的解空間：確定問題的可行解空間。

2. 確定解的表示：將解表示成某種數據結構（例如：數組、集合等）。

3. 確定約束條件：確定可行解的限制條件。

4. 確定搜索路徑：攤開搜索樹，決定搜索路徑。

5. 確定搜索順序：通常按照某種順序進行搜索，例如，深度優先、廣度優先、最小衝突等。

6. 遞歸搜索：根據搜索路徑和搜索順序進行遞歸搜索，直到找到可行解或者搜索完整個解空間。

7. 回溯：當發現不符合約束條件或無法執行下去時，則回溯到上一個選擇點且檢查下一個選擇。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Backtracking的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1.  全排列問題（Permutations Problem）
給定數字集合，找到所有可能的排列。
答案：https://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string/

2.  組合總數問題（Combination Sum Problem）
給定一組數字集合和一個目標數字，找到集合中所有相加等於目標數字的組合。
答案：https://leetcode.com/problems/combination-sum/

3.  單詞搜索問題（Word Search Problem）
給定一個矩陣和一個單詞，找到該單詞是否存在於矩陣中。單詞可以是垂直或水平排列的。
答案：https://leetcode.com/problems/word-search/

4.  n皇后問題（N-Queens Problem）
給定棋盤大小n和皇后的數量，找到皇后可以互相攻擊的所有可能排列位置。
答案：https://leetcode.com/problems/n-queens/

5.  數獨問題（Sudoku Problem）
給定一個9x9的空數獨遊戲板，找到解決方案。
答案：https://leetcode.com/problems/sudoku-solver/   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Backtracking的網路資料
{{< /ask_chatgpt >}}



1. Backtracking from GeeksforGeeks: 

This article from GeeksforGeeks offers a great introduction to backtracking with examples, pseudo-code, and explanations. The article covers key topics such as the basic backtrack algorithm, various optimization techniques, and how to implement backtracking in specific situations. 

Link: https://www.geeksforgeeks.org/backtracking-algorithms/

2. The Art of Backtracking from TopCoder: 

This article from TopCoder discusses the basics of backtracking, including how it works and how to implement it using recursion. It then goes into more advanced topics, such as pruning techniques, branch and bound algorithms, and how to optimize your code for better performance. 

Link: https://www.topcoder.com/thrive/articles/The%20Art%20of%20Backtracking

3. Backtracking Algorithms from TutorialsPoint: 

This tutorial from TutorialsPoint provides a step-by-step guide to understanding backtracking fundamentals. It covers topics such as how to generate all possible solutions, how to use constraints and rules to eliminate non-viable choices, and how to find the optimal solution. 

Link: https://www.tutorialspoint.com/Backtracking-Algorithm

4. Backtracking Explained from HackerEarth: 

This article from HackerEarth delves into the details of backtracking and explores its relationship to other algorithms, such as dynamic programming and divide and conquer. It also covers common problems that can be solved with backtracking, such as sudoku puzzles and n-queens problems. 

Link: https://www.hackerearth.com/practice/basic-programming/recursion/recursion-and-backtracking/tutorial/

5. Backtracking Algorithm and Its Applications from Educative: 

This article from Educative covers backtracking in-depth, including how to use it to solve common problems such as the knight's tour problem and the coin change problem. It also provides examples of how to optimize your code for better performance and how to use backtracking in real-world applications. 

Link: https://www.educative.io/blog/backtracking-algorithm-and-its-applications   

