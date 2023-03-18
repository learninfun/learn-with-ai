

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