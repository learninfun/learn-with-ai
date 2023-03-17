

Sass是一种CSS预处理器，它使得CSS的编写更加容易且高效。Sass相较于原本的CSS，可以更好地支持变数、巢状规则、混入（Mixin）、继承等高级特性，并支持更简洁明了的编写方式。

以下是一个简单的Sass示例：

```scss
// 定义变数
$primary-color: #007bff;

// 编写混入
@mixin btn-style {
  background-color: $primary-color;
  color: #fff;
  padding: 10px 20px;
  border-radius: 5px;
}

// 使用混入
.btn {
  @include btn-style;
}

// 编写巢状规则
.card {
  background-color: #f7f7f7;
  padding: 10px;

  .card-title {
    font-size: 24px;
    margin-bottom: 10px;
  }

  .card-body {
    font-size: 16px;
    line-height: 1.5;
  }
}
```

在上面的示例中，定义了一个名为`$primary-color`的变数，然后编写了一个名为`btn-style`的混入，将`$primary-color`用于简单的按钮样式中，然后在`.btn`的类名中调用此混入以使用按钮样式。

同时，示例中使用了巢状的规则，使设计更加易于理解。例如，`.card`的子元素`.card-title`和`.card-body`均可在`.card`规则中定义，且不需使用多个CSS层级以明确指定每个规则。