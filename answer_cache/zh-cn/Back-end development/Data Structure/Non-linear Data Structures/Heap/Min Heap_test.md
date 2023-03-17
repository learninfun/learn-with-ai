

1. 實作Min Heap的add方法，將一個數字加入到Min Heap中。
答案：

```python
def add(self, val):
    self.heap.append(val)
    curr = len(self.heap) - 1
    while curr > 0 and self.heap[curr] < self.heap[self.parent(curr)]:
        self.heap[curr], self.heap[self.parent(curr)] = self.heap[self.parent(curr)], self.heap[curr]
        curr = self.parent(curr)
```

2. 實作Min Heap的remove方法，將Min Heap中的最小值取出。
答案：

```python
def remove_min(self):
    if len(self.heap) == 1:
        return self.heap.pop()
    min_val = self.heap[0]
    self.heap[0] = self.heap.pop()
    self.min_heapify(0)
    return min_val

def min_heapify(self, i):
    left_child = self.left_child(i)
    right_child = self.right_child(i)
    smallest = i
    if left_child < len(self.heap) and self.heap[left_child] < self.heap[smallest]:
        smallest = left_child
    if right_child < len(self.heap) and self.heap[right_child] < self.heap[smallest]:
        smallest = right_child
    if smallest != i:
        self.heap[i], self.heap[smallest] = self.heap[smallest], self.heap[i]
        self.min_heapify(smallest)
```

3. 設計一個算法，用於在Min Heap中找到第k小的元素。
答案：

可以使用堆排序的思想，先建立一個大小為k的Min Heap，然後拉取剩餘的元素，如果當前元素大於Min Heap的root，則跳過該元素，否則將該元素加入到Min Heap中，並把Min Heap的root取出，直到遍歷完所有的元素為止。

```python
def find_kth_smallest(self, k):
    min_heap = []
    for i in range(k):
        min_heap.append(self.heap[i])
    heapq.heapify(min_heap)
    for i in range(k, len(self.heap)):
        if self.heap[i] > min_heap[0]:
            continue
        heapq.heappop(min_heap)
        heapq.heappush(min_heap, self.heap[i])
    return min_heap[0]
```

4. 設計一個算法，用於在Min Heap中找到第k大的元素。
答案：

可以使用Max Heap的思想，先建立一個大小為k的Max Heap，然後拉取剩餘的元素，如果當前元素小於Max Heap的root，則跳過該元素，否則將該元素加入到Max Heap中，並把Max Heap的root取出，直到遍歷完所有的元素為止。

```python
def find_kth_largest(self, k):
    max_heap = []
    for i in range(k):
        heapq.heappush(max_heap, -self.heap[i])
    for i in range(k, len(self.heap)):
        if self.heap[i] < -max_heap[0]:
            continue
        heapq.heappop(max_heap)
        heapq.heappush(max_heap, -self.heap[i])
    return -max_heap[0]
```

5. 設計一個算法，用於將一個已排序的數組轉換成Min Heap。
答案：

可以使用Min Heapify的思想，從最後一個有子節點的節點開始往前，對每一個節點都執行Min Heapify操作。

```python
def build_heap(self, arr):
    self.heap = arr
    for i in range(self.parent(len(self.heap) - 1), -1, -1):
        self.min_heapify(i)
```