

1. 如何利用媒體查詢(Media Queries)在不同的裝置上顯示不同的背景顏色？ 
答案：在 CSS 檔案中，可以使用以下的程式碼來實現：

@media (max-width: 767px) {
  body {
    background-color: red;
  }
}

@media (min-width: 768px) and (max-width: 1024px) {
  body {
    background-color: green;
  }
}

@media (min-width: 1025px) {
  body {
    background-color: blue;
  }
}

2. 如何在 Responsive web design 中讓一個元件水平置中顯示？
答案：可以在該元件外部包覆一個父元件，然後將父元件設為「text-align: center;」，該元件設置display: inline-block;。

例如，HTML 代碼如下：

<div class="parentDiv">
  <div class="centeredDiv">我是置中的元件</div>
</div>

然後在 CSS 中進行如下設置：

.parentDiv {
  text-align: center;
}

.centeredDiv {
  display: inline-block;
}

3. 如何在 RWD 中使用 flexbox 佈局？
答案：使用如下的 CSS 代碼：

.container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

4. 如何在不同的裝置上使用不同的字型大小？
答案：使用媒體查詢(Media Queries)：

@media screen and (max-width: 480px) {
  body {
    font-size: 14px;
  }
}

@media screen and (min-width: 481px) and (max-width: 768px) {
  body {
    font-size: 16px;
  }
}

@media screen and (min-width: 769px) {
  body {
    font-size: 18px;
  }
}

5. 如何在 Responsive web design 中選擇適當的圖片大小和解析度？
答案：使用不同大小的圖片，根據裝置大小進行切換。例如，可以使用 img srcset 屬性設置多個圖片尺寸和解析度，這樣瀏覽器可以根據裝置屏幕大小選擇最適合的圖片。代碼如下：

<img src="image.jpg" 
     alt="圖片" 
     srcset="image-small.jpg 480w, image-medium.jpg 768w, image-large.jpg 1024w"
     sizes="(max-width: 480px) 440px, (max-width: 768px) 728px, 960px">
     
這段代碼會在裝置寬度小於 480px 的情況下顯示 image-small.jpg，在 480-768px 之間的裝置上顯示 image-medium.jpg，在寬度大於 768px 的情況下顯示 image-large.jpg。sizes 屬性指定圖片在不同裝置大小下的顯示尺寸。