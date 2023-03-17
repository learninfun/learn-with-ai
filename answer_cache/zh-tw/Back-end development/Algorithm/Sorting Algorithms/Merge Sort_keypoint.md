

1. Divide and conquer: Merge Sort使用“分而治之”的策略，将一个大问题分解成较小的子问题，并且解决每个子问题，最后将它们合并成一个有序的问题。

2. Recursion: Merge Sort使用递归的形式实现分治法，每递归一层都会分解原问题并且排序子数组。

3. Merge: Merge Sort将两个已排序的子数组合并为一个已排序的数组，使得最终的数组保持正确的顺序。

4. Time complexity: Merge Sort的时间复杂度为nlogn，是一种高效的排序算法。

5. Space complexity: Merge Sort的空间复杂度是O(n)，因为必须在排序过程中创建一个与输入数组大小相同的临时数组。

6. Stability: Merge Sort是稳定的，因为它使用了相等的情况的规则，以保持同等大小的数字的相对位置。

7. Adaptive: Merge Sort也是自适应的，可以在原地排序的情况下使用。

8. Parallel: Merge Sort还有一个显著的特性，就是它很容易实现并发排序，因为它将大数组分成许多小的子数组进行排序，并且这些子数组可以并发处理