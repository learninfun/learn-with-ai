

Rabin-Karp算法是一種字符串匹配算法，它是基於哈希算法的。該算法通過對主字符串的各個子串和模式串進行哈希運算，來判斷它們是否相等。具體流程如下：

1. 計算模式串的哈希值。

2. 計算與模式串長度相同的第一個子串的哈希值。

3. 如果這兩個哈希值相等，那麼比較它們是否真的相等。如果相等，則返回子串在主字符串中的位置。

4. 如果哈希值不相等，則計算下一個子串的哈希值，並繼續比較。

舉個例子：

假設模式串是"ABCD"，主字符串是"BCDEABCD"。

1. 計算模式串的哈希值。

由於本例中的字符集較小，我們可以採用簡單的加法哈希。

hash("ABCD") = 'A' + 'B' + 'C' + 'D' = 65 + 66 + 67 + 68 = 266。

2. 計算第一個子串的哈希值。

hash("BCDE") = 'B' + 'C' + 'D' + 'E' = 66 + 67 + 68 + 69 = 270。

3. 比較哈希值。

由於哈希值不相等，我們需要計算下一個子串的哈希值。

hash("CDEA") = 'C' + 'D' + 'E' + 'A' = 67 + 68 + 69 + 65 = 269。

4. 繼續比較哈希值。

hash("DEAB") = 'D' + 'E' + 'A' + 'B' = 68 + 69 + 65 + 66 = 268。

5. 繼續比較哈希值。

hash("EABC") = 'E' + 'A' + 'B' + 'C' = 69 + 65 + 66 + 67 = 267。

6. 繼續比較哈希值。

hash("ABCD") = 'A' + 'B' + 'C' + 'D' = 65 + 66 + 67 + 68 = 266。

由於這兩個哈希值相等，我們需要比較它們是否真的相等。在本例中，它們確實相等，因此子串"ABCD"在主字符串中的起始位置是4。

總結：

Rabin-Karp算法的時間複雜度是O(n+m)，其中n是主字符串的長度，m是模式串的長度。雖然該算法的理論時間複雜度與暴力算法相同，但在實際應用中，Rabin-Karp算法通常比暴力算法更快，尤其是當主字符串和模式串非常大時。