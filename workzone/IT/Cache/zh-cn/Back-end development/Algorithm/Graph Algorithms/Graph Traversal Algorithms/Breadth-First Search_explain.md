

Breadth-First Search（BFS）是一种搜寻演算法，以广度优先的方式来遍历一个图形或树状结构。

其遍历的顺序是逐层往下，也就是先遍历所有的同一深度节点，再遍历下一深度的节点。在BFS遍历中，使用一个队列来维护已经被广度遍历的节点，以便按层访问下一阶段节点时使用。

举个例子：假设我们有一个有向图如下图所示：

<img src="https://i.imgur.com/Vhh9XfO.png" width="200"/>

我们从节点1开始进行BFS遍历，首先将节点1加入队列中。接下来，按照节点编号的大小顺序，先遍历节点2和节点3。

<img src="https://i.imgur.com/s8BDGdc.png" width="200"/>

然后，把节点2的相邻节点4，7加入队列中，把节点3的相邻节点5，6加入队列中。

<img src="https://i.imgur.com/gPXs4L4.png" width="200"/>

再遍历节点4和节点7，因为它们没有相邻节点可加入队列中，所以直接跳过。

<img src="https://i.imgur.com/yujfiFB.png" width="200"/>

最后，遍历节点5和节点6，发现节点5有一个相邻节点8，所以把节点8加入队列中。遍历完节点5和节点6，队列已经空了，此时遍历结束。

<img src="https://i.imgur.com/iCCvSTN.png" width="200"/>

这个例子中，我们首先访问距离起点节点最近的节点，然后是次近的节点，然后是更远的节点。这种BFS的访问顺序可以确保我们在最短时间内访问到所有节点。