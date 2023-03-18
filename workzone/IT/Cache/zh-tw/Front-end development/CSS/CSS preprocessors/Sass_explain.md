

Sass是一種CSS預處理器，它使得CSS的編寫更加容易且高效。Sass相較於原本的CSS，可以更好地支持變數、巢狀規則、混入（Mixin）、繼承等高級特性，並支持更簡潔明瞭的編寫方式。

以下是一個簡單的Sass示例：

```scss
// 定義變數
$primary-color: #007bff;

// 編寫混入
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

// 編寫巢狀規則
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

在上面的示例中，定義了一個名為`$primary-color`的變數，然後編寫了一個名為`btn-style`的混入，將`$primary-color`用於簡單的按鈕樣式中，然後在`.btn`的類名中調用此混入以使用按鈕樣式。

同時，示例中使用了巢狀的規則，使設計更加易於理解。例如，`.card`的子元素`.card-title`和`.card-body`均可在`.card`規則中定義，且不需使用多個CSS層級以明確指定每個規則。