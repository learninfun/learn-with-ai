

DOM manipulation 是指对于网页上的 DOM (Document Object Model) 元素进行操作和修改。这些操作可以是添加、删除或修改 DOM 元素的内容或样式。当网页需要根据用户的操作，或当需要对网页的内容进行动态修改时，DOM manipulation 是必要的技能。

以下是一些 DOM manipulation 的例子：

1. 改变元素内容：使用 JavaScript 可以使用 innerHTML() 方法来改变元素的内容。例如：

```javascript
var title = document.getElementById("title");
title.innerHTML = "New Title";
```

2. 改变元素样式：使用 JavaScript 可以使用 CSS 属性来修改元素的样式。例如：

```javascript
var title = document.getElementById("title");
title.style.color = "red";
```

3. 添加新的元素：使用 JavaScript 可以使用 createElement() 方法创建新元素，然后使用 appendChild() 方法将新元素添加到 DOM 中。例如：

```javascript
var newParagraph = document.createElement("p");
var paragraphText = document.createTextNode("This is a new paragraph.");
newParagraph.appendChild(paragraphText);

var articles = document.getElementById("articles");
articles.appendChild(newParagraph);
```

4. 删除元素：使用 JavaScript 可以使用 removeChild() 方法来删除 DOM 中的元素。例如：

```javascript
var title = document.getElementById("title");
var parent = title.parentNode;
parent.removeChild(title);
```