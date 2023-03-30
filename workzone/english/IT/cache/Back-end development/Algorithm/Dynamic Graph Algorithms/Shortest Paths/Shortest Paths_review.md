1. What is the difference between Dijkstra's algorithm and the Bellman-Ford algorithm for finding the shortest path in a graph? 
Answer: Dijkstra's algorithm works on graphs with non-negative edge weights, while Bellman-Ford can handle graphs with negative edge weights. 

2. How is the Floyd-Warshall algorithm able to find the shortest path between all pairs of vertices in a graph efficiently? 
Answer: The algorithm uses dynamic programming to compute the shortest path between all pairs of vertices by considering intermediate vertices along the way. 

3. In what situations might the A* algorithm be more efficient than Dijkstra's algorithm for finding the shortest path in a graph? 
Answer: A* is often faster than Dijkstra's when there is additional information available about the graph, such as a heuristic estimate of the distance to the target node. 

4. What is the time complexity of the Johnson's algorithm for finding the shortest paths in a graph? 
Answer: Johnson's algorithm has a time complexity of O(V²logV + VE), where V is the number of vertices and E is the number of edges in the graph. 

5. How can the idea of graph edge contraction be used in finding the shortest path in a graph? 
Answer: Edge contraction can simplify a graph by merging two nodes into a single supernode, reducing the number of edges in the graph. This can make it easier to compute the shortest path using algorithms such as Dijkstra's.