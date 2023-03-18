

The Rabin-Karp algorithm is a string searching algorithm that matches a pattern in a text by comparing hash values of substrings of the text and the pattern. It is considered a probabilistic algorithm as it may have false positives, but it has average-case time complexity of O(n+m), where n is the length of the text and m is the length of the pattern.

The algorithm works as follows:

1. Compute the hash value of the pattern.

2. Compute the hash values of all substrings of the text with the same length as the pattern.

3. Compare the hash value of the pattern with the hash values of the substrings. If there is a match, compare the actual characters of the pattern and substring to ensure it is not a false positive.

4. If the pattern is not found, slide the pattern over the text one character at a time and repeat steps 2-3.

Example:

Text: ABCDEFGHJKLFGFGFGFGFGFGFGFGFGFGFGFGHIJKLMN

Pattern: FGFGFGFGFG

Hash function: sum of ASCII values of characters

Step 1: Compute the hash value of the pattern

FGFGFGFGFG = 70+71+70+71+70+71+70+71+70+71 = 700

Step 2: Compute the hash values of all substrings of the text with the same length as the pattern

Substrings: ABCDEFGHJKL, BCDEFGHJKLF, CDEFGHJKLFG, DEFGHJKLFGF, EFGHJKLFGFG, FG...

Hash values: 65+66+67+68+69+70+71+72+74+75+76+70 = 804
66+67+68+69+70+71+72+74+75+76+70+71 = 810
67+68+69+70+71+72+74+75+76+70+71+70 = 812
...

Step 3: Compare the hash value of the pattern with the hash values of the substrings

The hash value of FGFGFGFGFG is 700. One of the substrings with the same hash value is FGFGFGFGFG. We then compare the actual characters of the pattern and substring to confirm it is not a false positive.

Step 4: If the pattern is not found, slide the pattern over the text one character at a time and repeat steps 2-3

We continue to slide the pattern over the text one character at a time until we find all occurrences of the pattern.