+++
title = "CSS 選擇器"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題CSS Selector的中文問題
{{< /ask_chatgpt >}}



1. 如何選擇所有的段落元素？

答案：可以使用 p 作為選擇器，例如：p { ... }

2. 如何選擇所有的超連結元素？

答案：可以使用 a 作為選擇器，例如：a { ... }

3. 如何選擇以特定顏色為背景的元素？

答案：可以使用 [style="background-color: #ff0000;"] 作為選擇器，例如：[style="background-color: #ff0000;"] { ... }

4. 如何選擇屬性名稱以特定字元開頭的元素？

答案：可以使用 [attribute^="value"] 作為選擇器，例如：[class^="box-"] { ... }

5. 如何選擇最後一個元素？

答案：可以使用 :last-child 作為選擇器，例如：div:last-child { ... }   

## 說明知識
{{< ask_chatgpt >}}
以中文說明CSS Selector並舉例
{{< /ask_chatgpt >}}



CSS Selector是一種CSS選擇器，可用於選擇HTML文檔中特定的元素，並對其應用樣式。以下是一些常用的CSS選擇器及其示例：

1. 元素選擇器（Element Selector）：

選擇所有p元素並應用樣式：

```CSS
p {
  color: red;
}
```

2. 類選擇器（Class Selector）：

選擇class為"example"的所有元素並應用樣式：

```CSS
.example {
  color: blue;
}
```

3. ID選擇器（ID Selector）：

選擇ID為"header"的元素並應用樣式：

```CSS
#header {
  background-color: gray;
}
```

4. 子元素選擇器（Child Selector）：

選擇所有div元素中的p元素並應用樣式：

```CSS
div > p {
  font-size: 16px;
}
```

5. 相鄰兄弟選擇器（Adjacent Sibling Selector）：

選擇class為"example"的元素之後的第一個p元素並應用樣式：

```CSS
.example + p {
  color: green;
}
```

6. 屬性選擇器（Attribute Selector）：

選擇所有包含title屬性的a元素並應用樣式：

```CSS
a[title] {
  text-decoration: underline;
}
```

以上是一些常用的CSS選擇器及其示例。這些選擇器可幫助您更好地選擇HTML文檔中的元素並對其應用樣式。   

## 彙總重點
{{< ask_chatgpt >}}
以中文條列CSS Selector的重點
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

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題CSS Selector的中等難度問題，並在後面列出答案
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

