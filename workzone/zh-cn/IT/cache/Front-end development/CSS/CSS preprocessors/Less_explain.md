

Less是一种CSS预处理器，它可以简化和加快CSS开发过程。

Less具有较CSS更多的功能，包括变量、嵌套、混合、运算符、函数等。这些功能使得Less代码更易于组织和维护，同时还可以帮助开发人员编写更干净，更可读的代码。

以下是一些Less的例子：

1. 定义变量

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

4. 运算符

```
@base-font-size: 16px;

h1 {
  font-size: @base-font-size + 5px;
}

p {
  font-size: (@base-font-size / 2);
}
```

5. 函数

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