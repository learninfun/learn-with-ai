+++
title = "单向链表"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Singly Linked List的问题
{{< /ask_chatgpt >}}



1. 如何在单链式列表的开始处添加一个元素？
2. 如何查找单链式列表的最后一个节点？
3. 如何删除单链式列表中的特定节点？
4. 如何将单链式列表逆转？
5. 如何在单链式列表中查找特定值的节点？   

## 说明知识
{{< ask_chatgpt >}}
说明Singly Linked List并举例
{{< /ask_chatgpt >}}



单向链结列表（Singly Linked List）是一种资料结构，由一连串节点组成，每个节点包含了数据和一个指向下一个节点的指标。第一个节点被称为头节点，最后一个节点则为尾节点，其下一个指针指向 NULL。

Singly Linked List 能够高效地遍历和添加/删除头部节点，但较难操作中间的节点。

以下是一个 Singly Linked List 的例子：

```
   Head    Node1   Node2    Node3    Node4     Tail
    ↓       ↓        ↓        ↓        ↓         ↓
┌───────┐┌───────┐┌───────┐┌───────┐┌─────────┐
│ data  ││ data  ││ data  ││ data  ││ data    │
├───────┤├───────┤├───────┤├───────┤├─────────┤
│  100  ││  200  ││  300  ││  400  ││  500    │
├───────┤├───────┤├───────┤├───────┤├─────────┤
│ next  ││ next  ││ next  ││ next  ││  next   │
├───────┤├───────┤├───────┤├───────┤├─────────┤
│ Node1 ││ Node2 ││ Node3 ││ Node4 ││  NULL   │
└───────┘└───────┘└───────┘└───────┘└─────────┘
```

在这个例子中，我们有一个包含 5 个节点的 Singly Linked List。每个节点都包含一个数据部分和一个指向下一个节点的指针。Head 节点包含数据 100，指针指向 Node1 节点。Node1 节点包含数据 200，指针指向 Node2 节点。以此类推直到 Tail 节点包含数据 500，指针指向 NULL，代表没有更多节点。   

## 汇总重点
{{< ask_chatgpt >}}
条列Singly Linked List的重点
{{< /ask_chatgpt >}}



1. 它是一种线性资料结构，由“节点”构成，每个节点至少包含一个资料元素和一个指向下一节点的指针。
2. 链表的头节点是唯一的，如果头节点为空，则链表是空的。
3. 链表的最后一个节点的指针指向空，表示这是链表的结束。
4. 插入和删除节点是链表的主要操作，通过调整指针来实现。
5. 链表的优点是可以动态分配内存空间，不需要预先指定大小，并且在插入和删除元素时不需要移动整个资料结构，效率较高。
6. 链表的不足之处是无法像数组那样根据下标来访问元素，必须从头节点开始遍历，效率较低。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Singly Linked List的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 实作一个反转Singly Linked List的函数
   答案：https://stackoverflow.com/questions/20662024/reverse-a-linked-list
2. 判断一个Singly Linked List是否存在环路
   答案：https://www.geeksforgeeks.org/detect-loop-in-a-linked-list/
3. 找到Singly Linked List中倒数第n个节点
   答案：https://www.geeksforgeeks.org/nth-node-from-the-end-of-a-linked-list/
4. 合并两个已排序的Singly Linked List成为一个新的已排序LinkedList
   答案：https://www.geeksforgeeks.org/merge-two-sorted-linked-lists/
5. 将Singly Linked List按照特定数值x当中的节点分为两部分，小于x的节点在前、大于等于x的节点在后
   答案：https://www.geeksforgeeks.org/partitioning-a-linked-list-around-a-given-value-and-keeping-the-original-order/   

