

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