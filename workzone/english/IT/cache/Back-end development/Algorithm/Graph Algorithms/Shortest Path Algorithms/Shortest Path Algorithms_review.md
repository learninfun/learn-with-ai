1. What is the difference between Dijkstra's Algorithm and Bellman-Ford Algorithm? 

Answer: Dijkstra's Algorithm is a single-source, unweighted shortest path algorithm, while Bellman-Ford Algorithm is a single-source, weighted shortest path algorithm that can handle negative edge weights.

2. What is the time complexity of Floyd-Warshall Algorithm? 

Answer: The time complexity of Floyd-Warshall Algorithm is O(V^3), where V is the number of vertices in the graph.

3. Describe the concept of a relaxation in the context of Shortest Path Algorithms.

Answer: Relaxation is the process of updating the tentative distance of a node when a shorter path to it is found during the execution of the algorithm. This is a key step in many Shortest Path Algorithms, including Dijkstra's Algorithm and Bellman-Ford Algorithm.

4. What is the role of a priority queue in Dijkstra's Algorithm? 

Answer: A priority queue is used in Dijkstra's Algorithm to maintain a list of nodes sorted by their tentative distance from the source vertex. This allows the algorithm to efficiently select the node with the shortest tentative distance in each step.

5. Can Shortest Path Algorithms be used to solve the Traveling Salesman Problem? 

Answer: No, Shortest Path Algorithms are not suitable for solving the Traveling Salesman Problem, since the objective of the problem is to find the shortest Hamiltonian cycle that visits every node exactly once, rather than finding the shortest path between two nodes. This is a much more complex problem that requires specialized algorithms, such as the Held-Karp Algorithm.