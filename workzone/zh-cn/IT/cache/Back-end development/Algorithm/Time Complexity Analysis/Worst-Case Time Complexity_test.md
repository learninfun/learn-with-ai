

1. Bubble Sort (冒泡排序) - Worst-Case Time Complexity: O(n^2)

2. Insertion Sort (插入排序) - Worst-Case Time Complexity: O(n^2)

3. Selection Sort (選擇排序) - Worst-Case Time Complexity: O(n^2)

4. Naive String Searching (字符串匹配搜索) - Worst-Case Time Complexity: O(m(n-m+1)), 其中m為模式串的長度，n為文本串的長度

5. Quick Sort (快速排序) - Worst-Case Time Complexity: O(n^2)

答案：

1. Bubble Sort: 按照從小到大的順序排序數組[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

時間複雜度為n^2，其中n為數組長度，這裡為10，最壞情況需要比較45次，移動45次

2. Insertion Sort: 按照從小到大的順序排序數組[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

時間複雜度為n^2，其中n為數組長度，這裡為10，最壞情況需要比較45次，移動45次

3. Selection Sort: 按照從小到大的順序排序數組[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

時間複雜度為n^2，其中n為數組長度，這裡為10，最壞情況需要比較45次，移動45次

4. Naive String Searching: 在字符串"ABABABABAB"中搜索模式串"ABAB", "ABAB"每個字符都匹配，但每次匹配都失配，時間複雜度為4*(10-4+1)=28

5. Quick Sort: 按照從小到大的順序排序數組[10, 9, 8, 7, 6, 5, 4, 3, 2, 1]

序列每次選擇第一個元素為樞軸，而序列本身是逆序的，所以每次劃分都到了最差情況，時間複雜度為n^2，其中n為數組長度，這裡為10，最壞情況需要比較45次，交換45次