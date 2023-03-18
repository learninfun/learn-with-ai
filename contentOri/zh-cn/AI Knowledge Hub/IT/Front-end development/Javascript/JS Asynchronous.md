+++
title = "JavaScript 非同步"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题JS Asynchronous的问题
{{< /ask_chatgpt >}}



1. 如何使用Js异步加载数据?
2. 什么是JavaScript的回调函数，为什么回调函数对于异步编程很重要？
3. 什么是Promise对象？如何使用Promise对象解决异步编程问题？
4. 如何使用JavaScript的async/await关键字编写异步代码？
5. 如何在Node.js中使用异步模块来处理并行任务？   

## 说明知识
{{< ask_chatgpt >}}
说明JS Asynchronous并举例
{{< /ask_chatgpt >}}



JS Asynchronous是指一種非同步執行方式，即在某些操作完成前，程式不會停止等待其完成，而是繼續執行後面的程式碼。

舉例來說，當你使用AJAX進行資料庫查詢時，程式不會停止等待查詢結果返回，而是繼續執行後面的代碼。

另一個例子是使用setTimeout()函數，該函數可以在一段時間後執行一段代碼，並不會使整個程式停止等待執行。

又例如Promise，當某個Promise還沒有得到結果時，程式不會停止等待，而是繼續執行後面的代碼，當Promise完成時，就會觸發回調函數。

總體而言，JS Asynchronous在處理複雜或耗時的操作時，可以讓程式更加流暢，不會阻塞或使程式失去響應。   

## 汇总重点
{{< ask_chatgpt >}}
条列JS Asynchronous的重点
{{< /ask_chatgpt >}}



1. JS Asynchronous 是一种程式设计方式，可以让程式在执行某些耗时操作时，同时执行其他程式码，提升程式效能。

2. JS Asynchronous 的核心是事件循环（Event Loop），它会不断检查事件队列中是否有任务需要处理，如果有就会执行它。

3. JS Asynchronous 的广泛应用包括网页开发、服务器编程、大数据处理、人工智慧等各种领域。

4. JS Asynchronous 的最常见实现方式包括回调函数（Callback），Promise、async/await 等。

5. 回调函数（Callback）是最早使用的 JS Asynchronous 方法，它会在某个事件发生时被呼叫，执行相关操作。

6. Promise 是 ES6 引入的一种 JS Asynchronous 技术，它能解决回调地狱（Callback Hell）等问题，让程式更容易理解和维护。

7. async/await 是 ES8 引入的一种 JS Asynchronous 技术，它则更加简洁、易读，也更容易捕获错误和处理异常情况。

8. JS Asynchronous 的实现必须遵循一定的顺序和流程，否则可能会产生意外的错误和结果。程式设计师应该熟练掌握相关知识和技能，以减少失败和重复工作。   

## 知识测验
{{< ask_chatgpt >}}
给我5题JS Asynchronous的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 写一个函式，用 Promise 实现一个延迟一定时间的效果。
答案：
```javascript
function delay(time) {
  return new Promise(resolve => setTimeout(resolve, time));
}
```

2. 写一个函式，用 Promise 实现一个简单的异步串行控制流，按顺序执行传入的异步函式。
答案：
```javascript
function serialFlow(tasks) {
  return tasks.reduce((promiseChain, task) => promiseChain.then(task), Promise.resolve());
}
```

3. 写一个函式，用 Promise 实现一个简单的异步并行控制流，并限制同时执行的数量。
答案：
```javascript
function parallelFlow(tasks, limit) {
  const allTasks = tasks.slice();
  return new Promise((resolve) => {
    let running = 0;
    let resolved = 0;
    const next = () => {
      while (running < limit && allTasks.length) {
        running++;
        const task = allTasks.shift();
        task().then(() => {
          if (++resolved >= tasks.length) {
            resolve();
          } else {
            running--;
            next();
          }
        });
      };
    };
    next();
  });
}
```

4. 写一个函式，用 Promise 实现一个简单的异步重试控制流，当异步函式执行失败时，会进行指定次数的重试，每次重试之间会暂停一段时间。
答案：
```javascript
function retryAsync(fn, retries = 3, delaySeconds = 1) {
  return new Promise((resolve, reject) => {
    const attempt = () => {
      fn().then(resolve).catch((error) => {
        if (retries <= 0) {
          reject(error);
        } else {
          retries--;
          setTimeout(attempt, delaySeconds * 1000);
        }
      });
    }
    attempt();
  });
}
```

5. 写一个函式，用 Promise 实现一个简单的异步 debounce 控制流，当连续多次调用函式时会等待指定时间后只执行一次。
答案：
```javascript
function debounceAsync(fn, delayMilliseconds) {
  let timeout;
  return async function() {
    const args = arguments;
    const later = () => {
      clearTimeout(timeout);
      fn.apply(this, args);
    };
    clearTimeout(timeout);
    timeout = setTimeout(later, delayMilliseconds);
  };
}
```   

## 网络数据
{{< ask_chatgpt >}}
给我5篇JS Asynchronous的网络数据
{{< /ask_chatgpt >}}

1. "Asynchronous JavaScript: How to Use Async and Await" by Mosh Hamedani on Freecodecamp
https://www.freecodecamp.org/news/asynchronous-javascript-how-to-use-async-and-await-8abc449032bb/

2. "Understanding Asynchronous JavaScript" on Medium by Brandon Morelli
https://medium.com/@BrandonMorelli/understanding-asynchronous-javascript-the-event-loop-74cd408419ff

3. "JavaScript Promises: An Introduction" on SitePoint by M. David Green
https://www.sitepoint.com/javascript-promises-an-introduction/

4. "Asynchronous Programming in JavaScript" on DigitalOcean by Kristopher Sandoval
https://www.digitalocean.com/community/tutorials/how-to-work-with-async-functions-in-javascript

5. "A Beginner’s Guide to JavaScript Asynchronous Programming" on Toptal by Carlos Villavicencio
https://www.toptal.com/javascript/asynchronous-javascript-async-await-tutorial   

