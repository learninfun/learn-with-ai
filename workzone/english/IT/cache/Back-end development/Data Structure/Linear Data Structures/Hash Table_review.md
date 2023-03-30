1. What is a hash table?
Answer: A hash table is a data structure that uses a hash function to map keys to values.

2. How does a hash table handle collisions?
Answer: When there is a collision (i.e., when two or more keys map to the same value), a hash table can either use chaining or open addressing to resolve the collision.

3. What is the worst-case time complexity for a hash table lookup operation?
Answer: The worst-case time complexity for a hash table lookup operation is O(n), where n is the size of the table.

4. How does the load factor affect the performance of a hash table?
Answer: The load factor (i.e., the ratio of the number of items in the hash table to the size of the table) can impact the performance of a hash table. As the load factor increases, the likelihood of collisions increases, which can slow down the lookup operation.

5. What is a good hash function?
Answer: A good hash function is one that generates a unique hash code for each key and distributes the hash codes evenly across the table. It should also be fast to compute and produce minimal collisions.