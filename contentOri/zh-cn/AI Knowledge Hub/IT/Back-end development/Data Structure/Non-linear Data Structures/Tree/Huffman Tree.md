+++
title = "哈夫曼树"
+++
## 习题预习
{{< ask_chatgpt >}}
给我5题Huffman Tree的问题
{{< /ask_chatgpt >}}



1. 如何建立一个Huffman Tree给以下字元的 frequencies: A(6), B(3), C(2), D(1), E(1)?
2. 如果一个Huffman Tree有1百万个字元，那么此树的深度是多少？
3. 如何在Huffman Tree中，找到某些特定字元的编码？
4. 如果一个Huffman Tree是压缩一个二进制字符流，需要多少位元才能储存此树的节点？
5. 什么是Huffman coding的时间和空间复杂度？   

## 说明知识
{{< ask_chatgpt >}}
说明Huffman Tree并举例
{{< /ask_chatgpt >}}



Huffman Tree（霍夫曼树）是一种二叉树，用于压缩数据和编码。它的结构和模样如下图所示：

[![huffman tree](https://upload.wikimedia.org/wikipedia/commons/thumb/0/02/Huffman_tree_2.svg/600px-Huffman_tree_2.svg.png)](https://en.wikipedia.org/wiki/File:Huffman_tree_2.svg)

Huffman Tree 由一个权值列表（通常是字符的出现频率）建立而成。透过一系列的操作，可以把权值最小的两个节点合并为一个新节点，新节点的权值等于它的两个子节点的权值之和。这个新节点会被插入回原权值列表中，并重复上述操作，直到只剩下一个节点，即为霍夫曼树的根节点。

对于一个待压缩的文本，可以透过霍夫曼编码将每个字符用另一个代表它的二进制码替代，使得密码簿的长度缩短，节省储存空间。具体方法是在霍夫曼树中，左子节点代表的二进制码为 0，右子节点为 1，将每个字符所对应的路径即可构成其二进制码。

举例来说，假设有一个文本 "aaabbcdddd"，每个字母的权值为：

- a: 3
- b: 2
- c: 1
- d: 4

则可以建立出以下的霍夫曼树：

[![example of huffman tree](https://i.imgur.com/fOCinF1.png)](https://i.imgur.com/fOCinF1.png)

从上图可知，字符 a 的二进制码为 0，字符 b 为 10，字符 c 为 110，字符 d 为 111。纪录压缩后的二进制码即可将原文本压缩。在解压缩时，透过霍夫曼树的路径，即可将每个二进制码替换成原本的字符，还原原文本。   

## 汇总重点
{{< ask_chatgpt >}}
条列Huffman Tree的重点
{{< /ask_chatgpt >}}



1. Huffman Tree 是一種二叉树，用於進行有效的資料壓縮和解壓縮。

2. 每個節點都有一個權值，該權值表示對應字符的出現頻率。

3. 民采用貪婪算法來構建Huffman Tree，即從權值最小的兩個節點開始合併，直到所有節點都結合成一個樹。

4. 每個節點都是由兩個子節點組成，左子樹表示第一位二進制數0，右子樹表示第一位二進制數1。

5. 通過對Huffman Tree的遍歷，可以編碼每個字符，使其代表的二進制序列最短。

6. 解碼時，根據遍歷Huffman Tree的路徑來查找相應的字符，並將它們組合成原始數據。

7. Huffman Tree是一種可變長編碼，字符的編碼長度根據出現頻率不同而不同，出現頻率越高的字符編碼越短。

8. Huffman Tree廣泛應用於數據壓縮和傳輸中，使數據傳輸更快，減少存儲空間，並提高數據使用效率。   

## 知识测验
{{< ask_chatgpt >}}
给我5题Huffman Tree的中等难度问题，并在后面列出答案
{{< /ask_chatgpt >}}



1. 將如下項目的頻率存儲在數組中: {a: 5, b: 9, c: 12, d: 13, e: 16, f: 45} 創建一個Huffman Tree，並求出“f”字符的Huffman編碼。
答案: Huffman Tree如下圖所示， “f”字符的編碼為“0”:
```
         ________0_______
        /                \
   ____1____          ___0___
  /         \        /       \
d:13         e:16  ____1     c:12
                 /     \ 
                b:9     a:5
```

2. 將如下項目的頻率存儲在數組中：{a: 12, b: 5, c: 15, d: 13, e: 11, f: 3}創建一個Huffman Tree，求出項目“b”和“e”的Huffman編碼。
答案: Huffman Tree如下圖所示，“b”字符的編碼為“111”，“e”字符的編碼為“00”:
```
         _________0_________
        /                   \
    ___1___             ____0___
   /       \           /        \
c:15      ___1___     d:13      e:11
         /       \  
        a:12     ___0___
                /       \
               b:5       f:3
```

3. 創建一個Huffman Tree並確定項目的Huffman編碼。該示例中給出項目和相應的頻率：{a: 50, b: 20, c: 40, d: 30, e: 10}
答案: Huffman Tree如下圖所示，“a”的編碼為“0”，“b”的編碼為“1101”，“c”的編碼為“10”，“d”的編碼為“111”，“e”的編碼為“1100”:
```
               __0__
            /        \
           /          \
        _1_            0
     /      \      /     \
    1      __0__   1      0
  /   \   /     \        /  \
c:40  d:30 a:50 e:10    b:20
```

4. 創建一個Huffman Tree並確定項目的Huffman編碼。該示例中給出項目和相應的頻率：{a: 20, b: 15, c: 25, d: 10, e: 30}
答案: Huffman Tree如下圖所示，“a”的編碼為“1011”，“b”的編碼為“1010”，“c”的編碼為“11”，“d”的編碼為“1101”，“e”的編碼為“0”:
```
                0
           /        \
         _1_        1
      /      \      
     0      __1__   
           /     \  
          0      _1_
               /   \
             a:20 b:15
             /    /   \
            c:25  d:10 e:30
```

5. 創建一個Huffman Tree並確定項目的Huffman編碼。該示例中給出項目和相應的頻率：{a: 40, b: 25, c: 15, d: 10, e: 5, f: 5}
答案: Huffman Tree如下圖所示，“a”的編碼為“0”，“b”的編碼為“11”，“c”的編碼為“101”，“d”的編碼為“1001”，“e”的編碼為“10001”，“f”的編碼為“10000”:
```
                __0__
             /        \
           _1_        _1_
         /     \    /     \
        a:40  b:25   __1__  f:5
                  /     \
                 c:15   __0__
                       /     \
                     d:10   e:5
```   

## 网络数据
{{< ask_chatgpt >}}
给我5篇Huffman Tree的网络数据
{{< /ask_chatgpt >}}



1. Huffman Tree - GeeksforGeeks
https://www.geeksforgeeks.org/huffman-coding-greedy-algo-3/

This article from GeeksforGeeks explains how to build a Huffman Tree, a binary tree used in Huffman coding for data compression. It provides a step-by-step algorithm for constructing a Huffman Tree from a given set of characters and their frequencies.

2. Huffman coding and Huffman trees - Khan Academy
https://www.khanacademy.org/computing/computer-science/informationtheory/moderninfotheory/v/huffman-coding

This video lesson from Khan Academy introduces the concept of Huffman coding and explains how to build a Huffman Tree. It also provides examples of how Huffman coding can be used to compress text and images.

3. How Huffman Coding Works - Computerphile
https://www.youtube.com/watch?v=JsTptu56GM8

This video from Computerphile gives a detailed explanation of Huffman coding and how it works, including how to build a Huffman Tree. It also discusses the advantages of Huffman coding compared to other compression methods.

4. Huffman coding - Wikipedia
https://en.wikipedia.org/wiki/Huffman_coding

The Wikipedia page on Huffman coding provides a thorough explanation of the algorithm used to build a Huffman Tree, as well as the mathematical principles that make it effective for compression. It also includes examples and diagrams to help illustrate the concept.

5. Building a Huffman tree - Rosetta Code
https://rosettacode.org/wiki/Huffman_coding#Python

This code snippet from Rosetta Code provides an example of how to implement the Huffman coding algorithm in Python, including how to build a Huffman Tree. It could be helpful for programmers looking to implement Huffman coding in their own projects.   
