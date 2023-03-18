

1. 如何利用媒体查询(Media Queries)在不同的装置上显示不同的背景颜色？ 
答案：在 CSS 档案中，可以使用以下的程式码来实现：

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

2. 如何在 Responsive web design 中让一个元件水平置中显示？
答案：可以在该元件外部包覆一个父元件，然后将父元件设为“text-align: center;”，该元件设置display: inline-block;。

例如，HTML 代码如下：

<div class="parentDiv">
  <div class="centeredDiv">我是置中的元件</div>
</div>

然后在 CSS 中进行如下设置：

.parentDiv {
  text-align: center;
}

.centeredDiv {
  display: inline-block;
}

3. 如何在 RWD 中使用 flexbox 布局？
答案：使用如下的 CSS 代码：

.container {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

4. 如何在不同的装置上使用不同的字型大小？
答案：使用媒体查询(Media Queries)：

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

5. 如何在 Responsive web design 中选择适当的图片大小和解析度？
答案：使用不同大小的图片，根据装置大小进行切换。例如，可以使用 img srcset 属性设置多个图片尺寸和解析度，这样浏览器可以根据装置屏幕大小选择最适合的图片。代码如下：

<img src="image.jpg" 
     alt="图片" 
     srcset="image-small.jpg 480w, image-medium.jpg 768w, image-large.jpg 1024w"
     sizes="(max-width: 480px) 440px, (max-width: 768px) 728px, 960px">
     
这段代码会在装置宽度小于 480px 的情况下显示 image-small.jpg，在 480-768px 之间的装置上显示 image-medium.jpg，在宽度大于 768px 的情况下显示 image-large.jpg。sizes 属性指定图片在不同装置大小下的显示尺寸。