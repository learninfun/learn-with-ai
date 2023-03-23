+++
title = "同源策略"
weight = "6"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Same-origin policy的问题
{{< /ask_chatgpt >}}



1. 什么是Same-origin policy? 

2. Same-origin policy对网站安全有什么影响？ 

3. 如何遵循Same-origin policy，以确保网站的安全？ 

4. Same-origin policy如何保护用户的隐私？ 

5. 在什么情况下，可能需要违反Same-origin policy，例如跨网站共享资源或与第三方服务集成？   

## 说明知识
{{< ask_chatgpt >}}
说明Same-origin policy并举例
{{< /ask_chatgpt >}}



Same-origin policy 是一个浏览器安全性机制，限制了一个网页文档（document）或脚本（script）在浏览器和同一个源（origin）下的资源交互，从而防止潜在的跨站脚本攻击（XSS）或跨站请求伪造（CSRF）等安全问题。

一个源通常由三部分组成：协议（Protocol）、域名（Domain）和端口（Port）。只有三个部分都相同的资源才被认为是同一个源，否则就被视为不同的源。 譬如：

- http://www.example.com 和 https://www.example.com 是不同的源。
- http://www.example.com 和 http://example.com 是不同的源。
- http://www.example.com:80 和 http://www.example.com:8080 是不同的源。

一些常见的Same-origin policy限制包括：

- JavaScript在同一源下的资源上想像可以自由使用，但他们无法访问其他源的资源。 假设有一个JavaScript应用程序在example.com上运行，那么就可以访问它同一源下的任何资源，例如example.com/about.html；但是，如果应用程式尝试访问example.net上的资源，则将会被同源政策阻止。

- 浏览器在送出不同源的 AJAX 请求或向 iframe 内载入内容时遵守同源政策。

- 不同域名下设置的 Cookie 不会被 JavaScript 访问，也不会被浏览器发送给不同域名。

- 一些HTML5 API （例如 Geolocation API 或者 Web Storage API）将遵循Same-origin policy，并禁止从非同一源的脚本中访问或修改数据。

简单地说，当Same-origin policy被强制执行时，浏览器只允许当前网页文档获得它自己的资源，而只能与同一个源相关的资源进行交互。这可以减少安全威胁，针对用户的浏览器提供更加安全的环境。   

## 汇总重点
{{< ask_chatgpt >}}
条列Same-origin policy的重点
{{< /ask_chatgpt >}}



1. Same-origin policy是一种网路安全机制，它限制了网页中的程式码只能访问它们自己的 origin（来源）。

2. origin是指协议（如http、https）、主机名和端口号的组合。如果两个URL的协议、主机名和端口号相同，那它们就是同源的。

3. Same-origin policy的目的是防止跨站脚本（Cross-site scripting，简称XSS）攻击和资料窃取等安全问题。

4. Same-origin policy的具体表现包括：禁止不同源的网页使用同一个浏览器存储（如cookie）；禁止不同源的网页使用同一个DOM（文档物件模型）。

5. Same-origin policy的限制可以通过CORS（跨来源资源共用）来打破，它允许不同源的网页之间进行资源共用。

6. Same-origin policy需要特别注意的是跨子域（Subdomain）需要特别处理，因为它们虽然不同源但是属于同一个顶级域名（例如a.example.com和b.example.com）。

7. Same-origin policy的实现是由浏览器负责的，为了保护用户的安全，浏览器不会允许跨域访问，开发者应该遵循Same-origin policy设计应用程式。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Same-origin policy的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. Same-origin policy 是什麼？它是如何保護使用者的安全性？

2. 當不同來源之間的資源需要進行交互時，Same-origin policy 會限制哪些行為？請舉例說明。

3. 請解釋 Same-origin policy 與 Cross-origin resource sharing (CORS) 的區別。

4. 如果一個網站想要跨域請求其他網站的資源，應該如何解決 Same-origin policy 的限制？

5. 除了在瀏覽器中執行時，Same-origin policy 也能在什麼情況下發揮作用？請舉例說明。

答案：

1. Same-origin policy 是瀏覽器安全機制之一，它強制限制網頁腳本等資源只能存取跟自身網頁來源相同的資源。這種限制減少了惡意網站惡意攻擊用戶的可能性。

2. Same-origin policy 通常限制了以下四種跨域操作：Cookie、LocalStorage 和 IndexedDB 存儲、讀取 DOM 元素內容、AJAX/Fetch 和 WebSocket 的發送和接收。例如，網站 A 的 JavaScript 不能使用 AJAX 向網站 B 發送請求，以防止惡意腳本盜取使用者的敏感信息。

3. Same-origin policy 是瀏覽器的內置安全特性，用於限制兩個不同源的網站之間的資源訪問；而 CORS 則是一種機制，允許網站解除跨域資源請求的限制。

4. 網站可以使用 CORS，以前端側的方式允許跨域請求。透過在請求標頭中添加特定設置，如 Access-Control-Allow-Origin，網站可以指示瀏覽器允許特定網站存取資源。

5. Same-origin policy 也可以在應用程式硬件層面發揮作用，例如瀏覽器插件可以通過使用 same-origin policy 防止第三方網站對插件的攻擊。另外， Content Security Policy (CSP) 也通過實施 same-origin policy 來限制網站的外部資源載入。   

