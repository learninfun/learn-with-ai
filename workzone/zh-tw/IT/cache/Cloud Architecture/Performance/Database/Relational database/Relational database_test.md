1. 在以下的資料表中，找出姓氏為"王"的員工的ID和姓名。
表格：員工表(Employee)
欄位：ID, 姓名, 姓氏, 職稱
答案：SELECT ID, 姓名 FROM Employee WHERE 姓氏='王'

2. 在以下的資料表中，計算每位學生的平均分數。
表格：成績表(Score)
欄位：學生ID, 科目, 分數
答案：SELECT 學生ID, AVG(分數) AS 平均分數 FROM Score GROUP BY 學生ID

3. 在以下的資料表中，找出最常出現的商品類別及其出現次數。
表格：商品表(Product)
欄位：商品ID, 商品名稱, 商品類別
答案：SELECT 商品類別, COUNT(*) AS 出現次數 FROM Product GROUP BY 商品類別 ORDER BY 出現次數 DESC LIMIT 1

4. 在以下的資料表中，找出有加入會員卻沒有下訂單的所有人的名字和電子郵件。
表格：會員表(Member), 訂單表(Order)
欄位（Member）: ID, 名字, 電子郵件
欄位（Order）: 會員ID, 訂單編號, 訂單日期
答案：SELECT 名字, 電子郵件 FROM Member WHERE ID NOT IN (SELECT 會員ID FROM Order)

5. 在以下的資料表中，找出訂單總金額最高的客戶姓名。
表格：客戶表(Customer), 訂單表(Order), 訂單明細表(OrderDetail)
欄位（Customer）: ID, 姓名
欄位（Order）: 訂單編號, 客戶ID, 訂單日期
欄位（OrderDetail）: 訂單編號, 商品ID, 商品數量, 商品單價
答案：SELECT 姓名 FROM Customer WHERE ID=(SELECT 客戶ID FROM Order WHERE 訂單編號=(SELECT 訂單編號 FROM OrderDetail GROUP BY 訂單編號 ORDER BY SUM(商品數量*商品單價) DESC LIMIT 1))