

1. What is Dijkstra's algorithm used for? 
Answer: Dijkstra's algorithm is used to find the shortest path between two nodes in a graph that has non-negative weight edges.

2. What is the difference between Dijkstra's algorithm and Bellman-Ford algorithm? 
Answer: Dijkstra's algorithm is faster and can only handle non-negative edge weights while Bellman-Ford algorithm can handle negative edge weights but takes longer to compute.

3. What is the time complexity of Dijkstra's algorithm? 
Answer: The time complexity of Dijkstra's algorithm is O(ElogV) where E is the number of edges and V is the number of vertices in the graph.

4. What is the difference between a directed and undirected graph in the context of shortest path algorithms? 
Answer: In an undirected graph, the shortest path between two nodes is the same regardless of the direction of traversal while in a directed graph, the shortest path can differ depending on the direction of travel.

5. How can negative edge weights affect the result of a shortest path algorithm? 
Answer: Negative edge weights can cause the result of a shortest path algorithm to be incorrect as some paths may appear to be shorter due to the negative weights but actually end up taking longer.