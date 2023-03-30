1. What is the maximum number of children a node in a B+ tree of order m can have?
Answer: A node in a B+ tree of order m can have a maximum of m children.

2. What is the minimum number of keys a non-root internal node in a B+ tree must have?
Answer: A non-root internal node in a B+ tree must have at least ceil(m/2) keys.

3. What is the height of a B+ tree with n keys and order m?
Answer: The height h of a B+ tree with n keys and order m is log_base((m+1)/2)((n+1)/2),rounded up to the nearest integer.

4. How do B+ trees ensure data retrieval efficiency?
Answer: B+ trees ensure data retrieval efficiency by keeping all data in leaf nodes, and by maintaining a balanced tree structure through the use of split and merge operations.

5. What is the difference between a B+ tree and a binary search tree (BST)?
Answer: A B+ tree is a multiway search tree with more than two children per node, while a BST has at most two children per node. Additionally, B+ trees are specifically designed for disk-based storage and optimal retrieval of data, while BSTs are more commonly used for in-memory operations.