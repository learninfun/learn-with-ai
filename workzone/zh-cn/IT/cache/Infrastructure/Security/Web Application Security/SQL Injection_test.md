

1. 請問下列SQL語法是否具備SQL Injection漏洞？

SELECT * FROM users WHERE username = 'John' AND password = '1234' OR '1'='1';

答案：是。以OR '1'='1'這種方式是否為真的條件，可以將整個WHERE子句視為真實，因此可以查詢到所有的使用者帳戶。

2. 請問下列SQL語法是否具備SQL Injection漏洞？

SELECT * FROM products WHERE id = 1 AND category = 'shoes';

答案：否。沒有針對使用者輸入內容的WHERE子句。

3. 請問下列SQL語法是否具備SQL Injection漏洞？

SELECT * FROM users WHERE username = '$username' AND password = '$password';

答案：是。因為使用者可以輸入任何想要的帳號和密碼，如果攻擊者能夠輸入SQL語句，就有可能進行SQL Injection攻擊。

4. 請問要如何在下列SQL語法中進行SQL Injection攻擊？

SELECT * FROM users WHERE username = 'John' AND password = '1234';

答案：攻擊者可以在帳戶名或密碼的輸入欄位中輸入特殊字符，如單引號、反斜杠等等，來觸發SQL Injection衝擊。

5. 請問下列SQL語法是否具備SQL Injection漏洞？

SELECT * FROM users WHERE username = 'John' OR 1=1;

答案：是。OR 1=1意味著任何帳戶都會被傳回，因此這種語句很容易受到SQL Injection攻擊。