

1. 如何在网页载入完成后执行一个函数？

答案：

```js
window.addEventListener('load', function() {
  // your code here
});
```

2. 如何防止点击一个连结时页面跳转？

答案：

```js
document.addEventListener('click', function(event) {
  // 防止连结默认行为
  event.preventDefault();
});
```

3. 如何在滑鼠移到一个元素时显示一个提示框？

答案：

```js
var element = document.getElementById('myElement');
element.addEventListener('mouseover', function() {
  // 显示提示框
  alert('Hello, World!');
});
```

4. 如何在按下键盘上的某个按键时执行一个函数？

答案：

```js
document.addEventListener('keydown', function(event) {
  // 判断按下的键是哪个
  if (event.key === 'Enter') {
    // 执行函数
    doSomething();
  }
});
```

5. 如何在拖动一个元素时修改它的位置？

答案：

```js
var element = document.getElementById('myElement');
var x, y;

element.addEventListener('mousedown', function(event) {
  // 记住滑鼠位置
  x = event.clientX - element.offsetLeft;
  y = event.clientY - element.offsetTop;
  
  document.addEventListener('mousemove', moveElement);
});

document.addEventListener('mouseup', function() {
  document.removeEventListener('mousemove', moveElement);
});

function moveElement(event) {
  // 计算新位置
  var newX = event.clientX - x;
  var newY = event.clientY - y;

  // 设定新位置
  element.style.left = newX + 'px';
  element.style.top = newY + 'px';
}
```