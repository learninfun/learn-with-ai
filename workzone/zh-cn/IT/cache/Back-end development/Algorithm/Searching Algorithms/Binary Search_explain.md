

二分搜寻（Binary Search）是一种在排好序的数列中搜寻目标值的演算法，也被称为折半搜寻。其基本思路为：

1. 以中间值为基准，将搜索范围分为左右两半。

2. 判断目标值位于哪一半，然后继续在该半部分中进行搜寻。

3. 如果中间值就是目标值，那么搜索结束，如果搜寻不到目标值，那么搜索范围会缩小到最小，并且通过返回值表示目标值是否存在。

以下是一个简单的二分搜寻算法的实现。

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

例：在一个有序的整数数组中查找目标值，如需查找的数组为【1,4,6,12,14,17,22,25,26,30】，目标值为17。那么首先将搜索范围定为整个数组，查找中间位置mid的数值14小于目标值17，因此缩小搜索范围至右半部分，再次查找中间位置mid的数值22，仍然比目标值大，继续缩小搜索范围至右半部分，至最后查找中间位置mid为目标值17，搜索成功，返回位置值5。