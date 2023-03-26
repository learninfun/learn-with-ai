+++
title = "鏈表"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Linked List的中文問題
{{< /ask_chatgpt >}}



1. 如何反轉一個單向鏈表？
2. 如何找到一個單向鏈表的中間節點？
3. 如何合併兩個已排序的單向鏈表？
4. 如何檢測一個單向鏈表是否存在環？
5. 如何刪除一個單向鏈表中指定的節點？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Linked List並舉例
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

## 條列重點
{{< ask_chatgpt >}}
以中文條列Linked List的重點
{{< /ask_chatgpt >}}



1. 鏈接點：Linked list由一系列節點（node）構成，每個節點包含了一個數值（value）和一個指向下一個節點的指針（pointer）。

2. 頭指針：Linked list中的第一個節點叫做頭節點（head），頭指針（head pointer）標誌著第一個節點。

3. 尾指針：Linked list中的最後一個節點叫做尾節點（tail），通常會使用一個特別的值（例如null或None）作為尾指針（tail pointer）。

4. 插入：向Linked list中插入一個新節點時，需要修改前一個節點的指針，讓它指向新節點，同時讓新節點指向原來的後一個節點。

5. 刪除：從Linked list中刪除一個節點時，需要修改前一個節點的指針，讓它指向後一個節點。

6. 查找：遍歷整個Linked list，逐個比較節點中的數值，查找特定的節點。

7. 修改：尋找Linked list中特定的節點，然後修改它的數值。

8. 鏈表的類型：單向鏈表、雙向鏈表、循環鏈表。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Linked List的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. Reverse Linked List II
2. Copy List with Random Pointer
3. Remove Nth Node From End of List
4. Partition List
5. Linked List Cycle II

1. Reverse Linked List II:
題目描述： 反轉從位置 m 到 n 的鏈表。請使用一趟掃瞄完成反轉。
示例: 
輸入: 1->2->3->4->5->NULL, m = 2, n = 4
輸出: 1->4->3->2->5->NULL
答案鏈接: https://leetcode.com/problems/reverse-linked-list-ii/

2. Copy List with Random Pointer:
題目描述： 給定一個鏈表，每個節點包含一個額外增加的隨機指針，
該指針可以指向鏈表中的任何節點或空節點。
要求返回這個鏈表的 深拷貝。 
示例：
輸入：
{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
解釋：
節點 1 的值是 1，它的下一個指針和隨機指針都指向節點 2 。
節點 2 的值是 2，它的下一個指針指向 null，隨機指針指向它本身。
答案鏈接：https://leetcode.com/problems/copy-list-with-random-pointer/

3. Remove Nth Node From End of List:
題目描述： 給定一個鏈表，刪除鏈表的倒數第 n 個節點，並且返回鏈表的頭結點。 
示例：
輸入: 1->2->3->4->5, n = 2
輸出: 1->2->3->5
答案鏈接： https://leetcode.com/problems/remove-nth-node-from-end-of-list/

4. Partition List:
題目描述： 給定一個鏈表和一個特定值 x，對鏈表進行分隔，
使得所有小於 x 的節點都在大於或等於 x 的節點之前。
你應當保留兩個分區中每個節點的初始相對位置。
示例：
輸入: head = 1->4->3->2->5->2, x = 3
輸出: 1->2->2->4->3->5
答案鏈接： https://leetcode.com/problems/partition-list/

5. Linked List Cycle II:
題目描述： 給定一個鏈表，返回鏈表開始入環的第一個節點。 如果鏈表無環，則返回 null。
示例：
輸入：head = [3,2,0,-4], pos = 1
輸出：tail connects to node index 1
答案鏈接： https://leetcode.com/problems/linked-list-cycle-ii/   

