+++
title = "米勒-拉宾质数判定算法"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Miller-Rabin Algorithm的问题
{{< /ask_chatgpt >}}



1. What is the Miller-Rabin algorithm used for? 
2. How does the Miller-Rabin algorithm differ from other primality testing algorithms? 
3. Can the Miller-Rabin algorithm determine the exact prime factorization of a number? 
4. How can the Miller-Rabin algorithm be implemented efficiently in code? 
5. Are there any limitations or drawbacks to using the Miller-Rabin algorithm compared to other primality testing algorithms?   

## 说明知识
{{< ask_chatgpt >}}
说明Miller-Rabin Algorithm并举例
{{< /ask_chatgpt >}}



Miller-Rabin Algorithm是一種概率性質質數測試法，用於檢測一個數字是否為質數。它基於Fermat's Little Theorem和二次剩餘的性質，在確定一個數字是否為質數時，比起其他傳統的質數測試算法，如暴力測試法或埃拉托色尼質數篩法，更加高效。

Miller-Rabin Algorithm的演算法流程如下：

1. 將欲判斷的數字n-1表示成 2^s * d 的形式，其中d為奇數。
2. 隨機選擇一個小於n的整數a作為底數。
3. 計算a^d mod n，若 a^d ≡ 1 (mod n) 或 a^(2^r * d) ≡ -1 (mod n) 就判定這是質數或可能是錯誤的。其中r∈[0,s-1]
4. 如果驗證不是質數，則繼續對a重新選擇底數，重複步驟2-3，若經過k次驗證皆為質數或可能是錯誤的，就判定n為質數；若經過k次驗證存在不合法的結果，就判定n為合數。

以下是一個Miller-Rabin Algorithm的例子：

我們想要確認60是否為質數。

1. 60-1可以表示成 2^2 * 15的形式。
2. 選擇一個隨機底數a = 2。
3. 計算2^15 mod 60，得到32。
4. 因為2^30 ≡ 1 (mod 60)，所以我們可以進行以下的判斷式：

  2^2 * 15 = 60-1
  2^15 ≡ 32 (mod 60)
  2^30 ≡ 1 (mod 60)
  
  因此，2^60 ≡ (2^30)^2 ≡ 1 (mod 60)，所以60可能是錯誤的質數，我們需要繼續進行驗證。
  
5. 我們再次選擇底數a = 7。
6. 計算7^15 mod 60，得到49。
7. 我們也可以進行以下的判斷式：

  7^2 ≡ 49 ≡ -11 (mod 60)
  7^4 ≡ (-11)^2 ≡ 121 ≡ 1 (mod 60)
  
  因此，7^15 ≡ (7^4) * (7^4) * 7^2 ≡ 1 (mod 60)，所以60可能是錯誤的質數，我們需要再次繼續驗證。
  
8. 我們再次選擇底數a = 11。
9. 計算11^15 mod 60，得到49。
10. 進行以下判斷：

  11^2 ≡ 1 (mod 60)
  11^4 ≡ 1 (mod 60)
  
  因此，11^15 ≡ (11^4) * (11^4) * 11^2 ≡ 1 (mod 60)，所以我們可以確定60是錯誤的合數。

因此，我們可以確定60是錯誤的合數，不是質數。 如此反覆執行，直到分別執行了k次測試都沒有成功，我們才能夠大致得到一個數字是否是質數。   

## 汇总重点
{{< ask_chatgpt >}}
条列Miller-Rabin Algorithm的重点
{{< /ask_chatgpt >}}



1. Miller-Rabin算法是一種用於快速判定一個數是否為素數的概率算法。
2. 這種算法基於費馬小定理以及阿基米德反演等數論知識，通過隨機選取若干個基數來檢測是否為素數。
3. Miller-Rabin算法針對偽素數的概率給出了一個極小值，因此在實際應用中已經被廣泛使用。
4. 通過設置適當的檢測次數，可以將檢測出錯的概率降到極低，從而確定一個數是否為素數。
5. 由於複雜度相對較低，Miller-Rabin算法被廣泛地應用在計算機理論、加密系統等領域。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Miller-Rabin Algorithm的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 对于一个给定的奇数 n，假设 a 是其中一个小于 n 的整数。如何使用 Miller-Rabin Algorithm 验证 a 是否是 n 的一个质数证人？

答案：以 k 次随机选取的 a 为例，将 n - 1 表示为 d * 2^s 的形式，其中 d 是奇数，s >= 1。接下来用快速幂算法计算 a^d mod n，如果得到的结果等于 1 或 n - 1，停止计算，并认为 a 是 n 的可能质数证人。如果结果不等于 1 或 n - 1，计算 a^(2^r * d) mod n，直到 r = s 或得到的结果等于 n - 1。如果最后得到的结果等于 n - 1，则仍认为 a 是 n 的可能质数证人。否则认为 a 不是 n 的质数证人。

2. 使用 Miller-Rabin Algorithm 求一个大质数（大于 10^10）的可能质数证人。 

答案：多次随机选取 a 进行测试，如果每次测试都得到 a 为 n 的可能质数证人，则有很高的概率认为 n 是质数。要注意选取的 a 需要小于 n，可使用随机数生成器达到这个目的。

3. 对于一个给定的奇数 n，使用 Miller-Rabin Algorithm 找到一个质数 p，使得 n - p 是一个平方数。

答案：如果 n 是质数，则 p = 2 是说是可以的。如果 n 不是质数，则可以随机选取 a 进行测试。如果得到 a 是 n 的一个可能质数证人，则有很高概率认为 n 是合数。此时可以计算 b = sqrt(n - 1)，如果 b 是一个整数，则 p = n - b^2 是一个质数。

4. 对于一个给定的质数 p，使用 Miller-Rabin Algorithm 判断是否存在一个 a，使得 a 是 p 的原根。

答案：对于 p 的每个质因子 q，如果 q = 2 或 p / q = 2，则 p 中不存在原根。否则，可以使用随机选取的 a 进行测试。如果得到 a 是 p 的可能质数证人，则有很高概率认为 a 是 p 的原根。此时计算 phi(p) 和其所有质因子的乘积，如果  a^((p-1) / q) mod p 不为 1，对于所有的质因子 q，则认为 a 是 p 的原根。

5. 对于一个给定的奇数 n，使用 Miller-Rabin Algorithm 判断是否存在一个奇数 g，使得 g 是 n 的原根。

答案：可以随机选取一个质数 p，使得 p 是 n 的一个质因子。然后选取一个小于 p 的整数 a，计算 g = a^((p-1) * (n-1) / phi(n)) mod n。如果 g 不等于 1，则有很高概率认为 g 是 n 的原根。否则，换一个 a 重新计算。如果 p 不能选出，则需要用更高级的算法解决问题。   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Miller-Rabin Algorithm的网络数据
{{< /ask_chatgpt >}}



1. GeeksforGeeks: Miller-Rabin Primality Test - A set of articles explaining the Miller-Rabin algorithm and its implementation in various programming languages. Includes a detailed explanation of the algorithm and its complexity, along with a Python code implementation.

2. Brilliant.org: Miller-Rabin Primality Test - A concise explanation of the Miller-Rabin algorithm along with a Java code implementation. Includes examples of how to use the algorithm to determine whether a large number is prime, and how to find the smallest prime number larger than a given number.

3. Khan Academy: Miller-Rabin Primality Test - A video lesson explaining the Miller-Rabin algorithm and its complexity. Includes examples of how to use the algorithm to determine whether a number is prime or composite.

4. Rosetta Code: Miller-Rabin Primality Test - A collection of programming examples in various languages demonstrating how to implement the Miller-Rabin algorithm. Includes C++, Java, Python, Ruby, and many other popular programming languages.

5. Wikipedia: Miller-Rabin Primality Test - An overview of the Miller-Rabin algorithm with a description of its history, complexity, and applications. Includes references to related topics in number theory and cryptography.   
