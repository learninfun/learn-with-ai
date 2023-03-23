+++
title = "DOM 操作"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题DOM manipulation的问题
{{< /ask_chatgpt >}}



1. 如何使用JavaScript获取DOM元素的属性值？

2. 如何使用JavaScript获取DOM元素的文本内容？

3. 如何使用JavaScript创建新的DOM元素？

4. 如何使用JavaScript更改DOM元素的属性值？

5. 如何使用JavaScript在DOM元素中添加或删除子元素？   

## 说明知识
{{< ask_chatgpt >}}
说明DOM manipulation并举例
{{< /ask_chatgpt >}}



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

## 汇总重点
{{< ask_chatgpt >}}
条列DOM manipulation的重点
{{< /ask_chatgpt >}}



1. 元素选择和查询：使用选择器和方法选择特定的元素，例如getElementById、querySelectorAll等。

2. 元素创建和添加：使用createElement和appendChild等方法创建和添加新元素。

3. 元素修改和删除：使用innerHTML和removeChild等方法修改和删除元素的内容和属性。

4. 属性修改和查询：使用getAttribute和setAttribute等方法修改和查询元素的属性。

5. 样式修改和查询：使用style和classList等方法修改和查询元素的样式。

6. 事件处理程序：使用addEventListener和removeEventListener等方法添加和删除事件处理程序。

7. AJAX：使用XMLHttpRequest对Web服务器进行非同步操作，从而更新网页内容。

8. 动态内容：使用innerHTML和createElement等方法动态更新网页内容、添加新元素和属性。

9. 数据交互：使用XMLHttpRequest和JSON等技术与Web服务器进行数据交互。

10. 动画效果：使用CSS和JavaScript创建交互性和动画效果。   

## 知识测验
{{< ask_chatgpt >}}
给我5题DOM manipulation的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 如何將一個特定的div元素中的所有子元素的文本內容加粗？（提示：使用for循環和CSS樣式屬性）
答案： 
```
const elem = document.querySelector('.my-div');
const children = elem.children;
for(let i = 0; i < children.length; i++){
  children[i].style.fontWeight = 'bold';
}
```

2. 如何使用JavaScript向已經存在的HTML列表中添加一個新的列表項目？
答案：
```
const list = document.querySelector('#my-list');
const newItem = document.createElement('li');
newItem.textContent = 'New Item';
list.appendChild(newItem);
```

3. 如何使用JavaScript在表格中插入一行新的數據行？
答案：
```
const table = document.querySelector('#my-table');
const newRow = table.insertRow();
const cell1 = newRow.insertCell(0);
const cell2 = newRow.insertCell(1);
cell1.textContent = 'John';
cell2.textContent = 'Doe';
```

4. 如何使用JavaScript替換一個div元素的背景圖像？
答案：
```
const elem = document.querySelector('.my-div');
elem.style.backgroundImage = 'url("new-image.jpg")';
```

5. 如何使用JavaScript將一個特定的div元素設置為可滾動的？
答案：
```
const elem = document.querySelector('.my-div');
elem.style.overflow = 'scroll';
```   

