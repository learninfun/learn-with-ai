

Memoization是一種常用的技術，它可以用來優化重複計算的過程。簡單來說，Memoization就是在計算結果後將其存儲起來，以便在未來需要時可以快速查找。這種方式能夠減少計算的時間複雜度，提高程式效率。

下面舉一個簡單的例子：

假設我們要計算費氏數列中第n個數，費氏數列的定義為：數列中的第一和第二個數都是1，從第三個數開始，每個數都是前面兩個數的和。也就是說，費氏數列的前幾項為：1, 1, 2, 3, 5, 8, 13, ...

我們可以使用遞迴來計算費氏數列中的第n個數，具體實現如下：

```
def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)
```

使用這個函數來計算費氏數列中第n個數的值時，當n較大時會出現計算時間過長的情況。因為計算過程中需要重複計算一些值，比如計算費氏數列中的第10項時，需要先計算出第9項和第8項，計算第9項時又需要先計算第8項和第7項，而計算第8項時還需要計算第7項和第6項，這樣就會浪費大量的計算時間。

使用Memoization可以避免這些重複計算的問題。具體實現方式如下：

```
cache = {}  # 定義一個字典，用來存儲計算結果

def fibonacci(n):
    if n < 2:
        return n
    if n in cache:  # 如果計算結果已經存在於cache中，直接返回該結果
        return cache[n]
    result = fibonacci(n-1) + fibonacci(n-2)
    cache[n] = result  # 將計算結果存儲到cache中
    return result
```

使用這個改進版的函數來計算費氏數列中的第n個數時，計算時間可以大幅降低。因為當需要計算一個已經計算過的值時，可以直接從cache中查找得到，而不需要重新計算。這樣就可以減少重複計算的次數，提高程式效率。