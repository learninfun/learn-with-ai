+++
title = "閉包"
weight = "3"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Closures的中文問題
{{< /ask_chatgpt >}}



1. 什麼是Closure？ 它在JavaScript中的作用是什麼？

2. 如何在JavaScript中創建一個Closure？可以舉一個實際的例子嗎？

3. Closer的作用域是什麼？它如何影響功能內變量的作用域？

4. 什麼時候應該使用Closure，而不是其他的函數？

5. 如果Closure導致內存洩露，你怎麼解決這個問題？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Closures並舉例
{{< /ask_chatgpt >}}



Closure是一種函數，它捕獲對外部變量的引用，當函數在外部被調用時，它可以訪問這些變量。簡而言之，closure是一個包含函數定義和指向它捕獲的變量的一個環境的對象。

下面是一個closure的示例代碼：

```python
def outer_func(x):
    def inner_func(y):
        return x + y
    return inner_func

closure = outer_func(10)
result = closure(5)
print(result)  # 15
```

在這個示例中，outer_func返回了inner_func，inner_func可以訪問outer_func中的變量x。在調用outer_func時傳遞的參數x是10，並且在之後將其捕獲在closure中。當closure被調用時，inner_func可以訪問x的值，並且將其加上將來傳遞的y值。在此示例中，closure的結果是15（即10 + 5）。

總的來說，closures是一種強大的python特性，可以使程序更加靈活和可讀性。它們可以捕獲變量，讓函數更加通用和易於重複使用。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列Closures的重點
{{< /ask_chatgpt >}}



1. Closures是一種函式內的函式，可以在函式中建立新的作用域。
2. Closures允許內部函式訪問其外部函式的變數和參數。
3. Closures允許內部函式保持其創建時外部作用域中的變數狀態，而不是在內部函式調用時重新定義這些變數。
4. Closures在JavaScript中廣泛應用，特別是用於創建私有變數和方法，以及面向對象編程模式。
5. 在Closures中，外部函式可以返回內部函式，因此Closures也被稱為函數工廠。
6. Closures和垃圾回收機制密切相關，不正確使用Closures可能會導致內存洩漏。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Closures的中等難度問題，並在後面列出答案
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

