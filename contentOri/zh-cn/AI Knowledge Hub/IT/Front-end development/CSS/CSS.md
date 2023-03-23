+++
title = "CSS"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题CSS的问题
{{< /ask_chatgpt >}}



1. CSS中的优先级是什么？如何计算优先级？
2. CSS如何实现响应式布局？
3. 谈一下CSS中的盒模型，包括标准盒模型和IE盒模型，以及它们的差异。
4. 如何实现CSS动画？有哪些常见的动画方式？
5. 在CSS中如何实现层叠效果？有哪些属性可以控制样式的层叠顺序？   

## 说明知识
{{< ask_chatgpt >}}
说明CSS并举例
{{< /ask_chatgpt >}}



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

## 汇总重点
{{< ask_chatgpt >}}
条列CSS的重点
{{< /ask_chatgpt >}}



1. 选择器：选择器是用来选取文档中特定元素的CSS选择器。
2. 盒模型：CSS盒模型描述了定位、外边距、边框和内边距等元素大小和位置的方式。
3. 样式属性：CSS属性（如颜色、字体大小、边框）是用来描述元素外观和行为的设置。
4. 布局/排版：CSS布局和排版定义了网页和文档中元素在屏幕上的位置和大小。
5. 动画和过渡效果：CSS可以通过动画和过渡效果添加互动性和动态效果，并改变元素的外观。
6. 响应式设计：CSS可以优化网页在不同屏幕大小和设备上的显示。
7. 浏览器兼容性：CSS的不同版本或浏览器可能会有不同的特性和支援性，需要考虑网页在不同浏览器中的显示和表现。   

## 知识测验
{{< ask_chatgpt >}}
给我5题CSS的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 如何让一个div置中又靠下，且不知道高度?

```
答案： 
div {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, 50%);
}
```

2. 如何让一个背景颜色区块产生光圈效果?

```
答案：
div {
  background-color: #333;
  box-shadow: 0 0 20px #fff;
}
```

3. 如何让一个圆圈镜像成为一个半圆?

```
答案：
div {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  overflow: hidden;
}
div:after {
  content: "";
  display: block;
  width: 100px;
  height: 50px;
  background-color: #333;
  border-radius: 0 0 50% 50%;
  transform: translateX(-50%);
}
```

4. 如何让一个图片hover放大?

```
答案：
div:hover img {
  transform: scale(1.2);
  transition: all 0.2s ease-in-out;
}
```

5. 如何让一个文字保持在固定位置，而不随着滚动而上下偏移?

```
答案：
div {
  position: fixed;
  top: 50px;
  left: 50px;
}
```   

