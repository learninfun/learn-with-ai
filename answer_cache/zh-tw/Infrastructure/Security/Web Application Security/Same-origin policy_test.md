

1. What is the purpose of Same-origin policy?
A: Same-origin policy is a security measure, which restricts the interaction between web pages from different origins. Its main purpose is to protect users from malicious scripts and potential data breaches.

2. How does Same-origin policy work?
A: Same-origin policy works by comparing the domain names, port numbers, and protocols of the web pages to determine if they are from the same origin. If they are not, the policy restricts the access to the resources of the other web page.

3. What are the exceptions to Same-origin policy?
A: Same-origin policy can be bypassed through several methods like cross-origin resource sharing (CORS), JSONP, and postMessage API. These methods allow the sharing of resources between web pages from different origins.

4. How can Same-origin policy impact web development?
A: Same-origin policy can impact web development by limiting the access to resources from different origins. This can cause issues with the integration of third-party APIs and libraries, requiring developers to use workarounds like CORS and JSONP.

5. How can Same-origin policy be enforced on a website?
A: Same-origin policy can be enforced on a website by setting the appropriate HTTP headers or using server-side languages like PHP and Node.js to restrict access to resources from different origins. Additionally, web developers can use content security policy (CSP) to reduce the risks of XSS attacks.