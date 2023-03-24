

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