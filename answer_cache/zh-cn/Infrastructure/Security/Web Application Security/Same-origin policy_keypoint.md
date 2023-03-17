

1. Same-origin policy是一种网路安全机制，它限制了网页中的程式码只能访问它们自己的 origin（来源）。

2. origin是指协议（如http、https）、主机名和端口号的组合。如果两个URL的协议、主机名和端口号相同，那它们就是同源的。

3. Same-origin policy的目的是防止跨站脚本（Cross-site scripting，简称XSS）攻击和资料窃取等安全问题。

4. Same-origin policy的具体表现包括：禁止不同源的网页使用同一个浏览器存储（如cookie）；禁止不同源的网页使用同一个DOM（文档物件模型）。

5. Same-origin policy的限制可以通过CORS（跨来源资源共用）来打破，它允许不同源的网页之间进行资源共用。

6. Same-origin policy需要特别注意的是跨子域（Subdomain）需要特别处理，因为它们虽然不同源但是属于同一个顶级域名（例如a.example.com和b.example.com）。

7. Same-origin policy的实现是由浏览器负责的，为了保护用户的安全，浏览器不会允许跨域访问，开发者应该遵循Same-origin policy设计应用程式。