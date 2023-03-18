+++
title = "线性伫列"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Linear Queue的问题
{{< /ask_chatgpt >}}



1. 请说明什么是 Linear Queue，以及它的特性是什么？
2. 在 Linear Queue 中，什么情况下会发生 Queue 满了的情况？如果发生了该怎么处理？
3. Linear Queue 的基本操作有哪些？请分别说明它们的意思。
4. 请写出一段程式码，实现 Linear Queue 的入队和出队操作。
5. Linear Queue 和 Circular Queue 有什么不同？在什么情况下应该选择使用 Circular Queue？   

## 说明知识
{{< ask_chatgpt >}}
说明Linear Queue并举例
{{< /ask_chatgpt >}}



Linear Queue是指一种基于先进先出(FIFO)原则的资料结构。如同一列在银行排队的方式，最先进入排队的人最先被处理，后进入排队的人就要等候前面的人处理完毕以后才能进行下一步操作。

在Linear Queue中，资料是线性排列的，并且在做入队(Enqueue)和出队(Dequeue)操作时，资料只能在头尾两端进行。一般来说，Linear Queue是用Array或Linked List实现的。

以下是Linear Queue的范例：

当一列人在银行排队时，最先进入排队的人(ID: 001)，会成为第一个进入Queue的元素。之后，第二个人(ID: 002)进入排队，成为Enqueue的元素。当第一个人(ID: 001)完成作业后，他成为Dequeue的元素，并且由第二个人(ID: 002)取代成为Head元素。

就像这个例子，当资料依照顺序进入Queue后，会依照先进先出的原则完成操作。所有在Queue中等候的元素会在适当时间被处理完毕。   

## 汇总重点
{{< ask_chatgpt >}}
条列Linear Queue的重点
{{< /ask_chatgpt >}}



1. Linear Queue是一種線性的資料結構，具有先進先出的特性。
2. 它通常包含一個前端(front)指針和一個後端(rear)指針，用來指示首尾兩端的位置。
3. 新元素在後端(rear)添加，舊元素在前端(front)刪除。
4. 元素依次排列，插入和刪除均按照先進先出(FIFO)的原則進行。
5. 當前端(front)和後端(rear)指針相等空間時，稱為空隊列，當尾指針加1等於佇列長度時，稱為滿隊列。
6. 可以使用數組或鏈表實現Linear Queue，一般選擇鏈表實現，具有彈性且節省空間。
7. 常見操作有：enqueue(添加元素)、dequeue(刪除元素)、isEmpty(判斷是否空隊列)、isFull(判斷是否滿隊列)等。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Linear Queue的中等难度问题，并在后面列出答案
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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Linear Queue的网络数据
{{< /ask_chatgpt >}}



1. "Linear Queue Implementation in C++" (https://www.geeksforgeeks.org/linear-queue-cpp/)
This article provides a step-by-step guide for implementing a linear queue data structure in C++ programming language. It includes a detailed explanation of the algorithms and code snippets for enqueue, dequeue, and display operations.

2. "Implementing a Linear Queue in Python" (https://www.tutorialspoint.com/python_data_structure/python_queue.htm)
This tutorial explains how to implement a linear queue data structure in Python programming language. The article provides code examples for the enqueue, dequeue, and display operations, along with an explanation of the time complexity of each operation.

3. "Linear Queue Data Structure" (https://www.javatpoint.com/data-structure-queue)
This article gives an overview of the linear queue data structure, including its definition, characteristics, and operations. The article also provides a comparison between a linear queue and a circular queue and explains when to use each data structure.

4. "Linear Queue using Array" (https://www.tutorialride.com/data-structures/linear-queue-in-c.htm)
This tutorial explains how to implement a linear queue using an array in C programming language. It includes code snippets for the enqueue, dequeue, and display operations and provides information about handling queue size and data overflow.

5. "Circular Queue vs Linear Queue" (https://www.guru99.com/circular-queue-vs-linear-queue.html)
This article compares the circular queue and linear queue data structures and explains when to use each. It provides examples of use cases for both data structures and highlights the advantages and disadvantages of each. The article also includes code snippets for implementing both data structures in C++ programming language.   

