

Bellman-Ford Algorithm is a shortest path algorithm used to find the shortest path from a source vertex to all other vertices in a graph with negative weight edges. The algorithm works by relaxing all the edges of the graph in a topological sorting order. It keeps track of the shortest distance from the source vertex to all other vertices in the graph by updating the distance in each relaxation step.

The process of relaxing an edge (u,v) is checking if the distance from the source vertex to v through u is shorter than the current distance from the source vertex to v. If it is, then the distance of v is updated.

The algorithm runs in O(V*E) time, where V is the number of vertices and E is the number of edges in the graph.

Example: Let us consider a graph with 4 vertices (A, B, C, D) and 5 edges as shown below:

         (A)--1-->(B)
         /^\    |
         |  \   2
         |   \  |
         3    \ v
         |    (C)
         |     ^
         1     |
         |     |
         v     |
         (D)--6--|

Let us assume that the source vertex is A. The algorithm works as follows:

1. Initially, the distance to A is 0 and for all other vertices, the distance is set to infinity.
   Distance: A-0, B-inf, C-inf, D-inf
   
2. Relaxation step 1: Relax all the edges of the graph in a topological manner.
   Distance: A-0, B-1, C-3, D-3
   
3. Relaxation step 2: Relax all the edges of the graph in a topological manner.
   Distance: A-0, B-1, C-3, D-3
   
4. Relaxation step 3: Relax all the edges of the graph in a topological manner.
   Distance: A-0, B-1, C-3, D-3
   
5. Relaxation step 4: Relax all the edges of the graph in a topological manner.
   Distance: A-0, B-1, C-3, D-3
   
6. Relaxation step 5: Relax all the edges of the graph in a topological manner.
   Distance: A-0, B-1, C-3, D-3
   
The final output of the algorithm is the distance of all vertices from the source vertex:

Distance: A-0, B-1, C-3, D-3

Therefore, the shortest path from A to other vertices is A-B (weight 1), A-C (weight 3), and A-D (weight 3).