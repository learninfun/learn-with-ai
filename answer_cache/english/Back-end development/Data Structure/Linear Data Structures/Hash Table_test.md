

1. What is the difference between separate chaining and open addressing in hash tables?

Answer: Separate chaining involves creating linked lists at each index of the hash table to store any elements that hash to that index. Open addressing, on the other hand, involves probing for an empty slot in the hash table when a collision occurs, and then storing the element in that slot.

2. In a hash table with linear probing, what is the danger of having a low "load factor"?

Answer: A low load factor means that the hash table is not fully utilized, and there are many empty slots. In a hash table with linear probing, if all the elements are clustered towards one end of the table, it can lead to a phenomenon called clustering, where successive collisions are more likely to fall in the same cluster, making the hash table inefficient.

3. How does a hash table handle collisions?

Answer: When a hash table encounters a collision, it uses a collision resolution mechanism to resolve it. Common methods include separate chaining (keeping linked lists at each index), open addressing (probing for an empty slot), and double hashing (using a second hash function to compute a different index).

4. What is the time complexity of inserting an element in a hash table?

Answer: The time complexity of inserting an element in a hash table is O(1) on average, assuming a good choice of hash function and a low load factor. However, in the worst case of all elements hashing to the same index, the time complexity can be O(n), where n is the number of elements in the hash table.

5. What happens if two elements have the same hash code in a hash table?

Answer: If two elements have the same hash code in a hash table, then a collision occurs. This means that both elements are trying to be stored at the same index. The hash table uses a collision resolution mechanism to resolve this, such as separate chaining or open addressing, as mentioned earlier.