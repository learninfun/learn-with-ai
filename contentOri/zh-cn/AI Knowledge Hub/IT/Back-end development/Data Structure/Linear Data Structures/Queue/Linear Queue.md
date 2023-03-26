+++
title = "线性伫列"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Linear Queue的中文问题
{{< /ask_chatgpt >}}



1. 请说明什么是 Linear Queue，以及它的特性是什么？
2. 在 Linear Queue 中，什么情况下会发生 Queue 满了的情况？如果发生了该怎么处理？
3. Linear Queue 的基本操作有哪些？请分别说明它们的意思。
4. 请写出一段程式码，实现 Linear Queue 的入队和出队操作。
5. Linear Queue 和 Circular Queue 有什么不同？在什么情况下应该选择使用 Circular Queue？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Linear Queue并举例
{{< /ask_chatgpt >}}



Linear Queue是指一种基于先进先出(FIFO)原则的资料结构。如同一列在银行排队的方式，最先进入排队的人最先被处理，后进入排队的人就要等候前面的人处理完毕以后才能进行下一步操作。

在Linear Queue中，资料是线性排列的，并且在做入队(Enqueue)和出队(Dequeue)操作时，资料只能在头尾两端进行。一般来说，Linear Queue是用Array或Linked List实现的。

以下是Linear Queue的范例：

当一列人在银行排队时，最先进入排队的人(ID: 001)，会成为第一个进入Queue的元素。之后，第二个人(ID: 002)进入排队，成为Enqueue的元素。当第一个人(ID: 001)完成作业后，他成为Dequeue的元素，并且由第二个人(ID: 002)取代成为Head元素。

就像这个例子，当资料依照顺序进入Queue后，会依照先进先出的原则完成操作。所有在Queue中等候的元素会在适当时间被处理完毕。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Linear Queue的重点
{{< /ask_chatgpt >}}



1. Linear Queue是一种线性的资料结构，具有先进先出的特性。
2. 它通常包含一个前端(front)指针和一个后端(rear)指针，用来指示首尾两端的位置。
3. 新元素在后端(rear)添加，旧元素在前端(front)删除。
4. 元素依次排列，插入和删除均按照先进先出(FIFO)的原则进行。
5. 当前端(front)和后端(rear)指针相等空间时，称为空队列，当尾指针加1等于伫列长度时，称为满队列。
6. 可以使用数组或链表实现Linear Queue，一般选择链表实现，具有弹性且节省空间。
7. 常见操作有：enqueue(添加元素)、dequeue(删除元素)、isEmpty(判断是否空队列)、isFull(判断是否满队列)等。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Linear Queue的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 实现一个简单的线性队列，要求包含以下操作：

- enqueue(item): 将元素 item 添加到队列的尾部。
- dequeue(): 返回并删除队列的前端元素，如果队列为空，则返回 None。
- is_empty(): 返回布尔值，用于判断队列是否为空。
- is_full(): 返回布尔值，用于判断队列是否已满。

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

2. 实现一个使用线性队列实现堆栈的类别，要求包含以下操作：

- push(item): 将元素 item 压入堆栈。
- pop(): 返回并删除堆栈顶部元素，如果堆栈为空，则返回 None。
- is_empty(): 返回布尔值，用于判断堆栈是否为空。

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

3. 实现一个线性队列，支持在任意位置插入和删除元素的操作，要求包含以下操作：

- insert(item, index): 在指定位置 index 插入元素 item。
- delete(index): 删除指定位置 index 上的元素，并返回它的值，如果该位置不存在元素，则返回 None。
- is_empty(): 返回布尔值，用于判断队列是否为空。

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

4. 实现一个线性队列，可以对队列中的元素进行反转操作，要求包含以下操作：

- reverse(): 将队列中的元素逆序排列。

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

5. 给定一个整数队列和一个整数 k，要求实现一个函数，将每个大小为 k 的子队列进行反转，不足 k 的子队列则保持原序列不变，要求时间复杂度为 O(n)。

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

