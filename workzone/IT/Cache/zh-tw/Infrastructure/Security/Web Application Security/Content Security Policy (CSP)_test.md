

1. CSP如何避免Cross-site scripting (XSS)攻擊？
答案：CSP可以限制執行JavaScript的來源，進而避免XSS攻擊。使用CSP時，可以設置白名單，只允許特定的域名下的JavaScript腳本執行。如果有其他來源的腳本試圖執行，就會被擋下來。例如，可以在Content-Security-Policy頭信息中設置「script-src https://example.com」，這樣只有來自example.com的JavaScript腳本才能執行。

2. 如何在CSP中設置嚴格的來源限制？
答案：可以使用「default-src」屬性設置CSP的嚴格來源限制。例如，可以使用「Content-Security-Policy: default-src 'none';」禁止任何外部資源的請求。

3. 如何在CSP中允許特定的iframe？
答案：可以使用"frame-src"屬性設置特定iframe的來源限制。例如，可以在Content-Security-Policy頭信息中設置「frame-src https://example.com」，這樣只有來自example.com的iframe才能載入。

4. 如何在CSP中設置允許inline樣式？
答案：可以使用「style-src」屬性設置允許inline樣式。例如，可以在Content-Security-Policy頭信息中設置「style-src 'self' 'unsafe-inline'」，這樣允許網頁中使用的inline樣式。

5. 如何在CSP中設置不允許外部圖片載入？
答案：可以使用「img-src」屬性設置是否允許外部圖片載入。例如，可以在Content-Security-Policy頭信息中設置「img-src 'self'」，這樣就只允許從同一域名下載入圖片。