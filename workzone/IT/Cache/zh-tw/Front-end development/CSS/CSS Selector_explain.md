

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