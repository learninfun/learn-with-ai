+++
title = "堆"
weight = "3"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Heap的问题
{{< /ask_chatgpt >}}



1. 如何实现Heap的排序？
2. 如何找到Heap中的最大值或最小值？
3. 如何向Heap中添加一个新值？
4. 如何将Heap中的一个元素删除？
5. 如何检查一个数据结构是否符合Heap的性质？   

## 说明知识
{{< ask_chatgpt >}}
说明Heap并举例
{{< /ask_chatgpt >}}



Heap是一种树形资料结构，其中每个节点都具有一个值，并且拥有一个父节点和零个或多个子节点。Heap可以分为最小堆和最大堆两种形式，最小堆中每个节点的值都比其子节点小，而最大堆中每个节点的值都比其子节点大。 

举例来说： 

假设有以下一个最小堆：

         1    
       /    \  
     2       3  
    /  \    /   \   
   4    5  6     7 

其中最小的元素是1，其它元素的值都比1大，而左子树的元素都比右子树的元素小，这就符合了最小堆的特性。

堆可以用来实现优先级队列，通常使用阵列来实现。在堆中，基本操作包括插入新元素、删除最小（或最大）元素、搜寻最小（或最大）元素等。堆排序也是一种使用堆的排序算法。   

## 汇总重点
{{< ask_chatgpt >}}
条列Heap的重点
{{< /ask_chatgpt >}}



1. Heap是一種樹狀結構，通常用於實現優先級隊列。
2. Heap通常被實現為二叉堆，其中每個節點的值小於或大於其子節點的值（稱為小根堆或大根堆）。
3. Heap支持插入新元素和將最小或最大元素彈出的操作，這些操作的時間複雜度都是O(log n)。
4. Heap可以通過將一個數字附加到堆底並再次構建堆來實現排序。
5. 小根堆（或大根堆）可以用於Dijkstra算法和Kruskal算法等算法中，它們都需要尋找最小（或最大）權重邊或最小（或最大）路徑。
6. Heap還可以用於動態編程和排序算法，如堆排序和快速排序。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Heap的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. LeetCode 215. Kth Largest Element in an Array
2. LeetCode 502. IPO
3. LeetCode 973. K Closest Points to Origin
4. LeetCode 347. Top K Frequent Elements
5. LeetCode 295. Find Median from Data Stream

答案不提供，希望试着自己解决问题。   

