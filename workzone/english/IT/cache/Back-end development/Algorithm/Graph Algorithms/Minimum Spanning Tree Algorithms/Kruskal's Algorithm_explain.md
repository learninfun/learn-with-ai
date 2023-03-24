

Kruskal's Algorithm is a greedy algorithm used for finding the minimum spanning tree (MST) of a connected, weighted graph. The MST is the tree that spans all the nodes in the graph with the minimum possible total edge weight.

The Kruskal's Algorithm works as follows:

1. Sort all the edges in non-decreasing order of their weight.
2. Create a forest of trees where each tree consists of one vertex (i.e., n disjoint sets).
3. Loop through each edge in the sorted list of edges and check if the two vertices of the edge belong to different sets (i.e., different trees). If yes, join the two trees and add the edge to the MST.
4. Continue until all the edges have been checked and the MST is obtained.

Example:
Consider the following graph with 6 vertices and 9 edges, where each edge is weighted as follows:

    A ----- 4 ----- B
    |               |
   2|              5|
    |               |
    C ----- 3 ----- D
     \             /
      1\          /2
        \       /
         \    /
          \ /
           E
           
To find the MST using Kruskal's Algorithm, we first sort the edges in non-decreasing order of their weight:

Edges:   (C, E)  (A, C)  (C, D)  (A, B)  (B, D)  (D, E)  (B, E)  (A, D)  (B, C)
Weights:   1       2       3       4       5       2       5       6       7

Next, we create a forest of trees consisting of each vertex as a separate tree:

    A       B       C       D       E

Then, we start checking the edges in the sorted order and adding them to the MST if they connect different trees. The first edge (C, E) joins tree C to tree E:

    A       B       C       D   
                            |   
                            E   

The second edge (A, C) joins tree A to tree C:

    B       C      
    |               
    A               
                  D                 
                  |   
                  E   

The third edge (C, D) joins tree C to tree D:

    B       C      
    |               
    A       D           
                  |         
                  E   

The fourth edge (A, B) joins tree A to tree B:

        A           
       /|\           
        |            
        C           
        |            
        B           
                  D                 
                  |         
                  E   

The fifth edge (B, D) joins tree B to tree D:

        A               
       /|\             
        |              
        C             
        |              
        B               
        |              
        D                
        |                
        E               

The sixth edge (D, E) joins tree D to tree E:

        A               
       /|\             
        |              
        C             
       /|\            
        |             
        B               
        |              
        D  --- E        

The seventh edge (B, E) cannot be added as it would create a cycle.

The eighth edge (A, D) cannot be added as it connects two nodes already in the same tree.

The ninth edge (B, C) cannot be added as it connects two nodes already in the same tree.

Therefore, the MST consists of the edges (C, E), (A, C), (C, D), (A, B), (B, D), and (D, E), with a total edge weight of 17.