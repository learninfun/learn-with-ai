

1) What is the time complexity of the Knuth-Morris-Pratt Algorithm?
Answer: The time complexity of the KMP algorithm is O(n).

2) How does the KMP algorithm improve on the brute-force algorithm for string matching?
Answer: The KMP algorithm improves on the brute-force algorithm by using the information derived from the matches so far to avoid recomparing overlapping substrings.

3) What are the two main steps of the KMP algorithm?
Answer: The two main steps of the KMP algorithm are the pre-processing step, which constructs the partial match table, and the matching step, which uses the table to find all occurrences of the pattern within the text.

4) How does the KMP algorithm handle the case where there are multiple occurrences of the pattern within the text?
Answer: The KMP algorithm uses the partial match table to keep track of where the previous matches occurred, and continues the search from the end of the previous match to find all occurrences of the pattern.

5) What is the partial match table, and how is it constructed?
Answer: The partial match table is a data structure used by the KMP algorithm to store information about the pattern that can be used to avoid unnecessary comparisons during string matching. It is constructed by analyzing the pattern itself and finding the longest proper prefix that is also a suffix of each prefix of the pattern.