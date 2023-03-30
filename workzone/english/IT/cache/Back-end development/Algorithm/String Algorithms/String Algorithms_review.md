1. What is the time complexity of the Boyer-Moore algorithm for string matching?
Answer: The worst-case time complexity of the Boyer-Moore algorithm is O(mn), where m is the length of the pattern and n is the length of the text.

2. What is the difference between Levenshtein distance and Hamming distance?
Answer: Hamming distance is the number of positions at which two strings differ, while Levenshtein distance is the minimum number of operations (insertions, deletions, or substitutions) required to transform one string into another.

3. What is the prefix function in the Knuth-Morris-Pratt algorithm?
Answer: The prefix function is a table that is computed for a pattern and is used to determine the length of the longest proper prefix of the pattern that is also a suffix of a substring. This information is then used to avoid unnecessary comparisons while searching for the pattern in a text.

4. How does the Rabin-Karp algorithm improve over the brute-force algorithm for string matching?
Answer: The Rabin-Karp algorithm uses hashing to compare the pattern and the text, which allows for constant time comparisons. This improves over the O(mn) time complexity of the brute-force algorithm.

5. What is the Burrows-Wheeler transform and how is it useful for compression?
Answer: The Burrows-Wheeler transform rearranges the characters in a string so that they become easier to compress using algorithms like move-to-front or run-length encoding. This is because the transformed string contains many runs of the same characters, which can be represented more compactly.