

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