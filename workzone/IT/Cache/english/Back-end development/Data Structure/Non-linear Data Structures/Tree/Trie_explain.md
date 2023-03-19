

Trie, also known as Prefix tree, is a tree-based data structure used for efficient searching and storing of strings. The name "Trie" comes from the word "Retrieval."

In a Trie, each node of the tree represents a string or a prefix of a string. The nodes are linked together to form a path that represents a complete string. The root is an empty string (''), and all the child nodes are characters that are part of the string. Each leaf node represents a complete word.

For example, consider the following list of words:

- bat
- bar
- bark
- bath
- badger

To create a Trie for these words, we start with an empty root node. Then, we add each word to the Trie by inserting its characters one by one. The resulting Trie would look like this:

       ('')
       /  |  |  |  \
     (b)(a)(b)(a)(d)
      |    |   |    |
     (a)  (r) (t)  (g)
          |    |
         (k)  (h)

In this Trie, the path from the root to the node 'ba' represents the word 'ba', and the path from the root to the node 'bar' represents the word 'bar'. Similarly, the path to the node 'bat' represents the word 'bat', and so on. The node 'bad' is not a leaf node as it only represents a prefix and does not form a complete word. The last 'g' node does not have any meaning as we do not have any word that continues from that node. 

Tries are commonly used in applications such as search engines, auto-complete features, and spelling checkers. They help in reducing the time complexity by eliminating redundant comparisons as we traverse the trie.