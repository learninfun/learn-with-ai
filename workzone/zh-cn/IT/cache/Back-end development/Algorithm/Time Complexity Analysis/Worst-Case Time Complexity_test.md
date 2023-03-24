

1. Bubble Sort (冒泡排序) - Worst-Case Time Complexity: O(n^2)

2. Insertion Sort (插入排序) - Worst-Case Time Complexity: O(n^2)

3. Selection Sort (选择排序) - Worst-Case Time Complexity: O(n^2)

4. Naive String Searching (字符串匹配搜索) - Worst-Case Time Complexity: O(m(n-m+1)), 其中m为模式串的长度，n为文本串的长度

5. Quick Sort (快速排序) - Worst-Case Time Complexity: O(n^2)

答案：

1. Bubble Sort: 按照从小到大的顺序排序数组[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

时间复杂度为n^2，其中n为数组长度，这里为10，最坏情况需要比较45次，移动45次

2. Insertion Sort: 按照从小到大的顺序排序数组[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

时间复杂度为n^2，其中n为数组长度，这里为10，最坏情况需要比较45次，移动45次

3. Selection Sort: 按照从小到大的顺序排序数组[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

时间复杂度为n^2，其中n为数组长度，这里为10，最坏情况需要比较45次，移动45次

4. Naive String Searching: 在字符串"ABABABABAB"中搜索模式串"ABAB", "ABAB"每个字符都匹配，但每次匹配都失配，时间复杂度为4*(10-4+1)=28

5. Quick Sort: 按照从小到大的顺序排序数组[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

序列每次选择第一个元素为枢轴，而序列本身是逆序的，所以每次划分都到了最差情况，时间复杂度为n^2，其中n为数组长度，这里为10，最坏情况需要比较45次，交换45次