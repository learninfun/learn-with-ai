

1. What type of graph does Dijkstra's algorithm work on?
Answer: It works on directed weighted graphs.

2. What is the time complexity of Dijkstra's algorithm?
Answer: The time complexity is O(E + V log V), where E is the number of edges and V is the number of vertices in the graph.

3. How does Dijkstra's algorithm find the shortest path in a graph?
Answer: It uses a priority queue to implement a greedy algorithm and explore the graph, always choosing the next vertex with the lowest cost.

4. Can Dijkstra's algorithm work on graphs with negative edge weights?
Answer: No, it cannot handle negative edge weights as the greedy nature of the algorithm might lead to incorrect shortest path calculations.

5. What is the main difference between Dijkstra's algorithm and the Bellman-Ford algorithm?
Answer: Dijkstra's algorithm uses a greedy approach, where it always takes the lowest cost option. In contrast, the Bellman-Ford algorithm uses dynamic programming to determine the shortest path, thus handling negative edge weights too.