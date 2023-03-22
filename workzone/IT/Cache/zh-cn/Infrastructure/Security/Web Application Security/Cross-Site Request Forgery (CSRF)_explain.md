

Cross-Site Request Forgery (CSRF) is a type of attack that allows an attacker to impersonate a legitimate user and obtain sensitive information or perform actions on behalf of the user without their consent. In this type of attack, the attacker is able to trick a victim's web browser into sending unwanted requests to a targeted website, using the user's session information that is stored in cookies.

For example, let's say a user is logged into their online banking account and has recently opened another tab to visit a different website. Meanwhile, the attacker sends an email to the user containing a link to a malicious website that contains a hidden form that performs a CSRF attack. When the user clicks on that link, the malicious code embedded in the form sends a request to the banking website, which is authorized by the user's session cookie. The request could, for example, transfer funds from the user's account to the attacker's account, as the user's session allows the action.

In this example, the victim's web browser was unaware of the malicious request being sent, and since the request was performed with the user's session cookie, it was authorized by the website. This type of attack can be prevented by using anti-forgery tokens, session cookies with HttpOnly and Secure flags, and verifying user actions before allowing them to perform actions that require authorization.