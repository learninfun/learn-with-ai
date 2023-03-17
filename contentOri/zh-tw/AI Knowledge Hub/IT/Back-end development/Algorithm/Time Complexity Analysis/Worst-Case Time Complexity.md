## 習題預習
{{< ask_chatgpt >}}
給我5題Worst-Case Time Complexity的問題
{{< /ask_chatgpt >}}



1. 一個未排序的陣列中找到最大的元素的時間複雜度是多少？
2. 在一個含有n個元素的已排序陣列中，找到一個元素的時間複雜度是多少？ 
3. 在一個樹形結構中，找到最深的葉子節點的時間複雜度是多少？
4. 在一個包含n個元素的無序二元搜尋樹中，找到最大或最小的元素的時間複雜度是多少？ 
5. 在一個n x n的矩陣中，找到最小值或最大值的時間複雜度是多少？   

## 說明知識
{{< ask_chatgpt >}}
說明Worst-Case Time Complexity並舉例
{{< /ask_chatgpt >}}



Worst-Case Time Complexity指的是在最劣情況下程序的時間複雜度。即在所有可能的輸入中，程序需要最長的時間才能完成。

舉例來說，如果我們要對一個數組進行排序，並且該數組已經是有序的，那麼一個時間複雜度較差的排序算法可能需要O(n^2)的時間才能完成排序。這是因為在有序數組中，該算法仍然需要比較和交換元素，但實際上這些操作是完全不必要的，因為數組已經排序好了。因此，我們可以說在這種情況下，該排序算法的最劣時間複雜度為O(n^2)。

另一個例子是搜索一個數組中是否存在一個元素。如果該元素在數組的開頭，則算法可以在常數時間內完成。但如果該元素位於數組的末尾，那麼該算法可能需要遍歷整個數組才能找到該元素。在這種情況下，該算法的最劣時間複雜度為O(n)。   

## 彙總重點
{{< ask_chatgpt >}}
條列Worst-Case Time Complexity的重點
{{< /ask_chatgpt >}}



Worst-Case Time Complexity是衡量算法效率的重要指標，指的是當算法面對最差情況時所需執行的最長時間。以下是Worst-Case Time Complexity的重點：

1. 最壞情況：Worst-Case Time Complexity是從理論上判斷算法效率的指標，它要考慮所有可能的輸入，並選取對算法耗時最長的那一種情況衡量效率。

2. 常見符號：通常用大O符號來表示Worst-Case Time Complexity，例如O(n)、O(logn)、O(n^2)等，它們分別表示當輸入規模為n時最壞情況下算法需要執行的時間；

3. 算法效率差距：同一個問題可以有不同的算法，它們的Worst-Case Time Complexity也可能截然不同，因此在選擇算法時必須特別注意該算法的時間複雜度；

4. 適用範圍：Worst-Case Time Complexity可以幫助我們找到最高效的算法，但也需要注意實際應用中的各種限制因素，例如硬體性能等；

5. 進一步評估算法表現：在實際情況下，算法的表現往往不僅與Worst-Case Time Complexity有關，還需要考慮其他因素，例如最好情況、平均情況等等。因此，對於一個算法的評估還需要綜合各種情況進行評估。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Worst-Case Time Complexity的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. Bubble Sort (冒泡排序) - Worst-Case Time Complexity: O(n^2)

2. Insertion Sort (插入排序) - Worst-Case Time Complexity: O(n^2)

3. Selection Sort (選擇排序) - Worst-Case Time Complexity: O(n^2)

4. Naive String Searching (字符串匹配搜索) - Worst-Case Time Complexity: O(m(n-m+1)), 其中m為模式串的長度，n為文本串的長度

5. Quick Sort (快速排序) - Worst-Case Time Complexity: O(n^2)

答案：

1. Bubble Sort: 按照從小到大的順序排序數組[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

時間複雜度為n^2，其中n為數組長度，這裡為10，最壞情況需要比較45次，移動45次

2. Insertion Sort: 按照從小到大的順序排序數組[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

時間複雜度為n^2，其中n為數組長度，這裡為10，最壞情況需要比較45次，移動45次

3. Selection Sort: 按照從小到大的順序排序數組[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

時間複雜度為n^2，其中n為數組長度，這裡為10，最壞情況需要比較45次，移動45次

4. Naive String Searching: 在字符串"ABABABABAB"中搜索模式串"ABAB", "ABAB"每個字符都匹配，但每次匹配都失配，時間複雜度為4*(10-4+1)=28

5. Quick Sort: 按照從小到大的順序排序數組[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

序列每次選擇第一個元素為樞軸，而序列本身是逆序的，所以每次劃分都到了最差情況，時間複雜度為n^2，其中n為數組長度，這裡為10，最壞情況需要比較45次，交換45次   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Worst-Case Time Complexity的網路資料
{{< /ask_chatgpt >}}



1. Worst-Case Time Complexity of Sorting Algorithms
https://www.geeksforgeeks.org/analysis-of-algorithms-set-2-asymptotic-analysis/
This article provides a detailed explanation of worst-case time complexity of various sorting algorithms such as Bubble Sort, Selection Sort, Insertion Sort, Merge Sort, Quick Sort, and Heap Sort. It also includes their best and average case complexities as well as space complexities.

2. Worst-Case Time Complexity of Graph Algorithms
https://www.tutorialspoint.com/data_structures_algorithms/worst_case_complexity_of_graph_algorithms.htm
This article discusses the worst-case time complexity of graph algorithms such as Breadth First Search (BFS), Depth First Search (DFS), Dijkstra's Algorithm, and Bellman-Ford Algorithm. It includes pseudocode for each algorithm and their complexity analysis.

3. Worst-Case Time Complexity of Dynamic Programming Algorithms
https://www.geeksforgeeks.org/worst-case-time-complexity-of-dynamic-programming-algorithms/
This article provides an overview of dynamic programming and its worst-case time complexity for various problems such as Fibonacci Series, Longest Common Subsequence, Knapsack Problem, and Rod Cutting Problem. It also includes code examples and complexity analysis for each problem.

4. Worst-Case Time Complexity of String Matching Algorithms
https://www.geeksforgeeks.org/worst-case-time-complexity-of-string-matching-algorithms/
This article discusses the worst-case time complexity of various string matching algorithms such as Naive Algorithm, Rabin-Karp Algorithm, Finite Automata Algorithm, and Knuth-Morris-Pratt Algorithm. It includes code examples and complexity analysis for each algorithm.

5. Worst-Case Time Complexity of Tree Algorithms
https://www.tutorialspoint.com/data_structures_algorithms/worst_case_complexity_of_tree_algorithms.htm
This article explains the worst-case time complexity of tree algorithms such as Binary Search Tree (BST), AVL Tree, B-Tree, and Red-Black Tree. It includes pseudocode and complexity analysis for each algorithm.   

