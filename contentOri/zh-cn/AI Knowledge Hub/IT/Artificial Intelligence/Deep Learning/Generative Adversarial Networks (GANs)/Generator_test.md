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