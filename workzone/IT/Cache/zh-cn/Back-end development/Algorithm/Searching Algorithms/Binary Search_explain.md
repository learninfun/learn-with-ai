

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