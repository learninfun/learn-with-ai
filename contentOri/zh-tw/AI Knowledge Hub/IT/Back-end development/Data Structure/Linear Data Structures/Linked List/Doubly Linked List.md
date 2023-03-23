+++
title = "雙向鏈表"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Doubly Linked List的中文問題
{{< /ask_chatgpt >}}



1. 如何在Doubly Linked List的開頭添加一個節點？
2. 如何在Doubly Linked List的末尾添加一個節點？
3. 如何在Doubly Linked List中查找一個特定節點？
4. 如何在Doubly Linked List中刪除一個特定節點？
5. 如何翻轉Doubly Linked List中的元素？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Doubly Linked List並舉例
{{< /ask_chatgpt >}}



雙向鏈結串列（Doubly Linked List），是一個資料結構，其中每個節點同時擁有指向前一個節點和後一個節點的指針，並可以在任意位置進行插入和刪除操作。因此，雙向鏈結串列不像單向鏈結串列那樣只能從頭到尾或從尾到頭遍歷。

一個雙向鏈結串列節點通常被組織為包含三個元素：

1. 值：存儲在節點中的數據元素
2. 指向前一個節點的指針（previous pointer）
3. 指向下一個節點的指針（next pointer）

對於第一個節點和最後一個節點，前一個節點指針和後一個節點指針可以為空指針（NULL）。

以下是一個節點的示例：

```
struct Node {
    int value;
    Node* prev;
    Node* next;
};
```

舉例：

假設有以下雙向鏈結串列：1 ⇄ 3 ⇄ 5 ⇄ 7，其中 1 是第一個節點，而 7 是最後一個節點。

- 如果要在第二個節點 3 後插入一個新節點 4，操作流程如下：

  1. 建立一個新節點 4；
  2. 使 4 的 prev 指向 3；
  3. 使 4 的 next 指向 5；
  4. 使 3 的 next 指向 4；
  5. 使 5 的 prev 指向 4；
  
  插入後，鏈結串列變為：1 ⇄ 3 ⇄ 4 ⇄ 5 ⇄ 7。

- 如果要在第四個節點 5 前插入一個新節點 2，操作流程如下：

  1. 建立一個新節點 2；
  2. 使 2 的 prev 指向 1；
  3. 使 2 的 next 指向 3；
  4. 使 1 的 next 指向 2；
  5. 使 3 的 prev 指向 2；
  
  插入後，鏈結串列變為：1 ⇄ 2 ⇄ 3 ⇄ 5 ⇄ 7。


- 如果要從鏈結串列中刪除節點 3，操作流程如下：

  1. 使 2 的 next 指向 5；
  2. 使 5 的 prev 指向 2；
  3. 釋放節點 3 的內存空間。
  
  刪除後，鏈結串列變為：1 ⇄ 2 ⇄ 5 ⇄ 7。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Doubly Linked List的重點
{{< /ask_chatgpt >}}



1. Doubly linked list是一種常見的資料結構，它與 singly linked list 相似，但每個節點有兩個指向（前一個和後一個）而不是一個。

2. Doubly linked list的主要優點是可以雙向遍歷列表。這意味著，可以通過列表的前面或後面尋找節點，而不必遞歸或反轉列表。

3. Doubly linked list需要更多的空間來存儲指向前面節點的指針，這會增加記憶體使用量。

4. 在Doubly linked list中，每個節點都有前一個節點和後一個節點的指針，這使得在插入或刪除節點時可以更加高效。

5. 在Doubly linked list中，頭尾節點的處理需要更多的注意，需要特別處理空列表的情況。

6. Doubly linked list支持正向和反向遍歷，這使得在某些算法中可以更加高效。

7. Doubly linked list的插入和刪除操作時需要更多的指針操作，可能比singly linked list要更複雜。

8. 在Doubly linked list中，每個節點都需要額外的空間來存儲前一個節點和後一個節點，這增加了節點的大小，可能會影響效率。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Doubly Linked List的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 實現一個Doubly Linked List的reverse方法

答案:

```java
public void reverse() {
    Node temp = null;
    Node current = head;
    while (current != null) {
        temp = current.prev;
        current.prev = current.next;
        current.next = temp;
        current = current.prev;
    }
    if (temp != null) {
        head = temp.prev;
    }
}
```

2. 實現一個Doubly Linked List的merge方法，將兩個有序的Doubly Linked List合併成一個新的有序Doubly Linked List

答案:

```java
public DoublyLinkedList merge(DoublyLinkedList list1, DoublyLinkedList list2) {
    DoublyLinkedList mergedList = new DoublyLinkedList();
    Node temp1 = list1.head;
    Node temp2 = list2.head;
    while (temp1 != null && temp2 != null) {
        if (temp1.value <= temp2.value) {
            mergedList.add(temp1.value);
            temp1 = temp1.next;
        } else {
            mergedList.add(temp2.value);
            temp2 = temp2.next;
        }
    }
    while (temp1 != null) {
        mergedList.add(temp1.value);
        temp1 = temp1.next;
    }
    while (temp2 != null) {
        mergedList.add(temp2.value);
        temp2 = temp2.next;
    }
    return mergedList;
}
```

3. 實現一個Doubly Linked List的removeDuplicates方法，刪除所有重複的元素

答案:

```java
public void removeDuplicates() {
    if (head == null) {
        return;
    }
    Node current = head;
    while (current != null) {
        Node temp = current.next;
        while (temp != null && temp.value == current.value) {
            temp = temp.next;
        }
        current.next = temp;
        temp.prev = current;
        current = current.next;
    }
}
```

4. 實現一個Doubly Linked List的findKthFromLast方法，查找倒數第k個元素

答案:

```java
public int findKthFromLast(int k) {
    Node current = head;
    Node runner = head;
    for (int i = 0; i < k; i++) {
        if (runner == null) {
            return -1; // 如果k大於列表的大小，返回-1
        }
        runner = runner.next;
    }
    while (runner != null) {
        current = current.next;
        runner = runner.next;
    }
    return current.value;
}
```

5. 實現一個Doubly Linked List的sort方法，將列表進行升序排列

答案:

```java
public void sort() {
    Node current = head;
    Node index = null;
    int temp;
    if (head == null) {
        return;
    } else {
        while (current != null) {
            index = current.next;
            while (index != null) {
                if (current.value > index.value) {
                    temp = current.value;
                    current.value = index.value;
                    index.value = temp;
                }
                index = index.next;
            }
            current = current.next;
        }
    }
}
```   

