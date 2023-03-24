

JDBC stands for Java Database Connectivity. It is a standard Java API mainly used for accessing a relational database from a Java application. With JDBC, developers can perform various database operations like querying data, inserting/updating data, and executing database transactions. It provides a set of classes and interfaces to interact with different databases using SQL statements.

Here's an example of JDBC implementation:

import java.sql.Connection;
import java.sql.DriverManager;

public class JdbcExample {
   public static void main(String[] args) {
      Connection conn = null;
      try {
         // Register JDBC driver
         Class.forName("com.mysql.jdbc.Driver");

         // Open connection
         System.out.println("Connecting to database...");
         conn = DriverManager.getConnection("jdbc:mysql://localhost/testdb","root", "password");

         // Execute a query
         System.out.println("Creating statement...");
         Statement stmt = conn.createStatement();
         String sql = "SELECT * FROM Employee";
         ResultSet rs = stmt.executeQuery(sql);

         // Extract data from result set
         while(rs.next()) {
            // Retrieve by column name
            int id = rs.getInt("id");
            String name = rs.getString("name");
            int age = rs.getInt("age");

            // Display values
            System.out.print("ID: " + id);
            System.out.print(", Name: " + name);
            System.out.println(", Age: " + age);
         }

         // Free resources
         rs.close();
         stmt.close();
         conn.close();
      } catch (SQLException se) {
         se.printStackTrace();
      } catch (Exception e) {
         e.printStackTrace();
      } finally {
         // Close resources
         try {
            if (stmt != null)
               stmt.close();
         } catch (SQLException se) {
            se.printStackTrace();
         }
         try {
            if (conn != null)
               conn.close();
         } catch (SQLException se) {
            se.printStackTrace();
         }
      }
   }
}

This example demonstrates how to connect to a MySQL database using JDBC, execute a SQL statement, and retrieve the records from the result set.