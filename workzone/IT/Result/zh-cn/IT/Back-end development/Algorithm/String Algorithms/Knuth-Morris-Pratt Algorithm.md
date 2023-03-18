+++
title = "Knuth-Morris-Pratt演算法"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Knuth-Morris-Pratt Algorithm的问题
{{< /ask_chatgpt >}}



1. 解释Knuth-Morris-Pratt Algorithm是什么，并描述其基本思想。

2. 给定一个文本字符串T和一个模式字符串P，请描述如何使用Knuth-Morris-Pratt Algorithm来查找P在T中的出现。

3. 如何实现Knuth-Morris-Pratt Algorithm的failure function？具体说明这个函数的用途和算法原理。

4. 如果文本字符串T非常长，而模式字符串P在其中出现的次数很少，如何优化Knuth-Morris-Pratt Algorithm以提高查找效率？

5. 考虑在字母表∑中搜索模式字符串P的所有循环位移的问题。请描述如何使用Knuth-Morris-Pratt Algorithm来解决这个问题。   

## 说明知识
{{< ask_chatgpt >}}
说明Knuth-Morris-Pratt Algorithm并举例
{{< /ask_chatgpt >}}



Knuth-Morris-Pratt (KMP)算法是解决字符串匹配问题的高效算法之一，用于在主串中查找模式串的匹配位置。

算法思想：

KMP算法的基本思想是不回溯主串指针i，而是通过在模式串中预先构建出一个next[]数组，来记录模式串的自匹配情况。当发生不匹配时，i指针不回溯，只需移动一定距离，这样既能提高匹配效率，又能减少匹配次数。

next[]数组的构建：

next[i]表示在i位置之前的子串中，前缀和后缀最长的匹配长度。例如：P = “ABAB”，在i=3(第四个字符)的时候，前缀是"ABA"，后缀是“BA”，并且两者长度相等，所以next[3]=2。

构建next[]数组的过程可以使用动态规划的思想，通过递推的方式求解。该过程可以大致分为以下几步：

1）next[0] = -1；next[1] = 0；

2）从i=2开始逐一计算，设i-1位置的next值为nextval，则有：

(1)如果P[i-1] == P[nextval]，则next[i] = nextval+1；

(2)如果P[i-1] != P[nextval]且nextval != -1，则需要回退，计算next[P[nextval]]；

(3)如果P[i-1] != P[nextval]且nextval = -1，则next[i]=0；

模式串匹配：

当模式串的next[]数组构建完成后，就可以在主串S中查找模式串P的匹配位置，这个过程相对简单：

1）初始化主串指针i=0，模式串指针j=0；

2）如果S[i] == P[j]，则i++,j++；

3）如果S[i] != P[j]，则需要移动模式串指针j，根据next[j]值来决定移动距离；

4）如果j等于模式串P的长度，说明匹配成功，返回主串当前位置i与模式串长度的差值即可；否则继续匹配。

示例：

主串S：BBC ABCDAB ABCDABCDABDE

模式串P：ABCDABD

next[]数组：[-1,0,0,0,0,1,2,0]

假设在主串S的第9个位置开始查找，匹配过程如下：

S  B  C  D  A  B  C  D  A  B  C  D  A  B  D  E
   P  A  B  C  D  A  B  D

1) S[9] != P[0]，移动模式串指针，j=next[0]= -1+1=0；

S  B  C  D  A  B  C  D  A  B  C  D  A  B  D  E
         P  A  B  C  D  A  B  D

2) S[9] == P[0], S[10] == P[1], S[11] == P[2]，经过三轮比较，匹配成功，返回i与模式串长度的差值，即9。

时间复杂度：

KMP算法的时间复杂度为O(n+m)，其中n和m分别是主串和模式串的长度。相比于朴素的字符串匹配算法，KMP算法具有更高的效率和更低的时间复杂度，适用于大规模字符串匹配问题的解决。   

## 汇总重点
{{< ask_chatgpt >}}
条列Knuth-Morris-Pratt Algorithm的重点
{{< /ask_chatgpt >}}



1. KMP算法是一种字符串匹配算法，用于在文本字符串中查找子串的出现位置。
2. 该算法利用已经匹配的信息来避免在搜索时无用的匹配操作，从而提高效率。
3. KMP算法中最重要的部分是构建一个表格，该表格可以告诉我们在匹配失败时应该跳到哪个位置继续搜索。
4. 构建这个表格的过程可以看作是一个自我匹配的过程，我们利用已经匹配好的信息来寻找一些规律并记录在表格中。
5. KMP算法的时间复杂度是O(n+m)，其中n和m分别是文本和子串的长度。
6. KMP算法的空间复杂度是O(m)，其中m是子串的长度。
7. KMP算法可以应用于多种数据结构和问题，例如字符串匹配、查找重复字串等。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Knuth-Morris-Pratt Algorithm的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 問題： 給定一個文本串s和一個模式串p，在s中查找是否存在一個子串與p匹配。
   答案: KMP算法
   解答: 
   KMP算法的基本想法是，當匹配錯誤發生時，已經匹配的那些字符很可能包含有關我們可以跳過的位置的信息，因此可以使用這些信息進行匹配，跳過一些不必要的字符。 
   
   
2. 問題：给定一个字符数组和一个模式串，返回模式在字符数组中的所有匹配位置（字符串可以重叠）。
   答案: KMP算法
   解答:
   KMP算法是一種用於查找字符串中子串的快速算法。基於該算法，可以實現一個函數或類，用於在文章中查找指定的關鍵字，或在編輯器中查找指定的代碼片段。 
   
3. 問題：给定一个文本串和一个模式串，将文本串中与模式串匹配的字符替换为另一指定字符。
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

## 网络数据
{{< ask_chatgpt >}}
给我5篇Knuth-Morris-Pratt Algorithm的网络数据
{{< /ask_chatgpt >}}



1. 维基百科 - Knuth-Morris-Pratt算法
https://zh.wikipedia.org/wiki/Knuth-Morris-Pratt%E7%AE%97%E6%B3%95

2. GeeksforGeeks - Knuth-Morris-Pratt (KMP) Algorithm
https://www.geeksforgeeks.org/kmp-algorithm-for-pattern-searching/

3. TutorialsPoint - Knuth-Morris-Pratt Algorithm
https://www.tutorialspoint.com/Knuth-Morris-Pratt-algorithm

4. Brilliant - Knuth-Morris-Pratt Algorithm
https://brilliant.org/wiki/knuth-morris-pratt-algorithm/

5. Programiz - Knuth-Morris-Pratt (KMP) Algorithm
https://www.programiz.com/dsa/knuth-morris-pratt-algorithm   

