+++
title = "Knuth-Morris-Pratt演算法"
+++
## 習題預習
{{< ask_chatgpt >}}
給我5題Knuth-Morris-Pratt Algorithm的問題
{{< /ask_chatgpt >}}



1. 解釋Knuth-Morris-Pratt Algorithm是什麼，並描述其基本思想。

2. 給定一個文本字符串T和一個模式字符串P，請描述如何使用Knuth-Morris-Pratt Algorithm來查找P在T中的出現。

3. 如何實現Knuth-Morris-Pratt Algorithm的failure function？具體說明這個函數的用途和算法原理。

4. 如果文本字符串T非常長，而模式字符串P在其中出現的次數很少，如何優化Knuth-Morris-Pratt Algorithm以提高查找效率？

5. 考慮在字母表Σ中搜索模式字符串P的所有循環位移的問題。請描述如何使用Knuth-Morris-Pratt Algorithm來解決這個問題。   

## 說明知識
{{< ask_chatgpt >}}
說明Knuth-Morris-Pratt Algorithm並舉例
{{< /ask_chatgpt >}}



Knuth-Morris-Pratt (KMP)算法是解決字符串匹配問題的高效算法之一，用於在主串中查找模式串的匹配位置。

算法思想：

KMP算法的基本思想是不回溯主串指針i，而是通過在模式串中預先構建出一個next[]數組，來記錄模式串的自匹配情況。當發生不匹配時，i指針不回溯，只需移動一定距離，這樣既能提高匹配效率，又能減少匹配次數。

next[]數組的構建：

next[i]表示在i位置之前的子串中，前綴和後綴最長的匹配長度。例如：P = 「ABAB」，在i=3(第四個字符)的時候，前綴是"ABA"，後綴是「BA」，並且兩者長度相等，所以next[3]=2。

構建next[]數組的過程可以使用動態規劃的思想，通過遞推的方式求解。該過程可以大致分為以下幾步：

1）next[0] = -1；next[1] = 0；

2）從i=2開始逐一計算，設i-1位置的next值為nextval，則有：

(1)如果P[i-1] == P[nextval]，則next[i] = nextval+1；

(2)如果P[i-1] != P[nextval]且nextval != -1，則需要回退，計算next[P[nextval]]；

(3)如果P[i-1] != P[nextval]且nextval = -1，則next[i]=0；

模式串匹配：

當模式串的next[]數組構建完成後，就可以在主串S中查找模式串P的匹配位置，這個過程相對簡單：

1）初始化主串指針i=0，模式串指針j=0；

2）如果S[i] == P[j]，則i++,j++；

3）如果S[i] != P[j]，則需要移動模式串指針j，根據next[j]值來決定移動距離；

4）如果j等於模式串P的長度，說明匹配成功，返回主串當前位置i與模式串長度的差值即可；否則繼續匹配。

示例：

主串S：BBC ABCDAB ABCDABCDABDE

模式串P：ABCDABD

next[]數組：[-1,0,0,0,0,1,2,0]

假設在主串S的第9個位置開始查找，匹配過程如下：

S  B  C  D  A  B  C  D  A  B  C  D  A  B  D  E
   P  A  B  C  D  A  B  D

1) S[9] != P[0]，移動模式串指針，j=next[0]= -1+1=0；

S  B  C  D  A  B  C  D  A  B  C  D  A  B  D  E
         P  A  B  C  D  A  B  D

2) S[9] == P[0], S[10] == P[1], S[11] == P[2]，經過三輪比較，匹配成功，返回i與模式串長度的差值，即9。

時間複雜度：

KMP算法的時間複雜度為O(n+m)，其中n和m分別是主串和模式串的長度。相比於樸素的字符串匹配算法，KMP算法具有更高的效率和更低的時間複雜度，適用於大規模字符串匹配問題的解決。   

## 彙總重點
{{< ask_chatgpt >}}
條列Knuth-Morris-Pratt Algorithm的重點
{{< /ask_chatgpt >}}



1. KMP算法是一種字符串匹配算法，用於在文本字符串中查找子串的出現位置。
2. 該算法利用已經匹配的信息來避免在搜索時無用的匹配操作，從而提高效率。
3. KMP算法中最重要的部分是構建一個表格，該表格可以告訴我們在匹配失敗時應該跳到哪個位置繼續搜索。
4. 構建這個表格的過程可以看作是一個自我匹配的過程，我們利用已經匹配好的信息來尋找一些規律並記錄在表格中。
5. KMP算法的時間複雜度是O(n+m)，其中n和m分別是文本和子串的長度。
6. KMP算法的空間複雜度是O(m)，其中m是子串的長度。
7. KMP算法可以應用於多種數據結構和問題，例如字符串匹配、查找重複字串等。   

## 知識測驗
{{< ask_chatgpt >}}
給我5題Knuth-Morris-Pratt Algorithm的中等難度問題，並在後面列出答案
{{< /ask_chatgpt >}}



1. 問題： 給定一個文本串s和一個模式串p，在s中查找是否存在一個子串與p匹配。
   答案: KMP算法
   解答: 
   KMP算法的基本想法是，當匹配錯誤發生時，已經匹配的那些字符很可能包含有關我們可以跳過的位置的信息，因此可以使用這些信息進行匹配，跳過一些不必要的字符。 
   
   
2. 問題：給定一個字符數組和一個模式串，返回模式在字符數組中的所有匹配位置（字符串可以重疊）。
   答案: KMP算法
   解答:
   KMP算法是一種用於查找字符串中子串的快速算法。基於該算法，可以實現一個函數或類，用於在文章中查找指定的關鍵字，或在編輯器中查找指定的代碼片段。 
   
3. 問題：給定一個文本串和一個模式串，將文本串中與模式串匹配的字符替換為另一指定字符。
   答案: KMP算法
   解答:
   KMP算法通常使用在字符串匹配問題上。但是，當使用KMP算法將一個字符串保存到另一個字符串中時，也可以將其視為字符串匹配問題。可以逐個字母遍歷字符串，將找到的字符替換成另一個字符。 
   
4. 問題：給定一個文本串和一個模式串，查找和計算文本串中所有匹配模式串的字符。
   答案: KMP算法
   解答:
   KMP算法是一種查找字符串中子串的快速算法。基於該算法，可以實現一個函數或類，用於在文章中查找指定的關鍵字，或在編輯器中查找指定的代碼片段。在計算每個匹配字符時，可以使用額外的數組。 
   
5. 問題：給定一個文本串和一個模式串，查找文本串中所有匹配模式串的字符，並返回匹配的字符個數。
   答案: KMP算法
   解答:
   KMP算法可以用於查找文本串中所有匹配模式串的字符。使用KMP算法時，可以選擇是否需要計算和返回所有匹配字符的個數。在使用KMP算法時，可以要求函數或類將計數和返回值分開處理。   

## 網路資料
{{< ask_chatgpt >}}
給我5篇Knuth-Morris-Pratt Algorithm的網路資料
{{< /ask_chatgpt >}}



1. 維基百科 - Knuth-Morris-Pratt算法
https://zh.wikipedia.org/wiki/Knuth-Morris-Pratt%E7%AE%97%E6%B3%95

2. GeeksforGeeks - Knuth-Morris-Pratt (KMP) Algorithm
https://www.geeksforgeeks.org/kmp-algorithm-for-pattern-searching/

3. TutorialsPoint - Knuth-Morris-Pratt Algorithm
https://www.tutorialspoint.com/Knuth-Morris-Pratt-algorithm

4. Brilliant - Knuth-Morris-Pratt Algorithm
https://brilliant.org/wiki/knuth-morris-pratt-algorithm/

5. Programiz - Knuth-Morris-Pratt (KMP) Algorithm
https://www.programiz.com/dsa/knuth-morris-pratt-algorithm   
