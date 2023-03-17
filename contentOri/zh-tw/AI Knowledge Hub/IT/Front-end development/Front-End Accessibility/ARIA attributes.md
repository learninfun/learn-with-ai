## 習題預習
{{< ask_chatgpt >}}
給我5題ARIA attributes的問題
{{< /ask_chatgpt >}}



1. 在一個網頁上的對話框內，你可以使用哪一個ARIA屬性來指示當前頁面內正在激活的控制項？
2. 在一個網頁上的表格中，你可以使用哪一個ARIA屬性來表明表格行的內容是數據還是標題？
3. 當一個網頁具有多個可視導航菜單時，你可以使用哪一個ARIA屬性來標識不同的菜單，以便更好地區分它們？
4. 在一個可展開的控件中，你可以使用哪一個ARIA屬性來指示該控件目前的狀態是已展開還是已折疊？
5. 在一個動態輪播中，你可以使用哪一個ARIA屬性來指示當前正在顯示的輪播內容，以便更好地向用戶傳達輪播信息？   

## 說明知識
{{< ask_chatgpt >}}
說明ARIA attributes並舉例
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

## 彙總重點
{{< ask_chatgpt >}}
條列ARIA attributes的重點
{{< /ask_chatgpt >}}



以下是ARIA属性的重点：

1. 角色(attribute-role)属性：定义了元素在页面中扮演的角色，如菜单、按钮、文本框等。

2. 描述(attribute-description)属性：提供了元素的额外描述信息，以便于屏幕阅读器能够更好的理解该元素的含义。

3. 标签(attribute-label)属性：为元素提供了更容易被理解的标签，通常用来替代没有语义的标签，如checkbox、radio等。

4. tab-index属性：定义了元素在页面中的tab键次序，以便于键盘导航。

5. 表示状态的属性（attribute-state）：主要有三种表示状态的属性aria-checked，aria-selected和aria-disabled，用于标记元素的状态。

6. 语义组(attribute-set)：用于将有关联性的元素如radiogroup、menu等组合到一个语义组中，以便于屏幕阅读器能够更好的理解相互关联的元素。

7. 辅助功能的提示信息（attribute-popup）：定义了与元素相关的提示信息，如弹出窗口、帮助文本等。这些提示信息可以帮助用户更好地理解元素的功能和用途。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題ARIA attributes的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 什麼是aria-current屬性？
答案：aria-current屬性指定元素的當前狀態。它可以用來指示哪些元素是當前選擇的或活動的。

2. 如何使用aria-describedby屬性？
答案：aria-describedby屬性可以闡述與元素相關的附加描述。它可以被用來作為輔助技術的提示，也可以用來提供更詳細的元素內容。

3. 什麼是aria-expanded屬性？
答案：aria-expanded屬性用於指示具有展開/收起功能的元素的當前狀態。它可以被設置為“true”或“false”，用於指示元素是否打開或關閉。

4. 如何使用aria-label屬性？
答案：aria-label屬性用於為元素提供可讀的名稱，通常被用於沒有文本描述的圖像、按鈕和表單控件等元素。這可以幫助輔助技術用戶更好地理解頁面內容。

5. 什麼是aria-hidden屬性？
答案：aria-hidden屬性可以用於隱藏屏幕閱讀器不應該讀取的元素。當aria-hidden屬性被設置為“true”時，元素將被隱藏並從輔助技術的焦點序列中移除。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇ARIA attributes的網路資料
{{< /ask_chatgpt >}}



1. "ARIA Attributes" by Mozilla:
https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA/ARIA_Techniques/Using_the_aria-label_attribute

This article by Mozilla covers several ARIA attributes, including the aria-label attribute. It explains how this attribute can be used to provide a more descriptive label for an element, improving its accessibility for people using screen readers and other assistive technologies.

2. "ARIA Landmarks" by WebAIM:
https://webaim.org/techniques/aria/landmarks/

This article by WebAIM explores ARIA landmarks, which are used to indicate the purpose and role of different areas of a webpage. The article explains how landmarks can help users navigate and understand a page's content, and provides examples of how to implement different types of landmarks using ARIA attributes.

3. "ARIA Attributes" by W3C:
https://www.w3.org/TR/wai-aria-1.2/#intro-aria

This page from the Web Accessibility Initiative (WAI) of the World Wide Web Consortium (W3C) provides an overview of ARIA attributes and their purpose. It also includes links to more detailed documentation about specific attributes, as well as guidance on how to use ARIA effectively to improve accessibility.

4. "ARIA Hidden and aria-describedby" by Deque:
https://www.deque.com/blog/aria-hidden-and-aria-describedby/

This article by Deque covers two common ARIA attributes: aria-hidden and aria-describedby. It explains how these attributes can be used to hide or describe content on a webpage, and provides examples of situations where they might be useful.

5. "ARIA Roles" by Accessible University:
https://accessibleuniversity.com/courses/the-what-and-why-of-aria/lessons/the-roles-of-aria/

This article by Accessible University focuses specifically on ARIA roles, which are used to describe the type or purpose of an element on a webpage. The article provides a list of common ARIA roles and explains how they can be used to improve accessibility for different types of users.   

