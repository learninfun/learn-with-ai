

Content Security Policy (CSP) is a security feature that helps websites to mitigate the risks of cross-site scripting (XSS) attacks, data injection, and other code injection attacks. CSP provides developers with a set of policies that restrict the resources a web page can load content from, and which types of code can be executed. Content Security Policy can be implemented using a HTTP response header, and it helps prevent attackers from injecting malicious code into a website.

An example of Content Security Policy would be to restrict a website to only load images from a specific domain, and to only run JavaScript code that is hosted locally. This can be achieved by setting the following CSP header in the server response:

Content-Security-Policy: default-src 'self'; img-src https://example.com

This CSP policy restricts the default source to the same origin (self), and allows images to be loaded only from the https://example.com domain. Thus, any attacker attempting to inject malicious code or images from a different domain will be blocked by the browser, reducing the risk of cross-site scripting and injection attacks.