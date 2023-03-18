+++
title = "Dijkstra演算法"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Dijkstra's Algorithm的问题
{{< /ask_chatgpt >}}



1. 在以下图形中，使用Dijkstra算法找到A到G的最短路径。
   ![Dijkstra Problem 1](https://i.imgur.com/SSRzWFL.png)
   
2. 假设您有以下路线和距离表，使用Dijkstra's算法找到从A到D的最短路径，可以将其路径和距离列出来吗？

|  路线  | 距离 |
|:------:|:----:|
| A - B  |  6   |
| A - C  |  3   |
| B - C  |  2   |
| B - D  |  2   |
| C - D  |  4   |
| C - E  |  5   |
| D - E  |  3   |

3. 在以下的图形中，使用Dijkstra's算法找到从A到J的最短路径。 如果有两条路径大小相同，您可以在回答中任选一条。 
   ![Dijkstra Problem 3](https://i.imgur.com/Q7V51bf.png)
   
4. 在以下图形中，使用Dijkstra's算法找到从A到G的最短路径。 如果有两条路径大小相同，您可以在回答中任选一条。 
   ![Dijkstra Problem 4](https://i.imgur.com/Gkp5XIk.png)
   
5. 在以下路线图中，使用Dijkstra's算法找到从A到F的最短路径。 如果有多个同样大小的路径，您可以在回答中任选一条。 
   ![Dijkstra Problem 5](https://i.imgur.com/dNR5Wav.png)   

## 说明知识
{{< ask_chatgpt >}}
说明Dijkstra's Algorithm并举例
{{< /ask_chatgpt >}}



Dijkstra算法是一种用于找到带有非负权重边的最短路径的算法。它根据优先级队列不断选择距离起点最近的节点，每次选中的节点都会类似贪心算法一样地更新起点到其他节点的距离。

以下是Dijkstra算法的步骤：

1.初始化：

   - 将起点的距离设为0，其他点的距离设为无限大

   - 将起点加入优先级队列

2. 开始遍历队列：

   - 从优先级队列中选择距离起点最近的节点

   - 对选择的节点进行松弛操作：即找到与这个节点相邻的边，并更新和其他节点的距离（如果新的路径比现有的路径更短）。

   - 将松弛过节点加入优先级队列

3. 重复步骤2，直到队列中没有节点。

以下是一个简单的例子，我们将在下图中找到起点到其他节点的最短路径

![Dijkstra_Example_Graph](https://i.imgur.com/fPofr0D.png)

1. 将起点A的距离设为0，其他的设为无限大。A加入优先级队列。

2. 从优先级队列中选择起点A，对A进行松弛操作。现在B、C、D、E的距离分别为5、3、2和6，B、C和D都加入优先级队列。

3. 从优先级队列中选择距离A最近的C节点，对C进行松弛操作。现在B、D、E的距离分别为7、2和6，D加入优先级队列。

4. 从优先级队列中选择距离A最近的D节点，对D进行松弛操作。现在B、E的距离分别为7和8，B加入优先级队列。

5. 从优先级队列中选择距离A最近的B节点，对B进行松弛操作。现在E的距离为9，E加入优先级队列。

6. 从优先级队列中选择距离A最近的E节点，对E进行松弛操作。现在没有新节点加入优先级队列。

7. 遍历已经没有节点加入优先级队列，算法结束。下图显示了每个节点到起点的最短距离结果。

![Dijkstra_Result](https://i.imgur.com/d9X4pVy.png)   

## 汇总重点
{{< ask_chatgpt >}}
条列Dijkstra's Algorithm的重点
{{< /ask_chatgpt >}}



以下是Dijkstra's Algorithm的重點：

1. 設定一個起點和終點來解決兩點之間的最短路徑。
2. 創建一個堆棧來儲存已訪問的節點和待訪問的節點。
3. 初始化起點節點。
4. 對於每一個節點，設置一個到達該節點的“最短距離”和“最短路徑”，然後添加到堆棧中。
5. 選擇到目前為止最短的節點，設置其為當前節點，從堆棧中移除。
6. 檢查該節點的相鄰節點，如果通過當前節點到達相鄰節點的距離短於之前所保存的“最短距離”，則更新該節點的“最短距離”和“最短路徑”。
7. 重複步驟5-6，直到目標節點被訪問或者堆棧被清空。

Dijkstra's Algorithm是一種解決圖形問題的有效方法，它可以找到兩點之間的最短路徑。該算法基於“貪心”的策略，通過不斷地選擇當前最短的路徑來搜尋該圖形，直到所有節點都被訪問過。該算法適用於沒有負權重圖的情況，並且可以通過最小堆棧實現高效率訪問。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Dijkstra's Algorithm的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 给定一个有向带权图，请问从起点到终点的最短路径总权值是多少？

起点：A，终点：E

|节点|A|B|C|D|E|
|---|---|---|---|---|---|
|A|0|4|1|0|0|
|B|0|0|0|5|0|
|C|0|2|0|0|4|
|D|0|0|0|0|1|
|E|0|0|0|0|0|

答案：最短路径总权值为3（A-C-E）。

2. 给定一个有向带权图，请问从起点到各个节点的最短路径总权值分别是多少？

起点：A

|节点|A|B|C|D|E|
|---|---|---|---|---|---|
|A|0|4|1|0|0|
|B|0|0|0|5|0|
|C|0|2|0|0|4|
|D|0|0|0|0|1|
|E|0|0|0|0|0|

答案：  
- 从起点A到B的最短路径总权值为4  
- 从起点A到C的最短路径总权值为1  
- 从起点A到D的最短路径总权值为无法到达  
- 从起点A到E的最短路径总权值为无法到达  

3. 给定一个加权无向图，请问从起点到终点的最短路径总权值是多少？

起点：A，终点：F

|节点|A|B|C|D|E|F|
|---|---|---|---|---|---|---|
|A|0|2|4|1|0|0|
|B|2|0|0|3|10|0|
|C|4|0|0|2|0|0|
|D|1|3|2|0|7|4|
|E|0|10|0|7|0|5|
|F|0|0|0|4|5|0|

答案：最短路径总权值为8（A-D-F）。

4. 给定一个加权有向图，请问从起点到终点的最短路径总权值是多少？

起点：A，终点：D

|节点|A|B|C|D|
|---|---|---|---|---|
|A|0|10|0|5|
|B|0|0|1|2|
|C|0|0|0|4|
|D|0|0|0|0|

答案：最短路径总权值为7（A-D）。

5. 给定一个加权有向图，请问从起点到终点的最短路径总权值是多少？

起点：A，终点：D

|节点|A|B|C|D|
|---|---|---|---|---|
|A|0|10|0|5|
|B|0|0|1|2|
|C|1|0|0|4|
|D|0|0|0|0|

答案：最短路径总权值为5（A-C-D）。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Dijkstra's Algorithm的网络数据
{{< /ask_chatgpt >}}



以下为5篇关于Dijkstra's Algorithm的网路资料：

1. Dijkstra’s Algorithm for Shortest Path: Explained (by Career Karma)
https://careerkarma.com/blog/dijkstras-algorithm-for-shortest-path-explained/

2. Dijkstra's Algorithm for finding Shortest Path with Example (by GeeksforGeeks)
https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-greedy-algo-7/

3. Dijkstra's algorithm (by Wikipedia)
https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm

4. Dijkstra Algorithm – Single Source Shortest Path (by Studytonight)
https://www.studytonight.com/data-structures/dijkstras-algorithm

5. Dijkstra’s Algorithm in Python (by Codecademy)
https://www.codecademy.com/learn/dijkstra-python   

