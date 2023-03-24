

CSS (Cascading Style Sheets) 是用來描述網頁上元素的樣式、佈局和特效的語言，它可以將網站的外觀裝扮完美地呈現出來。它與HTML共同組成了網頁的兩個主要元素。

以下是一些CSS的例子：

1. 樣式：

body {
  background-color: #f2f2f2;
  font-family: Arial, sans-serif;
}
h1 {
  color: #333;
  font-size: 36px;
}

這段 CSS 代碼定義了整個頁面的背景顏色和字體，以及標題元素的顏色和字體大小。

2. 佈局：

.container {
  max-width: 1200px;
  margin: 0 auto;
}
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  grid-gap: 20px;
}

這段 CSS 代碼設置了元素容器的寬度和水平居中，並利用Grid佈局設置了具有自適應列寬度和間距的元素網格。

3. 特效：

.btn {
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  padding: 10px 20px;
  text-decoration: none;
  transition: all .3s ease;
  cursor: pointer;
}
.btn:hover {
  background-color: #0056b3;
}

這段 CSS 代碼為按鈕元素添加了背景色、文字顏色、邊框、圓角、內邊距和動畫效果。當鼠標懸停在它上面時，背景顏色會漸變為不同的顏色。

綜上所述，CSS 為網站提供了豐富的樣式、佈局和特效選擇，讓網站外觀更加美觀、吸引人，也更具有互動性。