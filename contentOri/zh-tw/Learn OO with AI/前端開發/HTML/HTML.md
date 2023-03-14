## 習題預習
<<ASK_CHATGPT>>  
給我5題HTML的問題  
  
<<ANS_CHATGPT>>  


1. HTML是什麼？

2. 如何建立HTML文件？

3. HTML中有哪些基本元素？

4. 如何嵌套元素？

5. 如何使用樣式來設計網頁？  

## 說明知識
<<ASK_CHATGPT>>  
說明HTML並舉例  
  
<<ANS_CHATGPT>>  


HTML全名為HyperText Markup Language，是一種標準的網頁程式設計語言，用於創建網頁。它定義了網頁應該如何呈現，包括文本、圖像、連結、表格、音頻、視頻等內容的呈現方式。HTML是用標記語言的形式呈現，即使用標記指示瀏覽器顯示網頁的不同元素。

以下是一個簡單的HTML例子：

```html
<!DOCTYPE html>
<html>
  <head>
    <title>我的第一個網頁</title>
  </head>
  <body>
    <h1>歡迎來到我的網站</h1>
    <p>這是一段簡短的文字，用來演示HTML的基本結構。</p>
    <img src="https://picsum.photos/200" alt="一張隨機圖片">
    <a href="https://www.google.com">點擊這裡前往Google</a>
  </body>
</html>
```

這個HTML代碼表示一個簡單的網頁，包括標題、一段文字、一張圖片和一個連結。以下是對開頭和結尾標籤的解釋：

- `<!DOCTYPE html>`：聲明這是一個HTML5文檔。
- `<html>`：定義了整個HTML文件的開始。
- `<head>`：包含所有頁面的元數據，如標題、腳本和CSS。
- `<title>`：定義在瀏覽器標題欄中顯示的文本。
- `<body>`：包含在網頁上顯示的所有可見內容。
- `<h1>`：定義一個大標題。
- `<p>`：定義一個段落。
- `<img>`：定義一個圖像，包括它的來源和可選的替代文字。
- `<a>`：定義一個超鏈接，包括它的URL和顯示文本。  

## 彙總重點
<<ASK_CHATGPT>>  
條列HTML的重點  
  
<<ANS_CHATGPT>>  


1. HTML是一种标记语言，用于创建网页。

2. HTML标签用于指定网页内容的结构和样式。

3. HTML文档通常由以下几个部分组成: 页眉、页脚、导航栏、主体内容、图像和表格。

4. HTML标记通常由开始标记、内容和结束标记组成，例如：<p>这是一段文本</p>

5. HTML标记可以用于描述文本样式、页眉、页脚、图像、链接等等。

6. HTML可以与CSS和JavaScript等其他技术结合使用，以创建富有交互性和动态性的网页。

7. HTML5是最新的HTML标准，包括更多的功能和元素，以便更好地支持移动设备和多媒体。  

## 知識測驗
<<ASK_CHATGPT>>  
給我5題HTML的中等難度問題，並在後面列出答案  
  
<<ANS_CHATGPT>>  


1. 如何在HTML中插入圖片？並且讓圖片可以點擊跳轉到其他頁面

答案： 

```
<img src="圖片連結" alt="圖片文字說明" usemap="#mapname"> 

<map name="mapname">
   <area shape="rect" coords="x1,y1,x2,y2" href="跳轉頁面連結">
</map>
```



2. 如何在HTML中設置網頁的標題和網頁的icon？

答案：

```
<head>
   <title>網頁標題</title>
   <link rel="icon" href="favicon.ico">
</head>
```



3. 如何在HTML中創建一個表格？並且設置表格邊框和表頭等屬性？

答案：

```
<table border="1">
   <caption>表格標題</caption>
   <thead>
     <tr>
        <th>表頭1</th>
        <th>表頭2</th>
     </tr>
   </thead>
   <tbody>
      <tr>
         <td>單元格1</td>
         <td>單元格2</td>
      </tr>
      <tr>
         <td>單元格3</td>
         <td>單元格4</td>
      </tr>
   </tbody>
</table>
```



4. 如何在HTML中創建一個下拉菜單？並且設置選項和默認選項。

答案：

```
<select>
   <option value="option1">選項1</option>
   <option value="option2">選項2</option>
   <option value="option3" selected>默認選項</option>
</select>
```



5. 如何在HTML中創建一個可輸入的表單？並且設置表單的提交方式和內容。

答案：

```
<form action="提交頁面" method="post">
   <label>請輸入姓名：</label>
   <input type="text" name="name"><br>
   <label>請輸入郵箱：</label>
   <input type="email" name="email"><br>
   <label>請輸入密碼：</label>
   <input type="password" name="password"><br>
   <input type="submit" value="提交">
</form>
```  

## 網路資料
<<ASK_CHATGPT>>  
給我5篇HTML的網路資料  
  
<<ANS_CHATGPT>>  


1. W3Schools
W3Schools是一個線上的學習平台，提供各種網站開發相關的資源，包括HTML、CSS、JavaScript等。他們的網站也是用HTML建立的，而且很好看、很容易學習。

2. MDN Web Docs
MDN Web Docs是由Mozilla所建立的一個資源庫，提供各種網站開發相關的文檔和教程，包括HTML、CSS、JavaScript等。他們的網站也是用HTML建立的，而且非常詳細，對於入門者很友好。

3. GitHub Pages
GitHub Pages是由GitHub提供的一種網站托管服務，支援靜態網頁的建立和部署。你可以用HTML、CSS、JavaScript等建立自己的網站，然後通過GitHub Pages來展示給其他人看。

4. Stack Overflow
Stack Overflow是一個非常受歡迎的問答社區，開發者可以在這裡找到各種編程方面的解答和建議。他們的網站也是用HTML建立的，而且非常簡潔、易讀。

5. CodePen
CodePen是一個開發者社區，提供各種程式碼展示和分享平台。你可以在這裡發佈自己的HTML、CSS、JavaScript等作品，並與其他人互相分享、交流。他們的網站也是用HTML建立的，而且非常時髦、有趣。  

