## 習題預習
{{< ask_chatgpt >}}
給我5題Brute Force的問題
{{< /ask_chatgpt >}}



1. 問題：求一個數組中的最大值。解法：遍歷整個數組，找到最大的元素。

2. 問題：猜數字遊戲。解法：從 0 到 9 枚舉每一個可能的答案，直到猜中為止。

3. 問題：找到全排列中字典序最大的排列。解法：枚舉每一個排列，並找到字典序最大的排列。

4. 問題：找到一個整數是否為質數。解法：枚舉從 2 到該整數平方根的所有數字，判斷是否能整除該整數。

5. 問題：計算兩個有理數的和。解法：將兩個有理數表示為分數的形式，通分後求和。   

## 說明知識
{{< ask_chatgpt >}}
說明Brute Force並舉例
{{< /ask_chatgpt >}}



Brute Force（蠻力破解）是一種基於暴力嘗試所有可能性的攻擊方法，通常用於破解密碼或尋找密碼的漏洞。

在Brute Force攻擊中，攻擊者會盡可能多地嘗試所有可能的密碼，直到找到正確的密碼或密碼漏洞。這種攻擊方法可能需要花費大量的時間和計算資源，因此通常僅在重要資料或敏感系統上使用。

舉例來說，如果你忘記了你的電子郵件密碼，可以使用Brute Force方法來嘗試各種可能的密碼。攻擊者可以將所有可能的字母和數字組合成一個列表，然後使用程式將這些密碼一一嘗試，直到找到正確的密碼為止。然而，這種方法通常需要很長時間，因為可能會嘗試數以百萬計的密碼才能找到正確的密碼。   

## 彙總重點
{{< ask_chatgpt >}}
條列Brute Force的重點
{{< /ask_chatgpt >}}



1. Brute Force是一種暴力搜尋演算法，透過枚舉所有可能性來解決問題。

2. Brute Force通常是適用於資料量較小，且計算能力較強的場景。

3. Brute Force演算法的優點是簡單易懂，不需要複雜的數學推導，容易實現。

4. Brute Force演算法的缺點是時間複雜度高，當解空間很大時，計算量會非常龐大，甚至無法接受。

5. 為了解決Brute Force演算法的時間複雜度問題，可以採用分治、回溯等技術對其進行優化。

6. Brute Force演算法常用於排序、搜索、加密解密等場景，對於一些複雜的問題，則需要其他更加高效的演算法來求解。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Brute Force的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 題目：給定一個正整數n，求出1到n中所有質數的總和。
   答案：該問題可以使用埃氏篩法（Eratosphenes）進行解答，即對於每個數字，遍歷所有小於它的正整數，如果該正整數是其因數，就標記為非質數。最後將所有未被標記的數字相加即可。
   
2. 題目：給定一個由'a'和'b'組成的字符串s，請計算s中有多少個子串，滿足該子串中'a'的個數等於'b'的個數。
   答案：該問題可以使用暴力枚舉法進行解答，即對於s的每個子串，都計算其中'a'和'b'的個數，如果相等，就將答案加1。
   
3. 題目：給定一個由非負整數組成的數組nums，請找到其中的一個子串，使得該子串的和最大。
   答案：該問題可以使用暴力枚舉法進行解答，即對於nums的每個子串，都計算其和，找到最大的那個子串即可。
   
4. 題目：給定一個正整數n，求解1到n的所有排列，使得相鄰兩個數字之和是質數。
   答案：該問題可以使用回溯法進行解答，即從1開始遍歷所有排列，對於每個排列，檢查相鄰兩個數之和是否為質數，如果是，就加入答案列表。
   
5. 題目：給定一個由'a'和'b'組成的字符串s，請計算s中最長的連續子串，滿足其中'a'的個數等於'b'的個數。
   答案：該問題可以使用暴力枚舉法進行解答，即對於s的每個子串，都計算其中'a'和'b'的個數，如果相等，就繼續向後拓展，求出最長的符合條件的連續子串長度。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Brute Force的網路資料
{{< /ask_chatgpt >}}



1. "Brute Force Attack". OWASP Foundation. https://owasp.org/www-community/attacks/Brute_force_attack
2. "Brute-Force Attacks Explained". Imperva. https://www.imperva.com/learn/application-security/brute-force-attack/
3. "The Basics of Brute Force Attacks". ESET. https://www.eset.com/int/blog/threat-explainer/the-basics-of-brute-force-attacks/
4. "How Brute Force Attacks Work and How to Stop Them". Norton. https://us.norton.com/internetsecurity-malware-what-is-a-brute-force-attack.html
5. "What is Brute Force Hacking and How can you Prevent It?". Techradar. https://www.techradar.com/news/what-is-brute-force-hacking-and-how-can-you-prevent-it   

