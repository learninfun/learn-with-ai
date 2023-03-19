

1. What is the most commonly used method for primality testing?
Answer: The most commonly used method for primality testing is the Miller-Rabin algorithm.

2. How does the Fermat primality test work?
Answer: The Fermat primality test involves testing whether a given number n is prime by checking if a^(n-1) ≡ 1 (mod n) where a is any integer between 1 and n-1.

3. What is the complexity of the Miller-Rabin primality test?
Answer: The complexity of the Miller-Rabin primality test is O(k log^3 n) where k is the number of iterations required for a given accuracy.

4. Can the Miller-Rabin primality test produce false positives?
Answer: Yes, the Miller-Rabin primality test can produce false positives for composite numbers that behave like primes under the chosen bases used in the test.

5. What is the AKS algorithm for primality testing?
Answer: The AKS algorithm for primality testing is a deterministic algorithm that determines whether a given number n is prime or composite in polynomial time, specifically O(log^6 n).