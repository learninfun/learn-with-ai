

DOM manipulation 是指對於網頁上的 DOM (Document Object Model) 元素進行操作和修改。這些操作可以是添加、刪除或修改 DOM 元素的內容或樣式。當網頁需要根據用戶的操作，或當需要對網頁的內容進行動態修改時，DOM manipulation 是必要的技能。

以下是一些 DOM manipulation 的例子：

1. 改變元素內容：使用 JavaScript 可以使用 innerHTML() 方法來改變元素的內容。例如：

```javascript
var title = document.getElementById("title");
title.innerHTML = "New Title";
```

2. 改變元素樣式：使用 JavaScript 可以使用 CSS 屬性來修改元素的樣式。例如：

```javascript
var title = document.getElementById("title");
title.style.color = "red";
```

3. 添加新的元素：使用 JavaScript 可以使用 createElement() 方法創建新元素，然後使用 appendChild() 方法將新元素添加到 DOM 中。例如：

```javascript
var newParagraph = document.createElement("p");
var paragraphText = document.createTextNode("This is a new paragraph.");
newParagraph.appendChild(paragraphText);

var articles = document.getElementById("articles");
articles.appendChild(newParagraph);
```

4. 刪除元素：使用 JavaScript 可以使用 removeChild() 方法來刪除 DOM 中的元素。例如：

```javascript
var title = document.getElementById("title");
var parent = title.parentNode;
parent.removeChild(title);
```