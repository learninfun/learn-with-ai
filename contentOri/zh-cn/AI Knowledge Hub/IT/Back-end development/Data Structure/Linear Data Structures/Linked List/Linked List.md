+++
title = "链表"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Linked List的问题
{{< /ask_chatgpt >}}



1. 如何反轉一個單向鏈表？
2. 如何找到一個單向鏈表的中間節點？
3. 如何合併兩個已排序的單向鏈表？
4. 如何檢測一個單向鏈表是否存在環？
5. 如何刪除一個單向鏈表中指定的節點？   

## 说明知识
{{< ask_chatgpt >}}
说明Linked List并举例
{{< /ask_chatgpt >}}



Linked List是一種資料結構，其特點是每個節點(node)包含一個值(value)以及指向下一個節點的指標(pointer)，形成了一個"鏈"的結構，其中第一個節點稱為頭節點(head)，最後一個節點指向null。

Linked List的優點在於可以有效地處理添加、刪除節點等操作，並且可以動態調整大小。

以下是一個單向鏈表的範例：

```
class Node {
    int value;
    Node next;
  
    public Node(int value) {
        this.value = value;
        next = null;
    }
}

class LinkedList {
    Node head;
  
    public LinkedList() {
        head = null;
    }
  
    public void add(int value) {
        Node node = new Node(value);
        if (head == null) {
            head = node;
        } else {
            Node current = head;
            while (current.next != null) {
                current = current.next;
            }
            current.next = node;
        }
    }
  
    public void delete(int value) {
        if (head == null) {
            return; 
        }
        if (head.value == value) {
            head = head.next;
        } else {
            Node current = head;
            while (current.next != null && current.next.value != value) {
                current = current.next;
            }
            if (current.next != null) {
                current.next = current.next.next;
            }
        }
    }
}
```

這個LinkedList類包含了添加和刪除節點的方法。要添加節點，可以使用add方法。要刪除節點，可以使用delete方法。當然，還有很多其他方法可以實現在Linked List上的操作，這裡只是顯示了一個基本的範例。   

## 汇总重点
{{< ask_chatgpt >}}
条列Linked List的重点
{{< /ask_chatgpt >}}



1. 链接点：Linked list由一系列节点（node）构成，每个节点包含了一个数值（value）和一个指向下一个节点的指针（pointer）。

2. 头指针：Linked list中的第一个节点叫做头节点（head），头指针（head pointer）标志着第一个节点。

3. 尾指针：Linked list中的最后一个节点叫做尾节点（tail），通常会使用一个特别的值（例如null或None）作为尾指针（tail pointer）。

4. 插入：向Linked list中插入一个新节点时，需要修改前一个节点的指针，让它指向新节点，同时让新节点指向原来的后一个节点。

5. 删除：从Linked list中删除一个节点时，需要修改前一个节点的指针，让它指向后一个节点。

6. 查找：遍历整个Linked list，逐个比较节点中的数值，查找特定的节点。

7. 修改：寻找Linked list中特定的节点，然后修改它的数值。

8. 链表的类型：单向链表、双向链表、循环链表。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Linked List的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. Reverse Linked List II
2. Copy List with Random Pointer
3. Remove Nth Node From End of List
4. Partition List
5. Linked List Cycle II

1. Reverse Linked List II:
题目描述： 反转从位置 m 到 n 的链表。请使用一趟扫瞄完成反转。
示例: 
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
答案链接: https://leetcode.com/problems/reverse-linked-list-ii/

2. Copy List with Random Pointer:
题目描述： 给定一个链表，每个节点包含一个额外增加的随机指针，
该指针可以指向链表中的任何节点或空节点。
要求返回这个链表的 深拷贝。 
示例：
输入：
{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
解释：
节点 1 的值是 1，它的下一个指针和随机指针都指向节点 2 。
节点 2 的值是 2，它的下一个指针指向 null，随机指针指向它本身。
答案链接：https://leetcode.com/problems/copy-list-with-random-pointer/

3. Remove Nth Node From End of List:
题目描述： 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。 
示例：
输入: 1->2->3->4->5, n = 2
输出: 1->2->3->5
答案链接： https://leetcode.com/problems/remove-nth-node-from-end-of-list/

4. Partition List:
题目描述： 给定一个链表和一个特定值 x，对链表进行分隔，
使得所有小于 x 的节点都在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。
示例：
输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
答案链接： https://leetcode.com/problems/partition-list/

5. Linked List Cycle II:
题目描述： 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
示例：
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
答案链接： https://leetcode.com/problems/linked-list-cycle-ii/   

