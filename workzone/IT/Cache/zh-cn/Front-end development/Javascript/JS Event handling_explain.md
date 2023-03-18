

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