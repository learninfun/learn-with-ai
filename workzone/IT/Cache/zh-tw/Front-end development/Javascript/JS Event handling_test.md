

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