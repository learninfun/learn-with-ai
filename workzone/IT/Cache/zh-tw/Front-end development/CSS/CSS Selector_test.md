

1. 找出所有class名稱為"menu-item"的連結，但不包括在"
nav-header"裡面的。
```css
.menu-item:not(.nav-header) a
```

2. 找出所有屬於列表元素ol和ul的子元素li，但不包括子元素的子元素。
```css
ol>li, ul>li
```

3. 具有class名稱為"btn"，但不包括class名稱為"disabled"的按鈕。
```css
.btn:not(.disabled)
```

4. 查找最後一個class名稱為"post"的div元素中的所有直接子元素。
```css
div.post:last-child > *
```

5. 找出父元素為div，第一個子元素是header的段落元素p。
```css
div > header + p
```

答案僅供參考，可能不唯一。