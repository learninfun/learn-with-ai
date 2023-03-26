+++
title = "單向鏈表"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Singly Linked List的中文問題
{{< /ask_chatgpt >}}



1. 如何在單鏈式列表的開始處添加一個元素？
2. 如何查找單鏈式列表的最後一個節點？
3. 如何刪除單鏈式列表中的特定節點？
4. 如何將單鏈式列表逆轉？
5. 如何在單鏈式列表中查找特定值的節點？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Singly Linked List並舉例
{{< /ask_chatgpt >}}



單向鏈結列表（Singly Linked List）是一種資料結構，由一連串節點組成，每個節點包含了數據和一個指向下一個節點的指標。第一個節點被稱為頭節點，最後一個節點則為尾節點，其下一個指針指向 NULL。

Singly Linked List 能夠高效地遍歷和添加/刪除頭部節點，但較難操作中間的節點。

以下是一個 Singly Linked List 的例子：

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

在這個例子中，我們有一個包含 5 個節點的 Singly Linked List。每個節點都包含一個數據部分和一個指向下一個節點的指針。Head 節點包含數據 100，指針指向 Node1 節點。Node1 節點包含數據 200，指針指向 Node2 節點。以此類推直到 Tail 節點包含數據 500，指針指向 NULL，代表沒有更多節點。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Singly Linked List的重點
{{< /ask_chatgpt >}}



1. 它是一種線性資料結構，由「節點」構成，每個節點至少包含一個資料元素和一個指向下一節點的指針。
2. 鏈表的頭節點是唯一的，如果頭節點為空，則鏈表是空的。
3. 鏈表的最後一個節點的指針指向空，表示這是鏈表的結束。
4. 插入和刪除節點是鏈表的主要操作，通過調整指針來實現。
5. 鏈表的優點是可以動態分配內存空間，不需要預先指定大小，並且在插入和刪除元素時不需要移動整個資料結構，效率較高。
6. 鏈表的不足之處是無法像數組那樣根據下標來訪問元素，必須從頭節點開始遍歷，效率較低。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Singly Linked List的中等難度問題，並在後面列出答案
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

