+++
title = "CSS 选择器"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题CSS Selector的问题
{{< /ask_chatgpt >}}



1. 如何选择所有的段落元素？

答案：可以使用 p 作为选择器，例如：p { ... }

2. 如何选择所有的超连结元素？

答案：可以使用 a 作为选择器，例如：a { ... }

3. 如何选择以特定颜色为背景的元素？

答案：可以使用 [style="background-color: #ff0000;"] 作为选择器，例如：[style="background-color: #ff0000;"] { ... }

4. 如何选择属性名称以特定字元开头的元素？

答案：可以使用 [attribute^="value"] 作为选择器，例如：[class^="box-"] { ... }

5. 如何选择最后一个元素？

答案：可以使用 :last-child 作为选择器，例如：div:last-child { ... }   

## 说明知识
{{< ask_chatgpt >}}
说明CSS Selector并举例
{{< /ask_chatgpt >}}



CSS Selector是一种CSS选择器，可用于选择HTML文档中特定的元素，并对其应用样式。以下是一些常用的CSS选择器及其示例：

1. 元素选择器（Element Selector）：

选择所有p元素并应用样式：

```CSS
p {
  color: red;
}
```

2. 类选择器（Class Selector）：

选择class为"example"的所有元素并应用样式：

```CSS
.example {
  color: blue;
}
```

3. ID选择器（ID Selector）：

选择ID为"header"的元素并应用样式：

```CSS
#header {
  background-color: gray;
}
```

4. 子元素选择器（Child Selector）：

选择所有div元素中的p元素并应用样式：

```CSS
div > p {
  font-size: 16px;
}
```

5. 相邻兄弟选择器（Adjacent Sibling Selector）：

选择class为"example"的元素之后的第一个p元素并应用样式：

```CSS
.example + p {
  color: green;
}
```

6. 属性选择器（Attribute Selector）：

选择所有包含title属性的a元素并应用样式：

```CSS
a[title] {
  text-decoration: underline;
}
```

以上是一些常用的CSS选择器及其示例。这些选择器可帮助您更好地选择HTML文档中的元素并对其应用样式。   

## 汇总重点
{{< ask_chatgpt >}}
条列CSS Selector的重点
{{< /ask_chatgpt >}}



1. 選擇器：指定要樣式化的HTML元素。

2. 類別（Class）選擇器：透過`.classname`來選擇指定類別的HTML元素。

3. ID選擇器：透過`#idname`來選擇指定ID名稱的HTML元素。

4. 元素選擇器：透過元素名稱（如div、a、p等等）來選擇指定元素。

5. 屬性選擇器：透過`[attribute=value]`來選擇指定具有特定屬性的HTML元素。

6. 群組選擇器：透過使用逗號分隔多個選擇器，來同時樣式化這些元素。

7. 子元素選擇器：透過`parent>child`語法，選擇指定元素的直接子元素。

8. 後代元素選擇器：透過`ancestor descendant`語法，選擇指定元素裡的後代元素。

9. 兄弟元素選擇器：透過`el1~el2`語法，選擇el1元素之後的所有el2元素（二者都是同一父級元素的兄弟元素）。

10. 直接相鄰元素選擇器：透過`el1 + el2`語法，選擇距離el1最近的el2元素（二者都是同一父級元素的相鄰元素）。

11. 非選擇器：透過`:not(selector)`語法，選擇不符合指定條件的元素。

12. 整體選擇器：透過`*`語法，選擇所有元素。

13. 多重選擇器：透過使用逗號分隔多個條件，來選擇符合其中任意一個條件的元素。   

## 知识测验
{{< ask_chatgpt >}}
给我5题CSS Selector的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 找出所有class名稱為"menu-item"的連結，但不包括在"
nav-header"裡面的。
```css
.menu-item:not(.nav-header) a
```

2. 找出所有屬於列表元素ol和ul的子元素li，但不包括子元素的子元素。
```css
ol>li, ul>li
```

3. 具有class名稱為"btn"，但不包括class名稱為"disabled"的按鈕。
```css
.btn:not(.disabled)
```

4. 查找最後一個class名稱為"post"的div元素中的所有直接子元素。
```css
div.post:last-child > *
```

5. 找出父元素為div，第一個子元素是header的段落元素p。
```css
div > header + p
```

答案僅供參考，可能不唯一。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇CSS Selector的网络数据
{{< /ask_chatgpt >}}



1. "CSS Selectors" from W3schools.com: This comprehensive page explains CSS selectors in detail, including attribute selectors, class selectors, id selectors, and more.

2. "CSS Selectors Tutorial" by CSS-tricks.com: This tutorial covers the basics of selecting elements with CSS, including some advanced techniques.

3. "The 30 CSS Selectors You Must Memorize" by Codecademy.com: This list breaks down the most important CSS selectors, including pseudo-classes, pseudo-elements, and combinators.

4. "Advanced CSS Selectors" by Smashing Magazine: This article explores more advanced CSS selectors, including structural pseudo-classes, sibling selectors, and more.

5. "CSS Selectors Cheatsheet" by Hostinger.com: This handy cheatsheet provides quick reference for commonly used CSS selectors, including examples of how they can be used in code.   

