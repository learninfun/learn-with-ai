+++
title = "单向链表"
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



1. 實作一個反轉Singly Linked List的函數
   答案：https://stackoverflow.com/questions/20662024/reverse-a-linked-list
2. 判斷一個Singly Linked List是否存在環路
   答案：https://www.geeksforgeeks.org/detect-loop-in-a-linked-list/
3. 找到Singly Linked List中倒數第n個節點
   答案：https://www.geeksforgeeks.org/nth-node-from-the-end-of-a-linked-list/
4. 合併兩個已排序的Singly Linked List成為一個新的已排序LinkedList
   答案：https://www.geeksforgeeks.org/merge-two-sorted-linked-lists/
5. 將Singly Linked List按照特定數值x當中的節點分為兩部分，小於x的節點在前、大於等於x的節點在後
   答案：https://www.geeksforgeeks.org/partitioning-a-linked-list-around-a-given-value-and-keeping-the-original-order/   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Singly Linked List的网络数据
{{< /ask_chatgpt >}}



1. Singly Linked List - GeeksforGeeks

GeeksforGeeks 是一個專注於程式設計和計算機科學相關主題的網站，其關於 Singly Linked List 的文章非常詳細。這篇文章涵蓋了 Singly Linked List 概念、運作方式、插入、刪除以及反轉等操作。另外還提供了適當的程式碼範例，讓讀者能夠更深入了解 Singly Linked List。

連結：https://www.geeksforgeeks.org/data-structures/linked-list/singly-linked-list/

2. Singly Linked List - Wikipedia

Wikipedia 是一個由群眾貢獻而成的百科全書，其關於 Singly Linked List 的文章介紹了該資料結構的原理、結構以及實際應用。該文章有許多連結與參考資料，讓讀者能夠深入了解 Singly Linked List 相關的其他概念與技術。

連結：https://en.wikipedia.org/wiki/Linked_list

3. What is a Singly Linked List? - Study.com

Study.com 是一個在教育領域非常知名的網站，其關於 Singly Linked List 的文章寫得非常清晰易懂。該文章從需求開始，詳細介紹了 Singly Linked List 以及其常用的操作。此外，該文章還有一些適當的示範和測試問題，讓讀者能夠練習運用所學知識。

連結：https://study.com/academy/lesson/what-is-a-singly-linked-list.html

4. Singly Linked List in JavaScript - TutorialsPoint

TutorialsPoint 是一個專注於程式設計和資訊技術相關主題的網站，其關於 Singly Linked List 的文章是一個非常好的學習資源。該文章使用 JavaScript 程式示範了 Singly Linked List 相關的操作，如插入、刪除、反轉等。範例程式碼非常清晰易懂，適合 JavaScript 初學者閱讀練習。

連結：https://www.tutorialspoint.com/data_structures_algorithms/linked_list_program_in_javascript.htm

5. Different Types of Linked List - Includehelp

Includehelp 是一個關於程序設計和計算機科學的學習平台，其關於不同類型的 Link List 的文章值得一提。該文章介紹了 Singly Linked List 及其相關的其他類型，如 Doubly Linked List 和 Circular Linked List。此外，該文章還有適當的計算機代碼示範，讓讀者能夠更好地理解不同類型的 Link List 的運作方式。

連結：https://www.includehelp.com/ds/different-types-of-linked-list.aspx   

