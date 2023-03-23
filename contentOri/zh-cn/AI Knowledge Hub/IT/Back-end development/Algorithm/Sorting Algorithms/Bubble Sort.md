+++
title = "气泡排序"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Bubble Sort的问题
{{< /ask_chatgpt >}}



1. 什麼是 Bubble Sort ？ 它是如何工作的？

2. Bubble Sort 的時間複雜度是多少？它是如何算出來的？

3. 在實際應用中，Bubble Sort 有哪些限制？它在哪些情況下不適用？

4. 如何優化 Bubble Sort 的性能？有哪些可以提高效率的技巧？

5. Bubble Sort 和其他常見排序算法的比較。Bubble Sort 有哪些優點和缺點？相比之下，其他算法又有哪些不同之處？   

## 说明知识
{{< ask_chatgpt >}}
说明Bubble Sort并举例
{{< /ask_chatgpt >}}



Bubble sort是一種簡單的排序演算法，它重複地遍歷要排序的數列，比較相鄰的兩個數的大小，如果順序錯誤就進行交換。重複遍歷，直到沒有需要交換的元素為止，這時數列排序完成。

舉例來說，假設我們要對下列數列進行排序：[5, 8, 2, 1, 9]

Step 1：從數列的開頭開始比較相鄰的兩個數，5和8比較，順序正確，不交換；接著8和2比較，順序錯誤，進行交換。9不需要比較，因為已經是最後一個數，完成一次遍歷。目前數列為[5, 2, 8, 1, 9]

Step 2：從數列的開頭開始比較相鄰的兩個數，5和2比較，順序錯誤，進行交換。8和1比較，順序錯誤，進行交換。目前數列為[2, 5, 1, 8, 9]

Step 3：從數列的開頭開始比較相鄰的兩個數，2和5比較，順序正確，不交換。5和1比較，順序錯誤，進行交換。8和9比較，順序正確，不交換。目前數列為[2, 1, 5, 8, 9]

Step 4：從數列的開頭開始比較相鄰的兩個數，2和1比較，順序錯誤，進行交換。5和8比較，順序正確，不交換。8和9比較，順序正確，不交換。目前數列為[1, 2, 5, 8, 9]

現在數列已經排好序了，並且在第四次遍歷時沒有任何交換操作，所以排序完成。   

## 汇总重点
{{< ask_chatgpt >}}
条列Bubble Sort的重点
{{< /ask_chatgpt >}}



- Bubble Sort是一种基本的排序演算法，也是最简单的排序算法之一。
- 算法的核心概念是比较相邻的元素，如果它们的顺序错误就交换位置。
- 算法的运作过程是扫瞄整个待排序的序列，不断进行相邻元素的比较与交换操作，直到没有任何一对元素需要交换为止。
- Bubble Sort的时间复杂度为O(n^2)，效率较差。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Bubble Sort的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 将阵列中的偶数值递增排序，而奇数值则保持在原地。

```python
arr = [3, 4, 9, 1, 5, 2, 8, 7, 6]

for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if arr[j] % 2 == 0 and arr[j+1] % 2 == 0 and arr[j] > arr[j+1]:
            arr[j], arr[j+1] = arr[j+1], arr[j]
            
print(arr)
# Output: [3, 2, 4, 1, 5, 6, 8, 7, 9]
```

2. 将二维阵列按照其第二行递增排序。

```python
arr = [[3, 7], [9, 1], [5, 6], [2, 8], [4, 0]]

for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if arr[j][1] > arr[j+1][1]:
            arr[j], arr[j+1] = arr[j+1], arr[j]
        
print(arr)
# Output: [[4, 0], [9, 1], [5, 6], [3, 7], [2, 8]]
```

3. 将字串阵列按照字典顺序递减排序。

```python
arr = ["cat", "dog", "bird", "apple", "bug"]

for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if arr[j] < arr[j+1]:
            arr[j], arr[j+1] = arr[j+1], arr[j]
        
print(arr)
# Output: ['dog', 'cat', 'bug', 'bird', 'apple']
```

4. 找出阵列中第二小的元素。

```python
arr = [3, 4, 9, 1, 5, 2, 8, 7, 6]

for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if arr[j] > arr[j+1]:
            arr[j], arr[j+1] = arr[j+1], arr[j]
        
print(arr[1])
# Output: 2
```

5. 判断是否存在阵列中的任意连续子段，其元素均为递增序列。

```python
arr = [3, 4, 9, 1, 5, 2, 8, 7, 6]

for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if arr[j] > arr[j+1]:
            arr[j], arr[j+1] = arr[j+1], arr[j]

for i in range(len(arr)-1):
    if arr[i] < arr[i+1]:
        for j in range(i+1, len(arr)-1):
            if arr[j] > arr[j+1]:
                break
        else:
            print("True")
            break
else:
    print("False")
# Output: True
```   

