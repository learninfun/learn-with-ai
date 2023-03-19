

1. Bellman-Ford Algorithm is used to find the shortest path from a starting vertex to all other vertices in a weighted graph.
2. It is a dynamic programming algorithm that iteratively computes the shortest path distances by relaxing the edges of the graph in a topological order.
3. It can handle negative edge weights, which is not possible with Dijkstra's Algorithm.
4. If the graph contains a negative cycle, then Bellman-Ford Algorithm can detect it, and indicate that there is no shortest path.
5. The time complexity of the algorithm is O(V*E), where V is the number of vertices and E is the number of edges in the graph.
6. The algorithm returns a list of shortest path distances from the starting vertex to all other vertices in the graph.