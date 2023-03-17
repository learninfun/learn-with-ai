

Same-origin policy 是一个浏览器安全性机制，限制了一个网页文档（document）或脚本（script）在浏览器和同一个源（origin）下的资源交互，从而防止潜在的跨站脚本攻击（XSS）或跨站请求伪造（CSRF）等安全问题。

一个源通常由三部分组成：协议（Protocol）、域名（Domain）和端口（Port）。只有三个部分都相同的资源才被认为是同一个源，否则就被视为不同的源。 譬如：

- http://www.example.com 和 https://www.example.com 是不同的源。
- http://www.example.com 和 http://example.com 是不同的源。
- http://www.example.com:80 和 http://www.example.com:8080 是不同的源。

一些常见的Same-origin policy限制包括：

- JavaScript在同一源下的资源上想象可以自由使用，但他们无法访问其他源的资源。 假设有一个JavaScript应用程序在example.com上运行，那么就可以访问它同一源下的任何资源，例如example.com/about.html；但是，如果应用程式尝试访问example.net上的资源，则将会被同源政策阻止。

- 浏览器在送出不同源的 AJAX 请求或向 iframe 内载入内容时遵守同源政策。

- 不同域名下设置的 Cookie 不会被 JavaScript 访问，也不会被浏览器发送给不同域名。

- 一些HTML5 API （例如 Geolocation API 或者 Web Storage API）将遵循Same-origin policy，并禁止从非同一源的脚本中访问或修改数据。

简单地说，当Same-origin policy被强制执行时，浏览器只允许当前网页文档获得它自己的资源，而只能与同一个源相关的资源进行交互。这可以减少安全威胁，针对用户的浏览器提供更加安全的环境。