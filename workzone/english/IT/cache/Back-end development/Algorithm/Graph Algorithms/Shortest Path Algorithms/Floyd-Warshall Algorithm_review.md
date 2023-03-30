1. What is the time complexity of the Floyd-Warshall Algorithm?

Answer: The time complexity of the Floyd-Warshall Algorithm is O(n^3), where n is the number of vertices in the graph.

2. What is the purpose of using the Floyd-Warshall Algorithm?

Answer: The purpose of using the Floyd-Warshall Algorithm is to find the shortest path between all pairs of vertices in a weighted graph.

3. What is the significance of the diagonal of the distance matrix computed by the Floyd-Warshall Algorithm?

Answer: The diagonal of the distance matrix computed by the Floyd-Warshall Algorithm represents the shortest path between a vertex and itself, which is always 0.

4. Can the Floyd-Warshall Algorithm be used to detect negative cycles in a graph?

Answer: Yes, the Floyd-Warshall Algorithm can be used to detect negative cycles in a graph. If there is a negative cycle in the graph, the distance matrix computed by the algorithm will have negative values along the diagonal.

5. How does the Floyd-Warshall Algorithm handle disconnected graphs?

Answer: The Floyd-Warshall Algorithm handles disconnected graphs by treating each disconnected component as a separate graph and computing the shortest path between all pairs of vertices in each component separately. The resulting distance matrix will have infinite (or very large) values for pairs of vertices that are not connected by any path.