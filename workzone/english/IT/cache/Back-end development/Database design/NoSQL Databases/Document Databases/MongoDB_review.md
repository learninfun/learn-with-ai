1. QUESTION: How does MongoDB ensure high availability and data reliability?
   ANSWER: MongoDB uses a method called replication, which ensures that multiple copies of data are maintained across different servers. If one server goes down, another can take its place without any data loss.

2. QUESTION: What is sharding in MongoDB and why is it used?
   ANSWER: Sharding is the process of dividing large datasets into smaller ones and distributing them across different servers. It is used for scaling the database horizontally, allowing it to handle larger data volumes and more concurrent requests.

3. QUESTION: Can MongoDB be used for transactional processing?
   ANSWER: Yes, MongoDB offers support for multi-document transactions, enabling developers to perform complex, multi-step operations across multiple documents.

4. QUESTION: How does MongoDB handle data consistency in a distributed environment?
   ANSWER: MongoDB uses a protocol called two-phase commit (2PC) to ensure data consistency in a distributed environment. This ensures that all nodes involved in a transaction agree to commit or rollback a change in a distributed environment.

5. QUESTION: What is the difference between embedded documents and references in MongoDB?
   ANSWER: Embedded documents are documents that are nested within another document, whereas references are links to other documents in a separate collection. Embedded documents are useful for storing related data, while references are better suited for querying and retrieving large amounts of data.