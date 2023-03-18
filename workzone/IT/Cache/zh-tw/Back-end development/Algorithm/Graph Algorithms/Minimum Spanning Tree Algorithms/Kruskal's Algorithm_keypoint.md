

1. Kruskal's Algorithm is used to find the Minimum Spanning Tree (MST) of a given graph.

2. The algorithm works by sorting the edges of the graph by weight, and then adding the edges with the lowest weight first, as long as they do not create a cycle.

3. To determine whether an edge will create a cycle, the algorithm uses a Union-Find data structure.

4. The running time of Kruskal's Algorithm is O(E log E) where E is the number of edges in the graph.

5. Kruskal's Algorithm is a greedy algorithm, meaning that it makes the locally optimal choice at each step, in the hope that this will lead to a globally optimal solution.

6. The MST of a graph is a tree that spans all the vertices of the graph, while minimizing the total weight of the edges.

7. Kruskal's Algorithm can handle graphs with both positive and negative edge weights, but will only work on connected graphs.

8. Kruskal's Algorithm can be used in a variety of applications, including network design, clustering, and image segmentation.