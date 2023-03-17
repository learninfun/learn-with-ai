

1. 求出一個資料表中，每個不同的類型(type)的平均價格(price)，欄位名稱為avg_price。

答案：
SELECT type, AVG(price) AS avg_price
FROM table
GROUP BY type;

2. 找出一個資料表中，最小日期(date)和最大日期(date)之間的資料，以date為排序依據。

答案：
SELECT *
FROM table
WHERE date BETWEEN MIN(date) AND MAX(date)
ORDER BY date;

3. 針對一個資料表，找出所有price小於等於1000且type等於'A'的資料，以price為排序依據。

答案：
SELECT *
FROM table
WHERE price <= 1000 AND type = 'A'
ORDER BY price;

4. 找出一個資料表中，出現最多次的type，欄位名稱為most_common_type。

答案：
SELECT type AS most_common_type
FROM table
GROUP BY type
ORDER BY COUNT(*) DESC
LIMIT 1;

5. 假設有兩個資料表(table_1和table_2)，各自有id和name欄位，找出table_1中與table_2中都有的name，並按照id排序。

答案：
SELECT table_1.id, table_1.name
FROM table_1
INNER JOIN table_2 ON table_1.name = table_2.name
ORDER BY table_1.id;