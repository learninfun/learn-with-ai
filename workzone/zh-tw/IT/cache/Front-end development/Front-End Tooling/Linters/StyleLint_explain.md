

StyleLint是一種基於靜態分析的工具，用於檢查CSS、SCSS和LESS等CSS預處理器代碼的風格和錯誤。它可以幫助開發人員保持一致的CSS預設風格，並且可以找到一些語法錯誤、不良的用法、遺漏的樣式規則等問題。

下面舉個例子，假設我們有以下的CSS代碼：

```css
.foo {
  display: block;
    margin:10px;
  background-color:red; }

.bar{
    font-size: 12px;
}
```

如果使用StyleLint檢測這段代碼，會得到以下結果：

```
/path/to/file.css
3:3   error    Expected indentation of 2 spaces (indentation)
4:4   error    Expected single space after ":" (declaration-colon-space-after)
4:4   error    Expected a trailing semicolon (declaration-block-trailing-semicolon)
6:1   error    Expected indentation of 2 spaces (indentation)
6:19  error    Expected a leading zero (number-leading-zero)
```

可以發現，StyleLint能夠提示我們第3行的縮進不正確、第4行的冒號後沒有空格、缺少分號等錯誤，並且還能夠提示第6行字體大小應該以0開頭的問題。

透過StyleLint，我們可以避免不必要的代碼錯誤，維護良好的CSS代碼風格。