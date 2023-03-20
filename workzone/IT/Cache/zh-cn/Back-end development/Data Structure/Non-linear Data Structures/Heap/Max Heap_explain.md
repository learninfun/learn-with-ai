

Max Heap是一种二元树 (Binary Tree) 的数据结构，其中每个节点的值都大于其子树中的节点值。也就是说，树的根节点必须是树中所有节点的最大值。在Max Heap中，对于任意的节点 i，其左子节点为 2i，右子节点为 2i+1。

以下是一个Max Heap的例子：

     70
    / \
   50  60
  / \   \
 30  40  20

在这个例子中，根节点为 70，其左子节点为 50，右子节点为 60。左子节点 50 的左右子节点分别为 30 和 40，右子节点 60 只有一个右子节点 20。

Max Heap通常用于实现堆排列 (Heap Sort)、优先伫列 (Priority Queue) 等数据运算中，也可以用于找到最小 K 个数中的最大值。