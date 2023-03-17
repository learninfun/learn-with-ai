

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