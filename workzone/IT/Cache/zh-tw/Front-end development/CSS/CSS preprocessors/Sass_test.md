

1. 如何在Sass中定義一個可重複使用的Mixin (混合物)？
答：使用 @mixin 和 @include。例如：
```
@mixin box-shadow($shadow...) {
  -webkit-box-shadow: $shadow;
     -moz-box-shadow: $shadow;
          box-shadow: $shadow;
}

.box {
  @include box-shadow(0px 0px 10px rgba(0, 0, 0, 0.5));
}
```

2. 如何在Sass中定義一個變數，用於儲存顏色？
答：使用 $ 符號開頭，例如：
```
$primary-color: #007bff;
```

3. 如何使用Sass的運算功能計算兩個長度的和？
答：使用加號，例如：
```
width: 100px + 50px;
```

4. 如何使用Sass的流程控制功能（if、for、each、while）？
答：例如，使用 if ：
```
@if lightness($color) > 50% {
  background-color: black;
} @else {
  background-color: white;
}
```
或使用 for 循環：
```
@for $i from 1 through 3 {
  .class-#{$i} {
    width: 10px * $i;
  }
}
```

5. 如何使用 Sass 的 extend 擴展選擇器的屬性？
答： 使用 @extend, 用於擴展已有選擇器的屬性， 例如：
```
.sizing {
  font-size: 20px;
  padding: 10px;
}

.btn {
  @extend .sizing;
  background-color: $primary-color;
  color: white;
}
```