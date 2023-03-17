

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