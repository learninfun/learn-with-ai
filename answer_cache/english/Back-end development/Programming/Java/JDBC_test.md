

1. What is a JDBC driver?
Answer: A JDBC driver is a software component that enables communication between a Java application and a database management system.

2. What is the difference between a Type 1 and Type 4 JDBC driver?
Answer: A Type 1 JDBC driver translates JDBC calls into ODBC calls, while a Type 4 JDBC driver communicates directly with the database through a network protocol.

3. How do you connect to a database using JDBC?
Answer: To connect to a database using JDBC, you need to load the appropriate JDBC driver and then create a connection object using the DriverManager class.

4. What is a PreparedStatement in JDBC?
Answer: A PreparedStatement is a precompiled SQL statement that can be reused multiple times with different input parameters. It provides better performance and security compared to executing SQL statements directly.

5. How do you handle exceptions in JDBC?
Answer: In JDBC, you can handle exceptions using try-catch blocks or by letting the exception propagate to a higher level. Common exception types in JDBC include SQLException, BatchUpdateException, and DataTruncation.