

1. 假設有一個已排序好的陣列，它的值介於0到100之間，每個值都唯一。使用插值搜尋，找到25這個值的索引位置。
答案: 假設陣列名為arr ，索引位置為index。
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


2. 假設有一個已排序好的陣列，它的值介於0到100之間，每個值都唯一。使用插值搜尋，找到50這個值的索引位置。
答案: 假設陣列名為arr ，索引位置為index。
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


3. 假設有一個已排序好的陣列，它的值介於0到100之間，每個值都唯一。使用插值搜尋，找到80這個值的索引位置。
答案: 假設陣列名為arr ，索引位置為index。
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


4. 假設有一個已排序好的陣列，它的值介於0到100之間，每個值都唯一。使用插值搜尋，找到95這個值的索引位置。
答案: 假設陣列名為arr ，索引位置為index。
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


5. 假設有一個已排序好的陣列，它的值介於0到100之間，每個值都唯一。使用插值搜尋，找到10這個值的索引位置。
答案: 假設陣列名為arr ，索引位置為index。
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