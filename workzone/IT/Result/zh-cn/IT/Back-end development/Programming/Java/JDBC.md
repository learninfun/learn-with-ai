+++
title = "JDBC"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题JDBC的问题
{{< /ask_chatgpt >}}



1.  JDBC是什么？
2.  JDBC的API包括哪些？
3.  哪些不同的设置需求来连接资料库？
4.  怎样处理JDBC的交易？
5.  JDBC里的PreparedStatement和Statement之间有什么区别？   

## 说明知识
{{< ask_chatgpt >}}
说明JDBC并举例
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
- 使用 `ResultSet` 对象来处理查询结果
- 关闭所有资源，包括连接、语句和结果集，在 finally 区块中进行这些操作。   

## 汇总重点
{{< ask_chatgpt >}}
条列JDBC的重点
{{< /ask_chatgpt >}}



1. JDBC是Java Database Connectivity的縮寫，它提供了與關係型數據庫（如MySQL，Oracle等）進行交互的標準接口。
2. JDBC使用一組API來訪問數據庫的操作，這些API包括Connection，Statement或PreparedStatement，和ResultSet等。
3. Connection是 JDBC 中最基本的要素之一，它表示與數據庫之間的連接。該接口還包括其他許多重要功能，例如事務管理和自動提交。
4. Statement或PreparedStatement用於執行SQL查詢或更新數據庫記錄。PreparedStatement是一種比Statement更進一步的對象，它允許開發人員編寫可變参数的SQL語句。
5. ResultSet則用於處理查詢結果集，從中檢索數據以進行顯示或計算等操作。
6. JDBC程序實際上是通過建立Connection對象，創建Statement對象並執行SQL查詢來與數據庫進行交互。並且一定要在用完後關閉所有的資料庫資源，包括Connection，Statement和ResultSet。
7. 除了純JDBC之外，也可以使用Spring JDBC等框架來更簡化和加快JDBC的開發。   

## 知识测验
{{< ask_chatgpt >}}
给我5题JDBC的中等难度问题，并在后面列出答案
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

## 网络数据
{{< ask_chatgpt >}}
给我5篇JDBC的网络数据
{{< /ask_chatgpt >}}



1. "JDBC Tutorial"，Oracle官方网站，https://docs.oracle.com/en/database/oracle/oracle-database/21/jjdbc/toc.htm

这篇教学文章介绍了JDBC API的基本概念、连线与查询资料、使用PreparedStatement、处理结果集等等。

2. "JDBC Tutorial for Beginners with Examples"，JournalDev网站，https://www.journaldev.com/2471/jdbc-tutorial

这篇教学文章介绍从设定JDBC驱动程式到连线资料库、执行查询语句、以及使用PreparedStatement、CallableStatement等等的相关操作。

3. "JDBC Connection Pooling with HikariCP"，Baeldung网站，https://www.baeldung.com/hikaricp

这篇文章介绍了如何透过HikariCP进行JDBC连线池化以提高应用程式的效能和稳定性。

4. "JDBC Batch Processing using Java"，GeeksforGeeks网站，https://www.geeksforgeeks.org/jdbc-batch-processing-in-java/

这篇教学文章介绍了如何使用JDBC进行批次处理，在降低资料库负载的前提下提高效能。

5. "JDBC vs Hibernate: What’s the difference?"，Medium网站，https://medium.com/@natd/error-snapshot-fix-8ffdf2c69cbf

这篇文章比较了JDBC和Hibernate这两种不同的资料库存取方式，介绍了它们的优缺点和适用场景。   

