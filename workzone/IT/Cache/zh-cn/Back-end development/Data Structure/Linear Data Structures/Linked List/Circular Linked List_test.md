

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