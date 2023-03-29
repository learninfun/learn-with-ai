

Content Security Policy (CSP)是一種Web安全機制，可以限制網頁內容的來源和資源，防止惡意腳本攻擊和跨站腳本攻擊 (XSS)。CSP可以通過HTTP響應頭部的Content-Security-Policy標頭添加到網頁中，指定哪些資源是被允許取得的。

舉例來說，假設我們有一個網站，其中包含以下內容：

```html
<script src="https://cdn.example.com/script.js"></script>
<img src="https://example.com/image.jpg">
```

為了使用CSP，我們可以通過HTTP響應頭部添加以下設置：

```
Content-Security-Policy: default-src 'self' https://cdn.example.com
```

這個CSP設置指定了只允許自身網站 (self) 和 https://cdn.example.com 這個域名下的資源載入，所以如果有任何嘗試從其他的域名載入資源的行為，瀏覽器會擋下它們。

另外，CSP也可以用來防止XSS攻擊，例如：

```
Content-Security-Policy: script-src 'self' 'unsafe-inline' 'unsafe-eval'
```

這個CSP設置指定只允許自身網站執行JS程式碼，並且禁止直接插入JS程式碼。這樣就可以防止嘗試通過插入惡意腳本攻擊網站的攻擊行為。