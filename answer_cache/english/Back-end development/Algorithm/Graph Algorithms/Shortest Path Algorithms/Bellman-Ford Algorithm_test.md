

1. What is the Bellman-Ford algorithm used for?
Answer: The Bellman-Ford algorithm is a shortest path algorithm used to find the shortest path from a single source node to all other nodes in a weighted, directed graph.

2. What is the main difference between Dijkstra's algorithm and Bellman-Ford algorithm?
Answer: Dijkstra's algorithm is used to find the shortest path from a single source node to a single target node, while Bellman-Ford algorithm is used to find shortest path from a single source node to all other nodes in a graph.

3. What is the time complexity of the Bellman-Ford algorithm?
Answer: The time complexity of the Bellman-Ford algorithm is O(N * E), where N is the number of nodes and E is the number of edges in the graph.

4. How does the Bellman-Ford algorithm handle negative weight cycles?
Answer: The Bellman-Ford algorithm detects the presence of negative weight cycles in a graph and reports a negative cycle present in the graph, if any. It does not compute the shortest path in case negative weight cycles are present.

5. Can the Bellman-Ford algorithm be used on undirected graphs?
Answer: Yes, the Bellman-Ford algorithm can be used on undirected graphs by converting them to a directed graph by considering each edge as a pair of directed edges with opposite directions and same weight.