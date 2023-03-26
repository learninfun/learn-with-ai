+++
title = "氣泡排序"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Bubble Sort的中文問題
{{< /ask_chatgpt >}}



1. 什麼是 Bubble Sort ？ 它是如何工作的？

2. Bubble Sort 的時間複雜度是多少？它是如何算出來的？

3. 在實際應用中，Bubble Sort 有哪些限制？它在哪些情況下不適用？

4. 如何優化 Bubble Sort 的性能？有哪些可以提高效率的技巧？

5. Bubble Sort 和其他常見排序算法的比較。Bubble Sort 有哪些優點和缺點？相比之下，其他算法又有哪些不同之處？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Bubble Sort並舉例
{{< /ask_chatgpt >}}



Bubble sort是一種簡單的排序演算法，它重複地遍歷要排序的數列，比較相鄰的兩個數的大小，如果順序錯誤就進行交換。重複遍歷，直到沒有需要交換的元素為止，這時數列排序完成。

舉例來說，假設我們要對下列數列進行排序：[5, 8, 2, 1, 9]

Step 1：從數列的開頭開始比較相鄰的兩個數，5和8比較，順序正確，不交換；接著8和2比較，順序錯誤，進行交換。9不需要比較，因為已經是最後一個數，完成一次遍歷。目前數列為[5, 2, 8, 1, 9]

Step 2：從數列的開頭開始比較相鄰的兩個數，5和2比較，順序錯誤，進行交換。8和1比較，順序錯誤，進行交換。目前數列為[2, 5, 1, 8, 9]

Step 3：從數列的開頭開始比較相鄰的兩個數，2和5比較，順序正確，不交換。5和1比較，順序錯誤，進行交換。8和9比較，順序正確，不交換。目前數列為[2, 1, 5, 8, 9]

Step 4：從數列的開頭開始比較相鄰的兩個數，2和1比較，順序錯誤，進行交換。5和8比較，順序正確，不交換。8和9比較，順序正確，不交換。目前數列為[1, 2, 5, 8, 9]

現在數列已經排好序了，並且在第四次遍歷時沒有任何交換操作，所以排序完成。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Bubble Sort的重點
{{< /ask_chatgpt >}}



- Bubble Sort是一種基本的排序演算法，也是最簡單的排序算法之一。
- 算法的核心概念是比較相鄰的元素，如果它們的順序錯誤就交換位置。
- 算法的運作過程是掃瞄整個待排序的序列，不斷進行相鄰元素的比較與交換操作，直到沒有任何一對元素需要交換為止。
- Bubble Sort的時間複雜度為O(n^2)，效率較差。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Bubble Sort的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 將陣列中的偶數值遞增排序，而奇數值則保持在原地。

```python
arr = [3, 4, 9, 1, 5, 2, 8, 7, 6]

for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if arr[j] % 2 == 0 and arr[j+1] % 2 == 0 and arr[j] > arr[j+1]:
            arr[j], arr[j+1] = arr[j+1], arr[j]
            
print(arr)
# Output: [3, 2, 4, 1, 5, 6, 8, 7, 9]
```

2. 將二維陣列按照其第二行遞增排序。

```python
arr = [[3, 7], [9, 1], [5, 6], [2, 8], [4, 0]]

for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if arr[j][1] > arr[j+1][1]:
            arr[j], arr[j+1] = arr[j+1], arr[j]
        
print(arr)
# Output: [[4, 0], [9, 1], [5, 6], [3, 7], [2, 8]]
```

3. 將字串陣列按照字典順序遞減排序。

```python
arr = ["cat", "dog", "bird", "apple", "bug"]

for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if arr[j] < arr[j+1]:
            arr[j], arr[j+1] = arr[j+1], arr[j]
        
print(arr)
# Output: ['dog', 'cat', 'bug', 'bird', 'apple']
```

4. 找出陣列中第二小的元素。

```python
arr = [3, 4, 9, 1, 5, 2, 8, 7, 6]

for i in range(len(arr)):
    for j in range(len(arr)-i-1):
        if arr[j] > arr[j+1]:
            arr[j], arr[j+1] = arr[j+1], arr[j]
        
print(arr[1])
# Output: 2
```

5. 判斷是否存在陣列中的任意連續子段，其元素均為遞增序列。

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

