

Linear Search是一种基本的搜索算法，也叫顺序查找。它从序列的开头开始逐个查找，直到找到目标元素或者查找完整个序列为止。

例如，假设我们要查找一个数字3是否在一个给定的数组中，可以按以下过程实现线性搜索：

1.从数组的第一个元素开始遍历，即下标为0。

2.比较当前元素是否等于目标元素3。

3.如果相等，即找到目标元素，返回当前元素的下标。

4.如果不相等，继续遍历下一个元素。

5.如果全部元素遍历完后仍未找到目标元素，返回-1，表示未找到。

举例来说，假设给定一个数组arr={2,5,8,3,9,4}，要查找数字3是否在其中，那么按照以上过程，我们可以依次比较2、5、8、3，找到了目标元素3，返回它的下标3。

实际上，线性搜索算法可以应用于任何有序或无序的列表或数组中。但是，相对于其他高级的搜索算法，它的效率较低，因为它需要遍历整个序列才能找到目标元素。如果列表中有大量元素，这将会消耗大量的时间。