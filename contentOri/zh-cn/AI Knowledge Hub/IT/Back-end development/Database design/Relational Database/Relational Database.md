+++
title = "关联数据库"
weight = "1"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Relational Database的问题
{{< /ask_chatgpt >}}



1. 何謂關聯式資料庫？
2. 如何建立一個關聯式資料庫？ 
3. 什麼是資料表？何種資料可儲存於資料表中？ 
4. 如何透過SQL查詢來取得資料庫中特定的資料項目？ 
5. 如何在關聯式資料庫中建立關聯，以更有效率地儲存和查詢資料？   

## 说明知识
{{< ask_chatgpt >}}
说明Relational Database并举例
{{< /ask_chatgpt >}}



關聯式資料庫是由資料表組成的資料庫，每個表包含一個或多個欄位，每個欄位對應到一個特定的資料類型。表與表之間的關聯可以透過共同的欄位來建立，使得資料在多個表之間能夠交互使用。這種系統便於資料的查詢和管理，並且允許多個使用者同時使用。以下是一些可能需要使用關聯式資料庫的應用：

1. 網路商店：網路商店通常需要處理許多商品和客戶的資料。關聯式資料庫能夠輕易地儲存、更新、和查詢這些資料。

2. 學校管理系統：學校管理系統需要儲存學生和教師的資料。透過關聯式資料庫，可以輕鬆地跨不同的表來查詢和更新這些資料。

3. 銀行：銀行需要儲存客戶、帳戶和交易的資料。使用關聯式資料庫，銀行可以隨時追蹤個人和公司的資產狀況，以及處理各種交易、轉帳等事務。

4. 物流管理系統：物流公司需要儲存客戶、產品、訂單和出貨的資料。關聯式資料庫可以輕鬆處理這些資料，幫助物流公司追蹤貨物的運送進度，並確保產品能夠按時送達目的地。

總之，關聯式資料庫是一個有效的管理大量資料的途徑，它是商業和科技行業的核心。   

## 汇总重点
{{< ask_chatgpt >}}
条列Relational Database的重点
{{< /ask_chatgpt >}}



1. 资料以表格形式组织。
2. 主键（Primary Key）用于唯一识别每一笔记录。
3. 外键（Foreign Key）用于建立表格之间的关联。
4. 数据存储在记录中，每一列包含相同数量的字段。
5. 资料库的结构和关系可以通过ER图表达。
6. SQL是一种用于操作关系型数据库的语言。
7. ACID模型保证数据操作的原子性、一致性、隔离性和持久性。
8. 数据库能够提供高效的查询和快速的数据存取。
9. 数据库能够提供数据安全性和压缩储存等功能。
10. 数据库的设计应该考虑到性能、安全和可扩展性等因素。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Relational Database的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 在一个学校的学生资料库中，有两张表格：学生表(student)和成绩表(score)。请编写一个查询，显示出每个学生的姓名和他们的最高成绩。

答案：

```sql
SELECT student.name, MAX(score.score) 
FROM student 
JOIN score 
ON student.id = score.student_id 
GROUP BY student.id;
```

2. 在一个网上商店的资料库中，有两张表格：订单表(order)和产品表(product)。请编写一个查询，显示每个产品被订购的总数量。

答案：

```sql
SELECT product.name, SUM(order.quantity) 
FROM product 
JOIN order 
ON product.id = order.product_id 
GROUP BY product.id;
```

3. 在一个销售员和产品资料库中，有两张表格：销售员表(salesperson)和产品表(product)。请编写一个查询，显示销售员每个月的销售额。

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

4. 在一个学生资料库中，有一张表格：学生表(student)。每个学生都有一个班级(class)属性。请编写一个查询，显示每个班级有多少个学生以及平均年龄。

答案：

```sql
SELECT class, COUNT(*) AS num_students, AVG(age) AS avg_age 
FROM student 
GROUP BY class;
```

5. 在一个购物车资料库中，有三张表格：用户表(user)、产品表(product)和购物车表(cart)。用户可以将多个产品添加到购物车中。请编写一个查询，显示每个用户的购物车中有多少产品以及总价格。

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

