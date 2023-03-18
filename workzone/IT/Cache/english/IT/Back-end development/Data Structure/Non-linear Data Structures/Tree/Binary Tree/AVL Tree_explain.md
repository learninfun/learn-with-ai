

An AVL tree is a balanced binary search tree where the height of the left and right sub-tree of every node differs by at most one. It was named after the inventors Adelson-Velsky and Landis.

When a node is inserted or deleted from an AVL tree, the tree is rebalanced by performing rotations. A rotation is a tree operation that preserves the BST property and balances the height of the tree.

Here's an example of an AVL tree:

         10
       /    \
      6      15
     / \     /
    4   8  12
          /  \
         11  14

This tree is balanced because the height of the left sub-tree is 2 and the height of the right sub-tree is 3, which differs by at most 1.

If a new node is inserted with value 7, the tree would be rebalanced by performing a left rotation at node 6:

         10
       /    \
      8      15
     / \     /
    6   9  12
   / \
  4   7
        \
         8

After the rotation, the tree is still balanced with the height of the left sub-tree at 2 and the height of the right sub-tree at 3.