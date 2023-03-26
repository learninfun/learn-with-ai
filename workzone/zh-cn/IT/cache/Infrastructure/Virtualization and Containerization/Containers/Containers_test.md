

1. 给定一个由数字组成的二叉树，如何找到深度等于k的所有节点？
答案：使用深度优先搜索，记录当前节点的深度，并判断是否等于k。若等于k，则将该节点加入结果列表。继续遍历左子树和右子树。时间复杂度为O(n)。

2. 给定一个无序数组，如何使用单调栈来找到每个元素的下一个更大元素？
答案：使用单调递减的栈，并遍历数组。如果当前元素大于栈顶元素，则将栈顶元素弹出，并将当前元素作为栈顶元素的下一个更大元素。将当前元素入栈。时间复杂度为O(n)。

3. 给定一个字符串，找到其中最长的没有重复字符的连续子串长度？
答案：使用双指针，一个指向子串的起始位置，另一个不断向右移动。使用set记录子串中出现过的字符，如果发现子串中出现重复字符，就将左指针向右移动，并且从set中删除对应的字符。时间复杂度为O(n)。

4. 给定两个字符串s和t，判断t是否为s的一个子序列？
答案：使用双指针，一个指向字符串s的开头，另一个指向字符串t的开头。如果当前字符匹配，则都向右移动一位；否则只移动s指针。如果t指针到达末尾，说明t是s的子序列。时间复杂度为O(m+n)。

5. 给定一组任务和冷却时间n，每个任务执行需要一个单位的时间，相同任务之间需要等待n个单位的时间才能再次执行。计算完成所有任务所需的最短时间。
答案：使用桶排序，先统计每个任务出现的次数。将任务按照出现次数从大到小排序。将任务依次放入执行序列中，并在相应位置加入等待时间。时间复杂度为O(n)。