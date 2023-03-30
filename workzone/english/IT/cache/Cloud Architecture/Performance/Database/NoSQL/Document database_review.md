1. What is a key-value store in a document database?
A key-value store is a data model in a document database where each document is represented as a key-value pair. The key is a unique identifier for the document, and the value is the data associated with that document.

2. What is sharding in document databases, and what advantages does it offer?
Sharding is a technique used to distribute data across multiple nodes in a distributed system. In document databases, sharding allows for horizontal scalability and improved performance by spreading the data and workload across multiple machines instead of relying on a single server.

3. How are indexes used in document databases, and what benefits does indexing offer?
Indexes in document databases are used to speed up query performance by quickly locating data that matches the query criteria. They work by storing a sorted representation of the data, allowing for faster access and reducing the amount of data that needs to be scanned.

4. What is denormalization in document databases, and how does it impact performance?
Denormalization is the process of duplicating data across multiple collections or documents in a document database, typically to improve performance by reducing the need for joins or aggregations. It can improve read performance but may increase storage costs and complexity.

5. How are transactions handled in document databases, and what limitations do they have?
Document databases typically support atomic operations on single documents, but transactional consistency across multiple documents is more challenging. Some document databases offer transaction support through features like snapshot isolation, but these solutions may have limitations such as reduced performance or complexity.