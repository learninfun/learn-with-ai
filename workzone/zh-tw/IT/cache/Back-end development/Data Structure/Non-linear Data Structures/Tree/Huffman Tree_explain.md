

Huffman Tree（霍夫曼樹）是一種二叉樹，用於壓縮數據和編碼。它的結構和模樣如下圖所示：

[![huffman tree](https://upload.wikimedia.org/wikipedia/commons/thumb/0/02/Huffman_tree_2.svg/600px-Huffman_tree_2.svg.png)](https://en.wikipedia.org/wiki/File:Huffman_tree_2.svg)

Huffman Tree 由一個權值列表（通常是字符的出現頻率）建立而成。透過一系列的操作，可以把權值最小的兩個節點合併為一個新節點，新節點的權值等於它的兩個子節點的權值之和。這個新節點會被插入回原權值列表中，並重複上述操作，直到只剩下一個節點，即為霍夫曼樹的根節點。

對於一個待壓縮的文本，可以透過霍夫曼編碼將每個字符用另一個代表它的二進制碼替代，使得密碼簿的長度縮短，節省儲存空間。具體方法是在霍夫曼樹中，左子節點代表的二進制碼為 0，右子節點為 1，將每個字符所對應的路徑即可構成其二進制碼。

舉例來說，假設有一個文本 "aaabbcdddd"，每個字母的權值為：

- a: 3
- b: 2
- c: 1
- d: 4

則可以建立出以下的霍夫曼樹：

[![example of huffman tree](https://i.imgur.com/fOCinF1.png)](https://i.imgur.com/fOCinF1.png)

從上圖可知，字符 a 的二進制碼為 0，字符 b 為 10，字符 c 為 110，字符 d 為 111。紀錄壓縮後的二進制碼即可將原文本壓縮。在解壓縮時，透過霍夫曼樹的路徑，即可將每個二進制碼替換成原本的字符，還原原文本。