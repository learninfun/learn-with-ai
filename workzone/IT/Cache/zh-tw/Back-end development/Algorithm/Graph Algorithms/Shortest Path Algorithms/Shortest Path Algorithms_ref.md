

1. Dijkstra's Algorithm: 
Dijkstra's algorithm is a classic shortest path algorithm that can be used to find the shortest path between two nodes in a graph. It works by initially assigning a "tentative distance" to each node, which is then updated based on the smallest distance found so far. This process continues until the shortest path to the destination node is found. 

Source: https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-greedy-algo-7/

2. Bellman-Ford Algorithm: 
The Bellman-Ford algorithm is another classic shortest path algorithm that can be used to find the shortest path between two nodes in a graph. Unlike Dijkstra's algorithm, the Bellman-Ford algorithm can handle graphs with negative edge weights. It works by iteratively relaxing the edges in the graph until the shortest path to the destination node is found.

Source: https://www.geeksforgeeks.org/bellman-ford-algorithm-dp-23/

3. Floyd-Warshall Algorithm: 
The Floyd-Warshall algorithm is a dynamic programming algorithm that can be used to find the shortest path between all pairs of nodes in a weighted graph. It works by maintaining a matrix of the shortest distances between each pair of nodes, and updating this matrix iteratively until all pairs of nodes have been considered. 

Source: https://www.geeksforgeeks.org/floyd-warshall-algorithm-dp-16/

4. A* Search Algorithm: 
The A* search algorithm is a heuristic search algorithm that can be used to find the shortest path between two points in a graph. It works by assigning a "cost" to each node based on its distance from the starting node and its estimated distance to the destination node. This allows the algorithm to prioritize nodes that are more likely to lead to the shortest path. 

Source: https://www.geeksforgeeks.org/a-search-algorithm/

5. Johnson's Algorithm: 
Johnson's algorithm is a graph algorithm that can be used to find the shortest path between all pairs of nodes in a weighted graph. It works by first reweighting the edges in the graph using a technique called Bellman-Ford's algorithm, and then applying Dijkstra's algorithm to each node in the graph. 

Source: https://www.geeksforgeeks.org/johnsons-algorithm/