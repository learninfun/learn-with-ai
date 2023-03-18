+++
title = "环形伫列"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Circular Queue的问题
{{< /ask_chatgpt >}}



1. 什麼是Circular Queue？

2. Circular Queue和普通Queue的差異是什麼？

3. 如何實現一個Circular Queue？

4. Circular Queue在什麼場景中使用比較合適？

5. 如何保證Circular Queue的正確性和性能？   

## 说明知识
{{< ask_chatgpt >}}
说明Circular Queue并举例
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

## 汇总重点
{{< ask_chatgpt >}}
条列Circular Queue的重点
{{< /ask_chatgpt >}}



1. 循環性: Circular Queue是一種循環型資料結構，即當佇列頭到達佇列尾部時，會返回佇列頭部開始循環。

2. 優點: Circular Queue具有比普通Queue更高的效率，因為它可以利用佇列未使用的空間去存儲更多的元素。

3. 結構: Circular Queue有一個陣列buffer，在這個陣列中存儲元素，有一個front指針指向佇列頭，有一個rear指針指向佇列尾。

4. 操作: Circular Queue常用操作包括入佇列(enqueue)、出佇列(dequeue)、查看佇列頭部元素(peek)、查看佇列是否為空(isEmpty)、查看佇列是否已滿(isFull)。

5. 注意事項: 當復位front和rear指針時，必須確定佇列已經为空，否則有可能造成資料遺失或存儲佇列溢出。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Circular Queue的中等难度问题，并在后面列出答案
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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Circular Queue的网络数据
{{< /ask_chatgpt >}}



資料1：GeeksforGeeks - Circular Queue

網址：https://www.geeksforgeeks.org/circular-queue-set-1-introduction-array-implementation/

這篇文章從介紹Circular Queue的基本概念開始，說明了如何使用陣列實現Circular Queue，並介紹了相關的操作（如插入、刪除、檢查是否為空）。同時，該文章還討論了一些可能出現的問題，例如如何處理Queue滿了的情況。

資料2：Tutorialspoint - Circular Queue

網址：https://www.tutorialspoint.com/data_structures_algorithms/circular_queue_algorithm.htm

類似於GeeksforGeeks的文章，這篇教程也從介紹基本概念開始，接著提供了一些示例代碼。除此之外，該教程還介紹了使用循環缓存區（Circular Buffer）實現Circular Queue的方法，並討論了如何使用双端Queue（Deque）來解決一些特定的問題。

資料3：Programiz - Circular Queue

網址：https://www.programiz.com/dsa/circular-queue

Programiz提供了一個完整的教程，涵蓋Circular Queue的基本概念、如何實現、如何使用等。除了範例代碼外，該教程還提供了一些練習，讓讀者更好地理解所學知識。

資料4：Medium - Circular Queue

網址：https://medium.com/@sourabhv/understanding-circular-queue-implementation-in-java-90c9b9c8b31c

這篇Medium文章專注於Java中如何實現Circular Queue，並以代碼作為主要教學方法。該文章還討論了如何使用Java中的ArrayList來解決有限陣列帶來的問題。

資料5：Stack Abuse - Circular Queue

網址：https://stackabuse.com/stacks-and-queues-in-python/#circular-queue

這篇文章介紹了Python中如何實現Circular Queue。該文章旨在通過代碼示例來解釋概念，並提供了一些基本的操作。該篇文章還討論了如何解決一些可能出現的問題，例如在Queue滿了的情況下擴展陣列的方法。   

