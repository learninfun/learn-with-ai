+++
title = "環形佇列"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Circular Queue的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Circular Queue？

2. Circular Queue和普通Queue的差異是什麼？

3. 如何實現一個Circular Queue？

4. Circular Queue在什麼場景中使用比較合適？

5. 如何保證Circular Queue的正確性和性能？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Circular Queue並舉例
{{< /ask_chatgpt >}}



Circular Queue 是一種在介面上很像 Queue 的資料結構，但與一般的 Queue 不同的是，它允許在 Queue 的尾端插入元素，同時在 Queue 的頭端也同時可以移除元素，因此當 Queue 已塞滿元素時，它能夠 "循環 (circular)" 回到 Queue 的初始位置重新插入元素。

使用 Circular Queue 的優點是，當我們使用陣列來實現 Circular Queue 時，可以使用頭指標與尾指標來維護 Queue 的運行。同時，因為每當尾指標到達數組的尾端時，它都會重新回到陣列的起點，因此能夠節省空間的浪費。

舉例來說，如果我們有一個大小為 5 的 Circular Queue，它的運作方式可能如下：
- 開始時，head 與 tail 指標都指向位置 0。
- 插入元素 1，tail 指向位置 1。
- 插入元素 2，tail 指向位置 2。
- 插入元素 3，tail 指向位置 3。
- 移除 Queue 头元素，head 指向位置 1。
- 插入元素 4，tail 指向位置 4。
- 移除 Queue 头元素，head 指向位置 2。
- 移除 Queue 头元素，head 指向位置 3。
- 插入元素 5，tail 回到位置 0。

因此，從上面的例子中，我們可以看到，head 與 tail 指標就像一個圈一樣，當它們到達數組的尾端時，就會回到數組的頭部重新開始。這樣就能夠支持 Queue 的循環運行了。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Circular Queue的重點
{{< /ask_chatgpt >}}



1. 循環性: Circular Queue是一種循環型資料結構，即當佇列頭到達佇列尾部時，會返回佇列頭部開始循環。

2. 優點: Circular Queue具有比普通Queue更高的效率，因為它可以利用佇列未使用的空間去存儲更多的元素。

3. 結構: Circular Queue有一個陣列buffer，在這個陣列中存儲元素，有一個front指針指向佇列頭，有一個rear指針指向佇列尾。

4. 操作: Circular Queue常用操作包括入佇列(enqueue)、出佇列(dequeue)、查看佇列頭部元素(peek)、查看佇列是否為空(isEmpty)、查看佇列是否已滿(isFull)。

5. 注意事項: 當復位front和rear指針時，必須確定佇列已經为空，否則有可能造成資料遺失或存儲佇列溢出。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Circular Queue的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1.實現一個環形隊列，支持以下操作：
- 入隊
- 出隊
- 取隊首元素
- 檢查隊列是否已滿
- 檢查隊列是否為空

2.判斷一個給定的整數陣列nums能否通過環形旋轉得到另一個整數陣列circ。其中，環形旋轉表示把陣列最後一個元素放到最前面，其他元素向後移一位。
eg:
nums = [1, 2, 3, 4, 5]
circ = [4, 5, 1, 2, 3]
return True

3.設計一個支持環形緩存的線性LruCache。緩存有容量大小限制，當緩存滿了時，需要將最久未使用的元素移除。
eg:
cache = LruCache(3)
cache.set(1,1)
cache.set(2,2)
cache.set(3,3)
cache.get(1) # returns 1
cache.set(4,4)
cache.get(2) # returns -1

4.設計一個日期隊列，實現以下功能：
- 入隊一個日期
- 刪除最早的日期（即從隊首刪除）
- 取得最早日期（即取得隊首元素）
- 判斷日期隊列是否為空
- 判斷日期隊列是否已滿

5.實現一個環形隊列，支持環形迭代器。即可以從隊首迭代到隊尾，然後再從隊首開始迭代。   

