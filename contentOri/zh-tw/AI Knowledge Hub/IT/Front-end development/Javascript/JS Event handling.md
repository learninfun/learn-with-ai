+++
title = "JavaScript 事件處理"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題JS Event handling的問題
{{< /ask_chatgpt >}}



1. 什麼是事件？
2. 什麼是事件處理程序？
3. 如何在JS中綁定事件處理程序？
4. 如何防止事件冒泡？
5. 如何禁止默認事件行為？   

## 說明知識
{{< ask_chatgpt >}}
說明JS Event handling並舉例
{{< /ask_chatgpt >}}



JS Event handling是指處理網頁中各種事件（如按鈕點擊、滑鼠移動、按下鍵盤等）的能力，可以透過JavaScript代碼為HTML元素添加事件監聽器。當特定事件發生時，就執行對應的JavaScript代碼來處理該事件。

以下是一個簡單的例子：

HTML：

```html
<button id="myButton">Click me!</button>
```

JavaScript：

```javascript
var button = document.getElementById("myButton");

button.addEventListener("click", function() {
  alert("Button clicked!");
});
```

在此範例中，我們首先通過`document.getElementById()`方法獲取一個id為`myButton`的按鈕元素。然後，使用`addEventListener()`方法添加一個`click`事件監聽器，該監聽器會在按鈕被點擊時執行一個匿名函式，該函式會顯示一個彈出框，其中包含一條消息「Button clicked!」。

當按鈕被點擊時，事件系統會檢測到這個事件並觸發添加的監聽器，該監聽器執行了匿名函式，顯示了彈出框。這種方式的好處是，它可以讓開發者將相關邏輯和操作與HTML分離，使代碼更加模塊化和易於維護。   

## 彙總重點
{{< ask_chatgpt >}}
條列JS Event handling的重點
{{< /ask_chatgpt >}}



1. Understanding events
2. Event listener registration methods
3. Event bubbling and capturing
4. Event propagation and stopping
5. Event delegation
6. Event object and properties
7. Event types and their properties
8. Event handling best practices
9. Cross-browser compatibility issues
10. Debugging event handling issues   

## 知識測驗
{{< ask_chatgpt >}}
給我5題JS Event handling的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 如何在網頁載入完成後執行一個函數？

答案：

```js
window.addEventListener('load', function() {
  // your code here
});
```

2. 如何防止點擊一個連結時頁面跳轉？

答案：

```js
document.addEventListener('click', function(event) {
  // 防止連結默認行為
  event.preventDefault();
});
```

3. 如何在滑鼠移到一個元素時顯示一個提示框？

答案：

```js
var element = document.getElementById('myElement');
element.addEventListener('mouseover', function() {
  // 顯示提示框
  alert('Hello, World!');
});
```

4. 如何在按下鍵盤上的某個按鍵時執行一個函數？

答案：

```js
document.addEventListener('keydown', function(event) {
  // 判斷按下的鍵是哪個
  if (event.key === 'Enter') {
    // 執行函數
    doSomething();
  }
});
```

5. 如何在拖動一個元素時修改它的位置？

答案：

```js
var element = document.getElementById('myElement');
var x, y;

element.addEventListener('mousedown', function(event) {
  // 記住滑鼠位置
  x = event.clientX - element.offsetLeft;
  y = event.clientY - element.offsetTop;
  
  document.addEventListener('mousemove', moveElement);
});

document.addEventListener('mouseup', function() {
  document.removeEventListener('mousemove', moveElement);
});

function moveElement(event) {
  // 計算新位置
  var newX = event.clientX - x;
  var newY = event.clientY - y;

  // 設定新位置
  element.style.left = newX + 'px';
  element.style.top = newY + 'px';
}
```   

## 網路資料
{{< ask_chatgpt >}}
給我5篇JS Event handling的網路資料
{{< /ask_chatgpt >}}



1. JavaScript Event Handling – A Detailed Guide
https://www.edureka.co/blog/javascript-event-handling-tutorial/

2. Handling events in JavaScript – a comprehensive guide
https://www.javascripttutorial.net/javascript-dom/javascript-event-handling/

3. JavaScript Events – A Primer
https://www.sitepoint.com/javascript-events-primer/

4. The Ultimate Guide to JavaScript Event Handling
https://medium.com/better-programming/the-ultimate-guide-to-javascript-event-handling-7a1a01c84029

5. Events in JavaScript: How They Work and How to Use Them
https://www.digitalocean.com/community/tutorials/js-events-101   

