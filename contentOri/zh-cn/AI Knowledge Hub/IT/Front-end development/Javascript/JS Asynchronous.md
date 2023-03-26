+++
title = "JavaScript 非同步"
weight = "4"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題JS Asynchronous的中文問題
{{< /ask_chatgpt >}}



1. 如何使用Js异步加载数据?
2. 什么是JavaScript的回调函数，为什么回调函数对于异步编程很重要？
3. 什么是Promise对象？如何使用Promise对象解决异步编程问题？
4. 如何使用JavaScript的async/await关键字编写异步代码？
5. 如何在Node.js中使用异步模块来处理并行任务？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明JS Asynchronous並舉例
{{< /ask_chatgpt >}}



JS Asynchronous是指一種非同步執行方式，即在某些操作完成前，程式不會停止等待其完成，而是繼續執行後面的程式碼。

舉例來說，當你使用AJAX進行資料庫查詢時，程式不會停止等待查詢結果返回，而是繼續執行後面的代碼。

另一個例子是使用setTimeout()函數，該函數可以在一段時間後執行一段代碼，並不會使整個程式停止等待執行。

又例如Promise，當某個Promise還沒有得到結果時，程式不會停止等待，而是繼續執行後面的代碼，當Promise完成時，就會觸發回調函數。

總體而言，JS Asynchronous在處理複雜或耗時的操作時，可以讓程式更加流暢，不會阻塞或使程式失去響應。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列JS Asynchronous的重點
{{< /ask_chatgpt >}}



1. JS Asynchronous 是一種程式設計方式，可以讓程式在執行某些耗時操作時，同時執行其他程式碼，提升程式效能。

2. JS Asynchronous 的核心是事件循環（Event Loop），它會不斷檢查事件隊列中是否有任務需要處理，如果有就會執行它。

3. JS Asynchronous 的廣泛應用包括網頁開發、服務器編程、大數據處理、人工智慧等各種領域。

4. JS Asynchronous 的最常見實現方式包括回調函數（Callback），Promise、async/await 等。

5. 回調函數（Callback）是最早使用的 JS Asynchronous 方法，它會在某個事件發生時被呼叫，執行相關操作。

6. Promise 是 ES6 引入的一種 JS Asynchronous 技術，它能解決回調地獄（Callback Hell）等問題，讓程式更容易理解和維護。

7. async/await 是 ES8 引入的一種 JS Asynchronous 技術，它則更加簡潔、易讀，也更容易捕獲錯誤和處理異常情況。

8. JS Asynchronous 的實現必須遵循一定的順序和流程，否則可能會產生意外的錯誤和結果。程式設計師應該熟練掌握相關知識和技能，以減少失敗和重複工作。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題JS Asynchronous的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 寫一個函式，用 Promise 實現一個延遲一定時間的效果。
答案：
```javascript
function delay(time) {
  return new Promise(resolve => setTimeout(resolve, time));
}
```

2. 寫一個函式，用 Promise 實現一個簡單的異步串行控制流，按順序執行傳入的異步函式。
答案：
```javascript
function serialFlow(tasks) {
  return tasks.reduce((promiseChain, task) => promiseChain.then(task), Promise.resolve());
}
```

3. 寫一個函式，用 Promise 實現一個簡單的異步並行控制流，並限制同時執行的數量。
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

4. 寫一個函式，用 Promise 實現一個簡單的異步重試控制流，當異步函式執行失敗時，會進行指定次數的重試，每次重試之間會暫停一段時間。
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

5. 寫一個函式，用 Promise 實現一個簡單的異步 debounce 控制流，當連續多次調用函式時會等待指定時間後只執行一次。
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

