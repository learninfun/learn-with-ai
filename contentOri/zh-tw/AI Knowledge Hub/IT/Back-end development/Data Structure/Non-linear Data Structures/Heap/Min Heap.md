+++
title = "最小堆"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Min Heap的中文問題
{{< /ask_chatgpt >}}



1. 如何建立一個Min Heap？
2. 如何插入一個元素到Min Heap中？
3. 如何刪除Min Heap中的最小元素？
4. 如何取得Min Heap中的最小值？
5. 如何進行Min Heap的排序？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Min Heap並舉例
{{< /ask_chatgpt >}}



Min Heap是一種資料結構，它是一棵完全二元樹，即除了最後一層之外，每一層都是滿的，並且最後一層的節點都從左邊開始填入。在Min Heap中，每一個節點的值都小於或等於它的兩個子節點的值，因此根節點保存的是整個Heap中的最小值。

舉例來說，以下是一個Min Heap：

```
          3
       /     \
     12       5
    /  \     / \
   20  11   7   14
```

在這個Min Heap中，根節點的值為3，並且根節點的兩個子節點的值都大於等於3，即12和5。同樣地，12和5這兩個節點也滿足節點值小於或等於它們的子節點的值。

另外，Min Heap的一個重要特點是，當我們從Min Heap中刪除根節點時，Heap結構會自動調整，使得新的根節點還是整個Heap中的最小值。具體來說，我們可以先將最後一個節點複製到根節點的位置，然後不斷地比較新的根節點和它的兩個子節點的值，如果它的值大於其中某一個子節點的值，就交換它們的位置，直到找到了合適的位置為止。這樣一來，我們就可以在O(log n)的時間內刪除Min Heap中的最小值了。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Min Heap的重點
{{< /ask_chatgpt >}}



1. 它是一种堆数据结构，使用数组表示，并保持一定的顺序特征。
2. 堆通常分为两种类型：最大堆和最小堆。
3. 最小堆的根节点是具有最小值的节点，并且该属性在子树中递归地保持。
4. 最小堆的父节点始终小于或等于其子节点，而左子节点首先小于右子节点。
5. 插入新值时，可以将新值放到数组末尾，然后通过交换父节点和子节点的值来重新平衡堆。
6. 取出最小值时，从最小堆的根节点返回最小值，并从数组末尾将新值放入根节点位置，然后通过交换父节点和子节点的值来重新平衡堆。
7. 用于堆排序，优先队列等场合。
8. 时间复杂度为O(log n)，其中n是堆的大小。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Min Heap的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



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

