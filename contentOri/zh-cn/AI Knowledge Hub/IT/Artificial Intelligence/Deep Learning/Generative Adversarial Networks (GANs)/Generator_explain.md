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