+++
title = "最小堆"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Min Heap的中文问题
{{< /ask_chatgpt >}}



1. 如何建立一个Min Heap？
2. 如何插入一个元素到Min Heap中？
3. 如何删除Min Heap中的最小元素？
4. 如何取得Min Heap中的最小值？
5. 如何进行Min Heap的排序？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Min Heap并举例
{{< /ask_chatgpt >}}



Min Heap是一种资料结构，它是一棵完全二元树，即除了最后一层之外，每一层都是满的，并且最后一层的节点都从左边开始填入。在Min Heap中，每一个节点的值都小于或等于它的两个子节点的值，因此根节点保存的是整个Heap中的最小值。

举例来说，以下是一个Min Heap：

```
          3
       /     \
     12       5
    /  \     / \
   20  11   7   14
```

在这个Min Heap中，根节点的值为3，并且根节点的两个子节点的值都大于等于3，即12和5。同样地，12和5这两个节点也满足节点值小于或等于它们的子节点的值。

另外，Min Heap的一个重要特点是，当我们从Min Heap中删除根节点时，Heap结构会自动调整，使得新的根节点还是整个Heap中的最小值。具体来说，我们可以先将最后一个节点复制到根节点的位置，然后不断地比较新的根节点和它的两个子节点的值，如果它的值大于其中某一个子节点的值，就交换它们的位置，直到找到了合适的位置为止。这样一来，我们就可以在O(log n)的时间内删除Min Heap中的最小值了。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Min Heap的重点
{{< /ask_chatgpt >}}



1. 它是一种堆数据结构，使用数组表示，并保持一定的顺序特征。
2. 堆通常分为两种类型：最大堆和最小堆。
3. 最小堆的根节点是具有最小值的节点，并且该属性在子树中递归地保持。
4. 最小堆的父节点始终小于或等于其子节点，而左子节点首先小于右子节点。
5. 插入新值时，可以将新值放到数组末尾，然后通过交换父节点和子节点的值来重新平衡堆。
6. 取出最小值时，从最小堆的根节点返回最小值，并从数组末尾将新值放入根节点位置，然后通过交换父节点和子节点的值来重新平衡堆。
7. 用于堆排序，优先队列等场合。
8. 时间复杂度为O(log n)，其中n是堆的大小。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Min Heap的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 实作Min Heap的add方法，将一个数字加入到Min Heap中。
答案：

```python
def add(self, val):
    self.heap.append(val)
    curr = len(self.heap) - 1
    while curr > 0 and self.heap[curr] < self.heap[self.parent(curr)]:
        self.heap[curr], self.heap[self.parent(curr)] = self.heap[self.parent(curr)], self.heap[curr]
        curr = self.parent(curr)
```

2. 实作Min Heap的remove方法，将Min Heap中的最小值取出。
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

3. 设计一个算法，用于在Min Heap中找到第k小的元素。
答案：

可以使用堆排序的思想，先建立一个大小为k的Min Heap，然后拉取剩余的元素，如果当前元素大于Min Heap的root，则跳过该元素，否则将该元素加入到Min Heap中，并把Min Heap的root取出，直到遍历完所有的元素为止。

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

4. 设计一个算法，用于在Min Heap中找到第k大的元素。
答案：

可以使用Max Heap的思想，先建立一个大小为k的Max Heap，然后拉取剩余的元素，如果当前元素小于Max Heap的root，则跳过该元素，否则将该元素加入到Max Heap中，并把Max Heap的root取出，直到遍历完所有的元素为止。

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

5. 设计一个算法，用于将一个已排序的数组转换成Min Heap。
答案：

可以使用Min Heapify的思想，从最后一个有子节点的节点开始往前，对每一个节点都执行Min Heapify操作。

```python
def build_heap(self, arr):
    self.heap = arr
    for i in range(self.parent(len(self.heap) - 1), -1, -1):
        self.min_heapify(i)
```   

