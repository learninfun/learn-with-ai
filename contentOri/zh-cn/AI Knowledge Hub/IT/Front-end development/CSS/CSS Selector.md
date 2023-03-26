+++
title = "CSS 选择器"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题CSS Selector的中文问题
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
以中文说明CSS Selector并举例
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

## 条列重点
{{< ask_chatgpt >}}
以中文条列CSS Selector的重点
{{< /ask_chatgpt >}}



1. 选择器：指定要样式化的HTML元素。

2. 类别（Class）选择器：透过`.classname`来选择指定类别的HTML元素。

3. ID选择器：透过`#idname`来选择指定ID名称的HTML元素。

4. 元素选择器：透过元素名称（如div、a、p等等）来选择指定元素。

5. 属性选择器：透过`[attribute=value]`来选择指定具有特定属性的HTML元素。

6. 群组选择器：透过使用逗号分隔多个选择器，来同时样式化这些元素。

7. 子元素选择器：透过`parent>child`语法，选择指定元素的直接子元素。

8. 后代元素选择器：透过`ancestor descendant`语法，选择指定元素里的后代元素。

9. 兄弟元素选择器：透过`el1~el2`语法，选择el1元素之后的所有el2元素（二者都是同一父级元素的兄弟元素）。

10. 直接相邻元素选择器：透过`el1 + el2`语法，选择距离el1最近的el2元素（二者都是同一父级元素的相邻元素）。

11. 非选择器：透过`:not(selector)`语法，选择不符合指定条件的元素。

12. 整体选择器：透过`*`语法，选择所有元素。

13. 多重选择器：透过使用逗号分隔多个条件，来选择符合其中任意一个条件的元素。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题CSS Selector的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 找出所有class名称为"menu-item"的连结，但不包括在"
nav-header"里面的。
```css
.menu-item:not(.nav-header) a
```

2. 找出所有属于列表元素ol和ul的子元素li，但不包括子元素的子元素。
```css
ol>li, ul>li
```

3. 具有class名称为"btn"，但不包括class名称为"disabled"的按钮。
```css
.btn:not(.disabled)
```

4. 查找最后一个class名称为"post"的div元素中的所有直接子元素。
```css
div.post:last-child > *
```

5. 找出父元素为div，第一个子元素是header的段落元素p。
```css
div > header + p
```

答案仅供参考，可能不唯一。   

