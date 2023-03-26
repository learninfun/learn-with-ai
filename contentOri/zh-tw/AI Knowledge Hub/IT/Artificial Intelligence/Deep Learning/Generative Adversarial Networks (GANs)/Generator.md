+++
title = "生成器"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Generator的中文問題
{{< /ask_chatgpt >}}

1. 什麼是Generator？它有什麼作用？
2. 如何使用Python建立Generator？
3. Generator和List有什麼區別？
4. Generator和Iterator有什麼關係？
5. Generator在遍歷大數據集合時有什麼優勢？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Generator並舉例
{{< /ask_chatgpt >}}

Generator是一種Python語言中的迭代器（iterator），它可以讓我們生成一系列的值，並且在使用時可以加以控制，可以讓我們只需要生成需要的值，而不需要生成全部的值。Generator是使用yield語句來實現的，當我們使用next()函數來調用Generator時，它會執行yield語句將值返回給調用方，然後暫停執行，等待下一次調用。

以下是一個使用Generator生成斐波那契數列的例子：

```Python
def fib():
    a, b = 0, 1
    while True:
        yield a
        a, b = b, a + b
```
在這個例子中，我們使用while True來設定無限循環，然後使用yield a來生成數列中的每一個數字。當我們使用next()函數來調用這個Generator時，它會生成一個值並暫停執行，等待下一次調用。以下是使用這個Generator來生成數列的示例：

```Python
gen = fib()
for i in range(10):
    print(next(gen))
```
這段代碼將會輸出數列的前10個數字：

```
0
1
1
2
3
5
8
13
21
34
```   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Generator的重點
{{< /ask_chatgpt >}}

1. Generator是Python中的一種特殊類型，可以用來創建一個可以迭代的對象。
2. 使用yield關鍵字可以使函數返回一個Generator對象。
3. Generator對象可以在迭代時生成一系列的值，而不是一次性返回所有的值。
4. 與列表相比，Generator的主要優勢是可以減少內存的使用，因為只有在需要時才會生成值。
5. 在使用Generator時，可以使用for循環或者next()函數來遍歷生成的值。
6. Generator可以用來處理大型數據集，或者需要逐步計算的問題，例如計算斐波那契數列。
7. Generator的yield關鍵字可以有多個值，每次產生的值可以是一個元組、列表或字典等集合型別。
8. 一個Generator函數可以調用另一個Generator函數，並使用yield from語法遍歷產生的值。
9. 在使用Generator時，注意避免重復使用同一個Generator對象，否則會導致值的重複生成。
10. Generator是Python中較為高級的概念，但使用起來非常靈活和方便，可以大大提升代碼的可讀性和效率。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Generator的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}

1. 请编写一个生成器函数，用于生成斐波那契数列的前n项。
2. 请编写一个生成器函数，用于对一个字符串中的单词进行循环迭代。
3. 请编写一个生成器函数，用于生成范围内的全部质数。
4. 请编写一个生成器函数，用于生成n个范围内的随机整数。
5. 请编写一个生成器函数，用于对一个列表进行不断地随机位置交换，直到该列表变为有序状态。

答案：
1. 
```python
def fibonacci(n):
    a = 0
    b = 1
    count = 0
    while count < n:
        yield a
        a, b = b, a + b
        count += 1
```
2. 
```python
def word_iterate(s):
    for word in s.split():
        yield word
```
3. 
```python
def prime_number(range_num):
    for num in range(2, range_num + 1):
        prime = True
        for i in range(2, int(num ** 0.5) + 1):
            if num % i == 0:
                prime = False
                break
        if prime:
            yield num
```
4. 
```python
import random

def random_num(n, start, end):
    for i in range(n):
        yield random.randint(start, end)
```
5. 
```python
import random

def shuffle_list(l):
    shuffled = False
    while not shuffled:
        current_state = list(l)
        random.shuffle(l)
        if current_state == sorted(l):
            shuffled = True
        else:
            l = current_state
    return l
```   

