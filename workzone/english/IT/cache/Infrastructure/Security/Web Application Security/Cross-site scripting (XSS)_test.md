

1) What is Cross-site scripting (XSS)?
Answer: Cross-site scripting (XSS) is a type of security vulnerability that attackers can exploit to inject malicious code into a website, allowing them to execute unauthorized actions and steal sensitive information.

2) What are the common types of XSS attacks?
Answer: The common types of XSS attacks are reflected XSS, stored XSS, and DOM-based XSS. 

3) What is reflected XSS?
Answer: Reflected XSS is a type of XSS attack where a user is tricked into clicking on a link or visiting a web page that includes a malicious script, which then gets executed in the victim's browser.

4) How can you prevent XSS attacks?
Answer: To prevent XSS attacks, web developers can implement input validation, output encoding, and sanitization of user-generated content. Additionally, using HTTP-only cookies and implementing Content Security Policy (CSP) can also mitigate XSS vulnerabilities.

5) How can an attacker exploit a website using XSS?
Answer: An attacker can exploit a website using XSS by injecting malicious code into a vulnerable website's forms, comments, search boxes, or other areas where user-generated content can be submitted. Once the code is executed in a victim's browser, the attacker can steal sensitive information, hijack user sessions, or perform other unauthorized actions.