

1. What is the time complexity of Floyd-Warshall algorithm?
Answer: The time complexity of Floyd-Warshall algorithm is O(n^3), where n is the number of vertices in the graph.

2. What is the minimum possible distance between two vertices in a graph using Floyd-Warshall algorithm?
Answer: The minimum possible distance between two vertices using Floyd-Warshall algorithm is 0 if the vertices are the same, or the weight of the direct edge if there exists one, otherwise it is the sum of weights of the shortest path between them.

3. What is the optimal sequence of vertices in a path from 'A' to 'B' using Floyd-Warshall algorithm?
Answer: Floyd-Warshall algorithm does not provide the sequence of vertices in a path. Rather, it provides the shortest distance between any two vertices in the graph.

4. What is the output of Floyd-Warshall algorithm when applied on a disconnected graph?
Answer: The output of Floyd-Warshall algorithm on a disconnected graph is a matrix with infinite distances between disconnected vertices.

5. What is the advantage of using Floyd-Warshall algorithm over Dijkstra's algorithm for finding shortest path in a graph?
Answer: Floyd-Warshall algorithm can handle negative edges, while Dijkstra's algorithm cannot. Additionally, Floyd-Warshall algorithm can find the shortest paths between all pairs of vertices in the graph in a single run, while Dijkstra's algorithm requires multiple runs for each pair of vertices.