1. What is the Longest Common Subsequence (LCS) of the strings 'AGGTAB' and 'GXTXAYB'?
Answer: 'GTAB' (Length = 4)

2. What is the time complexity of the dynamic programming approach for finding the LCS of two strings?
Answer: O(mn), where m and n are the lengths of the two strings.

3. Can we use the LCS algorithm to compare more than two strings at once? If yes, how?
Answer: Yes, we can use the LCS algorithm to compare more than two strings by comparing any two strings at a time and then taking their LCS, and continuing this process until we have compared all the strings.

4. Is the LCS of two strings unique?
Answer: No, there can be multiple LCSs of two strings, but all will have the same length.

5. Can the LCS algorithm be used for comparing strings of different lengths? If yes, how?
Answer: Yes, the LCS algorithm can be adapted to compare strings of different lengths by padding the shorter string with spaces or a special character to make it equal in length to the longer string.