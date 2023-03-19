

1. What is Same-origin policy?
Answer: Same-origin policy is a security feature implemented by web browsers that restricts web pages from accessing resources from a different origin.

2. What is the purpose of Same-origin policy?
Answer: The main purpose of Same-origin policy is to prevent malicious websites from stealing sensitive data from other websites that a user may be logged into.

3. How does Same-origin policy work?
Answer: Same-origin policy works by comparing the origin (protocol, domain name, and port number) of the requesting web page with that of the requested resource. If the two origins are the same, the request is allowed, otherwise it is blocked.

4. What are some exceptions to Same-origin policy?
Answer: Some exceptions to Same-origin policy include cross-origin resource sharing (CORS), which allows a server to specify which origins are permitted to access its resources, and JSONP (JSON with padding), which allows data to be requested from a different origin using a script tag.

5. What are the potential risks of bypassing Same-origin policy?
Answer: By bypassing Same-origin policy, attackers may be able to steal sensitive data, perform actions on behalf of a user (such as sending emails or making purchases), or inject malicious code into a web page.