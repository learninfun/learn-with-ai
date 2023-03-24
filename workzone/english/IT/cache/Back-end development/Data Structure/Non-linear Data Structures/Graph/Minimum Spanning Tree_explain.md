

Minimum Spanning Tree (MST) is a graph algorithm that creates a tree from a weighted undirected graph. An MST is a subgraph of the original graph that connects all vertices with the minimum total edge weight. 

For example, consider the following weighted undirected graph with six vertices: 

    A---6---B
    |   /   |
    1  3    5
    | /     |
    C---2---D
    |       
    4       

One possible MST of this graph is: 

        B
      /   \
    A       D
      \   /
        C          

This MST has a total edge weight of 6+3+2+1 = 12, which is the minimum possible weight that connects all vertices. 

Note that there may be multiple MSTs for a given graph, but all MSTs will have the same total edge weight.