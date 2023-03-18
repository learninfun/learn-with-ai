

1. 如何在Sass中定义一个可重复使用的Mixin (混合物)？
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

2. 如何在Sass中定义一个变数，用于储存颜色？
答：使用 $ 符号开头，例如：
```
$primary-color: #007bff;
```

3. 如何使用Sass的运算功能计算两个长度的和？
答：使用加号，例如：
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
或使用 for 循环：
```
@for $i from 1 through 3 {
  .class-#{$i} {
    width: 10px * $i;
  }
}
```

5. 如何使用 Sass 的 extend 扩展选择器的属性？
答： 使用 @extend, 用于扩展已有选择器的属性， 例如：
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