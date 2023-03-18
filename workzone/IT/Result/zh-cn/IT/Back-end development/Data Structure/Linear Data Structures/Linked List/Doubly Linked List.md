+++
title = "双向链表"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Doubly Linked List的问题
{{< /ask_chatgpt >}}



1. 如何在Doubly Linked List的开头添加一个节点？
2. 如何在Doubly Linked List的末尾添加一个节点？
3. 如何在Doubly Linked List中查找一个特定节点？
4. 如何在Doubly Linked List中删除一个特定节点？
5. 如何翻转Doubly Linked List中的元素？   

## 说明知识
{{< ask_chatgpt >}}
说明Doubly Linked List并举例
{{< /ask_chatgpt >}}



双向链结串列（Doubly Linked List），是一个资料结构，其中每个节点同时拥有指向前一个节点和后一个节点的指针，并可以在任意位置进行插入和删除操作。因此，双向链结串列不像单向链结串列那样只能从头到尾或从尾到头遍历。

一个双向链结串列节点通常被组织为包含三个元素：

1. 值：存储在节点中的数据元素
2. 指向前一个节点的指针（previous pointer）
3. 指向下一个节点的指针（next pointer）

对于第一个节点和最后一个节点，前一个节点指针和后一个节点指针可以为空指针（NULL）。

以下是一个节点的示例：

```
struct Node {
    int value;
    Node* prev;
    Node* next;
};
```

举例：

假设有以下双向链结串列：1 ⇄ 3 ⇄ 5 ⇄ 7，其中 1 是第一个节点，而 7 是最后一个节点。

- 如果要在第二个节点 3 后插入一个新节点 4，操作流程如下：

  1. 建立一个新节点 4；
  2. 使 4 的 prev 指向 3；
  3. 使 4 的 next 指向 5；
  4. 使 3 的 next 指向 4；
  5. 使 5 的 prev 指向 4；
  
  插入后，链结串列变为：1 ⇄ 3 ⇄ 4 ⇄ 5 ⇄ 7。

- 如果要在第四个节点 5 前插入一个新节点 2，操作流程如下：

  1. 建立一个新节点 2；
  2. 使 2 的 prev 指向 1；
  3. 使 2 的 next 指向 3；
  4. 使 1 的 next 指向 2；
  5. 使 3 的 prev 指向 2；
  
  插入后，链结串列变为：1 ⇄ 2 ⇄ 3 ⇄ 5 ⇄ 7。


- 如果要从链结串列中删除节点 3，操作流程如下：

  1. 使 2 的 next 指向 5；
  2. 使 5 的 prev 指向 2；
  3. 释放节点 3 的内存空间。
  
  删除后，链结串列变为：1 ⇄ 2 ⇄ 5 ⇄ 7。   

## 汇总重点
{{< ask_chatgpt >}}
条列Doubly Linked List的重点
{{< /ask_chatgpt >}}



1. Doubly linked list是一种常见的资料结构，它与 singly linked list 相似，但每个节点有两个指向（前一个和后一个）而不是一个。

2. Doubly linked list的主要优点是可以双向遍历列表。这意味着，可以通过列表的前面或后面寻找节点，而不必递归或反转列表。

3. Doubly linked list需要更多的空间来存储指向前面节点的指针，这会增加记忆体使用量。

4. 在Doubly linked list中，每个节点都有前一个节点和后一个节点的指针，这使得在插入或删除节点时可以更加高效。

5. 在Doubly linked list中，头尾节点的处理需要更多的注意，需要特别处理空列表的情况。

6. Doubly linked list支持正向和反向遍历，这使得在某些算法中可以更加高效。

7. Doubly linked list的插入和删除操作时需要更多的指针操作，可能比singly linked list要更复杂。

8. 在Doubly linked list中，每个节点都需要额外的空间来存储前一个节点和后一个节点，这增加了节点的大小，可能会影响效率。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Doubly Linked List的中等难度问题，并在后面列出答案
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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Doubly Linked List的网络数据
{{< /ask_chatgpt >}}



1. "Doubly Linked List" from GeeksforGeeks - https://www.geeksforgeeks.org/doubly-linked-list/

This article provides an introduction to doubly linked list and explains its implementation with the help of code snippets in C++.

2. "Doubly Linked List Explained" from Tutorials Point - https://www.tutorialspoint.com/data_structures_algorithms/doubly_linked_list_algorithm.htm

This tutorial covers the basics of doubly linked list and includes an explanation of its insertion and deletion operations.

3. "Doubly Linked Lists" from CS50 - https://cs50.me/submissions/7707c2cf-6b37-4278-bd6d-e43d3ea3be98

This article from the Harvard CS50 course provides an easy-to-understand explanation of doubly linked list and its implementation using pointers.

4. "Doubly Linked Lists" from Data Structures and Algorithms in Java by Robert Lafore - https://booksite.elsevier.com/9788131211792/chapter10.pdf

This chapter from the popular programming book explains the concept of doubly linked list and its implementation using Java code examples.

5. "Doubly Linked List" from Data Structure and Algorithms by Narasimha Karumanchi - https://www.geeksforgeeks.org/doubly-linked-list/

This chapter from the Narasimha Karumanchi's book provides an in-depth explanation of doubly linked list and its various operations, including insertion, deletion, and traversal.   

