

A Doubly Linked List (DLL) is a type of data structure in which each node contains two links, one pointing to the next node and the other pointing to the previous node. This makes it possible to traverse the list in both forward and backward directions.

Here’s an example of a DLL:

```
    NULL <- head                               tail -> NULL
    NULL <- node0 <-> node1 <-> node2 <-> node3 -> NULL
    NULL <- tail                               head -> NULL
```

In this DLL, we have four nodes, with each node having value and two pointers – prev and next. The head of the list is the node0 and the tail is node3.

One advantage of DLL is that it allows for insertion and deletion of nodes from both ends efficiently. For example, to insert a node between node0 and node1 in the above example, we simply create a new node, set its prev pointer to node0, and its next pointer to node1. Then, we update the prev pointer of node1 to point to the new node, and the next pointer of node0 to point to the new node.

Overall, DLL is useful in scenarios where we need to traverse a list in both directions or perform frequent insertions/deletions at both ends of the list.