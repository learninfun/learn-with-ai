+++
title = "分支限界法"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Branch and Bound的问题
{{< /ask_chatgpt >}}



以下是5个Branch and Bound的问题:

1. 在路径规划中，你需要找到通过一系列城市的最短路径。使用Branch and Bound算法可以帮助你找到最短的路径。

2. 在资源分配问题中，你需要将可用资源分配给多个项目，以最大化全局效益。Branch and Bound算法可以帮助你找到最优解。

3. 在排程问题中，你需要在多个工作之间分配时间，以最小化完成时间。Branch and Bound算法可以帮助你减少总体完成时间。

4. 在0-1背包问题中，你需要填充一个有限大小的背包，以最大化价值。Branch and Bound算法可以帮助你找到最理想的装载方案。

5. 在数据压缩中，你需要找到最小的压缩方式，以减少数据量。Branch and Bound算法可以帮助你找到最佳的压缩方法。   

## 说明知识
{{< ask_chatgpt >}}
说明Branch and Bound并举例
{{< /ask_chatgpt >}}



Branch and Bound（分支定界法）是一种求解最优解问题的算法，通常用于组合优化，如旅行商问题、背包问题等。该算法利用了一个重要原则：通过一系列分支和限制条件的不断缩小搜索空间，同时利用上界和下界对每个搜索分支进行评估和剪枝，以找到最优解。 

具体来说，Branch和Bound的搜索过程分为以下几步：

1. 分支：从初始状态出发，不断将问题拆分为若干个子问题，即各种可能的解决途径。对于背包问题，一个可能的拆分方法是选或不选某个物品。

2. 限制条件：给每个子问题限制条件，从而使得搜索空间逐渐地收敛。对于背包问题，限制条件是背包容量必须小于等于总容量。

3. 上界估计：对每个子问题计算一个最大可行解的上界，用于和问题的实际解进行比较，确定是否需要进一步搜索。对于背包问题，上界可能是已经被选中的物品总价值加上剩余物品中价值最大的物品的价值。

4. 下界估计：对每个子问题计算一个最小可行解的下界，用于对搜索树进行剪枝，即丢弃那些估值低于当前已知最优解的搜索分支。对于背包问题，下界可能是已选择物品的总价值加上还能放入背包的最大价值。

5. 搜索回溯：在搜索分支全部完成后，输出最优解。如果搜索到某个分支的上界小于已知最优值，则剪枝退出。 

举例来说，假设我们要用Branch and Bound算法，解决一个旅行商问题，即从一个城市出发，依次经过其他城市回到原城市。已知每个城市之间的距离，求最短路程。

第一步是分支，我们从任意一个城市出发，可以分别选择下一个要访问的城市。例如，从城市A出发，我们可以选择去B、C、D等城市。

第二步是限制条件，我们需要给每个分支设置约束条件，例如当前已经经过的城市不能再次经过。

第三步是上界估计，对于每个分支，我们需要计算已经选择的路径长度与未选择的城市到起点的距离之和，同时加上一个下界估计。

第四步是下界估计，我们需要计算剩余城市之间的最小距离，假设路径已访问过的城市均为i1，i2， ...，ik，剩下未访问的城市为j1， j2， ...，jm，则下界估计为：

LB = Sum(D i j ) + min(D i1 k + D k j1 , D i2 k + D k j2 , ... , D ik k + D k jm )

其中，D i j 表示城市i到城市j的距离。我们选择未访问城市中到目前为止距离最近的一个城市，加入已知路径。以这种方式对每个分支进行下界估计。

第五步是回溯搜索，在已经遍历整个搜索树后，输出最优的路径及其长度。

总之，Branch and Bound是一种高效的求解组合优化问题的算法。通过不断缩小搜索空间并且对每个搜索分支进行剪枝，算法可以在较短的时间内找到最优解。   

## 汇总重点
{{< ask_chatgpt >}}
条列Branch and Bound的重点
{{< /ask_chatgpt >}}



1. 分支限界法是一种优化问题的求解方法，可以用于求解许多困难的问题，如最小化或最大化目标函数的线性规划问题、TSP问题等。
2. 分支限界法的基本思想是逐步构造一个搜索树，并使用启发函数将可行域切割成多个子域，以加快搜索速度。
3. 在搜索过程中，需要保存每个还未被扩展的节点的状态以及当前的最优解，并根据启发函数选择合适的节点进行扩展。
4. 分支限界法的剪枝策略是在搜索树的每个节点上进行，通常包括择优策略、可行性剪枝和界限剪枝等。
5. 分支限界法的时间复杂度较高，但通常可以得到最优解或近似最优解。为了减少时间消耗，可以使用一些技巧，如启发式搜索、并行搜索等。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Branch and Bound的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 在一個完全連通的無向圖中，找到最小權重生成樹。 

答案: Kruskal algorithm 和 Prim's algorithm 

2. 將$n$個項目分成$m$個集合，使得每個集合的總和盡可能接近一個給定的值$K$。 

答案: Knapsack問題的分支限界解法 

3. 在二元搜索樹中查找第$k$小的元素。 

答案: 計算每個節點的大小並使用分支限制法 

4. 旅行商問題：一個商業旅行者必須訪問$n$個城市，每個城市只能訪問一次，並在所有訪問城市的路徑中找到最短路徑。 

答案: 已知經過點集將事件分解為小問題的方法 

5. 在$0-1$背包中，找出總重量不超過$W$，但總價值最大的物品子集。 

答案: Branch and Bound解法   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Branch and Bound的网络数据
{{< /ask_chatgpt >}}



1. Branch and bound algorithm - GeeksforGeeks 
https://www.geeksforgeeks.org/branch-and-bound-algorithm/

这篇文章介绍了branch and bound algorithm (B&B)的概要和步骤，同时附有python程式码示范。文章中也提到一些B&B解决问题的例子，例如Traveling Salesman Problem (TSP)和Knapsack Problem等等。

2. Branch and Bound Algorithm - Brilliant.org 
https://brilliant.org/wiki/branch-and-bound/

这篇文章介绍了B&B的原理和实作方式。这篇的重点在于说明B&B算法如何运作以及如何应用它来求解最优解。文章中列举了一些范例，包括TSP和Maximum Clique Problem等等。

3. The Branch and Bound Algorithm - Princeton University Press 
https://press.princeton.edu/books/hardcover/9780691122218/the-branch-and-bound-algorithm

这个网站介绍了一本书，书名为The Branch and Bound Algorithm。该书是由Prof. Barnard Korte和Prof. Jens Vygen撰写，并且主要讲授了B&B算法的原理和在解决各种组合优化问题上的应用。

4. Branch and Bound - Tutorialspoint 
https://www.tutorialspoint.com/branch_and_bound_algorithm/index.htm

这篇文章介绍了B&B算法的结构和基本概念，以及如何透过这种算法求解最优解。文章中还提到如何使用B&B算法解决一些组合优化问题，例如Traveling Salesman Problem和Knapsack Problem等问题。

5. Branch and Bound Algorithm - Georgia Tech 
https://www.isye.gatech.edu/sites/default/files/documents/ipy2/Lecture6.pdf

这个网站是由Georgia Tech的师生撰写，介绍了B&B的运作方式，并且同时列举了一些关于BP问题的范例，以及如何使用B&B算法求解。简而言之，这个资源将B&B算法的重点视为如何在数学上解决组合优化问题。   

