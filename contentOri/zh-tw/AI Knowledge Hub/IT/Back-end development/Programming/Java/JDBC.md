## 習題預習
{{< ask_chatgpt >}}
給我5題JDBC的問題
{{< /ask_chatgpt >}}



1.  JDBC是什麼？
2.  JDBC的API包括哪些？
3.  哪些不同的設置需求來連接資料庫？
4.  怎樣處理JDBC的交易？
5.  JDBC裏的PreparedStatement和Statement之間有什麼區別？   

## 說明知識
{{< ask_chatgpt >}}
說明JDBC並舉例
{{< /ask_chatgpt >}}



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

## 彙總重點
{{< ask_chatgpt >}}
條列JDBC的重點
{{< /ask_chatgpt >}}



1. JDBC是Java Database Connectivity的縮寫，它提供了與關係型數據庫（如MySQL，Oracle等）進行交互的標準接口。
2. JDBC使用一組API來訪問數據庫的操作，這些API包括Connection，Statement或PreparedStatement，和ResultSet等。
3. Connection是 JDBC 中最基本的要素之一，它表示與數據庫之間的連接。該接口還包括其他許多重要功能，例如事務管理和自動提交。
4. Statement或PreparedStatement用於執行SQL查詢或更新數據庫記錄。PreparedStatement是一種比Statement更進一步的對象，它允許開發人員編寫可變參數的SQL語句。
5. ResultSet則用於處理查詢結果集，從中檢索數據以進行顯示或計算等操作。
6. JDBC程序實際上是通過建立Connection對象，創建Statement對象並執行SQL查詢來與數據庫進行交互。並且一定要在用完後關閉所有的資料庫資源，包括Connection，Statement和ResultSet。
7. 除了純JDBC之外，也可以使用Spring JDBC等框架來更簡化和加快JDBC的開發。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題JDBC的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 如何使用JDBC從數據庫中查詢特定日期之後的所有記錄?
答案：使用PreparedStatement對時間戳進行查詢，例如：
```
String sql = "SELECT * FROM myTable WHERE datetime > ?";
PreparedStatement pstmt = conn.prepareStatement(sql); 
pstmt.setTimestamp(1, new Timestamp(date.getTime())); 
```
其中，conn是已經建立好的數據庫連接對象，date是特定的日期對象。

2. 如何使用JDBC在數據庫中執行批量更新操作?
答案：使用PreparedStatement對象的addBatch()和executeBatch()方法，例如：
```
String sql = "UPDATE myTable SET name = ?, age = ? WHERE id = ?";
PreparedStatement pstmt = conn.prepareStatement(sql); 
for (int i = 0; i < data.size(); i++) { //data為需要更新的數據
    pstmt.setString(1, data.getName(i)); 
    pstmt.setInt(2, data.getAge(i)); 
    pstmt.setInt(3, data.getId(i)); 
    pstmt.addBatch(); 
}
int[] result = pstmt.executeBatch();
```
其中，conn是已經建立好的數據庫連接對象。

3. 如何使用JDBC實現分頁查詢?
答案：在SQL語句中使用LIMIT和OFFSET關鍵字實現分頁，例如：
```
String sql = "SELECT * FROM myTable LIMIT ? OFFSET ?";
PreparedStatement pstmt = conn.prepareStatement(sql);
pstmt.setInt(1, pageSize);
pstmt.setInt(2, (pageNum - 1) * pageSize);
```
其中，pageSize表示每頁顯示的記錄數量，pageNum表示要查詢的頁碼。

4. 如何使用JDBC在數據庫中執行事務?
答案：使用Connection對象的setAutoCommit(false)方法關閉自動提交，然後通過commit()和rollback()方法分別提交和回滾事務，例如：
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

5. 如何使用JDBC從數據庫中查詢有關聯的記錄?
答案：使用JOIN操作實現，例如：
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
其中，t1和t2分別代表要關聯的兩張表，table1Id是t1這張表中用來關聯的id欄位名稱。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇JDBC的網路資料
{{< /ask_chatgpt >}}



1. "JDBC Tutorial"，Oracle官方網站，https://docs.oracle.com/en/database/oracle/oracle-database/21/jjdbc/toc.htm

這篇教學文章介紹了JDBC API的基本概念、連線與查詢資料、使用PreparedStatement、處理結果集等等。

2. "JDBC Tutorial for Beginners with Examples"，JournalDev網站，https://www.journaldev.com/2471/jdbc-tutorial

這篇教學文章介紹從設定JDBC驅動程式到連線資料庫、執行查詢語句、以及使用PreparedStatement、CallableStatement等等的相關操作。

3. "JDBC Connection Pooling with HikariCP"，Baeldung網站，https://www.baeldung.com/hikaricp

這篇文章介紹了如何透過HikariCP進行JDBC連線池化以提高應用程式的效能和穩定性。

4. "JDBC Batch Processing using Java"，GeeksforGeeks網站，https://www.geeksforgeeks.org/jdbc-batch-processing-in-java/

這篇教學文章介紹了如何使用JDBC進行批次處理，在降低資料庫負載的前提下提高效能。

5. "JDBC vs Hibernate: What』s the difference?"，Medium網站，https://medium.com/@natd/error-snapshot-fix-8ffdf2c69cbf

這篇文章比較了JDBC和Hibernate這兩種不同的資料庫存取方式，介紹了它們的優缺點和適用場景。   

