

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