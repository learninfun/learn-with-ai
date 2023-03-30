1. What is Cross-Site Request Forgery (CSRF)? 

Answer: Cross-Site Request Forgery (CSRF) is a type of security vulnerability that allows an attacker to perform unwanted actions on behalf of the authenticated user, without their knowledge or consent.

2. How does CSRF attack work? 

Answer: CSRF attacks exploit the trust relationship between a user and a web application. The attacker can create a malicious page or exploit a vulnerability on a legitimate site to trick a user into performing an unintended action on a different site, where they have previously authenticated.

3. What are the consequences of a successful CSRF attack? 

Answer: The consequences of a successful CSRF attack can be severe. The attacker can perform actions such as changing a user's password, transferring funds, or making unauthorized purchases, depending on the functionality available to the user on the targeted site.

4. What are some techniques to prevent CSRF attacks? 

Answer: The most common technique to prevent CSRF attacks is to use a unique token for each user session, known as a CSRF token. Other techniques include using HTTP-only cookies, implementing strict SameSite cookie policies, and restricting CORS policies.

5. Can CSP (Content Security Policy) prevent CSRF attacks? 

Answer: CSP can help mitigate the impact of a successful CSRF attack by limiting the attacker's ability to execute arbitrary code. However, CSP alone cannot prevent CSRF attacks entirely, and other techniques such as CSRF tokens are still required.