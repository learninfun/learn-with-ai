

1. Bellman-Ford Algorithm is named after Richard Bellman and Lester Ford Jr.
2. It is a single-source shortest path algorithm that works with graphs that have negative edge weights.
3. The algorithm works on a graph with V vertices and E edges with arbitrary initial vertex s and updates the shortest path distance of every vertex from s.
4. The algorithm relaxes each edge for |V|-1 times to ensure to get the shortest path distance of every vertex.
5. If the algorithm detects a negative-edge cycle in the graph, it returns an error indicating it is impossible to find the shortest path under such circumstances.
6. The time complexity of the algorithm is O(VE) which is not as efficient as Dijkstra's algorithm if edge weights are non-negative.