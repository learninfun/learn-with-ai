1. What is JDBC and how does it operate with databases?
Answer: JDBC stands for Java Database Connectivity, which is a Java API that provides a standard way for Java applications to interact with relational databases. JDBC uses a driver manager and JDBC driver to connect to the database, execute SQL statements, and retrieve results.

2. What are the different types of JDBC drivers?
Answer: There are 4 types of JDBC drivers: 
- Type 1: JDBC-ODBC bridge driver (uses ODBC driver to connect to the database)
- Type 2: Native API driver (connects to the database using the database vendor's API)
- Type 3: Network protocol driver (uses a middleware server to translate JDBC calls into DBMS-specific protocols)
- Type 4: Pure Java driver (connects to the database using a pure Java implementation of the database protocol)

3. How do you establish a connection to a database using JDBC?
Answer: To establish a connection to a database using JDBC: 
- Load the JDBC driver class 
- Create a Connection object 
- Connect to the database using the URL, username, and password 
- Close the Connection object when finished.

4. What is a prepared statement in JDBC?
Answer: A prepared statement in JDBC is a precompiled SQL statement that can be used repeatedly with different parameter values. It is useful for executing statements with parameters, which enhances performance and improves security by preventing SQL injection attacks.

5. What is a ResultSet in JDBC?
Answer: A ResultSet in JDBC is an object that contains the results of a query. It is used to iterate over and retrieve the data from the query results. The ResultSet object is created when a Statement object is executed and is closed when either the Statement or Connection object is closed.