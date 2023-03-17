## 習題預習
{{< ask_chatgpt >}}
給我5題Circular Linked List的問題
{{< /ask_chatgpt >}}



1. 實作一個 Circular Linked List，使得最後一個節點的下一個節點指向第一個節點。
2. 查找 Circular Linked List 中是否存在某個特定元素，並回傳其位置。
3. 在 Circular Linked List 中插入一個節點，使其成為頭節點。
4. 刪除 Circular Linked List 中某個特定位置的節點。
5. 翻轉一個 Circular Linked List，使其從尾部插入節點，並將原來的頭節點變成尾節點。   

## 說明知識
{{< ask_chatgpt >}}
說明Circular Linked List並舉例
{{< /ask_chatgpt >}}



Circular Linked List是一種連結串列的資料結構，它與一般的單向鏈結串列不同的地方在於它結尾的節點會指向頭部節點，形成一個循環的連結串列。因此，當我們訪問最後一個節點時，它指向的是頭部節點而不是NULL。

以下是一個示範Circular Linked List的圖表：

```
1 -> 2 -> 3 -> 4 -> 1
```

在上面的例子中，我們可以看到最後一個節點指向第一個節點，因此它是一個圓形的連結串列。當我們查詢最後一個節點時，程式碼就會跳回第一個節點，進入一個循環，因為每個節點都指向下一個節點，直到最後一個節點再指向第一個節點。這使得操作和查詢Circular Linked List的時間複雜度相較於Linked List更高效。

在實現Circular Linked List時，我們需要在每個節點中維護指向下一個節點的指標，並且保證最後一個節點指向第一個節點。

以下是一個示範建立Circular Linked List的程式碼：

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

## 彙總重點
{{< ask_chatgpt >}}
條列Circular Linked List的重點
{{< /ask_chatgpt >}}



1. 它是一種連接成環形的鏈表，最後一個節點連接到第一個節點，形成循環路線。
2. 其中每個節點包含了數據以及指向下一個節點的指針，最後一個節點則指向第一個節點。
3. 與單向和雙向鏈表一樣，它支持基本操作，例如在頭部或尾部添加/刪除節點。
4. 在某些情況下，使用循環鏈表可以更有效的利用記憶體空間。
5. 在算法中，使用循環鏈表可以使代碼更簡潔易懂，例如在輸出環形標誌時，只需在節點指針上設定標誌即可完成。
6. 缺點是它易於產生死循環，因此需要仔細管理節點之間的指向關係，以避免程序陷入死循環的狀態。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Circular Linked List的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 實作將Circular Linked List反轉。

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

2. 實作從Circular Linked List中移除擁有特定數值的節點。

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

3. 判斷Circular Linked List是否為迴文（即正向與反向皆相同）。

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

4. 將兩個Circular Linked List合併（由小到大排序）。

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

5. 在Circular Linked List中找到最大值（若有重複數值，輸出第一個）。

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

## 網路資料
{{< ask_chatgpt >}}
給我5篇Circular Linked List的網路資料
{{< /ask_chatgpt >}}



1. Circular Linked List - GeeksforGeeks
https://www.geeksforgeeks.org/circular-linked-list/

這篇文章說明了Circular Linked List是什麼，以及如何建立和操作它。它也提供了一些程式碼和例子，以幫助讀者更好地理解這個概念。

2. Circular Linked List | Brilliant Math & Science Wiki
https://brilliant.org/wiki/circular-linked-list/

這個網站提供了一個非常基本的介紹，介紹了什麼是Circular Linked List、如何建立和操作它，以及一些應用案例。它的內容非常清晰明瞭，適合初學者閱讀。

3. Circular Linked List in Data Structure with Example
https://www.tutorialspoint.com/data_structures_algorithms/circular_linked_list_algorithm.htm

這篇文章通過示例演示了如何使用Circular Linked List來解決問題。它解釋了如何插入和刪除節點，以及如何訪問和更新它們。這個網站提供了許多其他的教學和例子，也是一個很好的資源。

4. Data Structures and Algorithms - Circular Linked List
https://www.javatpoint.com/circular-linked-list

這個網站提供了關於Circular Linked List的徹底的介紹，包括什麼是Circular Linked List、它有什麼優點和缺點、如何創建和操作它，以及如何解決常見的問題。它的內容非常豐富，是進階學習者學習的好資源。

5. Circular Linked List | Codecademy
https://www.codecademy.com/articles/circular-linked-list

Codecademy是一個出色的編程教學網站，這篇文章提供了一個簡單的介紹，解釋了如何使用Circular Linked List來解決問題。它提供了一些簡單的例子和程式碼，幫助學習者快速上手。   

