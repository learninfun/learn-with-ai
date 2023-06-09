+++
title = "内插搜寻"
weight = "3"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Interpolation Search的中文问题
{{< /ask_chatgpt >}}



1. 什么是Interpolation Search？
2. Interpolation Search的时间与空间复杂度是多少？
3. Interpolation Search在哪些情况下效益比较好？
4. 如何实作Interpolation Search？
5. Interpolation Search与Binary Search的区别是什么？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Interpolation Search并举例
{{< /ask_chatgpt >}}



Interpolation Search 是一种用于搜索有序数组中的元素的算法。其基本思想是根据要查找的元素在数组中可能出现的位置来进行查找，因为在有序数组中，较大的元素通常出现在数组的较后面，而较小的元素则出现在数组的较前面。因此，在找到一个比要查找的元素相近但略大的元素后，可以将搜索范围缩小到这个元素的前面。

以下是Interpolation Search的实现过程：

1. 假设有一个有序数组arr和要查找的元素x。
2. 初始化两个变量low和high，分别为数组的第一个元素和最后一个元素的下标。
3. 计算mid作为索引的估计值。mid = low + (high - low) * (x - arr[low]) / (arr[high] - arr[low])
4. 如果arr[mid]等于要查找的元素x，则返回mid。
5. 如果arr[mid]小于要查找的元素x，则要查找的元素x可能在mid的右边。因此更新low为mid + 1。
6. 如果arr[mid]大于要查找的元素x，则要查找的元素x可能出现在mid的左边。因此更新high为mid - 1。
7. 重复步骤3到6，直到找到目标元素或搜索区间缩小为空。

以下是Interpolation Search的一个例子：

假设有一个有序数组arr = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30]，要查找的元素x = 12。

初始时，low = 0，high = 14，mid = 0 + (14 - 0) * (12 - 2) / (30 - 2) = 4。此时，arr[4] = 10，小于要查找的元素x，因此更新low为5。

接下来，mid = 5 + (14 - 5) * (12 - 10) / (30 - 10) = 7。此时，arr[7] = 16，大于要查找的元素x，因此更新high为6。

然后，mid = 5 + (6 - 5) * (12 - 10) / (14 - 10) = 5。此时，arr[5] = 12，与要查找的元素相等，因此返回索引值5。

因此，Interpolation Search在此示例中成功找到了要查找的元素。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列Interpolation Search的重点
{{< /ask_chatgpt >}}



1. 描述：Interpolation Search是一种用于查找有序且均匀分布的数组中特定元素的搜索算法。

2. 工作原理：Interpolation Search基于线性搜索和二分搜索的想法。它通过估计元素的位置并缩小搜索范围来寻找特定元素。

3. 适用范围：Interpolation Search仅适用于有序且均匀分布的数组。

4. 时间复杂度：Interpolation Search的时间复杂度取决于数组中元素的分布。在最糟糕的情况下，它的时间复杂度为O(n)，在最好情况下，它的时间复杂度为O(1)。

5. 特点：Interpolation Search通常比二分搜索更快，因为它可以在平均O(loglog n)时间内找到元素，而二分搜索需要O(log n)时间。

6. 缺点：当数组中元素分布不均匀时，Interpolation Search的效率可能会受到影响，导致搜索时间更长。此外，它也不像二分搜索那么容易实现。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Interpolation Search的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 假设有一个已排序好的阵列，它的值介于0到100之间，每个值都唯一。使用插值搜寻，找到25这个值的索引位置。
答案: 假设阵列名为arr ，索引位置为index。
arr = [0, 5, 10, 20, 25, 30, 50, 75, 80, 90, 95, 100]
start_index = 0 
end_index = len(arr)-1
while start_index <= end_index:

  range_diff = arr[end_index] - arr[start_index]

  if range_diff == 0:
    mid_index = start_index 
  else:
    position = (25 - arr[start_index])/range_diff
    mid_index = int(start_index + (end_index - start_index) * position)

  if arr[mid_index] == 25:
    index = mid_index
    break
  elif arr[mid_index] < 25:
    start_index = mid_index + 1
  else:
    end_index = mid_index - 1
    
print(index)


2. 假设有一个已排序好的阵列，它的值介于0到100之间，每个值都唯一。使用插值搜寻，找到50这个值的索引位置。
答案: 假设阵列名为arr ，索引位置为index。
arr = [0, 5, 10, 20, 25, 30, 50, 75, 80, 90, 95, 100]
start_index = 0 
end_index = len(arr)-1
while start_index <= end_index:

  range_diff = arr[end_index] - arr[start_index]

  if range_diff == 0:
    mid_index = start_index 
  else:
    position = (50 - arr[start_index])/range_diff
    mid_index = int(start_index + (end_index - start_index) * position)

  if arr[mid_index] == 50:
    index = mid_index
    break
  elif arr[mid_index] < 50:
    start_index = mid_index + 1
  else:
    end_index = mid_index - 1
    
print(index)


3. 假设有一个已排序好的阵列，它的值介于0到100之间，每个值都唯一。使用插值搜寻，找到80这个值的索引位置。
答案: 假设阵列名为arr ，索引位置为index。
arr = [0, 5, 10, 20, 25, 30, 50, 75, 80, 90, 95, 100]
start_index = 0 
end_index = len(arr)-1
while start_index <= end_index:

  range_diff = arr[end_index] - arr[start_index]

  if range_diff == 0:
    mid_index = start_index 
  else:
    position = (80 - arr[start_index])/range_diff
    mid_index = int(start_index + (end_index - start_index) * position)

  if arr[mid_index] == 80:
    index = mid_index
    break
  elif arr[mid_index] < 80:
    start_index = mid_index + 1
  else:
    end_index = mid_index - 1
    
print(index)


4. 假设有一个已排序好的阵列，它的值介于0到100之间，每个值都唯一。使用插值搜寻，找到95这个值的索引位置。
答案: 假设阵列名为arr ，索引位置为index。
arr = [0, 5, 10, 20, 25, 30, 50, 75, 80, 90, 95, 100]
start_index = 0 
end_index = len(arr)-1
while start_index <= end_index:

  range_diff = arr[end_index] - arr[start_index]

  if range_diff == 0:
    mid_index = start_index 
  else:
    position = (95 - arr[start_index])/range_diff
    mid_index = int(start_index + (end_index - start_index) * position)

  if arr[mid_index] == 95:
    index = mid_index
    break
  elif arr[mid_index] < 95:
    start_index = mid_index + 1
  else:
    end_index = mid_index - 1
    
print(index)


5. 假设有一个已排序好的阵列，它的值介于0到100之间，每个值都唯一。使用插值搜寻，找到10这个值的索引位置。
答案: 假设阵列名为arr ，索引位置为index。
arr = [0, 5, 10, 20, 25, 30, 50, 75, 80, 90, 95, 100]
start_index = 0 
end_index = len(arr)-1
while start_index <= end_index:

  range_diff = arr[end_index] - arr[start_index]

  if range_diff == 0:
    mid_index = start_index 
  else:
    position = (10 - arr[start_index])/range_diff
    mid_index = int(start_index + (end_index - start_index) * position)

  if arr[mid_index] == 10:
    index = mid_index
    break
  elif arr[mid_index] < 10:
    start_index = mid_index + 1
  else:
    end_index = mid_index - 1
    
print(index)   

