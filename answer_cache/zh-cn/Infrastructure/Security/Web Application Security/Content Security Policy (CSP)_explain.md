

Content Security Policy (CSP)是一种Web安全机制，可以限制网页内容的来源和资源，防止恶意脚本攻击和跨站脚本攻击 (XSS)。CSP可以通过HTTP响应头部的Content-Security-Policy标头添加到网页中，指定哪些资源是被允许取得的。

举例来说，假设我们有一个网站，其中包含以下内容：

```html
<script src="https://cdn.example.com/script.js"></script>
<img src="https://example.com/image.jpg">
```

为了使用CSP，我们可以通过HTTP响应头部添加以下设置：

```
Content-Security-Policy: default-src 'self' https://cdn.example.com
```

这个CSP设置指定了只允许自身网站 (self) 和 https://cdn.example.com 这个域名下的资源载入，所以如果有任何尝试从其他的域名载入资源的行为，浏览器会挡下它们。

另外，CSP也可以用来防止XSS攻击，例如：

```
Content-Security-Policy: script-src 'self' 'unsafe-inline' 'unsafe-eval'
```

这个CSP设置指定只允许自身网站执行JS程式码，并且禁止直接插入JS程式码。这样就可以防止尝试通过插入恶意脚本攻击网站的攻击行为。