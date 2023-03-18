+++
title = "DOM 操作"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題DOM manipulation的問題
{{< /ask_chatgpt >}}



1. 如何使用JavaScript獲取DOM元素的屬性值？

2. 如何使用JavaScript獲取DOM元素的文本內容？

3. 如何使用JavaScript創建新的DOM元素？

4. 如何使用JavaScript更改DOM元素的屬性值？

5. 如何使用JavaScript在DOM元素中添加或刪除子元素？   

## 說明知識
{{< ask_chatgpt >}}
說明DOM manipulation並舉例
{{< /ask_chatgpt >}}



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

## 彙總重點
{{< ask_chatgpt >}}
條列DOM manipulation的重點
{{< /ask_chatgpt >}}



1. 元素選擇和查詢：使用選擇器和方法選擇特定的元素，例如getElementById、querySelectorAll等。

2. 元素創建和添加：使用createElement和appendChild等方法創建和添加新元素。

3. 元素修改和刪除：使用innerHTML和removeChild等方法修改和刪除元素的內容和屬性。

4. 屬性修改和查詢：使用getAttribute和setAttribute等方法修改和查詢元素的屬性。

5. 樣式修改和查詢：使用style和classList等方法修改和查詢元素的樣式。

6. 事件處理程序：使用addEventListener和removeEventListener等方法添加和刪除事件處理程序。

7. AJAX：使用XMLHttpRequest對Web服務器進行非同步操作，從而更新網頁內容。

8. 動態內容：使用innerHTML和createElement等方法動態更新網頁內容、添加新元素和屬性。

9. 數據交互：使用XMLHttpRequest和JSON等技術與Web服務器進行數據交互。

10. 動畫效果：使用CSS和JavaScript創建交互性和動畫效果。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題DOM manipulation的中等難度問題，並在後面列出答案
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

## 網路資料
{{< ask_chatgpt >}}
給我5篇DOM manipulation的網路資料
{{< /ask_chatgpt >}}



1. "Manipulating the DOM with JavaScript" by Mozilla Developer Network: This article explains how to interact with the Document Object Model (DOM) using JavaScript. It covers topics including selecting elements, modifying attributes and text, and creating new elements.

2. "Introduction to DOM Manipulation in JavaScript" by Scotch.io: This tutorial provides an overview of the basics of manipulating the DOM using JavaScript. It covers topics like selecting elements, modifying styles, and adding event listeners.

3. "10 JavaScript DOM Manipulation Techniques You Should Know" by SitePoint: This article covers ten useful techniques for manipulating the DOM using JavaScript. Topics include modifying text and attributes, adding and removing elements, and working with forms.

4. "Manipulating the DOM with jQuery" by W3Schools: This tutorial explains how to use the jQuery library to interact with the DOM. It covers selecting elements, modifying HTML and CSS, and creating animations.

5. "DOM Manipulation with Vanilla JavaScript" by JavaScript for Web Designers: This video tutorial covers basic DOM manipulation techniques using vanilla JavaScript. It includes examples of modifying text and attributes, and creating new elements.   

