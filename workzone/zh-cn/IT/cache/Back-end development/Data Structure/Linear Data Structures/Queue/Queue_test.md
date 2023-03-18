

1. 特定顺序遍历二叉树（Medium）
给定一个二叉树和一个正整数k，按照以下步骤打印出所有节点：
1. 首先打印以根节点为开始的前k个节点。
2. 接着打印所有由前k个节点的子节点分支出去的节点。
3. 重复步骤2，直到没有更多节点可以打印。
例如，如果k为2，下图中二叉树的节点遍历顺序为：1，2，3，4，5，6。

答案: https://leetcode.com/problems/print-binary-tree-in-order-of-levels/



2. 循环链表的环（Medium）
给定一个循环链表，查找并返回其中的环的起始节点。如果不存在环，则返回null。
例如，下图中的循环链表的环起始节点为3。

答案: https://leetcode.com/problems/linked-list-cycle-ii/


3. 正方形的填充（Medium）
给定一个大小为N×N的矩阵和一个起始点，从起始点开始按照特定规则将矩阵中的所有位置填充为相同的值。该规则是：将矩阵中每个与起始点相邻的位置的值改为起始点的值，直到所有相邻的位置都已被填充。
例如，下图中5×5矩阵的起始点为(1,1)，填充后的矩阵如图所示。

答案 : https://www.codewars.com/kata/rectangle-fill/train/python



4. 块状矩阵（Medium）
给定一个大小为N×N的矩阵和一个坐标(x,y)和一个大小为k的正方形区域。将指定区域中的每个元素加上指定值p，并返回更新后的矩阵。
例如，下图中4×4矩阵的(x,y)坐标为(2,2)，k大小为2，p值为3，更新后的矩阵如图所示。

答案: https://leetcode.com/problems/matrix-block-sum/


5. 充电器安排（Medium）
给定一个由n个非负整数表示的阵列，表示在不同的位置上有一系列充电器。每个充电器都有一定的充电范围，可以在指定的位置上充电器。请问至少需要安排多少个新的充电器，才能够在能达到所有充电需求的前提下，最小化充电器的使用量。
例如，下图中线段表示充电器的充电范围，黑点表示现有充电器的位置。最少需要安排2个新的充电器，位置分别在[1, 5] 和 [8, 12]，才能够在达到所有充电需求的前提下，最小化充电器的使用量。

答案: https://leetcode.com/problems/minimum-number-of-refueling-stops/