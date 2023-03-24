

1. CSP如何避免Cross-site scripting (XSS)攻击？
答案：CSP可以限制执行JavaScript的来源，进而避免XSS攻击。使用CSP时，可以设置白名单，只允许特定的域名下的JavaScript脚本执行。如果有其他来源的脚本试图执行，就会被挡下来。例如，可以在Content-Security-Policy头信息中设置“script-src https://example.com”，这样只有来自example.com的JavaScript脚本才能执行。

2. 如何在CSP中设置严格的来源限制？
答案：可以使用“default-src”属性设置CSP的严格来源限制。例如，可以使用“Content-Security-Policy: default-src 'none';”禁止任何外部资源的请求。

3. 如何在CSP中允许特定的iframe？
答案：可以使用"frame-src"属性设置特定iframe的来源限制。例如，可以在Content-Security-Policy头信息中设置“frame-src https://example.com”，这样只有来自example.com的iframe才能载入。

4. 如何在CSP中设置允许inline样式？
答案：可以使用“style-src”属性设置允许inline样式。例如，可以在Content-Security-Policy头信息中设置“style-src 'self' 'unsafe-inline'”，这样允许网页中使用的inline样式。

5. 如何在CSP中设置不允许外部图片载入？
答案：可以使用“img-src”属性设置是否允许外部图片载入。例如，可以在Content-Security-Policy头信息中设置“img-src 'self'”，这样就只允许从同一域名下载入图片。