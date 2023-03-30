1. What is the difference between authentication and authorization in web application security?
Answer: Authentication is the process of verifying the user's identity, while authorization is the process of granting or denying a user access to specific resources or functionalities based on their authenticated identity.

2. What is SQL injection and how can it be prevented?
Answer: SQL injection is a type of web application attack where malicious actors inject SQL code into an application's input fields to gain unauthorized access to databases. It can be prevented by applying input validation and parameterized queries, which filter and sanitize all user input before it is sent to the database.

3. What is cross-site scripting (XSS) and how can it be prevented?
Answer: XSS is a type of attack where malicious actors inject client-side scripts into web pages viewed by other users, in order to steal sensitive information or perform actions on their behalf. It can be prevented by applying input validation, sanitizing user input, and implementing secure coding practices, such as using Content Security Policy (CSP) directives.

4. What is clickjacking and how can it be prevented?
Answer: Clickjacking is a type of attack where malicious actors overlay a website or application with an invisible frame or link that tricks users into clicking on a button or link they didn't intend to. It can be prevented by using X-FRAME-OPTIONS HTTP response headers or implementing JavaScript frame-busting techniques that prevent framing of the application.

5. What is session fixation and how can it be prevented?
Answer: Session fixation is a type of attack where malicious actors exploit vulnerabilities in an application's session management process to hijack a user's session ID and access their account without authorization. It can be prevented by using secure session management techniques, such as session ID regeneration, storing session data on the server instead of the client, and implementing secure cookie settings.