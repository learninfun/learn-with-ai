+++
title = "關聯數據庫"
weight = "1"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Relational Database的中文問題
{{< /ask_chatgpt >}}



1. 何謂關聯式資料庫？
2. 如何建立一個關聯式資料庫？ 
3. 什麼是資料表？何種資料可儲存於資料表中？ 
4. 如何透過SQL查詢來取得資料庫中特定的資料項目？ 
5. 如何在關聯式資料庫中建立關聯，以更有效率地儲存和查詢資料？   

## 說明知識
{{< ask_chatgpt >}}
以中文說明Relational Database並舉例
{{< /ask_chatgpt >}}



關聯式資料庫是由資料表組成的資料庫，每個表包含一個或多個欄位，每個欄位對應到一個特定的資料類型。表與表之間的關聯可以透過共同的欄位來建立，使得資料在多個表之間能夠交互使用。這種系統便於資料的查詢和管理，並且允許多個使用者同時使用。以下是一些可能需要使用關聯式資料庫的應用：

1. 網路商店：網路商店通常需要處理許多商品和客戶的資料。關聯式資料庫能夠輕易地儲存、更新、和查詢這些資料。

2. 學校管理系統：學校管理系統需要儲存學生和教師的資料。透過關聯式資料庫，可以輕鬆地跨不同的表來查詢和更新這些資料。

3. 銀行：銀行需要儲存客戶、帳戶和交易的資料。使用關聯式資料庫，銀行可以隨時追蹤個人和公司的資產狀況，以及處理各種交易、轉帳等事務。

4. 物流管理系統：物流公司需要儲存客戶、產品、訂單和出貨的資料。關聯式資料庫可以輕鬆處理這些資料，幫助物流公司追蹤貨物的運送進度，並確保產品能夠按時送達目的地。

總之，關聯式資料庫是一個有效的管理大量資料的途徑，它是商業和科技行業的核心。   

## 條列重點
{{< ask_chatgpt >}}
以中文條列Relational Database的重點
{{< /ask_chatgpt >}}



1. 資料以表格形式組織。
2. 主鍵（Primary Key）用於唯一識別每一筆記錄。
3. 外鍵（Foreign Key）用於建立表格之間的關聯。
4. 數據存儲在記錄中，每一列包含相同數量的字段。
5. 資料庫的結構和關係可以通過ER圖表達。
6. SQL是一種用於操作關係型數據庫的語言。
7. ACID模型保證數據操作的原子性、一致性、隔離性和持久性。
8. 數據庫能夠提供高效的查詢和快速的數據存取。
9. 數據庫能夠提供數據安全性和壓縮儲存等功能。
10. 數據庫的設計應該考慮到性能、安全和可擴展性等因素。   

## 知識測驗
{{< ask_chatgpt >}}
以中文給我5題Relational Database的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 在一個學校的學生資料庫中，有兩張表格：學生表(student)和成績表(score)。請編寫一個查詢，顯示出每個學生的姓名和他們的最高成績。

答案：

```sql
SELECT student.name, MAX(score.score) 
FROM student 
JOIN score 
ON student.id = score.student_id 
GROUP BY student.id;
```

2. 在一個網上商店的資料庫中，有兩張表格：訂單表(order)和產品表(product)。請編寫一個查詢，顯示每個產品被訂購的總數量。

答案：

```sql
SELECT product.name, SUM(order.quantity) 
FROM product 
JOIN order 
ON product.id = order.product_id 
GROUP BY product.id;
```

3. 在一個銷售員和產品資料庫中，有兩張表格：銷售員表(salesperson)和產品表(product)。請編寫一個查詢，顯示銷售員每個月的銷售額。

答案：

```sql
SELECT salesperson.name, DATE_FORMAT(sales.date_sold, '%m-%Y') AS month, SUM(product.price * sales.quantity_sold) AS sales_total 
FROM sales 
JOIN salesperson 
ON sales.salesperson_id = salesperson.id 
JOIN product 
ON sales.product_id = product.id 
GROUP BY salesperson.id, DATE_FORMAT(sales.date_sold, '%m-%Y');
```

4. 在一個學生資料庫中，有一張表格：學生表(student)。每個學生都有一個班級(class)屬性。請編寫一個查詢，顯示每個班級有多少個學生以及平均年齡。

答案：

```sql
SELECT class, COUNT(*) AS num_students, AVG(age) AS avg_age 
FROM student 
GROUP BY class;
```

5. 在一個購物車資料庫中，有三張表格：用戶表(user)、產品表(product)和購物車表(cart)。用戶可以將多個產品添加到購物車中。請編寫一個查詢，顯示每個用戶的購物車中有多少產品以及總價格。

答案：

```sql
SELECT user.username, COUNT(*) AS num_products_in_cart, SUM(product.price) AS total_price 
FROM user 
JOIN cart 
ON user.id = cart.user_id 
JOIN product 
ON cart.product_id = product.id 
GROUP BY user.id;
```   

