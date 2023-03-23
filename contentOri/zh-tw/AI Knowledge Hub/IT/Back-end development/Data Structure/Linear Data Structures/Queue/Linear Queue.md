+++
title = "線性佇列"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Linear Queue的中文問題
{{< /ask_chatgpt >}}



1. 請說明什麼是 Linear Queue，以及它的特性是什麼？
2. 在 Linear Queue 中，什麼情況下會發生 Queue 滿了的情況？如果發生了該怎麼處理？
3. Linear Queue 的基本操作有哪些？請分別說明它們的意思。
4. 請寫出一段程式碼，實現 Linear Queue 的入隊和出隊操作。
5. Linear Queue 和 Circular Queue 有什麼不同？在什麼情況下應該選擇使用 Circular Queue？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Linear Queue並舉例
{{< /ask_chatgpt >}}



Linear Queue是指一種基於先進先出(FIFO)原則的資料結構。如同一列在銀行排隊的方式，最先進入排隊的人最先被處理，後進入排隊的人就要等候前面的人處理完畢以後才能進行下一步操作。

在Linear Queue中，資料是線性排列的，並且在做入隊(Enqueue)和出隊(Dequeue)操作時，資料只能在頭尾兩端進行。一般來說，Linear Queue是用Array或Linked List實現的。

以下是Linear Queue的範例：

當一列人在銀行排隊時，最先進入排隊的人(ID: 001)，會成為第一個進入Queue的元素。之後，第二個人(ID: 002)進入排隊，成為Enqueue的元素。當第一個人(ID: 001)完成作業後，他成為Dequeue的元素，並且由第二個人(ID: 002)取代成為Head元素。

就像這個例子，當資料依照順序進入Queue後，會依照先進先出的原則完成操作。所有在Queue中等候的元素會在適當時間被處理完畢。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Linear Queue的重點
{{< /ask_chatgpt >}}



1. Linear Queue是一種線性的資料結構，具有先進先出的特性。
2. 它通常包含一個前端(front)指針和一個後端(rear)指針，用來指示首尾兩端的位置。
3. 新元素在後端(rear)添加，舊元素在前端(front)刪除。
4. 元素依次排列，插入和刪除均按照先進先出(FIFO)的原則進行。
5. 當前端(front)和後端(rear)指針相等空間時，稱為空隊列，當尾指針加1等於佇列長度時，稱為滿隊列。
6. 可以使用數組或鏈表實現Linear Queue，一般選擇鏈表實現，具有彈性且節省空間。
7. 常見操作有：enqueue(添加元素)、dequeue(刪除元素)、isEmpty(判斷是否空隊列)、isFull(判斷是否滿隊列)等。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Linear Queue的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 實現一個簡單的線性隊列，要求包含以下操作：

- enqueue(item): 將元素 item 添加到隊列的尾部。
- dequeue(): 返回並刪除隊列的前端元素，如果隊列為空，則返回 None。
- is_empty(): 返回布爾值，用於判斷隊列是否為空。
- is_full(): 返回布爾值，用於判斷隊列是否已滿。

答案：

```
class Queue:
    def __init__(self, capacity):
        self.items = [None] * capacity
        self.capacity = capacity
        self.head = 0
        self.tail = -1
        self.size = 0

    def enqueue(self, item):
        if self.is_full():
            raise Exception("Queue is full")
        self.tail += 1
        self.items[self.tail % self.capacity] = item
        self.size += 1

    def dequeue(self):
        if self.is_empty():
            return None
        item = self.items[self.head % self.capacity]
        self.items[self.head % self.capacity] = None
        self.head += 1
        self.size -= 1
        return item

    def is_empty(self):
        return self.size == 0

    def is_full(self):
        return self.size == self.capacity

```

2. 實現一個使用線性隊列實現堆棧的類別，要求包含以下操作：

- push(item): 將元素 item 壓入堆棧。
- pop(): 返回並刪除堆棧頂部元素，如果堆棧為空，則返回 None。
- is_empty(): 返回布爾值，用於判斷堆棧是否為空。

答案：

```
class Stack:
    def __init__(self, capacity):
        self.queue = Queue(capacity)
        self.capacity = capacity

    def push(self, item):
        if self.queue.is_full():
            raise Exception("Stack is full")
        self.queue.enqueue(item)

    def pop(self):
        if self.queue.is_empty():
            return None
        for i in range(self.queue.size - 1):
            self.queue.enqueue(self.queue.dequeue())
        return self.queue.dequeue()

    def is_empty(self):
        return self.queue.is_empty()

```

3. 實現一個線性隊列，支持在任意位置插入和刪除元素的操作，要求包含以下操作：

- insert(item, index): 在指定位置 index 插入元素 item。
- delete(index): 刪除指定位置 index 上的元素，並返回它的值，如果該位置不存在元素，則返回 None。
- is_empty(): 返回布爾值，用於判斷隊列是否為空。

答案：

```
class Queue:
    def __init__(self):
        self.items = []

    def insert(self, item, index):
        self.items.insert(index, item)

    def delete(self, index):
        if index < 0 or index >= len(self.items):
            return None
        return self.items.pop(index)

    def is_empty(self):
        return len(self.items) == 0

```

4. 實現一個線性隊列，可以對隊列中的元素進行反轉操作，要求包含以下操作：

- reverse(): 將隊列中的元素逆序排列。

答案：

```
class Queue:
    def __init__(self):
        self.items = []

    def reverse(self):
        self.items.reverse()

    def enqueue(self, item):
        self.items.append(item)

    def dequeue(self):
        if not self.items:
            return None
        return self.items.pop(0)

    def is_empty(self):
        return len(self.items) == 0

```

5. 給定一個整數隊列和一個整數 k，要求實現一個函數，將每個大小為 k 的子隊列進行反轉，不足 k 的子隊列則保持原序列不變，要求時間複雜度為 O(n)。

答案：

```
class Queue:
    def __init__(self, items):
        self.items = items

    def reverse_k_subqueues(self, k):
        for i in range(0, len(self.items), k):
            self.items[i:i+k] = self.items[i:i+k][::-1]

    def is_empty(self):
        return len(self.items) == 0

q = Queue([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])
q.reverse_k_subqueues(3)
print(q.items)  # [3, 2, 1, 6, 5, 4, 9, 8, 7, 10]
```   

