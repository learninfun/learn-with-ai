

Same-origin policy 是一個瀏覽器安全性機制，限制了一個網頁文檔（document）或腳本（script）在瀏覽器和同一個源（origin）下的資源交互，從而防止潛在的跨站腳本攻擊（XSS）或跨站請求偽造（CSRF）等安全問題。

一個源通常由三部分組成：協議（Protocol）、域名（Domain）和端口（Port）。只有三個部分都相同的資源才被認為是同一個源，否則就被視為不同的源。 譬如：

- http://www.example.com 和 https://www.example.com 是不同的源。
- http://www.example.com 和 http://example.com 是不同的源。
- http://www.example.com:80 和 http://www.example.com:8080 是不同的源。

一些常見的Same-origin policy限制包括：

- JavaScript在同一源下的資源上想像可以自由使用，但他們無法訪問其他源的資源。 假設有一個JavaScript應用程序在example.com上運行，那麼就可以訪問它同一源下的任何資源，例如example.com/about.html；但是，如果應用程式嘗試訪問example.net上的資源，則將會被同源政策阻止。

- 瀏覽器在送出不同源的 AJAX 請求或向 iframe 內載入內容時遵守同源政策。

- 不同域名下設置的 Cookie 不會被 JavaScript 訪問，也不會被瀏覽器發送給不同域名。

- 一些HTML5 API （例如 Geolocation API 或者 Web Storage API）將遵循Same-origin policy，並禁止從非同一源的腳本中訪問或修改數據。

簡單地說，當Same-origin policy被強制執行時，瀏覽器只允許當前網頁文檔獲得它自己的資源，而只能與同一個源相關的資源進行交互。這可以減少安全威脅，針對用戶的瀏覽器提供更加安全的環境。