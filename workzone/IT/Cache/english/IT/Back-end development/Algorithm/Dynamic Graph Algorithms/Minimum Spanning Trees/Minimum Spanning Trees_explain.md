

Minimum Spanning Tree (MST) is a subset of edges in a weighted undirected graph that connects all the vertices (nodes) together without any loops and with the minimum possible total edge weight. 

Consider the following graph:

```
         2
   (1)--------(2)
   / \        / \
3/    4\   6/    \5
 /       \ /       \
(4)------(3)--------(5)
         7          8
```

To find the MST of this graph, we can use the Kruskal's algorithm, which follows the following steps:

1. Sort all the edges in ascending order of their weights.
2. Pick the smallest edge and check if it forms a cycle with the MST formed so far. If not, add the edge to the MST.
3. Repeat step 2 until all the vertices are included in the MST.

Using Kruskal's algorithm, we can find the MST of the above graph as:

```
   (1)--------(2)
             /   
          6/    
           /      
        (3)        
         | \      
         |5  \8            
         |   \     
        (5)---(4)         
```

Here, the MST has four edges with a total weight of 2+6+5+8=21. This is the minimum possible weight for a spanning tree of this graph.