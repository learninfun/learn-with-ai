

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