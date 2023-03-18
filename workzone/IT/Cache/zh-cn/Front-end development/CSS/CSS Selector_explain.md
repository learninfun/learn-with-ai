

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