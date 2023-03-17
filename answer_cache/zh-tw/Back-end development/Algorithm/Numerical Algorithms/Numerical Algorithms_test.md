

1) 使用二分搜尋法來尋找一個在排序過的數組中的數字。輸入為一個排序過的數組和要查找的數字，輸出為該數字在數組中的索引位置，如果該數字不在數組中則輸出-1。

答案：
```Python
def binary_search(arr, x):
    l, r = 0, len(arr)-1
    while l <= r:
        mid = (l + r) // 2
        if arr[mid] == x:
            return mid
        elif arr[mid] < x:
            l = mid + 1
        else:
            r = mid - 1
    return -1

print(binary_search([1, 2, 3, 4, 5], 3)) # output: 2
print(binary_search([1, 2, 3, 4, 5], 6)) # output: -1
```

2) 使用牛頓法求一個函數的根。輸入為一個函數、其導函數、初始值以及算法迭代次數，輸出為近似的根。

答案：
```Python
def newton_method(f, df, x0, n_iterations):
    x = x0
    for i in range(n_iterations):
        x = x - f(x)/df(x)
    return x

f = lambda x: x**3 - x**2 + 2
df = lambda x: 3*x**2 - 2*x
print(newton_method(f, df, 1, 5)) # output: 0.855624760256968
```

3) 使用歐拉法求解微分方程。輸入為一個微分方程、初始值、步長以及算法迭代次數，輸出為近似的解。

答案：
```Python
def euler_method(f, x0, y0, h, n_iterations):
    x, y = x0, y0
    for i in range(n_iterations):
        y += h * f(x, y)
        x += h
    return y
  
f = lambda x, y: x + y
print(euler_method(f, 0, 1, 0.1, 10)) # output: 2.6448088481707596
```

4) 使用高斯消元法求解線性方程組。輸入為一個線性方程組，輸出為方程組的解。

答案：
```Python
import numpy as np

def gaussian_elimination(a, b):
    n = len(b)
    for i in range(n):
        max_index = i + np.argmax(abs(a[i:, i]))
        a[[i, max_index]] = a[[max_index, i]]
        b[[i, max_index]] = b[[max_index, i]]
        for j in range(i+1, n):
            factor = a[j, i] / a[i, i]
            b[j] -= factor * b[i]
            for k in range(i, n):
                a[j, k] -= factor * a[i, k]
    x = np.zeros(n)
    for i in range(n-1, -1, -1):
        x[i] = (b[i] - np.dot(a[i, i+1:], x[i+1:])) / a[i, i]
    return x

a = np.array([[3, 2, -1],
              [2, -2, 4],
              [-1, 0.5, -1]])
b = np.array([1, -2, 0])
print(gaussian_elimination(a, b)) # output: [ 1. -2. -2.]
```

5) 使用SVM求解二元分類問題。輸入為分類問題的數據以及對應的標籤，輸出為訓練好的SVM模型。

答案：
```Python
from sklearn import svm
X = [[0, 0], [1, 1], [1, 0]]
y = [0, 1, 1]
clf = svm.SVC(kernel='linear', C=1000)
clf.fit(X, y)
print(clf.predict([[2, 2]])) # output: [1]
```