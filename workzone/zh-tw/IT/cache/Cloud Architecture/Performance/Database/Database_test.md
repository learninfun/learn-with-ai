1. 建立一個顧客的資料表，其中包含顧客姓名、電子郵件、國家和訂單數量。請寫一個SQL查詢，列出訂單數量最多的前5位顧客。
答案：SELECT name, email, country, order_count FROM customers ORDER BY order_count DESC LIMIT 5

2. 建立一個商品的資料表，其中包含產品名稱、類型、單位和庫存量。請寫一個SQL查詢，列出每種類型的商品庫存總量。
答案：SELECT type, SUM(stock) AS total_stock FROM products GROUP BY type

3. 建立一個訂單的資料表，其中包含訂單編號、顧客ID、訂單日期和總價格。請寫一個SQL查詢，列出每位顧客的平均訂單價格，並按照價格降序排列。
答案：SELECT c.name, AVG(o.total_price) AS avg_price FROM customers c JOIN orders o ON c.id = o.customer_id GROUP BY c.name ORDER BY avg_price DESC

4. 建立一個文章的資料表，其中包含文章標題、創建日期和作者ID。請寫一個SQL查詢，列出每位作者的文章數量，並按照文章數量降序排列。
答案：SELECT a.name, COUNT(p.id) AS article_count FROM authors a JOIN posts p ON a.id = p.author_id GROUP BY a.name ORDER BY article_count DESC

5. 建立一個學生的資料表，其中包含學生姓名、學號、email和出生日期。請寫一個SQL查詢，列出出生日期在1995年之前的學生的姓名和學號。
答案：SELECT name, student_id FROM students WHERE birthdate < '1995-01-01'