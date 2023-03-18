+++
title = "循环链表"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Circular Linked List的问题
{{< /ask_chatgpt >}}



1. 实作一个 Circular Linked List，使得最后一个节点的下一个节点指向第一个节点。
2. 查找 Circular Linked List 中是否存在某个特定元素，并回传其位置。
3. 在 Circular Linked List 中插入一个节点，使其成为头节点。
4. 删除 Circular Linked List 中某个特定位置的节点。
5. 翻转一个 Circular Linked List，使其从尾部插入节点，并将原来的头节点变成尾节点。   

## 说明知识
{{< ask_chatgpt >}}
说明Circular Linked List并举例
{{< /ask_chatgpt >}}



Circular Linked List是一种连结串列的资料结构，它与一般的单向链结串列不同的地方在于它结尾的节点会指向头部节点，形成一个循环的连结串列。因此，当我们访问最后一个节点时，它指向的是头部节点而不是NULL。

以下是一个示范Circular Linked List的图表：

```
1 -> 2 -> 3 -> 4 -> 1
```

在上面的例子中，我们可以看到最后一个节点指向第一个节点，因此它是一个圆形的连结串列。当我们查询最后一个节点时，程式码就会跳回第一个节点，进入一个循环，因为每个节点都指向下一个节点，直到最后一个节点再指向第一个节点。这使得操作和查询Circular Linked List的时间复杂度相较于Linked List更高效。

在实现Circular Linked List时，我们需要在每个节点中维护指向下一个节点的指标，并且保证最后一个节点指向第一个节点。

以下是一个示范建立Circular Linked List的程式码：

```c
typedef struct Node {
    int value;
    struct Node *next;
} Node;

Node *createCircularLinkedList(int values[], int n) {
    Node *head = NULL;
    Node *prev = NULL;
    Node *curr = NULL;
    for (int i = 0; i < n; i++) {
        curr = malloc(sizeof(Node));
        curr->value = values[i];
        curr->next = NULL;
        if (prev) {
            prev->next = curr;
        } else {
            head = curr;
        }
        prev = curr;
    }
    if (curr) {
        curr->next = head;
    }
    return head;
}
```   

## 汇总重点
{{< ask_chatgpt >}}
条列Circular Linked List的重点
{{< /ask_chatgpt >}}



1. 它是一種連接成環形的鏈表，最後一個節點連接到第一個節點，形成循環路線。
2. 其中每個節點包含了數據以及指向下一個節點的指針，最後一個節點則指向第一個節點。
3. 與單向和雙向鏈表一樣，它支持基本操作，例如在頭部或尾部添加/刪除節點。
4. 在某些情況下，使用循環鏈表可以更有效的利用記憶體空間。
5. 在算法中，使用循環鏈表可以使代碼更簡潔易懂，例如在輸出環形標誌時，只需在節點指針上設定標誌即可完成。
6. 缺點是它易於產生死循環，因此需要仔細管理節點之間的指向關係，以避免程序陷入死循環的狀態。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Circular Linked List的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 实作将Circular Linked List反转。

```
void reverse(Node** head) 
{ 
    if (*head == NULL) 
        return; 

    Node* prev = NULL; 
    Node* current = *head; 
    Node* next; 

    do 
    { 
        next = current->next; 
        current->next = prev; 
        prev = current; 
        current = next; 
    } while (current != *head); 

    (*head)->next = prev; 
    *head = prev; 
} 
```

2. 实作从Circular Linked List中移除拥有特定数值的节点。

```
void removeNode(Node** head, int key) 
{ 
    if (*head == NULL) 
        return; 

    Node* current = *head; 
    Node* prev; 

    do 
    { 
        if (current->data == key) 
        { 
            if (current == *head) 
                *head = current->next; 

            prev->next = current->next; 
            free(current); 
            current = prev->next; 
        } 
        else
        { 
            prev = current; 
            current = current->next; 
        } 
    } while (current != *head); 
} 
```

3. 判断Circular Linked List是否为回文（即正向与反向皆相同）。

```
bool isPalindrome(Node* head) 
{ 
    if (head == NULL) 
        return true; 

    Node *slow_ptr = head, *fast_ptr = head; 
    Node *prev_of_slow_ptr = head; 
    Node* midnode = NULL;  
    bool res = true; 

    if (head != NULL && head->next != NULL) 
    { 
        while (fast_ptr != NULL && fast_ptr->next != NULL) 
        { 
            fast_ptr = fast_ptr->next->next; 

            prev_of_slow_ptr = slow_ptr; 
            slow_ptr = slow_ptr->next; 
        } 

        if (fast_ptr != NULL) 
        { 
            midnode = slow_ptr; 
            slow_ptr = slow_ptr->next; 
        } 

        Node* second_half = slow_ptr; 
        prev_of_slow_ptr->next = NULL; 
        reverse(&second_half); 
        res = compareLists(head, second_half); 

        reverse(&second_half);  

        if (midnode != NULL) 
        { 

            prev_of_slow_ptr->next = midnode; 
            midnode->next = second_half; 
        } 
        else
            prev_of_slow_ptr->next = second_half; 
    } 
    return res; 
} 

bool compareLists(Node* head1, Node* head2) 
{ 
    Node* temp1 = head1; 
    Node* temp2 = head2; 

    while (temp1 && temp2) 
    { 
        if (temp1->data == temp2->data) 
        { 
            temp1 = temp1->next; 
            temp2 = temp2->next; 
        } 
        else
            return false; 
    } 

    if (temp1 == NULL && temp2 == NULL) 
        return true; 

    return false; 
} 
```

4. 将两个Circular Linked List合并（由小到大排序）。

```
Node* sortedMerge(Node* a, Node* b) 
{ 
    if (a == NULL) 
        return b; 
    if (b == NULL) 
        return a; 

    Node* result = NULL; 
    if (a->data <= b->data) 
    { 
        result = a; 
        result->next = sortedMerge(a->next, b); 
    } 
    else
    { 
        result = b; 
        result->next = sortedMerge(a, b->next); 
    } 

    return result; 
} 

Node* mergeSort(Node* head) 
{ 
    if (head == NULL || head->next == head) 
        return head; 

    Node *slow = head, *fast = head->next; 
    while (fast != head && fast->next != head) 
    { 
        slow = slow->next; 
        fast = fast->next->next; 
    } 

    Node* second_half = slow->next; 
    slow->next = head; 

    head = mergeSort(head); 
    second_half = mergeSort(second_half); 

    return sortedMerge(head, second_half); 
} 
```

5. 在Circular Linked List中找到最大值（若有重复数值，输出第一个）。

```
int getMax(Node* head) 
{ 
    if (head == NULL) 
        return -1; 

    int max = head->data; 
    Node* temp = head->next; 
    while (temp != head) 
    { 
        if (temp->data > max) 
            max = temp->data; 
        temp = temp->next; 
    } 
    return max; 
} 
```   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Circular Linked List的网络数据
{{< /ask_chatgpt >}}



1. Circular Linked List - GeeksforGeeks
https://www.geeksforgeeks.org/circular-linked-list/

这篇文章说明了Circular Linked List是什么，以及如何建立和操作它。它也提供了一些程式码和例子，以帮助读者更好地理解这个概念。

2. Circular Linked List | Brilliant Math & Science Wiki
https://brilliant.org/wiki/circular-linked-list/

这个网站提供了一个非常基本的介绍，介绍了什么是Circular Linked List、如何建立和操作它，以及一些应用案例。它的内容非常清晰明了，适合初学者阅读。

3. Circular Linked List in Data Structure with Example
https://www.tutorialspoint.com/data_structures_algorithms/circular_linked_list_algorithm.htm

这篇文章通过示例演示了如何使用Circular Linked List来解决问题。它解释了如何插入和删除节点，以及如何访问和更新它们。这个网站提供了许多其他的教学和例子，也是一个很好的资源。

4. Data Structures and Algorithms - Circular Linked List
https://www.javatpoint.com/circular-linked-list

这个网站提供了关于Circular Linked List的彻底的介绍，包括什么是Circular Linked List、它有什么优点和缺点、如何创建和操作它，以及如何解决常见的问题。它的内容非常丰富，是进阶学习者学习的好资源。

5. Circular Linked List | Codecademy
https://www.codecademy.com/articles/circular-linked-list

Codecademy是一个出色的编程教学网站，这篇文章提供了一个简单的介绍，解释了如何使用Circular Linked List来解决问题。它提供了一些简单的例子和程式码，帮助学习者快速上手。   

