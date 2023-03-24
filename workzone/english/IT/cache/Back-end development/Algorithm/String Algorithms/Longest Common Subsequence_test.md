

1. What is a Longest Common Subsequence (LCS)?

Answer: A LCS is a sequence of characters that are found in two or more given strings in the same relative order, but not necessarily contiguous or the same length.

2. What is the difference between LCS and Longest Common Substring (LCS)?

Answer: LCS is based on finding the longest string of characters that are common to two or more strings, while LCS is based on finding the longest subsequence of characters that are common to two or more strings.

3. How can you compute the LCS of two given strings in a dynamic programming approach?

Answer: You can compute the LCS of two strings using a dynamic programming approach that involves creating a matrix of the lengths of common substrings between the two strings, and then tracing back the matrix to find the LCS.

4. What is the time complexity of computing the LCS using the dynamic programming approach?

Answer: The time complexity of computing the LCS using the dynamic programming approach is O(mn), where m and n are the lengths of the two input strings.

5. How can you extend the LCS problem to include more than two input strings?

Answer: You can extend the LCS problem to include more than two input strings by creating a multi-dimensional matrix that includes one dimension for each input string, and then computing the LCS using a recursive formula that compares all possible common subsequences across all input strings.