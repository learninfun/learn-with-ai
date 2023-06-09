

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