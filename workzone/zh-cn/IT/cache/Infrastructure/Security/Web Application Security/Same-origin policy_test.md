

1. Same-origin policy 是什么？它是如何保护使用者的安全性？

2. 当不同来源之间的资源需要进行交互时，Same-origin policy 会限制哪些行为？请举例说明。

3. 请解释 Same-origin policy 与 Cross-origin resource sharing (CORS) 的区别。

4. 如果一个网站想要跨域请求其他网站的资源，应该如何解决 Same-origin policy 的限制？

5. 除了在浏览器中执行时，Same-origin policy 也能在什么情况下发挥作用？请举例说明。

答案：

1. Same-origin policy 是浏览器安全机制之一，它强制限制网页脚本等资源只能存取跟自身网页来源相同的资源。这种限制减少了恶意网站恶意攻击用户的可能性。

2. Same-origin policy 通常限制了以下四种跨域操作：Cookie、LocalStorage 和 IndexedDB 存储、读取 DOM 元素内容、AJAX/Fetch 和 WebSocket 的发送和接收。例如，网站 A 的 JavaScript 不能使用 AJAX 向网站 B 发送请求，以防止恶意脚本盗取使用者的敏感信息。

3. Same-origin policy 是浏览器的内置安全特性，用于限制两个不同源的网站之间的资源访问；而 CORS 则是一种机制，允许网站解除跨域资源请求的限制。

4. 网站可以使用 CORS，以前端侧的方式允许跨域请求。透过在请求标头中添加特定设置，如 Access-Control-Allow-Origin，网站可以指示浏览器允许特定网站存取资源。

5. Same-origin policy 也可以在应用程式硬件层面发挥作用，例如浏览器插件可以通过使用 same-origin policy 防止第三方网站对插件的攻击。另外， Content Security Policy (CSP) 也通过实施 same-origin policy 来限制网站的外部资源载入。