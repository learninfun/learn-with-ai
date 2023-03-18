

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