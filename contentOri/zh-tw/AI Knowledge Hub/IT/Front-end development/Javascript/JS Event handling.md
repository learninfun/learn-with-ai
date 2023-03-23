+++
title = "JavaScript 事件處理"
weight = "2"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題JS Event handling的中文問題
{{< /ask_chatgpt >}}



1. 什麼是事件？
2. 什麼是事件處理程序？
3. 如何在JS中綁定事件處理程序？
4. 如何防止事件冒泡？
5. 如何禁止默認事件行為？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明JS Event handling並舉例
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
以中文條列JS Event handling的重點
{{< /ask_chatgpt >}}



1. 事件：JavaScript中的事件指的是可以在瀏覽器或文檔中發生的操作或行為，例如單擊、雙擊、鍵盤按鍵等。

2. 事件處理程序：事件處理程序是指在事件發生時運行的函數。

3. 事件監聽器：事件監聽器是指用於監聽當特定事件發生時自動調用函數的方法。它可以通過 addEventListener() 函數來實現。

4. 事件對像：事件對象是在事件發生時傳遞給函數的參數，它包含有關事件的各種信息，例如事件類型、目標元素、鼠標位置等。

5. 事件冒泡和事件捕獲：事件冒泡和事件捕獲是DOM事件模型中的兩種級別的事件傳播機制。事件捕獲是指從父級元素到目標元素的事件傳播，而事件冒泡是指從目標元素到父級元素的事件傳播。

6. 阻止事件默認行為：通過調用事件對象的 preventDefault() 方法，可以阻止事件的默認行為，例如禁用鏈接的跳轉或表單提交等。

7. 停止事件傳播：通過調用事件對象的 stopPropagation() 方法，可以停止事件的傳播，也就是阻止事件冒泡或事件捕獲。

8. 事件委託：事件委託是指將事件處理程序綁定在父元素上，然後利用事件冒泡的機制來處理子元素的事件。

9. DOMContentLoaded 事件：DOMContentLoaded 事件是指在文檔加載完成後觸發的事件，用於執行需要在文檔載入完成之後才能執行的代碼。

10. resize 事件：resize 事件是指當瀏覽器窗口的大小發生改變時觸發的事件。

11. scroll 事件：scroll 事件是指當文檔滾動時觸發的事件。

12. mouseover 和 mouseout 事件：mouseover 和 mouseout 事件是指當鼠標移動到元素上方（mouseover）或離開元素（mouseout）時觸發的事件。

13. click 事件：click 事件是指當鼠標單擊元素時觸發的事件。

14. keydown、keyup 和 keypress 事件：keydown、keyup 和 keypress 事件是指當用戶按下或鬆開鍵盤鍵時觸發的事件。其中，keydown 和 keyup 事件傳遞的參數包含有關鍵盤按鍵的信息，而 keypress 事件則不包含。

15. load 事件：load 事件是指當文檔或圖片等資源加載完成時觸發的事件。

16. submit 事件：submit 事件是指當表單提交時觸發的事件，通常用於驗證表單輸入並防止表單重複提交。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題JS Event handling的中等難度問題，並在後面列出答案
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

