

1. Reverse Linked List II
2. Copy List with Random Pointer
3. Remove Nth Node From End of List
4. Partition List
5. Linked List Cycle II

1. Reverse Linked List II:
题目描述： 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
示例: 
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
答案链接: https://leetcode.com/problems/reverse-linked-list-ii/

2. Copy List with Random Pointer:
题目描述： 给定一个链表，每个节点包含一个额外增加的随机指针，
该指针可以指向链表中的任何节点或空节点。
要求返回这个链表的 深拷贝。 
示例：
输入：
{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
解释：
节点 1 的值是 1，它的下一个指针和随机指针都指向节点 2 。
节点 2 的值是 2，它的下一个指针指向 null，随机指针指向它本身。
答案链接：https://leetcode.com/problems/copy-list-with-random-pointer/

3. Remove Nth Node From End of List:
题目描述： 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。 
示例：
输入: 1->2->3->4->5, n = 2
输出: 1->2->3->5
答案链接： https://leetcode.com/problems/remove-nth-node-from-end-of-list/

4. Partition List:
题目描述： 给定一个链表和一个特定值 x，对链表进行分隔，
使得所有小于 x 的节点都在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。
示例：
输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5
答案链接： https://leetcode.com/problems/partition-list/

5. Linked List Cycle II:
题目描述： 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
示例：
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
答案链接： https://leetcode.com/problems/linked-list-cycle-ii/