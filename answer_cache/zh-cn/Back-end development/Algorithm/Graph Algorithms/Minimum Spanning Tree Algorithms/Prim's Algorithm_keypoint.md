

1. Prim's Algorithm is used to find the minimum spanning tree of a connected weighted undirected graph.

2. It starts with a single vertex and then expands by adding the closest vertex to the tree until all vertices are included.

3. At each step, the algorithm selects the minimum-weight edge that connects a vertex in the tree to a vertex outside the tree.

4. The algorithm continues until all vertices are either in the tree or have been considered for inclusion.

5. Prim's Algorithm is guaranteed to find the minimum spanning tree for a connected graph with non-negative edge weights.

6. The time complexity of Prim's Algorithm is O(ElogV) using a priority queue or heap, where E is the number of edges and V is the number of vertices in the graph.

7. The output of Prim's Algorithm is a set of edges that forms the minimum spanning tree of the graph.