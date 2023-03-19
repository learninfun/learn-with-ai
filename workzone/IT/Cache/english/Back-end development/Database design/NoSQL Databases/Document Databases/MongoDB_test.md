

1. What is MongoDB?
Answer: MongoDB is a highly scalable and flexible NoSQL document-oriented database system designed to handle large volumes of unstructured data.

2. What is the difference between MongoDB and relational databases?
Answer: MongoDB is a NoSQL database, while relational databases are SQL-based. MongoDB stores data in documents, which are flexible and allow for nested data structures. Relational databases store data in tables with static columns and rows.

3. How does MongoDB ensure high performance?
Answer: MongoDB uses indexes extensively to ensure quick access to data. Additionally, it uses sharding to distribute data across multiple nodes for high availability and scalability.

4. What is the syntax for creating a collection in MongoDB?
Answer: db.createCollection(name, options) where "name" is the name of the collection and "options" are optional parameters to customize the collection's behavior.

5. How can you import data into MongoDB from a CSV file?
Answer: You can use the mongoimport command-line tool to import data from a CSV file into a MongoDB collection. The syntax is as follows: mongoimport --db DATABASE_NAME --collection COLLECTION_NAME --type csv --file FILE_NAME --headerline.