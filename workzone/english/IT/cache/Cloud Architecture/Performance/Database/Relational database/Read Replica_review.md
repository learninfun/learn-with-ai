1. What is a Read Replica?

Answer: A Read Replica is a relational database instance that is created to offload read traffic from a primary instance.

2. What is the difference between a Read Replica and a backup of a database?

Answer: A Read Replica is an active copy of a database that can handle read-only queries, while a backup is a passive copy of a database that does not accept queries.

3. Can a Read Replica be used for failover purposes?

Answer: No, Read Replica cannot be used for failover purposes. It is used for scaling and offloading read traffic from the primary instance.

4. How does a Read Replica keep up with the changes in the primary instance?

Answer: A Read Replica uses asynchronous replication to keep up with changes in the primary instance. Changes are made to the primary instance and then replicated to the Read Replica.

5. Can a Read Replica be in a different region or availability zone from the primary instance?

Answer: Yes, it is possible to have a Read Replica in a different region or availability zone from the primary instance. This provides redundancy and can help with disaster recovery.