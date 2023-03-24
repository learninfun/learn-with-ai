1. 定义：Topological Sort 是一种对有向无环图（DAG）进行节点排序的演算法。

2. 应用：Topological Sort 常用于寻找项目间的先后关系，例如工作流程中的先后顺序、课程选修的先后等。

3. 步骤：Topological Sort 的步骤为寻找 DAG 的一个非循环路径，并依照该路径的顺序将节点排序。

4. 实现方法：Topological Sort 可以使用 DFS 或 BFS 两种方式来实现，其中 BFS 的时间复杂度较低。

5. 结果：Topological Sort 的结果并不唯一，可能存在多种排序结果。

6. 特殊情况：如果 DAG 中存在环路则无法进行 Topological Sort，这时需要进行环路检测或者使用其他方法进行排序。

7. 应用范例：如下图所示，该 DAG 表示六个项目之间的先后关系，使用 Topological Sort 可以得到的一个排序结果为 B, D, A, C, F, E。

   ![image.png](https://i.imgur.com/XnoZwqa.png)