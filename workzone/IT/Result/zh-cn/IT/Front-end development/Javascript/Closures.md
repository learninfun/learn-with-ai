+++
title = "闭包"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Closures的问题
{{< /ask_chatgpt >}}



1. 什么是Closure？ 它在JavaScript中的作用是什么？

2. 如何在JavaScript中创建一个Closure？可以举一个实际的例子吗？

3. Closer的作用域是什么？它如何影响功能内变量的作用域？

4. 什么时候应该使用Closure，而不是其他的函数？

5. 如果Closure导致内存泄露，你怎么解决这个问题？   

## 说明知识
{{< ask_chatgpt >}}
说明Closures并举例
{{< /ask_chatgpt >}}



Closure是一种函数，它捕获对外部变量的引用，当函数在外部被调用时，它可以访问这些变量。简而言之，closure是一个包含函数定义和指向它捕获的变量的一个环境的对象。

下面是一个closure的示例代码：

```python
def outer_func(x):
    def inner_func(y):
        return x + y
    return inner_func

closure = outer_func(10)
result = closure(5)
print(result)  # 15
```

在这个示例中，outer_func返回了inner_func，inner_func可以访问outer_func中的变量x。在调用outer_func时传递的参数x是10，并且在之后将其捕获在closure中。当closure被调用时，inner_func可以访问x的值，并且将其加上将来传递的y值。在此示例中，closure的结果是15（即10 + 5）。

总的来说，closures是一种强大的python特性，可以使程序更加灵活和可读性。它们可以捕获变量，让函数更加通用和易于重复使用。   

## 汇总重点
{{< ask_chatgpt >}}
条列Closures的重点
{{< /ask_chatgpt >}}



1. Closures是一种函式内的函式，可以在函式中建立新的作用域。
2. Closures允许内部函式访问其外部函式的变数和参数。
3. Closures允许内部函式保持其创建时外部作用域中的变数状态，而不是在内部函式调用时重新定义这些变数。
4. Closures在JavaScript中广泛应用，特别是用于创建私有变数和方法，以及面向对象编程模式。
5. 在Closures中，外部函式可以返回内部函式，因此Closures也被称为函数工厂。
6. Closures和垃圾回收机制密切相关，不正确使用Closures可能会导致内存泄漏。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Closures的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 實作一個createElement函數，讓該函數能接收一個HTML tag（如 "div"），並回傳一個新的函數，該函數可以接收一個文字字串，並使用createElement中傳入的tag創建一個新的DOM元素，然後將該字串添加為元素的文字內容。

答案：

```javascript
function createElement(tag) {
  return function(content) {
    var element = document.createElement(tag);
    element.textContent = content;
    return element;
  }
}

var createH1 = createElement("h1");
var heading = createH1("Hello World");
document.body.appendChild(heading);
```

2. 建立一個能計算次方的函數 "powerFunc"，該函數可以接收一個數字作為底數並回傳一個新的函數，該函數可以接收一個指數並回傳底數的指數次方。

答案：

```javascript
function powerFunc(base) {
  return function(exponent) {
    return Math.pow(base, exponent);
  }
}

var square = powerFunc(2);
console.log(square(3)); // 8
```

3. 實現一個 curry 函數，使其能夠 currying 外部函數。例如：

```javascript
function add(a, b, c) {
  return a + b + c;
}
var curriedAdd = curry(add);
curriedAdd(1)(2)(3) // 6
```

答案：

```javascript
function curry(func) {
  return function curried(...args) {
    if (args.length >= func.length) {
      return func.apply(null, args);
    } else {
      return function(...args2) {
        return curried.apply(null, args.concat(args2));
      }
    }
  };
}

var add = function (a, b, c) {
  return a + b + c;
};

var curriedAdd = curry(add);
console.log(curriedAdd(1)(2)(3)); // 6
```

4. 写一个函数 sequence，让它可以按顺序执行一个或多个异步任务，如下所示：

```javascript
function asyncTask1(callback) {
  console.log('asyncTask1');
  callback();
}

function asyncTask2(callback) {
  console.log('asyncTask2');
  callback();
}

function asyncTask3(callback) {
  console.log('asyncTask3');
  callback();
}

sequence(asyncTask1, asyncTask2, asyncTask3, function() {
  console.log('done!');
});
```

输出结果：

```javascript
asyncTask1
asyncTask2
asyncTask3
done!
```

答案：

```javascript
function sequence(...tasks) {
  const last = tasks.pop();
  function handleCallbacks(tasks, callback) {
    const task = tasks.shift();
    if (typeof task === 'function') {
      task(function() {
        handleCallbacks(tasks, callback);
      });
    } else {
      callback();
    }
  }
  handleCallbacks(tasks.slice(), last);
}
```

5. 编写一个函数createKeyValueStore，可以创建一个简单的键值存储对象，并返回一个新的函数。当调用该函数时，键值存储应该更新为指定的键和值。每当更新时，应该在控制台上打印出新的键值存储。例如：

```javascript
const keyValues = createKeyValueStore();
keyValues('name', 'John');
keyValues('age', 30);

// 控制台输出：
// { name: 'John' }
// { name: 'John', age: 30 }
```

答案：

```javascript
function createKeyValueStore() {
  const keyValueStore = {};
  return function(key, value) {
    keyValueStore[key] = value;
    console.log(keyValueStore);
  }
}

const keyValues = createKeyValueStore();
keyValues('name', 'John');
keyValues('age', 30);
```   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Closures的网络数据
{{< /ask_chatgpt >}}



1. Understanding JavaScript Closures

This article from SitePoint gives a comprehensive overview of closures in JavaScript, including how they work, why they are useful, and examples of how to use closures in your code. It's a great starting point for anyone looking to understand closures better.

Link: https://www.sitepoint.com/javascript-closures/

2. Mastering JavaScript Closures

This article from Medium provides a more in-depth look at closures, including practical examples of how to use them in your code. It covers topics such as lexical scoping, the "this" keyword, and closures and memory management.

Link: https://medium.com/@prashantramnyc/mastering-javascript-closures-9d7e5eb5d5df

3. JavaScript Closures: A Comprehensive Guide

This guide from Toptal goes into great detail about closures in JavaScript, including how they work, common use cases, and best practices for using them. It also includes multiple code examples and exercises to help you practice using closures in your own code.

Link: https://www.toptal.com/javascript/closures

4. The Magic of JavaScript Closures

This article from Codeburst provides a more conceptual overview of closures in JavaScript, discussing how they allow for encapsulation, creating private variables, and maintaining state in your code. It includes helpful code examples and visuals to illustrate the concepts.

Link: https://codeburst.io/the-magic-of-javascript-closures-2b22bebe69af

5. Unraveling Closure Functions in JavaScript

This article from Scotch.io provides a beginner-friendly look at closures in JavaScript, including how they work, basic syntax, and common patterns for using closures in your code. It also includes helpful visuals and code examples to illustrate the concepts.

Link: https://scotch.io/tutorials/unraveling-closure-functions-in-javascript   

