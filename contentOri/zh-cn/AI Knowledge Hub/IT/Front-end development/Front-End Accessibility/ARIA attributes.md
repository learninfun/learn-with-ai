+++
title = "ARIA属性"
weight = "2"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题ARIA attributes的问题
{{< /ask_chatgpt >}}



1. 在一個網頁上的對話框內，你可以使用哪一個ARIA屬性來指示當前頁面內正在激活的控制項？
2. 在一個網頁上的表格中，你可以使用哪一個ARIA屬性來表明表格行的內容是數據還是標題？
3. 當一個網頁具有多個可視導航菜單時，你可以使用哪一個ARIA屬性來標識不同的菜單，以便更好地區分它們？
4. 在一個可展開的控件中，你可以使用哪一個ARIA屬性來指示該控件目前的狀態是已展開還是已折疊？
5. 在一個動態輪播中，你可以使用哪一個ARIA屬性來指示當前正在顯示的輪播內容，以便更好地向用戶傳達輪播信息？   

## 说明知识
{{< ask_chatgpt >}}
说明ARIA attributes并举例
{{< /ask_chatgpt >}}



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

## 汇总重点
{{< ask_chatgpt >}}
条列ARIA attributes的重点
{{< /ask_chatgpt >}}



以下是ARIA属性的重点：

1. 角色(attribute-role)属性：定义了元素在页面中扮演的角色，如菜单、按钮、文本框等。

2. 描述(attribute-description)属性：提供了元素的额外描述信息，以便于屏幕阅读器能够更好的理解该元素的含义。

3. 标签(attribute-label)属性：为元素提供了更容易被理解的标签，通常用来替代没有语义的标签，如checkbox、radio等。

4. tab-index属性：定义了元素在页面中的tab键次序，以便于键盘导航。

5. 表示状态的属性（attribute-state）：主要有三种表示状态的属性aria-checked，aria-selected和aria-disabled，用于标记元素的状态。

6. 语义组(attribute-set)：用于将有关联性的元素如radiogroup、menu等组合到一个语义组中，以便于屏幕阅读器能够更好的理解相互关联的元素。

7. 辅助功能的提示信息（attribute-popup）：定义了与元素相关的提示信息，如弹出窗口、帮助文本等。这些提示信息可以帮助用户更好地理解元素的功能和用途。   

## 知识测验
{{< ask_chatgpt >}}
给我5题ARIA attributes的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 什么是aria-current属性？
答案：aria-current属性指定元素的当前状态。它可以用来指示哪些元素是当前选择的或活动的。

2. 如何使用aria-describedby属性？
答案：aria-describedby属性可以阐述与元素相关的附加描述。它可以被用来作为辅助技术的提示，也可以用来提供更详细的元素内容。

3. 什么是aria-expanded属性？
答案：aria-expanded属性用于指示具有展开/收起功能的元素的当前状态。它可以被设置为“true”或“false”，用于指示元素是否打开或关闭。

4. 如何使用aria-label属性？
答案：aria-label属性用于为元素提供可读的名称，通常被用于没有文本描述的图像、按钮和表单控件等元素。这可以帮助辅助技术用户更好地理解页面内容。

5. 什么是aria-hidden属性？
答案：aria-hidden属性可以用于隐藏屏幕阅读器不应该读取的元素。当aria-hidden属性被设置为“true”时，元素将被隐藏并从辅助技术的焦点序列中移除。   

