

Dynamic MST (Minimum Spanning Tree) is an algorithm that updates the minimum spanning tree of a graph in response to some changes in the graph such as insertions or deletions of edges. The algorithm only updates the affected portions of the tree and avoids recomputing in areas that are unchanged.

For example, consider a graph G with 5 vertices {A, B, C, D, E} and the following edges:

AB = 2
AC = 1
BC = 3
CD = 4
CE = 5
DE = 6

The minimum spanning tree of G is {A->B, A->C, B->C, C->D, C->E}. Now, suppose that the edge from B to C is removed from the graph. We can use a Dynamic MST algorithm to update the tree. The algorithm will compute the new minimum spanning tree for G' (the graph with the edge removed) as {A->B, A->C, C->D, C->E}.

The key insight of Dynamic MST is that the minimum spanning tree of a graph is a tree that includes all the vertices and a subset of the edges that connect them. When an edge is added or removed, the minimum spanning tree can be updated using the existing sub-tree and the new or deleted edge. This enables the algorithm to update the MST efficiently without recomputing the entire tree.

Dynamic MST has many applications such as network routing, image processing, and data compression.