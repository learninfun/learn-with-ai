+++
title = "线性数据结构"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Linear Data Structures的问题
{{< /ask_chatgpt >}}



1. 请问何谓Linear Data Structures？其特色是什么？
2. 请问Stack和Queue分别是哪种Linear Data Structures？它们的差异在哪里？
3. 请问Linked List、Array和Vector都是哪种Linear Data Structures？它们的优缺点分别是什么？
4. 请问Hash Table和Heap分别是哪种Linear Data Structures？它们的特色和使用情境分别是什么？
5. 描述一下Binary Search Tree的特性及其常见操作。   

## 说明知识
{{< ask_chatgpt >}}
说明Linear Data Structures并举例
{{< /ask_chatgpt >}}



线性资料结构是指其资料元素按照一定的顺序排列，并且每个资料元素都只有一个前驱元素（第一个资料元素除外）和一个后继元素（最后一个资料元素除外），即资料元素之间存在一对一的前后关系。常见的线性资料结构包括阵列，链表，伫列，堆栈等。

以下是几种常见的线性资料结构：

1. 阵列：阵列是一个在内存中分配连续记忆体的资料结构，它通过索引来访问和操作元素。阵列的特点是能够快速访问元素，但在插入和删除操作时需要移动阵列中其他元素的位置。例如，int nums [5] = {1, 2, 3, 4, 5}。

2. 链表：链表是由节点组成的资料结构，每个节点包含资料和指向下一个节点的指针。链表的特点是在插入和删除记录时不需要移动其他元素位置，但是访问元素时需要遍历整个链表。例如，单链表、双向链表、循环链表等。

3. 伫列：伫列是具有先进先出（FIFO）特点的资料结构，类似于排队。在伫列的一端添加元素，在另一端删除元素。例如，等待列。

4. 堆栈：堆栈是具有后进先出（LIFO）特点的资料结构，类似于一叠盘子，最后放上去的最先被取下来。在堆栈中添加元素的操作称为推入(push)，删除元素的操作称为弹出(pop)。例如，浏览器的后退按钮。   

## 汇总重点
{{< ask_chatgpt >}}
条列Linear Data Structures的重点
{{< /ask_chatgpt >}}



1. 线性结构：线性数据结构是数据元素按线性次序排列的结构，具有先后顺序，且每个元素之间均具有一对一的直接前驱和直接后继关系。

2. 数组Array：数组是一种最常见的线性数据结构，它可以连续存储多个相同类型的数据，并可以通过下标快速访问和修改数据。

3. 链表Linked List：链表是一种线性数据结构，它由一个个节点按顺序连接而成，每个节点包含数据域和指向下一个节点的指针。

4. 堆栈Stack：堆栈是一种先进后出（LIFO）的线性数据结构，它支持push和pop操作，用于处理临时性数据。

5. 队列Queue：队列是一种先进先出（FIFO）的线性数据结构，它支持enqueue和dequeue操作，用于实现消息伫列等场景。

6. 双向链表Doubly Linked List： 双向链表在链表的基础上增加了prev指针，实现快速遍历以及反向遍历，但同时也带来了额外的空间和时间复杂度开销。

7. 循环链表Circular Linked List： 循环链表是一种特殊的链表，尾节点的next指向头节点，形成一个循环。

8. 阵列堆积Array Heap： 堆是一种特殊的数据结构，它是一个完全二叉树，且每个节点都大于或等于（或小于或等于）其子节点，常用于实现优先级伫列等场景。

9. 优先级伫列Priority Queue： 优先级伫列是一种特殊的队列，每个元素都有一个优先级，优先级高的元素先出队列。

10. 哈希表Hash Table： 哈希表是一种基于散列表实现的映射结构，它能够实现高效的数据查询和修改操作。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Linear Data Structures的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 将阵列中的元素循环向右移动k位。例如，当阵列为[1,2,3,4,5,6]且k=2时，阵列变为[5,6,1,2,3,4]。
答案：https://leetcode.com/problems/rotate-array/
2. 实现一个简单的LIFO栈，支持推入、弹出和取顶操作。
答案：https://stackoverflow.com/questions/19141648/how-to-implement-a-stack-in-python
3. 找到阵列中缺少的最小正整数。例如，当阵列为[1,2,0]时，缺少的最小正整数为3。
答案：https://leetcode.com/problems/first-missing-positive/
4. 计算阵列中每个元素左侧和右侧的最大值之差的绝对值的和。例如，当阵列为[2,1,3,4,1]时，计算结果为8。左侧最大值为[0,2,2,3,3]，右侧最大值为[4,4,4,4,5]。
答案：https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/
5. 实现一个字典数据结构，支持插入、查询和删除操作。
答案：https://github.com/grantjenks/python-diskcached   

