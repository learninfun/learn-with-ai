

1. Reverse Linked List II
2. Copy List with Random Pointer
3. Remove Nth Node From End of List
4. Partition List
5. Linked List Cycle II

1. Reverse Linked List II:
題目描述： 反轉從位置 m 到 n 的鏈表。請使用一趟掃瞄完成反轉。
示例: 
輸入: 1->2->3->4->5->NULL, m = 2, n = 4
輸出: 1->4->3->2->5->NULL
答案鏈接: https://leetcode.com/problems/reverse-linked-list-ii/

2. Copy List with Random Pointer:
題目描述： 給定一個鏈表，每個節點包含一個額外增加的隨機指針，
該指針可以指向鏈表中的任何節點或空節點。
要求返回這個鏈表的 深拷貝。 
示例：
輸入：
{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
解釋：
節點 1 的值是 1，它的下一個指針和隨機指針都指向節點 2 。
節點 2 的值是 2，它的下一個指針指向 null，隨機指針指向它本身。
答案鏈接：https://leetcode.com/problems/copy-list-with-random-pointer/

3. Remove Nth Node From End of List:
題目描述： 給定一個鏈表，刪除鏈表的倒數第 n 個節點，並且返回鏈表的頭結點。 
示例：
輸入: 1->2->3->4->5, n = 2
輸出: 1->2->3->5
答案鏈接： https://leetcode.com/problems/remove-nth-node-from-end-of-list/

4. Partition List:
題目描述： 給定一個鏈表和一個特定值 x，對鏈表進行分隔，
使得所有小於 x 的節點都在大於或等於 x 的節點之前。
你應當保留兩個分區中每個節點的初始相對位置。
示例：
輸入: head = 1->4->3->2->5->2, x = 3
輸出: 1->2->2->4->3->5
答案鏈接： https://leetcode.com/problems/partition-list/

5. Linked List Cycle II:
題目描述： 給定一個鏈表，返回鏈表開始入環的第一個節點。 如果鏈表無環，則返回 null。
示例：
輸入：head = [3,2,0,-4], pos = 1
輸出：tail connects to node index 1
答案鏈接： https://leetcode.com/problems/linked-list-cycle-ii/