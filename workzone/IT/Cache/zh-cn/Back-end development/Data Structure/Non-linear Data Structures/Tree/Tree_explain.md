

Tree是一种常见的资料结构，它由许多组成称为节点，节点之间可以有连接称为边的二元关系，因此Tree也称为二元树。Tree以分层的方式呈现，每一层节点与它上层的节点之间，都有唯一的连接，而与同一层的其他节点之间则没有连接。树状结构的节点具有父节点和子节点之分，根节点位于顶层，没有父节点，而叶节点则没有子节点。

树的节点上可以附加任意的资料，一棵树可以用这些附加的资料来表达一个超级多的信息，例如家族树，以个人为节点，附加的资料可以是姓名，性别，年龄，电话等信息，通过树的分层结构，实现每个人之间的关系。

下面举一个二元树的例子：

     1
    / \
   2   3
  / \   \
 4   5   6

上面的示例是一个深度为3的树，含有6个节点和5个边，其中节点1是根节点，它有两个子节点2和3，节点2又有两个子节点4和5，节点3有一个子节点6。节点1到节点2之间有一条边，节点2到节点4之间有一条边，节点2到节点5之间有一条边，节点1到节点3之间有一条边，节点3到节点6之间有一条边。

透过这些节点之间的连接，可以实现很多不同的应用，例如档案系统、编译器、协议分析等。