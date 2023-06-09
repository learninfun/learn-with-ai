+++
title = "JDBC"
weight = "5"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题JDBC的中文问题
{{< /ask_chatgpt >}}



1.  JDBC是什么？
2.  JDBC的API包括哪些？
3.  哪些不同的设置需求来连接资料库？
4.  怎样处理JDBC的交易？
5.  JDBC里的PreparedStatement和Statement之间有什么区别？   

## 说明知识
{{< ask_chatgpt >}}
以中文说明JDBC并举例
{{< /ask_chatgpt >}}



JDBC（Java Database Connectivity）是一种Java API（应用程式介面），用于与关联式资料库系统进行交互。 它定义了一组Java 类和介面，使开发人员可以编写Java应用程式，以便通过标准SQL语言与各种RDBMS进行通信。

下面是一个简单的JDBC例子，使用JDBC API与MySQL数据库进行交互：

```java
import java.sql.*;

public class JDBCTest {
   static final String JDBC_DRIVER = "com.mysql.jdbc.Driver";
   static final String DB_URL = "jdbc:mysql://localhost/sample_db";

   static final String USER = "username";
   static final String PASS = "password";

   public static void main(String[] args) {
      Connection conn = null;
      Statement stmt = null;
      try{
         Class.forName(JDBC_DRIVER);

         System.out.println("Connecting to database...");
         conn = DriverManager.getConnection(DB_URL,USER,PASS);

         System.out.println("Creating statement...");
         stmt = conn.createStatement();
         String sql;
         sql = "SELECT id, name, age FROM employees";
         ResultSet rs = stmt.executeQuery(sql);

         while(rs.next()){
            int id = rs.getInt("id");
            String name = rs.getString("name");
            int age = rs.getInt("age");

            System.out.print("ID: " + id);
            System.out.print(", Name: " + name);
            System.out.println(", Age: " + age);
         }

         rs.close();
         stmt.close();
         conn.close();
      } catch(SQLException se) {
         se.printStackTrace();
      } catch(Exception e) {
         e.printStackTrace();
      } finally {
         try {
            if(stmt!=null) stmt.close();
         } catch(SQLException se2) {
         }
         try {
            if(conn!=null) conn.close();
         } catch(SQLException se) {
            se.printStackTrace();
         }
      }
   }
}
```

简要解释：

- 我们首先要注册JDBC驱动程序，通过使用 `Class.forName()`
- 创建 JDBC连接并使用 `DriverManager.getConnection()` 连接到数据库
- 使用 `Connection.createStatement()` 创建一个 `Statement` 对象，并使用它来执行SQL查询
- 使用 `ResultSet` 对像来处理查询结果
- 关闭所有资源，包括连接、语句和结果集，在 finally 区块中进行这些操作。   

## 条列重点
{{< ask_chatgpt >}}
以中文条列JDBC的重点
{{< /ask_chatgpt >}}



1. JDBC是Java Database Connectivity的缩写，它提供了与关系型数据库（如MySQL，Oracle等）进行交互的标准接口。
2. JDBC使用一组API来访问数据库的操作，这些API包括Connection，Statement或PreparedStatement，和ResultSet等。
3. Connection是 JDBC 中最基本的要素之一，它表示与数据库之间的连接。该接口还包括其他许多重要功能，例如事务管理和自动提交。
4. Statement或PreparedStatement用于执行SQL查询或更新数据库记录。PreparedStatement是一种比Statement更进一步的对象，它允许开发人员编写可变参数的SQL语句。
5. ResultSet则用于处理查询结果集，从中检索数据以进行显示或计算等操作。
6. JDBC程序实际上是通过建立Connection对象，创建Statement对象并执行SQL查询来与数据库进行交互。并且一定要在用完后关闭所有的资料库资源，包括Connection，Statement和ResultSet。
7. 除了纯JDBC之外，也可以使用Spring JDBC等框架来更简化和加快JDBC的开发。   

## 知识测验
{{< ask_chatgpt >}}
以中文给我5题JDBC的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 如何使用JDBC从数据库中查询特定日期之后的所有记录?
答案：使用PreparedStatement对时间戳进行查询，例如：
```
String sql = "SELECT * FROM myTable WHERE datetime > ?";
PreparedStatement pstmt = conn.prepareStatement(sql); 
pstmt.setTimestamp(1, new Timestamp(date.getTime())); 
```
其中，conn是已经建立好的数据库连接对象，date是特定的日期对象。

2. 如何使用JDBC在数据库中执行批量更新操作?
答案：使用PreparedStatement对象的addBatch()和executeBatch()方法，例如：
```
String sql = "UPDATE myTable SET name = ?, age = ? WHERE id = ?";
PreparedStatement pstmt = conn.prepareStatement(sql); 
for (int i = 0; i < data.size(); i++) { //data为需要更新的数据
    pstmt.setString(1, data.getName(i)); 
    pstmt.setInt(2, data.getAge(i)); 
    pstmt.setInt(3, data.getId(i)); 
    pstmt.addBatch(); 
}
int[] result = pstmt.executeBatch();
```
其中，conn是已经建立好的数据库连接对象。

3. 如何使用JDBC实现分页查询?
答案：在SQL语句中使用LIMIT和OFFSET关键字实现分页，例如：
```
String sql = "SELECT * FROM myTable LIMIT ? OFFSET ?";
PreparedStatement pstmt = conn.prepareStatement(sql);
pstmt.setInt(1, pageSize);
pstmt.setInt(2, (pageNum - 1) * pageSize);
```
其中，pageSize表示每页显示的记录数量，pageNum表示要查询的页码。

4. 如何使用JDBC在数据库中执行事务?
答案：使用Connection对象的setAutoCommit(false)方法关闭自动提交，然后通过commit()和rollback()方法分别提交和回滚事务，例如：
```
Connection conn = DriverManager.getConnection(url, user, password);
conn.setAutoCommit(false);
try {
    Statement stmt = conn.createStatement();
    // some sql operations here
    stmt.execute("INSERT INTO myTable(name, age) VALUES('Tom', 20)");
    stmt.execute("UPDATE myTable SET age = 22 WHERE name = 'Tom'");
    conn.commit();
} catch (SQLException e) {
    conn.rollback();
} finally {
    conn.setAutoCommit(true);
    conn.close();
}
```

5. 如何使用JDBC从数据库中查询有关联的记录?
答案：使用JOIN操作实现，例如：
```
String sql = "SELECT t1.*, t2.age FROM table1 t1 JOIN table2 t2 ON t1.id = t2.table1Id";
PreparedStatement pstmt = conn.prepareStatement(sql);
ResultSet rs = pstmt.executeQuery();
while (rs.next()) {
    String name = rs.getString("name");
    int age1 = rs.getInt("age1");
    int age2 = rs.getInt("age2");
    System.out.println(name + ", " + age1 + ", " + age2);
}
```
其中，t1和t2分别代表要关联的两张表，table1Id是t1这张表中用来关联的id栏位名称。   

