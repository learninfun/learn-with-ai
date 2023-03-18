

A binary tree is a tree-based data structure in which each node has at most two children, known as left child and right child. These children are ordered for a parent node, meaning one child precedes the other. The first node of the tree, known as the root, only has references to its two children, while the other nodes can have references to their parent, left child, and right child.

Here is an example of a binary tree:

```
       1
     /   \
    2     3
   / \   / \
  4   5 6   7
```

In the above example, `1` is the root node, and its left child is `2`, and the right child is `3`. The left subtree of `1` contains nodes `2`, `4`, and `5`, while the right subtree contains nodes `3`, `6`, and `7`. Every child node in the tree has, at most, two children. For instance, `4` and `5` have no children on their own, but are both children of `2`. `6` and `7` have no children, and they are the children of `3`.