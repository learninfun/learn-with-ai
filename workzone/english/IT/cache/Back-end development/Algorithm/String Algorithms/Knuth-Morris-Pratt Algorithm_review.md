1. Q: What is the time complexity of the Knuth-Morris-Pratt Algorithm?
A: The time complexity of the KMP algorithm is O(m+n), where m is the length of the pattern and n is the length of the string.

2. Q: How does the KMP algorithm avoid unnecessary comparisons?
A: The KMP algorithm constructs a partial match table, which is used to determine the maximum possible shift when a mismatch occurs. This avoids unnecessary comparisons by skipping over already matched characters.

3. Q: Does the KMP algorithm work for patterns containing repeating characters?
A: Yes, the KMP algorithm can handle patterns containing repeating characters, as it only relies on the length of the longest proper suffix that matches the proper prefix.

4. Q: What happens if the pattern is longer than the string in KMP algorithm?
A: If the pattern is longer than the string in the KMP algorithm, the algorithm will return that the pattern is not found in the string.

5. Q: What is the space complexity of the KMP algorithm?
A: The space complexity of the KMP algorithm is O(m), where m is the length of the pattern. This is due to the partial match table that is constructed, which requires m space.