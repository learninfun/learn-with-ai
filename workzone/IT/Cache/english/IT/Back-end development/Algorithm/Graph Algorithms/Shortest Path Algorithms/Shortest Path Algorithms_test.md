

1. Which shortest path algorithm is commonly used for finding the shortest route between two points in a road network? 

Answer: Dijkstra's algorithm 

2. What is the main factor that determines the performance of a shortest path algorithm? 

Answer: It is generally the number of nodes and edges in the graph that the algorithm needs to process. 

3. Can Bellman-Ford algorithm handle negative edge weights? 

Answer: Yes, but the graph must not have any negative cycle, else the algorithm will not output the correct result. 

4. Which is faster in dense graphs, Dijkstra's or Floyd-Warshall algorithm? 

Answer: Floyd-Warshall algorithm is faster in dense graphs as it has a time complexity of O(n^3) whereas Dijkstra's has a time complexity of O((n+m) log n), where n is the number of nodes and m the number of edges. 

5. Which algorithm is preferred for finding the shortest paths from a single source node to all other nodes in a graph? 

Answer: Dijkstra's algorithm is not optimal for this case, whereas Bellman-Ford algorithm is a good choice especially if the graph has negative edge weights.