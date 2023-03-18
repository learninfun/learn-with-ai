

ARIA是Accessible Rich Internet Applications的縮寫，是一種協助開發者打造無障礙網站的技術。ARIA屬性可讓開發者為HTML元素添加功能，以提供螢幕閱讀器、自動化工具等輔助技術對網站的可訪問性。

以下是幾個ARIA屬性的示例及其用法：

1. aria-label：定義元素的可讀標籤，以提供更詳細的描述。例如：

```html
<button aria-label="關閉對話框">X</button>
```

2. aria-describedby：定義與元素相關的文本說明，以給予更多詳情。例如：

```html
<input type="text" aria-describedby="username-help" />

<p id="username-help">請輸入3到6個字符的用戶名</p>
```

3. aria-disabled：表示元素是否啟用。例如：

```html
<button aria-disabled="true">已停用的按鈕</button>
```

4. aria-haspopup：定義元素是否具有下拉式選單、彈出式對話框等彈出式選單功能。例如：

```html
<a href="#" aria-haspopup="true">下拉式菜單</a>
```

5. aria-checked：表示元素是否被勾選。例如：

```html
<input type="checkbox" aria-checked="true" />
```

6. aria-expanded：表示元素是否展開。例如：

```html
<button aria-expanded="false" aria-controls="menu">展開選單</button>

<ul id="menu">
  <li><a href="#">菜單項目一</a></li>
  <li><a href="#">菜單項目二</a></li>
</ul>
```

以上是幾個常用的ARIA屬性示例，這些屬性能協助開發者使網站更加無障礙。