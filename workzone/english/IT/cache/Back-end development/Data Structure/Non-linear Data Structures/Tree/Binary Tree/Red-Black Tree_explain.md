

Red-Black Tree is a self-balancing Binary Search Tree in which each node has one additional color attribute, either red or black, to ensure that the tree is balanced over time. The balance is achieved by applying specific rules during the insertion and deletion operations that ensure that the height of the tree remains logarithmic, O(log n), where n is the number of elements in the tree.

The main properties of Red-Black Tree are:

1. Every node is either red or black.
2. The root is always black.
3. Every leaf (NULL) is black.
4. If a node is red, then both its children are black.
5. Every simple path from a node to a descendant leaf contains the same number of black nodes.

An example of a Red-Black Tree:

Let's say we want to insert the following elements in a Red-Black Tree: 30, 40, 50, 20, 60, 70, 10, 25. 

1. We start by creating a new node for 30 and color it black since it is the root of the tree.
       30(B)

2. Insert 40 as its right child and color it red.
       30(B)
            \
            40(R)

3. Insert 50 as the right child of 40 and color it black.
       30(B)
            \
            40(B)
                \
                50(R)

4. Insert 20 as the left child of 30 and color it black.
       30(B)
      /    \
  20(B)  40(B)
              \
              50(R)

5. Insert 60 and color it red as the right child of 50. Then, apply a Left Rotation on node 40 to fix the violation of the Black Height property.
       30(B)
      /    \
  20(B)  50(B)
        /    \
    40(R) 60(R)

6. Insert 70 as the right child of 60 and color it black.
            50(B)
           /    \
       30(R) 60(B)
      /     \     \
  20(B)  40(B) 70(R)
  
7. Insert 10 as the left child of 20 and color it red. Then, apply a Right Rotation on node 30 and a Left Rotation on node 50 to fix the double violation of the Red-Black Tree properties.
             50(B)
            /     \
          30(R)  60(B)
         /    \     \
     20(B) 40(B)   70(R)
    /  
 10(R)