

1. Linear Queue是一种线性的资料结构，具有先进先出的特性。
2. 它通常包含一个前端(front)指针和一个后端(rear)指针，用来指示首尾两端的位置。
3. 新元素在后端(rear)添加，旧元素在前端(front)删除。
4. 元素依次排列，插入和删除均按照先进先出(FIFO)的原则进行。
5. 当前端(front)和后端(rear)指针相等空间时，称为空队列，当尾指针加1等于伫列长度时，称为满队列。
6. 可以使用数组或链表实现Linear Queue，一般选择链表实现，具有弹性且节省空间。
7. 常见操作有：enqueue(添加元素)、dequeue(删除元素)、isEmpty(判断是否空队列)、isFull(判断是否满队列)等。