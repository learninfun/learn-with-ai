

1. Chinese Remainder Theorem: 对于一个正整数m，如果它是两个正整数a和b（不一定互素）的最小公倍数，那么可以通过以下方式解决下列同余方程组：

 x ≡ a1 (mod m1)
 x ≡ a2 (mod m2)
 …
 x ≡ an (mod mn)

2. Chinese Remainder Theorem的解是唯一的，且可以通过以下方式构造：

 x = a1M1y1 + a2M2y2 + … + anMnyn mod m

其中Mi = m / mi，yi是Mi模mi的乘法逆元。

3. 如果m1，m2，…，mn是互不相同的质数，那么通过前两个重点中提到的方法解决同余方程组的计算很快，因为每个Mi都只有一个质因数。

4. Chinese Remainder Theorem广泛用于加密和数学上的问题解决，例如RSA加密算法。