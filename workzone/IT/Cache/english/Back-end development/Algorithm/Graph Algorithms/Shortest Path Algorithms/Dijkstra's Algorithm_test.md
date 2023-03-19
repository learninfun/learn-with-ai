

1. What is Dijkstra's algorithm?
Answer: Dijkstra's algorithm is a shortest path algorithm that finds the shortest path between a source node and all other nodes in a weighted graph.

2. What is the time complexity of Dijkstra's algorithm?
Answer: The time complexity of Dijkstra's algorithm is O(E log V), where E is the number of edges and V is the number of vertices in the graph.

3. How does Dijkstra's algorithm handle negative weight edges in a graph?
Answer: Dijkstra's algorithm does not work with negative weight edges. It assumes that all edges have non-negative weights. However, there are modified versions of Dijkstra's algorithm that can handle negative weight edges, such as Bellman-Ford algorithm.

4. What is the main difference between Dijkstra's algorithm and A* algorithm?
Answer: Dijkstra's algorithm finds the shortest path between a source node and all other nodes in the graph, while A* algorithm uses heuristics to find the shortest path between a source node and a target node in the graph.

5. Can Dijkstra's algorithm find the shortest path in a directed graph?
Answer: Yes, Dijkstra's algorithm can find the shortest path in a directed graph. However, the shortest path may not be unique if there are multiple paths with the same length.