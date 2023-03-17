

1. What is the Bellman-Ford algorithm used for?
Answer: The algorithm is used to find the shortest path between two nodes in a weighted graph.

2. What is the time complexity of the Bellman-Ford algorithm?
Answer: The time complexity of the algorithm is O(EV), where E is the number of edges and V is the number of vertices in the graph.

3. How does the Bellman-Ford algorithm handle negative weight edges?
Answer: The algorithm can detect the presence of negative weight cycles in a graph and reports an error in case it is present. If there is no negative cycle, then it can report the shortest path even for graphs with negative-weight edges.

4. How does the Bellman-Ford algorithm differ from Dijkstra's algorithm?
Answer: Dijkstra's algorithm is only applicable to graphs without negative weight edges, while Bellman-Ford algorithm can handle negative weights. Also, Dijkstra's algorithm uses a priority queue to determine the minimum distance vertex to visit next, while Bellman-Ford algorithm simply relaxes all edges in each iteration.

5. What is the relaxation process in the Bellman-Ford algorithm?
Answer: The relaxation process is essentially the algorithm's method of updating the distance of each vertex from the source vertex. It is achieved by checking if the distance of a vertex can be improved by relaxing an edge that is connected to it. If the distance can be improved, then the updated distance is recorded.