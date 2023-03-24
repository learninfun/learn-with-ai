

1. What is the basic idea behind the Rabin-Karp algorithm?
Answer: The Rabin-Karp algorithm is a string matching algorithm that uses hashing to compare the pattern to the text. It performs pattern matching by comparing the hash values of the pattern and substrings of the text.

2. What is the worst-case time complexity of the Rabin-Karp algorithm?
Answer: The worst-case time complexity of the Rabin-Karp algorithm is O(mn), where m is the length of the pattern and n is the length of the text.

3. How does the Rabin-Karp algorithm handle collisions in hash values?
Answer: The Rabin-Karp algorithm uses a rolling hash function that allows it to handle collisions by rehashing only the affected substrings, rather than recomputing the hash values of all substrings.

4. What is the role of modular arithmetic in the Rabin-Karp algorithm?
Answer: The Rabin-Karp algorithm uses modular arithmetic to reduce the hash values to a fixed range, which helps to prevent overflow and ensures that the hash values are uniformly distributed.

5. How does the Rabin-Karp algorithm compare to other string matching algorithms, such as the Knuth-Morris-Pratt algorithm?
Answer: The Rabin-Karp algorithm is generally slower than the Knuth-Morris-Pratt algorithm for small patterns, but it can be faster for larger patterns or when multiple patterns need to be matched simultaneously. Additionally, the Rabin-Karp algorithm is more flexible because it can be easily adapted to handle different types of patterns (e.g., regular expressions).