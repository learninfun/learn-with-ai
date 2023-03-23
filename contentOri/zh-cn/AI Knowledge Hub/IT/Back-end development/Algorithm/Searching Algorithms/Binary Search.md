+++
title = "二分搜寻"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Binary Search的问题
{{< /ask_chatgpt >}}



1. 在一個排序好的整數數組中，查找一個特定的元素。
2. 在一個二維排序矩陣中，查找一個特定的元素。矩陣每行都按遞增順序排序，每列都按遞增順序排序。
3. 在升序排列的旋轉數組中，查找一個特定的元素。例如，{5, 6, 7, 8, 9, 10, 1, 2, 3}是一個旋轉了的排序數組，我們要查找元素3。
4. 查找旋轉排序數組中的最小元素。例如，在{5, 6, 7, 8, 9, 10, 1, 2, 3}中，最小元素是1。
5. 在多個排序數組中查找一個共同的元素。例如，在{2, 4, 6, 8, 10}和{1, 3, 4, 5, 7, 9, 10}這兩個排序數組中，查找共同的元素4。   

## 说明知识
{{< ask_chatgpt >}}
说明Binary Search并举例
{{< /ask_chatgpt >}}



二分搜尋（Binary Search）是一種在排好序的數列中搜尋目標值的演算法，也被稱為折半搜尋。其基本思路為：

1. 以中間值為基準，將搜索範圍分為左右兩半。

2. 判斷目標值位於哪一半，然後繼續在該半部分中進行搜尋。

3. 如果中間值就是目標值，那麼搜索結束，如果搜尋不到目標值，那麼搜索範圍會縮小到最小，並且通過返回值表示目標值是否存在。

以下是一個簡單的二分搜尋算法的實現。

```
public static int binarySearch(int[] nums, int target) {
    int left = 0, right = nums.length - 1;
    while (left <= right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] == target) {
            return mid;
        } else if (nums[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return -1;
}
```

例：在一個有序的整數數組中查找目標值，如需查找的數組為【1,4,6,12,14,17,22,25,26,30】，目標值為17。那麼首先將搜索範圍定為整個數組，查找中間位置mid的數值14小於目標值17，因此縮小搜索範圍至右半部分，再次查找中間位置mid的數值22，仍然比目標值大，繼續縮小搜索範圍至右半部分，至最後查找中間位置mid為目標值17，搜索成功，返回位置值5。   

## 汇总重点
{{< ask_chatgpt >}}
条列Binary Search的重点
{{< /ask_chatgpt >}}



1. Binary Search是一種搜索算法，通常用於有序數組或列表中的元素查找。
2. 該算法的基本思想是將數組或列表分為兩半並查找中間元素，如果該元素等於目標元素，則返回其索引，否則繼續搜索相應的區域。
3. 在每個步驟中，都可以將搜索範圍大小縮小一半，因此，其時間複雜度是O(log n)。
4. Binary Search要求數組或列表必須是有序的，這樣才能進行區間比較。
5. 該算法的實現可以使用遞歸或迭代的方式，具體實現方法取決於像語言和代碼習慣等因素。
6. Binary Search還可應用於查找插入位置，以及在數組中查找出現次數等操作。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Binary Search的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 在一個已排序的整數數組中，查找第一個大於等於目標值的數字，如果目標值不存在，則返回-1。例如：arr = [1，3，4，4，4，6，7]，目標值為5，則返回6.

答案： https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

2. 你有一個正整數數組nums，請找出一個索引所在位置，使得該位置左邊所有數字的和等於該位置右邊所有數字的和。如果這樣的位置不存在，則返回-1。例如：nums = [1，7，3，6，5，6]，則返回3.

答案： https://leetcode.com/problems/find-pivot-index/

3. 在一個由n個元素的旋轉排序數組中查找目標值。旋轉排序數組是指通過將一個有序數組的一部分旋轉而得到的。例如，數組[0，1，2，4，5，6，7]可能會變為[4，5，6，7，0，1，2]。注意，這個數組不包括重複值。如果目標值在數組中，則返回它的索引，否則返回-1。例如：nums = [4，5，6，7，0，1，2]，target = 0，則返回4.

答案： https://leetcode.com/problems/search-in-rotated-sorted-array/

4. 在一个大小为n的排序数组中寻找缺失的数字。假设数组中每个元素都是独一无二的并且在给定范围内。例如：nums = [0，1，3]，期望输出为2。

答案： https://leetcode.com/problems/missing-number/

5. 在一个排序并且可能包含重复元素的数组nums中，找到target第一次和最后一次出现的位置。如果在数组中不存在 target ，则返回[-1, -1]。例如：nums = [5，7，7，8，8，10]，target = 8，期望输出为[3, 4]。

答案： https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/   

