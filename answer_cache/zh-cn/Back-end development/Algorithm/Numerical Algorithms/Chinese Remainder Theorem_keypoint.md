

1. Chinese Remainder Theorem: 對於一個正整數m，如果它是兩個正整數a和b（不一定互素）的最小公倍數，那麼可以通過以下方式解決下列同餘方程組：

 x ≡ a1 (mod m1)
 x ≡ a2 (mod m2)
 …
 x ≡ an (mod mn)

2. Chinese Remainder Theorem的解是唯一的，且可以通過以下方式構造：

 x = a1M1y1 + a2M2y2 + … + anMnyn mod m

其中Mi = m / mi，yi是Mi模mi的乘法逆元。

3. 如果m1，m2，…，mn是互不相同的質數，那麼通過前兩個重點中提到的方法解決同餘方程組的計算很快，因為每個Mi都只有一個質因數。

4. Chinese Remainder Theorem廣泛用於加密和數學上的問題解決，例如RSA加密算法。