

1. 循環性: Circular Queue是一種循環型資料結構，即當佇列頭到達佇列尾部時，會返回佇列頭部開始循環。

2. 優點: Circular Queue具有比普通Queue更高的效率，因為它可以利用佇列未使用的空間去存儲更多的元素。

3. 結構: Circular Queue有一個陣列buffer，在這個陣列中存儲元素，有一個front指針指向佇列頭，有一個rear指針指向佇列尾。

4. 操作: Circular Queue常用操作包括入佇列(enqueue)、出佇列(dequeue)、查看佇列頭部元素(peek)、查看佇列是否為空(isEmpty)、查看佇列是否已滿(isFull)。

5. 注意事項: 當復位front和rear指針時，必須確定佇列已經为空，否則有可能造成資料遺失或存儲佇列溢出。