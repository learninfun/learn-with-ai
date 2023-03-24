

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