+++
title = "JavaScript 事件处理"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题JS Event handling的问题
{{< /ask_chatgpt >}}



1. 什么是事件？
2. 什么是事件处理程序？
3. 如何在JS中绑定事件处理程序？
4. 如何防止事件冒泡？
5. 如何禁止默认事件行为？   

## 说明知识
{{< ask_chatgpt >}}
说明JS Event handling并举例
{{< /ask_chatgpt >}}



JS Event handling是指处理网页中各种事件（如按钮点击、滑鼠移动、按下键盘等）的能力，可以透过JavaScript代码为HTML元素添加事件监听器。当特定事件发生时，就执行对应的JavaScript代码来处理该事件。

以下是一个简单的例子：

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

在此范例中，我们首先通过`document.getElementById()`方法获取一个id为`myButton`的按钮元素。然后，使用`addEventListener()`方法添加一个`click`事件监听器，该监听器会在按钮被点击时执行一个匿名函式，该函式会显示一个弹出框，其中包含一条消息“Button clicked!”。

当按钮被点击时，事件系统会检测到这个事件并触发添加的监听器，该监听器执行了匿名函式，显示了弹出框。这种方式的好处是，它可以让开发者将相关逻辑和操作与HTML分离，使代码更加模块化和易于维护。   

## 汇总重点
{{< ask_chatgpt >}}
条列JS Event handling的重点
{{< /ask_chatgpt >}}



1. 事件：JavaScript中的事件指的是可以在浏览器或文档中发生的操作或行为，例如单击、双击、键盘按键等。

2. 事件处理程序：事件处理程序是指在事件发生时运行的函数。

3. 事件监听器：事件监听器是指用于监听当特定事件发生时自动调用函数的方法。它可以通过 addEventListener() 函数来实现。

4. 事件对像：事件对象是在事件发生时传递给函数的参数，它包含有关事件的各种信息，例如事件类型、目标元素、鼠标位置等。

5. 事件冒泡和事件捕获：事件冒泡和事件捕获是DOM事件模型中的两种级别的事件传播机制。事件捕获是指从父级元素到目标元素的事件传播，而事件冒泡是指从目标元素到父级元素的事件传播。

6. 阻止事件默认行为：通过调用事件对象的 preventDefault() 方法，可以阻止事件的默认行为，例如禁用链接的跳转或表单提交等。

7. 停止事件传播：通过调用事件对象的 stopPropagation() 方法，可以停止事件的传播，也就是阻止事件冒泡或事件捕获。

8. 事件委托：事件委托是指将事件处理程序绑定在父元素上，然后利用事件冒泡的机制来处理子元素的事件。

9. DOMContentLoaded 事件：DOMContentLoaded 事件是指在文档加载完成后触发的事件，用于执行需要在文档载入完成之后才能执行的代码。

10. resize 事件：resize 事件是指当浏览器窗口的大小发生改变时触发的事件。

11. scroll 事件：scroll 事件是指当文档滚动时触发的事件。

12. mouseover 和 mouseout 事件：mouseover 和 mouseout 事件是指当鼠标移动到元素上方（mouseover）或离开元素（mouseout）时触发的事件。

13. click 事件：click 事件是指当鼠标单击元素时触发的事件。

14. keydown、keyup 和 keypress 事件：keydown、keyup 和 keypress 事件是指当用户按下或松开键盘键时触发的事件。其中，keydown 和 keyup 事件传递的参数包含有关键盘按键的信息，而 keypress 事件则不包含。

15. load 事件：load 事件是指当文档或图片等资源加载完成时触发的事件。

16. submit 事件：submit 事件是指当表单提交时触发的事件，通常用于验证表单输入并防止表单重复提交。   

## 知识测验
{{< ask_chatgpt >}}
给我5题JS Event handling的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



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

