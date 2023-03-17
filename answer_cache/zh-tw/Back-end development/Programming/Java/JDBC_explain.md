

JDBC（Java Database Connectivity）是一種Java API（應用程式介面），用於與關聯式資料庫系統進行交互。 它定義了一組Java 類和介面，使開發人員可以編寫Java應用程式，以便通過標準SQL語言與各種RDBMS進行通信。

下面是一個簡單的JDBC例子，使用JDBC API與MySQL數據庫進行交互：

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

簡要解釋：

- 我們首先要註冊JDBC驅動程序，通過使用 `Class.forName()`
- 創建 JDBC連接並使用 `DriverManager.getConnection()` 連接到數據庫
- 使用 `Connection.createStatement()` 創建一個 `Statement` 對象，並使用它來執行SQL查詢
- 使用 `ResultSet` 對像來處理查詢結果
- 關閉所有資源，包括連接、語句和結果集，在 finally 區塊中進行這些操作。