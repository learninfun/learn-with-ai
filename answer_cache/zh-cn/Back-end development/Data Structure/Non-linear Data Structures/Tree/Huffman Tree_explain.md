

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