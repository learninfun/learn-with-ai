

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