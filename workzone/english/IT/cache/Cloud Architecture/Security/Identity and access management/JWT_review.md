1. What is JWT?
Answer: JWT stands for JSON Web Token. It is a JSON object that is used to securely transmit information between two parties.

2. How does JWT work?
Answer: When a user logs in, a JWT containing a set of claims is generated and signed with a secret key. The JWT is then sent to the user and is included in subsequent requests. The server uses the secret key to verify the JWT and extract the claims.

3. What are the three parts of a JWT?
Answer: The three parts of a JWT are the header, payload, and signature. The header contains information about the algorithm used to sign the token. The payload contains the claims, or information about the user. The signature is used to verify the integrity of the token.

4. What is the difference between a JWT and a session cookie?
Answer: A JWT is a self-contained token that does not require server-side storage, whereas a session cookie relies on server-side storage. JWTs are also stateless, meaning that each request includes all the necessary information. Session cookies require the server to look up the session data on each request.

5. What are some common vulnerabilities with JWT?
Answer: Some common vulnerabilities with JWT include insufficient signing algorithms, insecure storage of secrets, and insufficient validation of token signatures. It is important to securely store JWT secrets and use strong signing algorithms to avoid these vulnerabilities.