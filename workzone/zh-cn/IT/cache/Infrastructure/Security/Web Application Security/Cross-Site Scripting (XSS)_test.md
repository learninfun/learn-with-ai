

1. 假设有一个网站使用非安全的Cookie储存用户讯息，请问攻击者可否透过XSS攻击偷取此Cookie？

答案：是。攻击者可以透过注入一段含有恶意的JavaScript程式码，使用户浏览器执行此程式码，从而窃取Cookie储存的用户讯息。

2. 请问DOM-Based XSS攻击与传统的反射型XSS攻击有何区别？

答案：DOM-Based XSS攻击是针对使用JavaScript动态修改HTML DOM的网页，而非直接注入HTML标签内容的XSS攻击。攻击方式类似反射型XSS攻击，但注入点在JavaScript，而非在HTTP回应中。

3. 如果一个网站使用HTTP而不是HTTPS连线，请问攻击者可否透过中间人攻击进行XSS注入？

答案：是。由于HTTP传输的讯息是明文的，攻击者可在传输过程中修改HTTP回应内容，进行XSS注入攻击。

4. 请问在提交表单时，为何需要使用CSRF Token保护？如何预防XSS攻击？

答案：提交表单时需要使用CSRF Token保护，因为攻击者可透过XSS攻击获取用户的身份验证Cookie，进而伪造请求，对该用户的资料进行恶意操作。在预防XSS攻击方面，可使用输入验证、过滤输入内容、安全储存敏感资料等方式。

5. 请问如何使用Content Security Policy (CSP)保护网站免受XSS攻击？

答案：可透过设定Content Security Policy (CSP)限制网站可以载入的资源，防止攻击者的脚本注入到网页中。具体的措施可以包括限制脚本的来源、禁用内联脚本等。另外，也可使用跨站点脚本攻击防御框架，如X-XSS-Protection等。