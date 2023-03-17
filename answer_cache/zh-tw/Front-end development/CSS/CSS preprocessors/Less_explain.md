

Less是一種CSS預處理器，它可以簡化和加快CSS開發過程。

Less具有較CSS更多的功能，包括變量、嵌套、混合、運算符、函數等。這些功能使得Less代碼更易於組織和維護，同時還可以幫助開發人員編寫更乾淨，更可讀的代碼。

以下是一些Less的例子：

1. 定義變量

```
@primary-color: #007bff;
@secondary-color: #6c757d;

.header {
  background-color: @primary-color;
  color: @secondary-color;
}
```

2. 嵌套

```
header {
  h1 {
    font-size: 3em;
    margin-bottom: 0.5em;
  }
  p {
    font-size: 1.2em;
    margin: 0;
  }
}
```

3. 混合

```
.gradient(@start-color: #fff, @end-color: #000) {
  background: @start-color;
  background: linear-gradient(to bottom, @start-color, @end-color);
}

.header {
  .gradient(#fff, #007bff);
}
```

4. 運算符

```
@base-font-size: 16px;

h1 {
  font-size: @base-font-size + 5px;
}

p {
  font-size: (@base-font-size / 2);
}
```

5. 函數

```
@base-font-size: 16px;

@em-base: (@base-font-size);

body {
  font-size: (@base-font-size);
}

h1 {
  font-size: (@base-font-size * 2.5);
}

h2 {
  font-size: (@base-font-size * 2);
}
```