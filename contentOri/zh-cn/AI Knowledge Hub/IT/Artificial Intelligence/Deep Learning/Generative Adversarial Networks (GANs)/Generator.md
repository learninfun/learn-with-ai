+++
title = "Generator"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Generator的中文问题
{{< /ask_chatgpt >}}

1. 什么是Generator？它有什么作用？
2. 如何使用Python建立Generator？
3. Generator和List有什么区别？
4. Generator和Iterator有什么关系？
5. Generator在遍历大数据集合时有什么优势？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明Generator并举例
{{< /ask_chatgpt >}}

Generator是一种Python语言中的迭代器（iterator），它可以让我们生成一系列的值，并且在使用时可以加以控制，可以让我们只需要生成需要的值，而不需要生成全部的值。Generator是使用yield语句来实现的，当我们使用next()函数来调用Generator时，它会执行yield语句将值返回给调用方，然后暂停执行，等待下一次调用。

以下是一个使用Generator生成斐波那契数列的例子：

```Python
def fib():
    a, b = 0, 1
    while True:
        yield a
        a, b = b, a + b
```
在这个例子中，我们使用while True来设定无限循环，然后使用yield a来生成数列中的每一个数字。当我们使用next()函数来调用这个Generator时，它会生成一个值并暂停执行，等待下一次调用。以下是使用这个Generator来生成数列的示例：

```Python
gen = fib()
for i in range(10):
    print(next(gen))
```
这段代码将会输出数列的前10个数字：

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

## 条列重点
{{< ask_chatgpt >}}
以中文条列Generator的重点
{{< /ask_chatgpt >}}

1. Generator是Python中的一种特殊类型，可以用来创建一个可以迭代的对象。
2. 使用yield关键字可以使函数返回一个Generator对象。
3. Generator对象可以在迭代时生成一系列的值，而不是一次性返回所有的值。
4. 与列表相比，Generator的主要优势是可以减少内存的使用，因为只有在需要时才会生成值。
5. 在使用Generator时，可以使用for循环或者next()函数来遍历生成的值。
6. Generator可以用来处理大型数据集，或者需要逐步计算的问题，例如计算斐波那契数列。
7. Generator的yield关键字可以有多个值，每次产生的值可以是一个元组、列表或字典等集合型别。
8. 一个Generator函数可以调用另一个Generator函数，并使用yield from语法遍历产生的值。
9. 在使用Generator时，注意避免重复使用同一个Generator对象，否则会导致值的重复生成。
10. Generator是Python中较为高级的概念，但使用起来非常灵活和方便，可以大大提升代码的可读性和效率。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题Generator的中等难度问题，并在后面列出答案
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

