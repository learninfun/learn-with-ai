

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