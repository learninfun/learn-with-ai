

1. Same-origin policy是一種網路安全機制，它限制了網頁中的程式碼只能訪問它們自己的 origin（來源）。

2. origin是指協議（如http、https）、主機名和端口號的組合。如果兩個URL的協議、主機名和端口號相同，那它們就是同源的。

3. Same-origin policy的目的是防止跨站腳本（Cross-site scripting，簡稱XSS）攻擊和資料竊取等安全問題。

4. Same-origin policy的具體表現包括：禁止不同源的網頁使用同一個瀏覽器存儲（如cookie）；禁止不同源的網頁使用同一個DOM（文檔物件模型）。

5. Same-origin policy的限制可以通過CORS（跨來源資源共用）來打破，它允許不同源的網頁之間進行資源共用。

6. Same-origin policy需要特別注意的是跨子域（Subdomain）需要特別處理，因為它們雖然不同源但是屬於同一個頂級域名（例如a.example.com和b.example.com）。

7. Same-origin policy的實現是由瀏覽器負責的，為了保護用戶的安全，瀏覽器不會允許跨域訪問，開發者應該遵循Same-origin policy設計應用程式。