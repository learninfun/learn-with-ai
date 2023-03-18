+++
title = "內插搜尋"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Interpolation Search的問題
{{< /ask_chatgpt >}}



1. 什麼是Interpolation Search？
2. Interpolation Search的時間與空間複雜度是多少？
3. Interpolation Search在哪些情況下效益比較好？
4. 如何實作Interpolation Search？
5. Interpolation Search與Binary Search的區別是什麼？   

## 說明知識
{{< ask_chatgpt >}}
說明Interpolation Search並舉例
{{< /ask_chatgpt >}}



Interpolation Search 是一種用於搜索有序數組中的元素的算法。其基本思想是根據要查找的元素在數組中可能出現的位置來進行查找，因為在有序數組中，較大的元素通常出現在數組的較後面，而較小的元素則出現在數組的較前面。因此，在找到一個比要查找的元素相近但略大的元素後，可以將搜索範圍縮小到這個元素的前面。

以下是Interpolation Search的實現過程：

1. 假設有一個有序數組arr和要查找的元素x。
2. 初始化兩個變量low和high，分別為數組的第一個元素和最後一個元素的下標。
3. 計算mid作為索引的估計值。mid = low + (high - low) * (x - arr[low]) / (arr[high] - arr[low])
4. 如果arr[mid]等於要查找的元素x，則返回mid。
5. 如果arr[mid]小於要查找的元素x，則要查找的元素x可能在mid的右邊。因此更新low為mid + 1。
6. 如果arr[mid]大於要查找的元素x，則要查找的元素x可能出現在mid的左邊。因此更新high為mid - 1。
7. 重複步驟3到6，直到找到目標元素或搜索區間縮小為空。

以下是Interpolation Search的一個例子：

假設有一個有序數組arr = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30]，要查找的元素x = 12。

初始時，low = 0，high = 14，mid = 0 + (14 - 0) * (12 - 2) / (30 - 2) = 4。此時，arr[4] = 10，小於要查找的元素x，因此更新low為5。

接下來，mid = 5 + (14 - 5) * (12 - 10) / (30 - 10) = 7。此時，arr[7] = 16，大於要查找的元素x，因此更新high為6。

然後，mid = 5 + (6 - 5) * (12 - 10) / (14 - 10) = 5。此時，arr[5] = 12，與要查找的元素相等，因此返回索引值5。

因此，Interpolation Search在此示例中成功找到了要查找的元素。   

## 彙總重點
{{< ask_chatgpt >}}
條列Interpolation Search的重點
{{< /ask_chatgpt >}}



1. 描述：Interpolation Search是一種用於查找有序且均勻分佈的數組中特定元素的搜索算法。

2. 工作原理：Interpolation Search基於線性搜索和二分搜索的想法。它通過估計元素的位置並縮小搜索範圍來尋找特定元素。

3. 適用範圍：Interpolation Search僅適用於有序且均勻分佈的數組。

4. 時間複雜度：Interpolation Search的時間複雜度取決於數組中元素的分佈。在最糟糕的情況下，它的時間複雜度為O(n)，在最好情況下，它的時間複雜度為O(1)。

5. 特點：Interpolation Search通常比二分搜索更快，因為它可以在平均O(loglog n)時間內找到元素，而二分搜索需要O(log n)時間。

6. 缺點：當數組中元素分佈不均勻時，Interpolation Search的效率可能會受到影響，導致搜索時間更長。此外，它也不像二分搜索那麼容易實現。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Interpolation Search的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



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

## 網路資料
{{< ask_chatgpt >}}
給我5篇Interpolation Search的網路資料
{{< /ask_chatgpt >}}



1. GeeksforGeeks：Interpolation Search
https://www.geeksforgeeks.org/interpolation-search/

這篇來自於知名的資訊科技網站GeeksforGeeks，介紹了Interpolation Search的應用、工作原理、優點、缺點等相關知識。

2. 演算法筆記：插值搜尋
http://www.csie.ntnu.edu.tw/~u91029/Search.html#1

這篇來自於臺灣大學資訊工程學系，是臺灣比較知名的資訊科技網站之一，介紹了Interpolation Search的原理、程式碼實現與範例。

3. Code Faster：Interpolation Search in Python
https://codefaster.substack.com/p/interpolation-search-in-python

這篇是來自於Code Faster的文章，介紹了如何用Python語言實現Interpolation Search，同時還說明了為什麼Interpolation Search比二分搜索更快。

4. 那些年我們一起查過的算法：插值查找(Interpolation Search)
https://zhuanlan.zhihu.com/p/100816770

這篇來自於知名問答網站知乎，作者對Interpolation Search這個算法進行了簡單的介紹，同時還分享了相關的程式碼實現。

5. CSDN：Interpolation Search
https://blog.csdn.net/XDU0325/article/details/92263810

這篇來自於CSDN，CSDN是一個面向中國軟件開發者的平台，這篇文章介紹了Interpolation Search的原理，並且附上了C語言和Java語言的程式碼實現。   

